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

func (s Book) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.ID // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.MediaID // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "media_id",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Images // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "images",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Tags {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tags",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.NumPages // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "num_pages",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.NumFavorites // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "num_favorites",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Image) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.W // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "w",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.H // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "h",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Images) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Pages {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pages",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Cover // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "cover",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumbnail // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumbnail",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SearchByTagIDOKApplicationJSON) Validate() error {
	return nil
}
func (s SearchOKApplicationJSON) Validate() error {
	return nil
}
func (s SearchResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Result {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Tag) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.ID // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Type // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
