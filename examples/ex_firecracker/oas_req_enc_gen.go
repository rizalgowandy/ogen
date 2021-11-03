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

func encodeCreateSnapshotRequest(req SnapshotCreateParams, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCreateSyncActionRequest(req InstanceActionInfo, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeLoadSnapshotRequest(req SnapshotLoadParams, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeMmdsConfigPutRequest(req MmdsConfig, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeMmdsPatchRequest(req MmdsPatchReq, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeMmdsPutRequest(req MmdsPutReq, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchBalloonRequest(req BalloonUpdate, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchBalloonStatsIntervalRequest(req BalloonStatsUpdate, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchGuestDriveByIDRequest(req PartialDrive, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchGuestNetworkInterfaceByIDRequest(req PartialNetworkInterface, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchMachineConfigurationRequest(req MachineConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePatchVmRequest(req VM, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutBalloonRequest(req Balloon, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutGuestBootSourceRequest(req BootSource, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutGuestDriveByIDRequest(req Drive, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutGuestNetworkInterfaceByIDRequest(req NetworkInterface, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutGuestVsockRequest(req Vsock, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutLoggerRequest(req Logger, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutMachineConfigurationRequest(req MachineConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePutMetricsRequest(req Metrics, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}
