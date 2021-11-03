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

func decodeFoobarGetParams(r *http.Request) (FoobarGetParams, error) {
	var params FoobarGetParams
	// Decode param "inlinedParam" located in "Query".
	{
		values, ok := r.URL.Query()["inlinedParam"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.InlinedParam = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `parse parameter inlinedParam located in query`)
			}
		} else {
			return params, errors.New(`query parameter inlinedParam not specified`)
		}
	}
	// Decode param "skip" located in "Query".
	{
		values, ok := r.URL.Query()["skip"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(s)
				if err != nil {
					return err
				}

				params.Skip = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `parse parameter skip located in query`)
			}
		} else {
			return params, errors.New(`query parameter skip not specified`)
		}
	}
	return params, nil
}

func decodePetFriendsNamesByIDParams(r *http.Request) (PetFriendsNamesByIDParams, error) {
	var params PetFriendsNamesByIDParams
	// Decode param "id" located in "Path".
	{
		param := chi.URLParam(r, "id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter id not specified`)
		}
	}
	return params, nil
}

func decodePetGetParams(r *http.Request) (PetGetParams, error) {
	var params PetGetParams
	// Decode param "petID" located in "Query".
	{
		values, ok := r.URL.Query()["petID"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `parse parameter petID located in query`)
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:       true,
					Min:          1337,
					MaxSet:       false,
					Max:          0,
					MinExclusive: false,
					MaxExclusive: false,
				}).Validate(int64(params.PetID)); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `invalid parameter petID (query`)
			}
		} else {
			return params, errors.New(`query parameter petID not specified`)
		}
	}
	// Decode param "x-tags" located in "Header".
	{
		param := r.Header.Get("x-tags")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Explode: false,
			})

			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var ParamsXTagsItem uuid.UUID
					if err := func() error {
						s, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(s)
						if err != nil {
							return err
						}

						ParamsXTagsItem = c
						return nil
					}(); err != nil {
						return err
					}
					params.XTags = append(params.XTags, ParamsXTagsItem)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, `parse header: param 'x-tags`)
			}
		} else {
			return params, errors.New(`header parameter x-tags not specified`)
		}
	}
	// Decode param "x-scope" located in "Header".
	{
		param := r.Header.Get("x-scope")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Explode: false,
			})

			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var ParamsXScopeItem string
					if err := func() error {
						s, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(s)
						if err != nil {
							return err
						}

						ParamsXScopeItem = c
						return nil
					}(); err != nil {
						return err
					}
					params.XScope = append(params.XScope, ParamsXScopeItem)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, `parse header: param 'x-scope`)
			}
		} else {
			return params, errors.New(`header parameter x-scope not specified`)
		}
	}
	// Decode param "token" located in "Query".
	{
		values, ok := r.URL.Query()["token"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Token = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `parse parameter token located in query`)
			}
		} else {
			return params, errors.New(`query parameter token not specified`)
		}
	}
	return params, nil
}

func decodePetGetByNameParams(r *http.Request) (PetGetByNameParams, error) {
	var params PetGetByNameParams
	// Decode param "name" located in "Path".
	{
		param := chi.URLParam(r, "name")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter name not specified`)
		}
	}
	return params, nil
}

func decodePetNameByIDParams(r *http.Request) (PetNameByIDParams, error) {
	var params PetNameByIDParams
	// Decode param "id" located in "Path".
	{
		param := chi.URLParam(r, "id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter id not specified`)
		}
	}
	return params, nil
}
