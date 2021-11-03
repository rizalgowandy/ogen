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

func decodeCreateSnapshotResponse(resp *http.Response, span trace.Span) (res CreateSnapshotRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &CreateSnapshotNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeCreateSyncActionResponse(resp *http.Response, span trace.Span) (res CreateSyncActionRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &CreateSyncActionNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDescribeBalloonConfigResponse(resp *http.Response, span trace.Span) (res DescribeBalloonConfigRes, err error) {
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

			var response Balloon
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDescribeBalloonStatsResponse(resp *http.Response, span trace.Span) (res DescribeBalloonStatsRes, err error) {
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

			var response BalloonStats
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDescribeInstanceResponse(resp *http.Response, span trace.Span) (res DescribeInstanceRes, err error) {
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

			var response InstanceInfo
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetExportVmConfigResponse(resp *http.Response, span trace.Span) (res GetExportVmConfigRes, err error) {
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

			var response FullVmConfiguration
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMachineConfigurationResponse(resp *http.Response, span trace.Span) (res GetMachineConfigurationRes, err error) {
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

			var response MachineConfiguration
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeLoadSnapshotResponse(resp *http.Response, span trace.Span) (res LoadSnapshotRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &LoadSnapshotNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsConfigPutResponse(resp *http.Response, span trace.Span) (res MmdsConfigPutRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &MmdsConfigPutNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsGetResponse(resp *http.Response, span trace.Span) (res MmdsGetRes, err error) {
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

			var response MmdsGetOK
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsPatchResponse(resp *http.Response, span trace.Span) (res MmdsPatchRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &MmdsPatchNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsPutResponse(resp *http.Response, span trace.Span) (res MmdsPutRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &MmdsPutNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchBalloonResponse(resp *http.Response, span trace.Span) (res PatchBalloonRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchBalloonNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchBalloonStatsIntervalResponse(resp *http.Response, span trace.Span) (res PatchBalloonStatsIntervalRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchBalloonStatsIntervalNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchGuestDriveByIDResponse(resp *http.Response, span trace.Span) (res PatchGuestDriveByIDRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchGuestDriveByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchGuestNetworkInterfaceByIDResponse(resp *http.Response, span trace.Span) (res PatchGuestNetworkInterfaceByIDRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchMachineConfigurationResponse(resp *http.Response, span trace.Span) (res PatchMachineConfigurationRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchMachineConfigurationNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchVmResponse(resp *http.Response, span trace.Span) (res PatchVmRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PatchVmNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutBalloonResponse(resp *http.Response, span trace.Span) (res PutBalloonRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutBalloonNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestBootSourceResponse(resp *http.Response, span trace.Span) (res PutGuestBootSourceRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutGuestBootSourceNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestDriveByIDResponse(resp *http.Response, span trace.Span) (res PutGuestDriveByIDRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutGuestDriveByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestNetworkInterfaceByIDResponse(resp *http.Response, span trace.Span) (res PutGuestNetworkInterfaceByIDRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestVsockResponse(resp *http.Response, span trace.Span) (res PutGuestVsockRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutGuestVsockNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutLoggerResponse(resp *http.Response, span trace.Span) (res PutLoggerRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutLoggerNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutMachineConfigurationResponse(resp *http.Response, span trace.Span) (res PutMachineConfigurationRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutMachineConfigurationNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutMetricsResponse(resp *http.Response, span trace.Span) (res PutMetricsRes, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 204:
		return &PutMetricsNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
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
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}
