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

func NewCreateSnapshotHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `CreateSnapshot`,
			trace.WithAttributes(otelogen.OperationID(`createSnapshot`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeCreateSnapshotRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.CreateSnapshot(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeCreateSnapshotResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewCreateSyncActionHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `CreateSyncAction`,
			trace.WithAttributes(otelogen.OperationID(`createSyncAction`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeCreateSyncActionRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.CreateSyncAction(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeCreateSyncActionResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewDescribeBalloonConfigHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `DescribeBalloonConfig`,
			trace.WithAttributes(otelogen.OperationID(`describeBalloonConfig`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.DescribeBalloonConfig(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeDescribeBalloonConfigResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewDescribeBalloonStatsHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `DescribeBalloonStats`,
			trace.WithAttributes(otelogen.OperationID(`describeBalloonStats`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.DescribeBalloonStats(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeDescribeBalloonStatsResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewDescribeInstanceHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `DescribeInstance`,
			trace.WithAttributes(otelogen.OperationID(`describeInstance`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.DescribeInstance(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeDescribeInstanceResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewGetExportVmConfigHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `GetExportVmConfig`,
			trace.WithAttributes(otelogen.OperationID(`getExportVmConfig`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.GetExportVmConfig(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeGetExportVmConfigResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewGetMachineConfigurationHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `GetMachineConfiguration`,
			trace.WithAttributes(otelogen.OperationID(`getMachineConfiguration`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.GetMachineConfiguration(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeGetMachineConfigurationResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewLoadSnapshotHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `LoadSnapshot`,
			trace.WithAttributes(otelogen.OperationID(`loadSnapshot`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeLoadSnapshotRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.LoadSnapshot(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeLoadSnapshotResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewMmdsConfigPutHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `MmdsConfigPut`,
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeMmdsConfigPutRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.MmdsConfigPut(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeMmdsConfigPutResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewMmdsGetHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `MmdsGet`,
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		response, err := s.MmdsGet(ctx)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeMmdsGetResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewMmdsPatchHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `MmdsPatch`,
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeMmdsPatchRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.MmdsPatch(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeMmdsPatchResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewMmdsPutHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `MmdsPut`,
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodeMmdsPutRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.MmdsPut(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeMmdsPutResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchBalloonHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchBalloon`,
			trace.WithAttributes(otelogen.OperationID(`patchBalloon`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePatchBalloonRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchBalloon(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchBalloonResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchBalloonStatsIntervalHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchBalloonStatsInterval`,
			trace.WithAttributes(otelogen.OperationID(`patchBalloonStatsInterval`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePatchBalloonStatsIntervalRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchBalloonStatsInterval(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchBalloonStatsIntervalResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchGuestDriveByIDHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchGuestDriveByID`,
			trace.WithAttributes(otelogen.OperationID(`patchGuestDriveByID`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		params, err := decodePatchGuestDriveByIDParams(r)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}
		request, err := decodePatchGuestDriveByIDRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchGuestDriveByID(ctx, request, params)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchGuestDriveByIDResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchGuestNetworkInterfaceByIDHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchGuestNetworkInterfaceByID`,
			trace.WithAttributes(otelogen.OperationID(`patchGuestNetworkInterfaceByID`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		params, err := decodePatchGuestNetworkInterfaceByIDParams(r)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}
		request, err := decodePatchGuestNetworkInterfaceByIDRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchGuestNetworkInterfaceByID(ctx, request, params)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchMachineConfigurationHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchMachineConfiguration`,
			trace.WithAttributes(otelogen.OperationID(`patchMachineConfiguration`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePatchMachineConfigurationRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchMachineConfiguration(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchMachineConfigurationResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPatchVmHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PatchVm`,
			trace.WithAttributes(otelogen.OperationID(`patchVm`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePatchVmRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PatchVm(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePatchVmResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutBalloonHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutBalloon`,
			trace.WithAttributes(otelogen.OperationID(`putBalloon`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutBalloonRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutBalloon(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutBalloonResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutGuestBootSourceHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutGuestBootSource`,
			trace.WithAttributes(otelogen.OperationID(`putGuestBootSource`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutGuestBootSourceRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutGuestBootSource(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutGuestBootSourceResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutGuestDriveByIDHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutGuestDriveByID`,
			trace.WithAttributes(otelogen.OperationID(`putGuestDriveByID`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		params, err := decodePutGuestDriveByIDParams(r)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}
		request, err := decodePutGuestDriveByIDRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutGuestDriveByID(ctx, request, params)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutGuestDriveByIDResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutGuestNetworkInterfaceByIDHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutGuestNetworkInterfaceByID`,
			trace.WithAttributes(otelogen.OperationID(`putGuestNetworkInterfaceByID`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		params, err := decodePutGuestNetworkInterfaceByIDParams(r)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}
		request, err := decodePutGuestNetworkInterfaceByIDRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutGuestNetworkInterfaceByID(ctx, request, params)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutGuestVsockHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutGuestVsock`,
			trace.WithAttributes(otelogen.OperationID(`putGuestVsock`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutGuestVsockRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutGuestVsock(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutGuestVsockResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutLoggerHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutLogger`,
			trace.WithAttributes(otelogen.OperationID(`putLogger`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutLoggerRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutLogger(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutLoggerResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutMachineConfigurationHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutMachineConfiguration`,
			trace.WithAttributes(otelogen.OperationID(`putMachineConfiguration`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutMachineConfigurationRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutMachineConfiguration(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutMachineConfigurationResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func NewPutMetricsHandler(s Server, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	cfg := newConfig(opts...)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := cfg.Tracer.Start(r.Context(), `PutMetrics`,
			trace.WithAttributes(otelogen.OperationID(`putMetrics`)),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()
		request, err := decodePutMetricsRequest(r, span)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PutMetrics(ctx, request)
		if err != nil {
			span.RecordError(err)
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePutMetricsResponse(response, w, span); err != nil {
			span.RecordError(err)
			return
		}
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
