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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/errors"
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

func encodeGetBookResponse(response GetBookRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Book:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := json.GetEncoder()
		defer json.PutEncoder(e)
		more := json.NewMore(e)
		defer more.Reset()
		more.More()
		response.WriteJSON(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *GetBookForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/gallery/{book_id}: unexpected response type: %T`, response)
	}
}

func encodeGetPageCoverImageResponse(response GetPageCoverImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageCoverImageOKImage:
		w.Header().Set("Content-Type", "image/*")
		w.WriteHeader(200)
		return errors.New(`image/* encoder not implemented`)
	case *GetPageCoverImageForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/galleries/{media_id}/cover.{format}: unexpected response type: %T`, response)
	}
}

func encodeGetPageImageResponse(response GetPageImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageImageOKImage:
		w.Header().Set("Content-Type", "image/*")
		w.WriteHeader(200)
		return errors.New(`image/* encoder not implemented`)
	case *GetPageImageForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/galleries/{media_id}/{page}.{format}: unexpected response type: %T`, response)
	}
}

func encodeGetPageThumbnailImageResponse(response GetPageThumbnailImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageThumbnailImageOKImage:
		w.Header().Set("Content-Type", "image/*")
		w.WriteHeader(200)
		return errors.New(`image/* encoder not implemented`)
	case *GetPageThumbnailImageForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/galleries/{media_id}/{page}t.{format}: unexpected response type: %T`, response)
	}
}

func encodeSearchResponse(response SearchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := json.GetEncoder()
		defer json.PutEncoder(e)
		more := json.NewMore(e)
		defer more.Reset()
		// Unsupported kind "alias".
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/galleries/search: unexpected response type: %T`, response)
	}
}

func encodeSearchByTagIDResponse(response SearchByTagIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchByTagIDOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := json.GetEncoder()
		defer json.PutEncoder(e)
		more := json.NewMore(e)
		defer more.Reset()
		// Unsupported kind "alias".
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchByTagIDForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/galleries/tagged: unexpected response type: %T`, response)
	}
}
