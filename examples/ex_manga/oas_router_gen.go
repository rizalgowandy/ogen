// Code generated by ogen, DO NOT EDIT.

package api

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
		case "api": // -> 1
			// Edge: 1, path: "api".
			elem, p = nextElem(p)
			switch string(elem) {
			case "gallery": // -> 2
				// Edge: 2, path: "gallery".
				elem, p = nextElem(p)
				switch string(elem) {
				default:
					if args == nil {
						args = make(map[string]string)
					}
					args["book_id"] = string(elem)
					// GET /api/gallery/{book_id}
					s.handleGetBookRequest(args, w, r)
					return
				}
			case "galleries": // -> 14
				// Edge: 14, path: "galleries".
				elem, p = nextElem(p)
				switch string(elem) {
				case "search": // -> 15
					// GET /api/galleries/search
					s.handleSearchRequest(args, w, r)
					return
				case "tagged": // -> 16
					// GET /api/galleries/tagged
					s.handleSearchByTagIDRequest(args, w, r)
					return
				default:
					s.notFound(w, r)
					return
				}
			default:
				s.notFound(w, r)
				return
			}
		case "galleries": // -> 4
			// Edge: 4, path: "galleries".
			elem, p = nextElem(p)
			switch string(elem) {
			default:
				if args == nil {
					args = make(map[string]string)
				}
				args["media_id"] = string(elem)
				// Edge: 5, path: "".
				elem, p = nextElem(p)
				switch string(elem) {
				case "cover.": // -> 6
					// Edge: 6, path: "cover.".
					elem, p = nextElem(p)
					switch string(elem) {
					default:
						if args == nil {
							args = make(map[string]string)
						}
						args["format"] = string(elem)
						// GET /galleries/{media_id}/cover.{format}
						s.handleGetPageCoverImageRequest(args, w, r)
						return
					}
				case "": // -> 8
					// Edge: 8, path: "".
					elem, p = nextElem(p)
					switch string(elem) {
					default:
						if args == nil {
							args = make(map[string]string)
						}
						args["page"] = string(elem)
						// Edge: 9, path: "".
						elem, p = nextElem(p)
						switch string(elem) {
						case ".": // -> 10
							// Edge: 10, path: ".".
							elem, p = nextElem(p)
							switch string(elem) {
							default:
								if args == nil {
									args = make(map[string]string)
								}
								args["format"] = string(elem)
								// GET /galleries/{media_id}/{page}.{format}
								s.handleGetPageImageRequest(args, w, r)
								return
							}
						case "t.": // -> 12
							// Edge: 12, path: "t.".
							elem, p = nextElem(p)
							switch string(elem) {
							default:
								if args == nil {
									args = make(map[string]string)
								}
								args["format"] = string(elem)
								// GET /galleries/{media_id}/{page}t.{format}
								s.handleGetPageThumbnailImageRequest(args, w, r)
								return
							}
						default:
							s.notFound(w, r)
							return
						}
					}
				default:
					s.notFound(w, r)
					return
				}
			}
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}
