// Binary ogen generates go source code from OAS.
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen"
	"github.com/ogen-go/ogen/gen/genfs"
	"github.com/ogen-go/ogen/internal/location"
	"github.com/ogen-go/ogen/internal/ogenversion"
	"github.com/ogen-go/ogen/internal/ogenzap"
	"github.com/ogen-go/ogen/internal/urlpath"
)

func cleanDir(targetDir string, files []os.DirEntry) (rerr error) {
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		if !(strings.HasSuffix(name, "_gen.go") || strings.HasSuffix(name, "_gen_test.go")) {
			continue
		}
		if !(strings.HasPrefix(name, "openapi") || strings.HasPrefix(name, "oas")) {
			continue
		}
		// Do not return error if file does not exist.
		if err := os.Remove(filepath.Join(targetDir, name)); err != nil && !os.IsNotExist(err) {
			// Do not stop on first error, try to remove all files.
			rerr = multierr.Append(rerr, err)
		}
	}
	return rerr
}

func generate(f file, packageName, targetDir string, clean bool, opts gen.Options) error {
	data := f.data
	log := opts.Logger

	spec, err := ogen.Parse(data)
	if err != nil {
		// For pretty error message, we need to pass location.File.
		return &location.Error{
			File: f.location(),
			Err:  errors.Wrap(err, "parse spec"),
		}
	}

	start := time.Now()
	g, err := gen.NewGenerator(spec, opts)
	if err != nil {
		return errors.Wrap(err, "build IR")
	}
	log.Debug("Build IR", zap.Duration("took", time.Since(start)))

	// Clean target dir only after flag parsing, spec parsing and IR building.
	switch files, err := os.ReadDir(targetDir); {
	case os.IsNotExist(err):
		if err := os.MkdirAll(targetDir, 0o750); err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		if clean {
			if err := cleanDir(targetDir, files); err != nil {
				return errors.Wrap(err, "clean")
			}
		}
	}

	fs := genfs.FormattedSource{
		// FIXME(tdakkota): write source uses imports.Process which also uses go/format.
		// 	So, there is no reason to format source twice or provide a flag to disable formatting.
		Format: false,
		Root:   targetDir,
	}
	start = time.Now()
	if err := g.WriteSource(fs, packageName); err != nil {
		return errors.Wrap(err, "write")
	}
	log.Debug("Write", zap.Duration("took", time.Since(start)))

	return nil
}

func handleGenerateError(w io.Writer, color bool, err error) (r bool) {
	defer func() {
		// Add trailing newline to the error message if error is handled.
		if r {
			_, _ = fmt.Fprintln(w)
		}
	}()

	if location.PrintPrettyError(w, color, err) {
		return true
	}

	if notImplErr, ok := errors.Into[*gen.ErrNotImplemented](err); ok {
		_, _ = fmt.Fprintf(w, `
Feature %[1]q is not implemented yet.
Try to run ogen with --debug.ignoreNotImplemented %[1]q or with --debug.noerr to skip unsupported operations.
`, notImplErr.Name)
		return true
	}

	if ctErr, ok := errors.Into[*gen.ErrUnsupportedContentTypes](err); ok {
		_, _ = fmt.Fprintf(w, `
Content types [%s] are unsupported.
Try to run ogen with --debug.ignoreNotImplemented %q or with --debug.noerr to skip unsupported operations.
Also, you can use --ct-alias to map content types to supported ones.
`,
			strings.Join(ctErr.ContentTypes, ", "),
			"unsupported content type",
		)
		return true
	}

	return false
}

type file struct {
	data     []byte
	fileName string
	source   string
	rootURL  *url.URL
}

func (f file) location() location.File {
	return location.NewFile(f.fileName, f.source, f.data)
}

func parseSpecPath(
	p string,
	client *http.Client,
	readFile func(string) ([]byte, error),
	logger *zap.Logger,
) (f file, opts gen.RemoteOptions, _ error) {
	// FIXME(tdakkota): pass context.
	containsDrive := runtime.GOOS == "windows" && filepath.VolumeName(p) != ""
	if u, _ := url.Parse(p); u != nil && !containsDrive && u.Scheme != "" {
		switch u.Scheme {
		case "http", "https":
			_, fileName := path.Split(u.Path)

			resp, err := client.Get(p)
			if err != nil {
				return f, opts, err
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return f, opts, err
			}

			f = file{
				data:     data,
				fileName: fileName,
				source:   p,
				rootURL:  u,
			}
			opts = gen.RemoteOptions{
				HTTPClient: client,
				ReadFile: func(p string) ([]byte, error) {
					return nil, errors.New("local files are not supported in remote mode")
				},
				Logger: logger,
			}
			return f, opts, nil
		case "file":
			converted, err := urlpath.URLToFilePath(u)
			if err != nil {
				return f, opts, errors.Wrap(err, "convert url to file path")
			}
			p = converted
		default:
			return f, opts, errors.Errorf("unsupported scheme %q", u.Scheme)
		}
	}
	p = filepath.Clean(p)

	abs, err := filepath.Abs(p)
	if err != nil {
		return f, opts, err
	}
	_, fileName := filepath.Split(p)

	data, err := readFile(p)
	if err != nil {
		return f, opts, err
	}

	u, err := urlpath.URLFromFilePath(abs)
	if err != nil {
		return f, opts, errors.Wrap(err, "convert file path to url")
	}

	f = file{
		data:     data,
		fileName: fileName,
		source:   p,
		rootURL:  u,
	}
	opts = gen.RemoteOptions{
		HTTPClient: client,
		ReadFile:   readFile,
		Logger:     logger,
	}

	return f, opts, nil
}

