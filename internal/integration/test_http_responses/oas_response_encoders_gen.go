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
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeAnyContentTypeBinaryStringSchemaDefaultResponse(response AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	st := http.StatusText(code)
	if code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	writer := w
	if _, err := io.Copy(writer, response.Response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeCombinedResponse(response CombinedRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CombinedOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *Combined2XXStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		e.Int(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *Combined5XXStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		e.Bool(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *CombinedDefStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		e.ArrStart()
		for _, elem := range response.Response {
			e.Str(elem)
		}
		e.ArrEnd()
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeHeaders200Response(response Headers200OK, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "TestHeader" header.
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
	case *HeadersCombinedOK:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "TestHeader" header.
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

	case *HeadersCombined4XX:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "TestHeader" header.
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
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		return nil

	case *HeadersCombinedDef:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "TestHeader" header.
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
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeHeadersDefaultResponse(response HeadersDefaultDef, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "TestHeader" header.
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
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	st := http.StatusText(code)
	if code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	return nil
}

func encodeHeadersPatternResponse(response HeadersPattern4XX, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "TestHeader" header.
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
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	st := http.StatusText(code)
	if code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	return nil
}

func encodeIntersectPatternCodeResponse(response IntersectPatternCodeRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *IntersectPatternCodeOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *IntersectPatternCode2XXStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		e.Int(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeMultipleGenericResponsesResponse(response MultipleGenericResponsesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *NilInt:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
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
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOctetStreamBinaryStringSchemaResponse(response OctetStreamBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeOctetStreamEmptySchemaResponse(response OctetStreamEmptySchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeTextPlainBinaryStringSchemaResponse(response TextPlainBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}
