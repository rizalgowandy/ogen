// Code generated by ogen, DO NOT EDIT.

package techempower

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

// WriteJSON implements json.Marshaler.
func (s HelloWorld) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	e.ObjField("message")
	e.Str(s.Message)
	e.ObjEnd()
}

// ReadJSON reads HelloWorld from json stream.
func (s *HelloWorld) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode HelloWorld to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			v, err := d.Str()
			s.Message = string(v)
			if err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// WriteJSON implements json.Marshaler.
func (s WorldObject) WriteJSON(e *json.Encoder) {
	e.ObjStart()
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	e.ObjField("id")
	e.Int64(s.ID)
	more.More()
	e.ObjField("randomNumber")
	e.Int64(s.RandomNumber)
	e.ObjEnd()
}

// ReadJSON reads WorldObject from json stream.
func (s *WorldObject) ReadJSON(d *json.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode WorldObject to nil`)
	}
	return d.ObjBytes(func(d *json.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			v, err := d.Int64()
			s.ID = int64(v)
			if err != nil {
				return err
			}
		case "randomNumber":
			v, err := d.Int64()
			s.RandomNumber = int64(v)
			if err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

func (WorldObjects) WriteJSON(e *json.Encoder)      {}
func (WorldObjects) ReadJSON(d *json.Decoder) error { return nil }
