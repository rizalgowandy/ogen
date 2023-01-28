// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
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
		case '/': // Prefix: "/event"
			if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf node.
				switch r.Method {
				case "POST":
					s.handlePublishEventRequest([0]string{}, w, r)
				default:
					s.notAllowed(w, r, "POST")
				}

				return
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
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

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/event"
			if l := len("/event"); len(elem) >= l && elem[0:l] == "/event" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "POST":
					// Leaf: PublishEvent
					r.name = "PublishEvent"
					r.operationID = "publishEvent"
					r.pathPattern = "/event"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
		}
	}
	return r, false
}

// Handle handles webhook request.
//
// Returns true if there is a webhook handler for given name and requested method.
func (s *WebhookServer) Handle(webhookName string, w http.ResponseWriter, r *http.Request) bool {
	switch webhookName {
	case "status":
		switch r.Method {
		case "GET":
			s.handleStatusWebhookRequest([0]string{}, w, r)
		default:
			return false
		}
		return true
	case "update":
		switch r.Method {
		case "DELETE":
			s.handleUpdateDeleteRequest([0]string{}, w, r)
		case "POST":
			s.handleUpdateWebhookRequest([0]string{}, w, r)
		default:
			return false
		}
		return true
	default:
		return false
	}
}

// Handler returns http.Handler for webhook.
//
// Returns NotFound handler if spec doesn't contain webhook with given name.
//
// Returned handler calls MethodNotAllowed handler if webhook doesn't define requested method.
func (s *WebhookServer) Handler(webhookName string) http.Handler {
	switch webhookName {
	case "status":
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// We know that webhook exists, so false means wrong method.
			if !s.Handle(webhookName, w, r) {
				s.notAllowed(w, r, "GET")
			}
		})
	case "update":
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// We know that webhook exists, so false means wrong method.
			if !s.Handle(webhookName, w, r) {
				s.notAllowed(w, r, "DELETE,POST")
			}
		})
	default:
		return http.HandlerFunc(s.notFound)
	}
}
