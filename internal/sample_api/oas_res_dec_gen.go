// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
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
	"time"

	"github.com/go-chi/chi/v5"
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
	_ = chi.Context{}
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
)

func decodeFoobarGetResponse(resp *http.Response, span trace.Span) (res FoobarGetRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeFoobarPostResponse(resp *http.Response, span trace.Span) (res FoobarPostRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeFoobarPutResponse(resp *http.Response, span trace.Span) (res FoobarPutDefStatusCode, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	default:
		return FoobarPutDefStatusCode{StatusCode: resp.StatusCode}, nil
	}
}

func decodePetCreateResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetFriendsNamesByIDResponse(resp *http.Response, span trace.Span) (res []string, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response []string
			if err := func() error {
				response = nil
				if err := i.Array(func(i *json.Iter) error {
					var elem string
					v, err := i.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetGetResponse(resp *http.Response, span trace.Span) (res PetGetRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetDefStatusCode
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePetGetByNameResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetNameByIDResponse(resp *http.Response, span trace.Span) (res string, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIter()
			defer json.PutIter(i)
			i.ResetBytes(buf.Bytes())

			var response string
			if err := func() error {
				v, err := i.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetUpdateNameAliasPostResponse(resp *http.Response, span trace.Span) (res PetUpdateNameAliasPostDefStatusCode, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	default:
		return PetUpdateNameAliasPostDefStatusCode{StatusCode: resp.StatusCode}, nil
	}
}

func decodePetUpdateNamePostResponse(resp *http.Response, span trace.Span) (res PetUpdateNamePostDefStatusCode, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	default:
		return PetUpdateNamePostDefStatusCode{StatusCode: resp.StatusCode}, nil
	}
}
