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
	args := [1]string{}

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
			case 'c': // Prefix: "co"
				origElem := elem
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					origElem := elem
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleComplicatedParameterNameGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'n': // Prefix: "ntentParameters/"
					origElem := elem
					if l := len("ntentParameters/"); len(elem) >= l && elem[0:l] == "ntentParameters/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "path"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleContentParametersRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'o': // Prefix: "okieParameter"
					origElem := elem
					if l := len("okieParameter"); len(elem) >= l && elem[0:l] == "okieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleCookieParameterRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "headerParameter"
				origElem := elem
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHeaderParameterRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'o': // Prefix: "object"
				origElem := elem
				if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'C': // Prefix: "CookieParameter"
					origElem := elem
					if l := len("CookieParameter"); len(elem) >= l && elem[0:l] == "CookieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleObjectCookieParameterRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'Q': // Prefix: "QueryParameter"
					origElem := elem
					if l := len("QueryParameter"); len(elem) >= l && elem[0:l] == "QueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleObjectQueryParameterRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "pathParameter/"
				origElem := elem
				if l := len("pathParameter/"); len(elem) >= l && elem[0:l] == "pathParameter/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "value"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handlePathParameterRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 's': // Prefix: "s"
				origElem := elem
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "ame_name/"
					origElem := elem
					if l := len("ame_name/"); len(elem) >= l && elem[0:l] == "ame_name/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "param"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleSameNameRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'i': // Prefix: "imilarNames"
					origElem := elem
					if l := len("imilarNames"); len(elem) >= l && elem[0:l] == "imilarNames" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleSimilarNamesRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
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
	args        [1]string
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
			case 'c': // Prefix: "co"
				origElem := elem
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					origElem := elem
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ComplicatedParameterNameGet
							r.name = "ComplicatedParameterNameGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/complicatedParameterName"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'n': // Prefix: "ntentParameters/"
					origElem := elem
					if l := len("ntentParameters/"); len(elem) >= l && elem[0:l] == "ntentParameters/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "path"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ContentParameters
							r.name = "ContentParameters"
							r.summary = ""
							r.operationID = "contentParameters"
							r.pathPattern = "/contentParameters/{path}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'o': // Prefix: "okieParameter"
					origElem := elem
					if l := len("okieParameter"); len(elem) >= l && elem[0:l] == "okieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: CookieParameter
							r.name = "CookieParameter"
							r.summary = ""
							r.operationID = "cookieParameter"
							r.pathPattern = "/cookieParameter"
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
			case 'h': // Prefix: "headerParameter"
				origElem := elem
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: HeaderParameter
						r.name = "HeaderParameter"
						r.summary = ""
						r.operationID = "headerParameter"
						r.pathPattern = "/headerParameter"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'o': // Prefix: "object"
				origElem := elem
				if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'C': // Prefix: "CookieParameter"
					origElem := elem
					if l := len("CookieParameter"); len(elem) >= l && elem[0:l] == "CookieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ObjectCookieParameter
							r.name = "ObjectCookieParameter"
							r.summary = ""
							r.operationID = "objectCookieParameter"
							r.pathPattern = "/objectCookieParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'Q': // Prefix: "QueryParameter"
					origElem := elem
					if l := len("QueryParameter"); len(elem) >= l && elem[0:l] == "QueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ObjectQueryParameter
							r.name = "ObjectQueryParameter"
							r.summary = ""
							r.operationID = "objectQueryParameter"
							r.pathPattern = "/objectQueryParameter"
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
			case 'p': // Prefix: "pathParameter/"
				origElem := elem
				if l := len("pathParameter/"); len(elem) >= l && elem[0:l] == "pathParameter/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "value"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: PathParameter
						r.name = "PathParameter"
						r.summary = ""
						r.operationID = "pathParameter"
						r.pathPattern = "/pathParameter/{value}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 's': // Prefix: "s"
				origElem := elem
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "ame_name/"
					origElem := elem
					if l := len("ame_name/"); len(elem) >= l && elem[0:l] == "ame_name/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "param"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: SameName
							r.name = "SameName"
							r.summary = "parameters with different location, but with the same name"
							r.operationID = "sameName"
							r.pathPattern = "/same_name/{param}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'i': // Prefix: "imilarNames"
					origElem := elem
					if l := len("imilarNames"); len(elem) >= l && elem[0:l] == "imilarNames" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: SimilarNames
							r.name = "SimilarNames"
							r.summary = "parameters with different location, but with similar names"
							r.operationID = "similarNames"
							r.pathPattern = "/similarNames"
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

			elem = origElem
		}
	}
	return r, false
}
