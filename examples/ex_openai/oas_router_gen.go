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
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'n': // Prefix: "nswers"
					if l := len("nswers"); len(elem) >= l && elem[0:l] == "nswers" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateAnswerRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'u': // Prefix: "udio/trans"
					if l := len("udio/trans"); len(elem) >= l && elem[0:l] == "udio/trans" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "criptions"
						if l := len("criptions"); len(elem) >= l && elem[0:l] == "criptions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleCreateTranscriptionRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'l': // Prefix: "lations"
						if l := len("lations"); len(elem) >= l && elem[0:l] == "lations" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleCreateTranslationRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 'c': // Prefix: "c"
				if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'h': // Prefix: "hat/completions"
					if l := len("hat/completions"); len(elem) >= l && elem[0:l] == "hat/completions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateChatCompletionRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'l': // Prefix: "lassifications"
					if l := len("lassifications"); len(elem) >= l && elem[0:l] == "lassifications" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateClassificationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'o': // Prefix: "ompletions"
					if l := len("ompletions"); len(elem) >= l && elem[0:l] == "ompletions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateCompletionRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				}
			case 'e': // Prefix: "e"
				if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "dits"
					if l := len("dits"); len(elem) >= l && elem[0:l] == "dits" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateEditRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'm': // Prefix: "mbeddings"
					if l := len("mbeddings"); len(elem) >= l && elem[0:l] == "mbeddings" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateEmbeddingRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'n': // Prefix: "ngines"
					if l := len("ngines"); len(elem) >= l && elem[0:l] == "ngines" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListEnginesRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "engine_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleRetrieveEngineRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/search"
							if l := len("/search"); len(elem) >= l && elem[0:l] == "/search" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleCreateSearchRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					}
				}
			case 'f': // Prefix: "fi"
				if l := len("fi"); len(elem) >= l && elem[0:l] == "fi" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "les"
					if l := len("les"); len(elem) >= l && elem[0:l] == "les" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListFilesRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handleCreateFileRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "file_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "DELETE":
								s.handleDeleteFileRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							case "GET":
								s.handleRetrieveFileRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/content"
							if l := len("/content"); len(elem) >= l && elem[0:l] == "/content" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleDownloadFileRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				case 'n': // Prefix: "ne-tunes"
					if l := len("ne-tunes"); len(elem) >= l && elem[0:l] == "ne-tunes" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListFineTunesRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handleCreateFineTuneRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "fine_tune_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleRetrieveFineTuneRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "cancel"
								if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleCancelFineTuneRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "POST")
									}

									return
								}
							case 'e': // Prefix: "events"
								if l := len("events"); len(elem) >= l && elem[0:l] == "events" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleListFineTuneEventsRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					}
				}
			case 'i': // Prefix: "images/"
				if l := len("images/"); len(elem) >= l && elem[0:l] == "images/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "edits"
					if l := len("edits"); len(elem) >= l && elem[0:l] == "edits" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateImageEditRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'g': // Prefix: "generations"
					if l := len("generations"); len(elem) >= l && elem[0:l] == "generations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateImageRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'v': // Prefix: "variations"
					if l := len("variations"); len(elem) >= l && elem[0:l] == "variations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateImageVariationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				}
			case 'm': // Prefix: "mode"
				if l := len("mode"); len(elem) >= l && elem[0:l] == "mode" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "ls"
					if l := len("ls"); len(elem) >= l && elem[0:l] == "ls" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleListModelsRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "model"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "DELETE":
								s.handleDeleteModelRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							case "GET":
								s.handleRetrieveModelRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET")
							}

							return
						}
					}
				case 'r': // Prefix: "rations"
					if l := len("rations"); len(elem) >= l && elem[0:l] == "rations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleCreateModerationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

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
	name        string
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
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'n': // Prefix: "nswers"
					if l := len("nswers"); len(elem) >= l && elem[0:l] == "nswers" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateAnswer
							r.name = "CreateAnswer"
							r.operationID = "createAnswer"
							r.pathPattern = "/answers"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'u': // Prefix: "udio/trans"
					if l := len("udio/trans"); len(elem) >= l && elem[0:l] == "udio/trans" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "criptions"
						if l := len("criptions"); len(elem) >= l && elem[0:l] == "criptions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: CreateTranscription
								r.name = "CreateTranscription"
								r.operationID = "createTranscription"
								r.pathPattern = "/audio/transcriptions"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'l': // Prefix: "lations"
						if l := len("lations"); len(elem) >= l && elem[0:l] == "lations" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: CreateTranslation
								r.name = "CreateTranslation"
								r.operationID = "createTranslation"
								r.pathPattern = "/audio/translations"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'c': // Prefix: "c"
				if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'h': // Prefix: "hat/completions"
					if l := len("hat/completions"); len(elem) >= l && elem[0:l] == "hat/completions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateChatCompletion
							r.name = "CreateChatCompletion"
							r.operationID = "createChatCompletion"
							r.pathPattern = "/chat/completions"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'l': // Prefix: "lassifications"
					if l := len("lassifications"); len(elem) >= l && elem[0:l] == "lassifications" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateClassification
							r.name = "CreateClassification"
							r.operationID = "createClassification"
							r.pathPattern = "/classifications"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'o': // Prefix: "ompletions"
					if l := len("ompletions"); len(elem) >= l && elem[0:l] == "ompletions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateCompletion
							r.name = "CreateCompletion"
							r.operationID = "createCompletion"
							r.pathPattern = "/completions"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'e': // Prefix: "e"
				if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "dits"
					if l := len("dits"); len(elem) >= l && elem[0:l] == "dits" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateEdit
							r.name = "CreateEdit"
							r.operationID = "createEdit"
							r.pathPattern = "/edits"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'm': // Prefix: "mbeddings"
					if l := len("mbeddings"); len(elem) >= l && elem[0:l] == "mbeddings" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateEmbedding
							r.name = "CreateEmbedding"
							r.operationID = "createEmbedding"
							r.pathPattern = "/embeddings"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'n': // Prefix: "ngines"
					if l := len("ngines"); len(elem) >= l && elem[0:l] == "ngines" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListEngines"
							r.operationID = "listEngines"
							r.pathPattern = "/engines"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "engine_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "RetrieveEngine"
								r.operationID = "retrieveEngine"
								r.pathPattern = "/engines/{engine_id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/search"
							if l := len("/search"); len(elem) >= l && elem[0:l] == "/search" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: CreateSearch
									r.name = "CreateSearch"
									r.operationID = "createSearch"
									r.pathPattern = "/engines/{engine_id}/search"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'f': // Prefix: "fi"
				if l := len("fi"); len(elem) >= l && elem[0:l] == "fi" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "les"
					if l := len("les"); len(elem) >= l && elem[0:l] == "les" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListFiles"
							r.operationID = "listFiles"
							r.pathPattern = "/files"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "CreateFile"
							r.operationID = "createFile"
							r.pathPattern = "/files"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "file_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								r.name = "DeleteFile"
								r.operationID = "deleteFile"
								r.pathPattern = "/files/{file_id}"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								r.name = "RetrieveFile"
								r.operationID = "retrieveFile"
								r.pathPattern = "/files/{file_id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/content"
							if l := len("/content"); len(elem) >= l && elem[0:l] == "/content" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: DownloadFile
									r.name = "DownloadFile"
									r.operationID = "downloadFile"
									r.pathPattern = "/files/{file_id}/content"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				case 'n': // Prefix: "ne-tunes"
					if l := len("ne-tunes"); len(elem) >= l && elem[0:l] == "ne-tunes" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListFineTunes"
							r.operationID = "listFineTunes"
							r.pathPattern = "/fine-tunes"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "CreateFineTune"
							r.operationID = "createFineTune"
							r.pathPattern = "/fine-tunes"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "fine_tune_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "RetrieveFineTune"
								r.operationID = "retrieveFineTune"
								r.pathPattern = "/fine-tunes/{fine_tune_id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "cancel"
								if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										// Leaf: CancelFineTune
										r.name = "CancelFineTune"
										r.operationID = "cancelFineTune"
										r.pathPattern = "/fine-tunes/{fine_tune_id}/cancel"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							case 'e': // Prefix: "events"
								if l := len("events"); len(elem) >= l && elem[0:l] == "events" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: ListFineTuneEvents
										r.name = "ListFineTuneEvents"
										r.operationID = "listFineTuneEvents"
										r.pathPattern = "/fine-tunes/{fine_tune_id}/events"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
							}
						}
					}
				}
			case 'i': // Prefix: "images/"
				if l := len("images/"); len(elem) >= l && elem[0:l] == "images/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "edits"
					if l := len("edits"); len(elem) >= l && elem[0:l] == "edits" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateImageEdit
							r.name = "CreateImageEdit"
							r.operationID = "createImageEdit"
							r.pathPattern = "/images/edits"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'g': // Prefix: "generations"
					if l := len("generations"); len(elem) >= l && elem[0:l] == "generations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateImage
							r.name = "CreateImage"
							r.operationID = "createImage"
							r.pathPattern = "/images/generations"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'v': // Prefix: "variations"
					if l := len("variations"); len(elem) >= l && elem[0:l] == "variations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateImageVariation
							r.name = "CreateImageVariation"
							r.operationID = "createImageVariation"
							r.pathPattern = "/images/variations"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'm': // Prefix: "mode"
				if l := len("mode"); len(elem) >= l && elem[0:l] == "mode" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "ls"
					if l := len("ls"); len(elem) >= l && elem[0:l] == "ls" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "ListModels"
							r.operationID = "listModels"
							r.pathPattern = "/models"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "model"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								// Leaf: DeleteModel
								r.name = "DeleteModel"
								r.operationID = "deleteModel"
								r.pathPattern = "/models/{model}"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								// Leaf: RetrieveModel
								r.name = "RetrieveModel"
								r.operationID = "retrieveModel"
								r.pathPattern = "/models/{model}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				case 'r': // Prefix: "rations"
					if l := len("rations"); len(elem) >= l && elem[0:l] == "rations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: CreateModeration
							r.name = "CreateModeration"
							r.operationID = "createModeration"
							r.pathPattern = "/moderations"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
