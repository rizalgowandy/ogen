// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

type HeadersCombinedParams struct {
	Default bool
}

func decodeHeadersCombinedParams(args [0]string, r *http.Request) (params HeadersCombinedParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: default.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "default",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToBool(val)
				if err != nil {
					return err
				}

				params.Default = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: default: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}
