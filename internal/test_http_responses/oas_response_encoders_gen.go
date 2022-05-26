// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeAnyContentTypeBinaryStringSchemaResponse(response AnyContentTypeBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "*/*")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	if _, err := io.Copy(w, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeAnyContentTypeBinaryStringSchemaDefaultResponse(response AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "*/*")
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	if _, err := io.Copy(w, response.Response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeHeaders200Response(response Headers200OKHeaders, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode 'TestHeader' header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.TestHeader))
			}); err != nil {
				return errors.Wrap(err, "encode TestHeader header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}
func encodeHeadersCombinedResponse(response HeadersCombinedRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *HeadersCombinedOKHeaders:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode 'TestHeader' header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "TestHeader",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.TestHeader))
				}); err != nil {
					return errors.Wrap(err, "encode TestHeader header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		return nil

	case *HeadersCombinedDefStatusCodeWithHeaders:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode 'TestHeader' header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "TestHeader",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.TestHeader))
				}); err != nil {
					return errors.Wrap(err, "encode TestHeader header")
				}
			}
		}
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		return nil

	default:
		return errors.Errorf("/headersCombined"+`: unexpected response type: %T`, response)
	}
}
func encodeHeadersDefaultResponse(response HeadersDefaultDefStatusCodeWithHeaders, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode 'TestHeader' header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.TestHeader))
			}); err != nil {
				return errors.Wrap(err, "encode TestHeader header")
			}
		}
	}
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	return nil

}
func encodeMultipleGenericResponsesResponse(response MultipleGenericResponsesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *NilInt:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *NilString:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/multipleGenericResponses"+`: unexpected response type: %T`, response)
	}
}
func encodeOctetStreamBinaryStringSchemaResponse(response OctetStreamBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	if _, err := io.Copy(w, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeOctetStreamEmptySchemaResponse(response OctetStreamEmptySchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	if _, err := io.Copy(w, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
