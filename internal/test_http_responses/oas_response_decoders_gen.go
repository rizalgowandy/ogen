// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeAnyContentTypeBinaryStringSchemaResponse(resp *http.Response, span trace.Span) (res AnyContentTypeBinaryStringSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("*/*", ct):
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := AnyContentTypeBinaryStringSchemaOK{Data: bytes.NewReader(b)}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeAnyContentTypeBinaryStringSchemaDefaultResponse(resp *http.Response, span trace.Span) (res AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, err error) {
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ht.MatchContentType("*/*", ct):
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		response := AnyContentTypeBinaryStringSchemaDefaultDef{Data: bytes.NewReader(b)}
		return AnyContentTypeBinaryStringSchemaDefaultDefStatusCode{
			StatusCode: resp.StatusCode,
			Response:   response,
		}, nil
	default:
		return res, validate.InvalidContentType(ct)
	}
}

func decodeCombinedResponse(resp *http.Response, span trace.Span) (res CombinedRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response CombinedOK
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
	}
	switch resp.StatusCode / 100 {
	case 2:
		// Pattern 2XX.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response int
			if err := func() error {
				v, err := d.Int()
				response = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &Combined2XXStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 5:
		// Pattern 5XX.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response bool
			if err := func() error {
				v, err := d.Bool()
				response = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &Combined5XXStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
		var response []string
		if err := func() error {
			response = make([]string, 0)
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem string
				v, err := d.Str()
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
		return &CombinedDefStatusCode{
			StatusCode: resp.StatusCode,
			Response:   response,
		}, nil
	default:
		return res, validate.InvalidContentType(ct)
	}
}

func decodeHeaders200Response(resp *http.Response, span trace.Span) (res Headers200OK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		var wrapper Headers200OK
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse 'TestHeader' header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				wrapper.TestHeader = c
				return nil
			}); err != nil {
				return res, errors.Wrap(err, "parse TestHeader header")
			}
		}
		return wrapper, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeHeadersCombinedResponse(resp *http.Response, span trace.Span) (res HeadersCombinedRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		var wrapper HeadersCombinedOK
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse 'TestHeader' header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				wrapper.TestHeader = c
				return nil
			}); err != nil {
				return res, errors.Wrap(err, "parse TestHeader header")
			}
		}
		return &wrapper, nil
	}
	switch resp.StatusCode / 100 {
	case 4:
		// Pattern 4XX.
		var wrapper HeadersCombined4XX
		wrapper.StatusCode = resp.StatusCode
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse 'TestHeader' header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				wrapper.TestHeader = c
				return nil
			}); err != nil {
				return res, errors.Wrap(err, "parse TestHeader header")
			}
		}
		return &wrapper, nil
	}
	// Default response.
	var wrapper HeadersCombinedDef
	wrapper.StatusCode = resp.StatusCode
	h := uri.NewHeaderDecoder(resp.Header)
	// Parse 'TestHeader' header.
	{
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "TestHeader",
			Explode: false,
		}
		if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
			val, err := d.DecodeValue()
			if err != nil {
				return err
			}

			c, err := conv.ToString(val)
			if err != nil {
				return err
			}

			wrapper.TestHeader = c
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "parse TestHeader header")
		}
	}
	return &wrapper, nil
}

func decodeHeadersDefaultResponse(resp *http.Response, span trace.Span) (res HeadersDefaultDef, err error) {
	// Default response.
	var wrapper HeadersDefaultDef
	wrapper.StatusCode = resp.StatusCode
	h := uri.NewHeaderDecoder(resp.Header)
	// Parse 'TestHeader' header.
	{
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "TestHeader",
			Explode: false,
		}
		if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
			val, err := d.DecodeValue()
			if err != nil {
				return err
			}

			c, err := conv.ToString(val)
			if err != nil {
				return err
			}

			wrapper.TestHeader = c
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "parse TestHeader header")
		}
	}
	return wrapper, nil
}

func decodeHeadersPatternResponse(resp *http.Response, span trace.Span) (res HeadersPattern4XX, err error) {
	switch resp.StatusCode / 100 {
	case 4:
		// Pattern 4XX.
		var wrapper HeadersPattern4XX
		wrapper.StatusCode = resp.StatusCode
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse 'TestHeader' header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "TestHeader",
				Explode: false,
			}
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				wrapper.TestHeader = c
				return nil
			}); err != nil {
				return res, errors.Wrap(err, "parse TestHeader header")
			}
		}
		return wrapper, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIntersectPatternCodeResponse(resp *http.Response, span trace.Span) (res IntersectPatternCodeRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response IntersectPatternCodeOKApplicationJSON
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
	}
	switch resp.StatusCode / 100 {
	case 2:
		// Pattern 2XX.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response int
			if err := func() error {
				v, err := d.Int()
				response = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &IntersectPatternCode2XXStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeMultipleGenericResponsesResponse(resp *http.Response, span trace.Span) (res MultipleGenericResponsesRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response NilInt
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
	case 201:
		// Code 201.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response NilString
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
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOctetStreamBinaryStringSchemaResponse(resp *http.Response, span trace.Span) (res OctetStreamBinaryStringSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := OctetStreamBinaryStringSchemaOK{Data: bytes.NewReader(b)}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOctetStreamEmptySchemaResponse(resp *http.Response, span trace.Span) (res OctetStreamEmptySchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := OctetStreamEmptySchemaOK{Data: bytes.NewReader(b)}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeTextPlainBinaryStringSchemaResponse(resp *http.Response, span trace.Span) (res TextPlainBinaryStringSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "text/plain":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := TextPlainBinaryStringSchemaOK{Data: bytes.NewReader(b)}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
