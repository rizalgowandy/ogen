// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "allRequestBodies"
				origElem := elem
				if l := len("allRequestBodies"); len(elem) >= l && elem[0:l] == "allRequestBodies" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "POST":
						s.handleAllRequestBodiesRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
				switch elem[0] {
				case 'O': // Prefix: "Optional"
					origElem := elem
					if l := len("Optional"); len(elem) >= l && elem[0:l] == "Optional" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAllRequestBodiesOptionalRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'b': // Prefix: "base64Request"
				origElem := elem
				if l := len("base64Request"); len(elem) >= l && elem[0:l] == "base64Request" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleBase64RequestRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			case 'm': // Prefix: "maskContentType"
				origElem := elem
				if l := len("maskContentType"); len(elem) >= l && elem[0:l] == "maskContentType" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "POST":
						s.handleMaskContentTypeRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
				switch elem[0] {
				case 'O': // Prefix: "Optional"
					origElem := elem
					if l := len("Optional"); len(elem) >= l && elem[0:l] == "Optional" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleMaskContentTypeOptionalRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "streamJSON"
				origElem := elem
				if l := len("streamJSON"); len(elem) >= l && elem[0:l] == "streamJSON" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleStreamJSONRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "allRequestBodies"
				origElem := elem
				if l := len("allRequestBodies"); len(elem) >= l && elem[0:l] == "allRequestBodies" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						r.name = "AllRequestBodies"
						r.summary = ""
						r.operationID = "allRequestBodies"
						r.pathPattern = "/allRequestBodies"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case 'O': // Prefix: "Optional"
					origElem := elem
					if l := len("Optional"); len(elem) >= l && elem[0:l] == "Optional" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "AllRequestBodiesOptional"
							r.summary = ""
							r.operationID = "allRequestBodiesOptional"
							r.pathPattern = "/allRequestBodiesOptional"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'b': // Prefix: "base64Request"
				origElem := elem
				if l := len("base64Request"); len(elem) >= l && elem[0:l] == "base64Request" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "Base64Request"
						r.summary = ""
						r.operationID = "base64Request"
						r.pathPattern = "/base64Request"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'm': // Prefix: "maskContentType"
				origElem := elem
				if l := len("maskContentType"); len(elem) >= l && elem[0:l] == "maskContentType" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						r.name = "MaskContentType"
						r.summary = ""
						r.operationID = "maskContentType"
						r.pathPattern = "/maskContentType"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case 'O': // Prefix: "Optional"
					origElem := elem
					if l := len("Optional"); len(elem) >= l && elem[0:l] == "Optional" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "MaskContentTypeOptional"
							r.summary = ""
							r.operationID = "maskContentTypeOptional"
							r.pathPattern = "/maskContentTypeOptional"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "streamJSON"
				origElem := elem
				if l := len("streamJSON"); len(elem) >= l && elem[0:l] == "streamJSON" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "StreamJSON"
						r.summary = ""
						r.operationID = "streamJSON"
						r.pathPattern = "/streamJSON"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
