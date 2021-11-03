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

func decodePatchGuestDriveByIDParams(r *http.Request) (PatchGuestDriveByIDParams, error) {
	var params PatchGuestDriveByIDParams
	// Decode param "drive_id" located in "Path".
	{
		param := chi.URLParam(r, "drive_id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
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

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter drive_id not specified`)
		}
	}
	return params, nil
}

func decodePatchGuestNetworkInterfaceByIDParams(r *http.Request) (PatchGuestNetworkInterfaceByIDParams, error) {
	var params PatchGuestNetworkInterfaceByIDParams
	// Decode param "iface_id" located in "Path".
	{
		param := chi.URLParam(r, "iface_id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
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

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter iface_id not specified`)
		}
	}
	return params, nil
}

func decodePutGuestDriveByIDParams(r *http.Request) (PutGuestDriveByIDParams, error) {
	var params PutGuestDriveByIDParams
	// Decode param "drive_id" located in "Path".
	{
		param := chi.URLParam(r, "drive_id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
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

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter drive_id not specified`)
		}
	}
	return params, nil
}

func decodePutGuestNetworkInterfaceByIDParams(r *http.Request) (PutGuestNetworkInterfaceByIDParams, error) {
	var params PutGuestNetworkInterfaceByIDParams
	// Decode param "iface_id" located in "Path".
	{
		param := chi.URLParam(r, "iface_id")
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
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

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path parameter iface_id not specified`)
		}
	}
	return params, nil
}
