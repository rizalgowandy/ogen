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

func decodeCreatePetResponse(resp *http.Response, span trace.Span) (res CreatePetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetCreate
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 409:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R409
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeCreatePetCategoriesResponse(resp *http.Response, span trace.Span) (res CreatePetCategoriesRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetCategoriesCreate
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 409:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R409
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeCreatePetFriendsResponse(resp *http.Response, span trace.Span) (res CreatePetFriendsRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetFriendsCreate
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 409:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R409
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeCreatePetOwnerResponse(resp *http.Response, span trace.Span) (res CreatePetOwnerRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetOwnerCreate
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 409:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R409
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeDeletePetResponse(resp *http.Response, span trace.Span) (res DeletePetRes, err error) {
	switch resp.StatusCode {
	case 204:
		return &DeletePetNoContent{}, nil
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeDeletePetOwnerResponse(resp *http.Response, span trace.Span) (res DeletePetOwnerRes, err error) {
	switch resp.StatusCode {
	case 204:
		return &DeletePetOwnerNoContent{}, nil
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeListPetResponse(resp *http.Response, span trace.Span) (res ListPetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ListPetOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeListPetCategoriesResponse(resp *http.Response, span trace.Span) (res ListPetCategoriesRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ListPetCategoriesOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeListPetFriendsResponse(resp *http.Response, span trace.Span) (res ListPetFriendsRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ListPetFriendsOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeReadPetResponse(resp *http.Response, span trace.Span) (res ReadPetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetRead
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeReadPetOwnerResponse(resp *http.Response, span trace.Span) (res ReadPetOwnerRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetOwnerRead
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeUpdatePetResponse(resp *http.Response, span trace.Span) (res UpdatePetRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response PetUpdate
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R400
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R404
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response R500
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