func run() error {
	set := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	set.Usage = func() {
		_, toolName := filepath.Split(os.Args[0])
		_, _ = fmt.Fprintf(set.Output(), "Usage: %s [options] <spec>\n", toolName)
		set.PrintDefaults()
	}

	var (
		// Generator options.
		targetDir         = set.String("target", "api", "Path to target dir")
		packageName       = set.String("package", "api", "Target package name")
		inferTypes        = set.Bool("infer-types", false, "Infer schema types, if type is not defined explicitly")
		clean             = set.Bool("clean", false, "Clean generated files before generation")
		generateTests     = set.Bool("generate-tests", false, "Generate tests encode-decode/based on schema examples")
		allowRemote       = set.Bool("allow-remote", false, "Enables remote references resolving")
		skipTestsRegex    = set.String("skip-tests", "", "Skip tests matched by regex")
		skipUnimplemented = set.Bool("skip-unimplemented", false, "Disables generation of UnimplementedHandler")
		noClient          = set.Bool("no-client", false, "Disables client generation")
		noServer          = set.Bool("no-server", false, "Disables server generation")
		// Debug options.
		debugIgnoreNotImplemented = set.String("debug.ignoreNotImplemented", "",
			"Ignore methods having functionality which is not implemented")
		debugNoerr = set.Bool("debug.noerr", false, "Ignore errors")

		// Logging options.
		logOptions ogenzap.Options

		// Profile options.
		cpuProfile     = set.String("cpuprofile", "", "Write cpu profile to file")
		memProfile     = set.String("memprofile", "", "Write memory profile to this file")
		memProfileRate = set.Int("memprofilerate", -1, "If > 0, sets runtime.MemProfileRate")

		// Version option.
		version = set.Bool("version", false, "Print version and exit")
	)
	logOptions.RegisterFlags(set)

	var (
		ctAliases gen.ContentTypeAliases

		filterPath    *regexp.Regexp
		filterMethods []string
	)
	set.Var(&ctAliases, "ct-alias", "Content type alias, e.g. text/x-markdown=text/plain")
	set.Func("filter-path", "Filter operations by path regex", func(s string) (err error) {
		filterPath, err = regexp.Compile(s)
		return err
	})
	set.Func("filter-methods", "Filter operations by HTTP methods (comma-separated)", func(s string) error {
		for _, m := range strings.Split(s, ",") {
			m = strings.TrimSpace(m)
			if m == "" {
				continue
			}
			filterMethods = append(filterMethods, m)
		}
		return nil
	})

	if err := set.Parse(os.Args[1:]); err != nil {
		return err
	}

	if *version {
		info, _ := ogenversion.GetInfo()
		fmt.Println(info)
		return nil
	}

	specPath := set.Arg(0)
	if set.NArg() == 0 || specPath == "" {
		set.Usage()
		return errors.New("no spec provided")
	}

	logger, err := ogenzap.Create(logOptions)
	if err != nil {
		return err
	}
	defer func() {
		_ = logger.Sync()
	}()

	if f := *cpuProfile; f != "" {
		f, err := os.Create(f)
		if err != nil {
			return errors.Wrap(err, "create cpu profile")
		}
		defer func() {
			_ = f.Close()
		}()

		if err := pprof.StartCPUProfile(f); err != nil {
			logger.Error("Start CPU profiling", zap.Error(err))
		} else {
			defer pprof.StopCPUProfile()
		}
	}
	if f := *memProfile; f != "" {
		f, err := os.Create(f)
		if err != nil {
			return errors.Wrap(err, "create memory profile")
		}
		defer func() {
			_ = f.Close()
		}()

		if *memProfileRate > 0 {
			runtime.MemProfileRate = *memProfileRate
		}
		defer func() {
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				logger.Error("Write memory profile", zap.Error(err))
			}
		}()
	}

	f, remoteOpts, err := parseSpecPath(
		specPath,
		&http.Client{Timeout: time.Minute},
		os.ReadFile,
		logger.Named("remote"),
	)
	if err != nil {
		return err
	}

	opts := gen.Options{
		NoClient:             *noClient,
		NoServer:             *noServer,
		GenerateExampleTests: *generateTests,
		SkipTestRegex:        nil, // Set below.
		SkipUnimplemented:    *skipUnimplemented,
		InferSchemaType:      *inferTypes,
		AllowRemote:          *allowRemote,
		RootURL:              f.rootURL,
		Remote:               remoteOpts,
		Filters: gen.Filters{
			PathRegex: filterPath,
			Methods:   filterMethods,
		},
		IgnoreNotImplemented: strings.Split(*debugIgnoreNotImplemented, ","),
		ContentTypeAliases:   ctAliases,
		File:                 f.location(),
		Logger:               logger,
	}
	if expr := *skipTestsRegex; expr != "" {
		r, err := regexp.Compile(expr)
		if err != nil {
			return errors.Wrap(err, "skipTestsRegex")
		}
		opts.SkipTestRegex = r
	}
	if *debugNoerr {
		opts.IgnoreNotImplemented = []string{"all"}
	} else {
		// Normalize ignore rules.
		for i := range opts.IgnoreNotImplemented {
			opts.IgnoreNotImplemented[i] = strings.TrimSpace(opts.IgnoreNotImplemented[i])
		}
	}

	if err := generate(f, *packageName, *targetDir, *clean, opts); err != nil {
		if handleGenerateError(os.Stderr, logOptions.Color, err) {
			return errors.New("generation failed")
		}
		return errors.Wrap(err, "generate")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
