// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
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

func decodeCreatePetsResponse(resp *http.Response, span trace.Span) (res CreatePetsRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 201:
		return &CreatePetsCreated{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeListPetsResponse(resp *http.Response, span trace.Span) (res ListPetsRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pets
			if err := func() error {
				if err := fmt.Errorf(`decoding of "Pets" (alias) is not implemented`); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeShowPetByIdResponse(resp *http.Response, span trace.Span) (res ShowPetByIdRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Pet
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}
