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
			case 'm': // Prefix: "market/"
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketBondsGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
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
					case 'a': // Prefix: "andles"
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketCandlesGetRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'u': // Prefix: "urrencies"
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketCurrenciesGetRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 'e': // Prefix: "etfs"
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketEtfsGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'o': // Prefix: "orderbook"
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleMarketOrderbookGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleMarketSearchByFigiGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 't': // Prefix: "ticker"
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleMarketSearchByTickerGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					case 't': // Prefix: "tocks"
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleMarketStocksGetRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'o': // Prefix: "o"
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOperationsGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'r': // Prefix: "rders"
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleOrdersGetRequest([0]string{}, w, r)
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
									s.handleOrdersCancelPostRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						case 'l': // Prefix: "limit-order"
							if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleOrdersLimitOrderPostRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						case 'm': // Prefix: "market-order"
							if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleOrdersMarketOrderPostRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					}
				}
			case 'p': // Prefix: "portfolio"
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handlePortfolioGetRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handlePortfolioCurrenciesGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 's': // Prefix: "sandbox/"
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
					case 'l': // Prefix: "lear"
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxClearPostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'u': // Prefix: "urrencies/balance"
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxCurrenciesBalancePostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				case 'p': // Prefix: "positions/balance"
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleSandboxPositionsBalancePostRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'r': // Prefix: "re"
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxRegisterPostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'm': // Prefix: "move"
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSandboxRemovePostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 'u': // Prefix: "user/accounts"
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleUserAccountsGetRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
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
			case 'm': // Prefix: "market/"
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: MarketBondsGet
							r.name = "MarketBondsGet"
							r.operationID = ""
							r.pathPattern = "/market/bonds"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
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
					case 'a': // Prefix: "andles"
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: MarketCandlesGet
								r.name = "MarketCandlesGet"
								r.operationID = ""
								r.pathPattern = "/market/candles"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'u': // Prefix: "urrencies"
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: MarketCurrenciesGet
								r.name = "MarketCurrenciesGet"
								r.operationID = ""
								r.pathPattern = "/market/currencies"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'e': // Prefix: "etfs"
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: MarketEtfsGet
							r.name = "MarketEtfsGet"
							r.operationID = ""
							r.pathPattern = "/market/etfs"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'o': // Prefix: "orderbook"
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: MarketOrderbookGet
							r.name = "MarketOrderbookGet"
							r.operationID = ""
							r.pathPattern = "/market/orderbook"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: MarketSearchByFigiGet
									r.name = "MarketSearchByFigiGet"
									r.operationID = ""
									r.pathPattern = "/market/search/by-figi"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 't': // Prefix: "ticker"
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: MarketSearchByTickerGet
									r.name = "MarketSearchByTickerGet"
									r.operationID = ""
									r.pathPattern = "/market/search/by-ticker"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					case 't': // Prefix: "tocks"
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: MarketStocksGet
								r.name = "MarketStocksGet"
								r.operationID = ""
								r.pathPattern = "/market/stocks"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'o': // Prefix: "o"
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OperationsGet
							r.name = "OperationsGet"
							r.operationID = ""
							r.pathPattern = "/operations"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'r': // Prefix: "rders"
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "OrdersGet"
							r.operationID = ""
							r.pathPattern = "/orders"
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
									// Leaf: OrdersCancelPost
									r.name = "OrdersCancelPost"
									r.operationID = ""
									r.pathPattern = "/orders/cancel"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 'l': // Prefix: "limit-order"
							if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: OrdersLimitOrderPost
									r.name = "OrdersLimitOrderPost"
									r.operationID = ""
									r.pathPattern = "/orders/limit-order"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 'm': // Prefix: "market-order"
							if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: OrdersMarketOrderPost
									r.name = "OrdersMarketOrderPost"
									r.operationID = ""
									r.pathPattern = "/orders/market-order"
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
			case 'p': // Prefix: "portfolio"
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "PortfolioGet"
						r.operationID = ""
						r.pathPattern = "/portfolio"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: PortfolioCurrenciesGet
							r.name = "PortfolioCurrenciesGet"
							r.operationID = ""
							r.pathPattern = "/portfolio/currencies"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 's': // Prefix: "sandbox/"
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
					case 'l': // Prefix: "lear"
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SandboxClearPost
								r.name = "SandboxClearPost"
								r.operationID = ""
								r.pathPattern = "/sandbox/clear"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'u': // Prefix: "urrencies/balance"
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SandboxCurrenciesBalancePost
								r.name = "SandboxCurrenciesBalancePost"
								r.operationID = ""
								r.pathPattern = "/sandbox/currencies/balance"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'p': // Prefix: "positions/balance"
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: SandboxPositionsBalancePost
							r.name = "SandboxPositionsBalancePost"
							r.operationID = ""
							r.pathPattern = "/sandbox/positions/balance"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'r': // Prefix: "re"
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SandboxRegisterPost
								r.name = "SandboxRegisterPost"
								r.operationID = ""
								r.pathPattern = "/sandbox/register"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'm': // Prefix: "move"
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SandboxRemovePost
								r.name = "SandboxRemovePost"
								r.operationID = ""
								r.pathPattern = "/sandbox/remove"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'u': // Prefix: "user/accounts"
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: UserAccountsGet
						r.name = "UserAccountsGet"
						r.operationID = ""
						r.pathPattern = "/user/accounts"
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
	return r, false
}
