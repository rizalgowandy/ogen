// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeGetBookResponse(response GetBookRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Book:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetBookForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
func encodeGetPageCoverImageResponse(response GetPageCoverImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageCoverImageOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetPageCoverImageForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
func encodeGetPageImageResponse(response GetPageImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageImageOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetPageImageForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
func encodeGetPageThumbnailImageResponse(response GetPageThumbnailImageRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPageThumbnailImageOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetPageThumbnailImageForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
func encodeSearchResponse(response SearchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SearchForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
func encodeSearchByTagIDResponse(response SearchByTagIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchByTagIDOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *SearchByTagIDForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
