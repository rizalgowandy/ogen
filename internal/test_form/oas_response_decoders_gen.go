// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeTestFormURLEncodedResponse(resp *http.Response, span trace.Span) (res TestFormURLEncodedOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestFormURLEncodedOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestMultipartResponse(resp *http.Response, span trace.Span) (res TestMultipartOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestMultipartOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestMultipartUploadResponse(resp *http.Response, span trace.Span) (res TestMultipartUploadOK, err error) {
	switch resp.StatusCode {
	case 200:
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
			var response TestMultipartUploadOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestShareFormSchemaResponse(resp *http.Response, span trace.Span) (res TestShareFormSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestShareFormSchemaOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}