// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func skipSlash(p []byte) []byte {
	if len(p) > 0 && p[0] == '/' {
		return p[1:]
	}
	return p
}

// nextElem return next path element from p and forwarded p.
func nextElem(p []byte) (elem, next []byte) {
	p = skipSlash(p)
	idx := bytes.IndexByte(p, '/')
	if idx < 0 {
		idx = len(p)
	}
	return p[:idx], p[idx:]
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := []byte(r.URL.Path)
	if len(p) == 0 {
		s.notFound(w, r)
		return
	}

	var (
		elem []byte            // current element, without slashes
		args map[string]string // lazily initialized
	)

	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		// Root edge.
		elem, p = nextElem(p)
		switch string(elem) {
		case "cached-worlds": // -> 1
			// GET /cached-worlds
			s.handleCachingRequest(args, w, r)
			return
		case "db": // -> 2
			// GET /db
			s.handleDBRequest(args, w, r)
			return
		case "json": // -> 3
			// GET /json
			s.handleJSONRequest(args, w, r)
			return
		case "queries": // -> 4
			// GET /queries
			s.handleQueriesRequest(args, w, r)
			return
		case "updates": // -> 5
			// GET /updates
			s.handleUpdatesRequest(args, w, r)
			return
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}
