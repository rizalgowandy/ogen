// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleMultipleGenericResponsesRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'a': // Prefix: "anyContentTypeBinaryStringSchema"
				if l := len("anyContentTypeBinaryStringSchema"); len(elem) >= l && elem[0:l] == "anyContentTypeBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleAnyContentTypeBinaryStringSchemaRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: AnyContentTypeBinaryStringSchemaDefault
						s.handleAnyContentTypeBinaryStringSchemaDefaultRequest([0]string{}, w, r)

						return
					}
				}
			case 'm': // Prefix: "multipleGenericResponses"
				if l := len("multipleGenericResponses"); len(elem) >= l && elem[0:l] == "multipleGenericResponses" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: MultipleGenericResponses
					s.handleMultipleGenericResponsesRequest([0]string{}, w, r)

					return
				}
			case 'o': // Prefix: "octetStream"
				if l := len("octetStream"); len(elem) >= l && elem[0:l] == "octetStream" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleOctetStreamEmptySchemaRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'B': // Prefix: "BinaryStringSchema"
					if l := len("BinaryStringSchema"); len(elem) >= l && elem[0:l] == "BinaryStringSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OctetStreamBinaryStringSchema
						s.handleOctetStreamBinaryStringSchemaRequest([0]string{}, w, r)

						return
					}
				case 'E': // Prefix: "EmptySchema"
					if l := len("EmptySchema"); len(elem) >= l && elem[0:l] == "EmptySchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OctetStreamEmptySchema
						s.handleOctetStreamEmptySchemaRequest([0]string{}, w, r)

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [0]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [0]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "MultipleGenericResponses"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'a': // Prefix: "anyContentTypeBinaryStringSchema"
				if l := len("anyContentTypeBinaryStringSchema"); len(elem) >= l && elem[0:l] == "anyContentTypeBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "AnyContentTypeBinaryStringSchema"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: AnyContentTypeBinaryStringSchemaDefault
						r.name = "AnyContentTypeBinaryStringSchemaDefault"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			case 'm': // Prefix: "multipleGenericResponses"
				if l := len("multipleGenericResponses"); len(elem) >= l && elem[0:l] == "multipleGenericResponses" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: MultipleGenericResponses
					r.name = "MultipleGenericResponses"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'o': // Prefix: "octetStream"
				if l := len("octetStream"); len(elem) >= l && elem[0:l] == "octetStream" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "OctetStreamEmptySchema"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'B': // Prefix: "BinaryStringSchema"
					if l := len("BinaryStringSchema"); len(elem) >= l && elem[0:l] == "BinaryStringSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OctetStreamBinaryStringSchema
						r.name = "OctetStreamBinaryStringSchema"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'E': // Prefix: "EmptySchema"
					if l := len("EmptySchema"); len(elem) >= l && elem[0:l] == "EmptySchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OctetStreamEmptySchema
						r.name = "OctetStreamEmptySchema"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			}
		}
	}
	return r, false
}
