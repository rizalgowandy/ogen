// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeAddPetRequest(r *http.Request, span trace.Span) (
	req NewPet,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request NewPet
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
