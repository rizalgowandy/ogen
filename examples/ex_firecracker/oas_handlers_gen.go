// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// HandleCreateSnapshotRequest handles createSnapshot operation.
//
// PUT /snapshot/create
func (s *Server) handleCreateSnapshotRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateSnapshot",
		trace.WithAttributes(otelogen.OperationID("createSnapshot")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateSnapshotRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleCreateSyncActionRequest handles createSyncAction operation.
//
// PUT /actions
func (s *Server) handleCreateSyncActionRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateSyncAction",
		trace.WithAttributes(otelogen.OperationID("createSyncAction")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateSyncActionRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateSyncAction(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateSyncActionResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDescribeBalloonConfigRequest handles describeBalloonConfig operation.
//
// GET /balloon
func (s *Server) handleDescribeBalloonConfigRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeBalloonConfig",
		trace.WithAttributes(otelogen.OperationID("describeBalloonConfig")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.DescribeBalloonConfig(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeBalloonConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDescribeBalloonStatsRequest handles describeBalloonStats operation.
//
// GET /balloon/statistics
func (s *Server) handleDescribeBalloonStatsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeBalloonStats",
		trace.WithAttributes(otelogen.OperationID("describeBalloonStats")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.DescribeBalloonStats(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeBalloonStatsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDescribeInstanceRequest handles describeInstance operation.
//
// GET /
func (s *Server) handleDescribeInstanceRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeInstance",
		trace.WithAttributes(otelogen.OperationID("describeInstance")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.DescribeInstance(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeInstanceResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleGetExportVmConfigRequest handles getExportVmConfig operation.
//
// GET /vm/config
func (s *Server) handleGetExportVmConfigRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetExportVmConfig",
		trace.WithAttributes(otelogen.OperationID("getExportVmConfig")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.GetExportVmConfig(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetExportVmConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleGetMachineConfigurationRequest handles getMachineConfiguration operation.
//
// GET /machine-config
func (s *Server) handleGetMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetMachineConfiguration",
		trace.WithAttributes(otelogen.OperationID("getMachineConfiguration")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.GetMachineConfiguration(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleLoadSnapshotRequest handles loadSnapshot operation.
//
// PUT /snapshot/load
func (s *Server) handleLoadSnapshotRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "LoadSnapshot",
		trace.WithAttributes(otelogen.OperationID("loadSnapshot")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeLoadSnapshotRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.LoadSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeLoadSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleMmdsConfigPutRequest handles  operation.
//
// PUT /mmds/config
func (s *Server) handleMmdsConfigPutRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsConfigPut",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsConfigPutRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.MmdsConfigPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsConfigPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleMmdsGetRequest handles  operation.
//
// GET /mmds
func (s *Server) handleMmdsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsGet",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.MmdsGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleMmdsPatchRequest handles  operation.
//
// PATCH /mmds
func (s *Server) handleMmdsPatchRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsPatch",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsPatchRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.MmdsPatch(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsPatchResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleMmdsPutRequest handles  operation.
//
// PUT /mmds
func (s *Server) handleMmdsPutRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsPut",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsPutRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.MmdsPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchBalloonRequest handles patchBalloon operation.
//
// PATCH /balloon
func (s *Server) handlePatchBalloonRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchBalloon",
		trace.WithAttributes(otelogen.OperationID("patchBalloon")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchBalloonRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchBalloonStatsIntervalRequest handles patchBalloonStatsInterval operation.
//
// PATCH /balloon/statistics
func (s *Server) handlePatchBalloonStatsIntervalRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchBalloonStatsInterval",
		trace.WithAttributes(otelogen.OperationID("patchBalloonStatsInterval")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchBalloonStatsIntervalRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchBalloonStatsInterval(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchBalloonStatsIntervalResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchGuestDriveByIDRequest handles patchGuestDriveByID operation.
//
// PATCH /drives/{drive_id}
func (s *Server) handlePatchGuestDriveByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchGuestDriveByID",
		trace.WithAttributes(otelogen.OperationID("patchGuestDriveByID")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePatchGuestDriveByIDParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePatchGuestDriveByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchGuestNetworkInterfaceByIDRequest handles patchGuestNetworkInterfaceByID operation.
//
// PATCH /network-interfaces/{iface_id}
func (s *Server) handlePatchGuestNetworkInterfaceByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchGuestNetworkInterfaceByID",
		trace.WithAttributes(otelogen.OperationID("patchGuestNetworkInterfaceByID")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePatchGuestNetworkInterfaceByIDParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePatchGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchMachineConfigurationRequest handles patchMachineConfiguration operation.
//
// PATCH /machine-config
func (s *Server) handlePatchMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchMachineConfiguration",
		trace.WithAttributes(otelogen.OperationID("patchMachineConfiguration")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchMachineConfigurationRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePatchVmRequest handles patchVm operation.
//
// PATCH /vm
func (s *Server) handlePatchVmRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchVm",
		trace.WithAttributes(otelogen.OperationID("patchVm")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchVmRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PatchVm(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchVmResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutBalloonRequest handles putBalloon operation.
//
// PUT /balloon
func (s *Server) handlePutBalloonRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutBalloon",
		trace.WithAttributes(otelogen.OperationID("putBalloon")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutBalloonRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutGuestBootSourceRequest handles putGuestBootSource operation.
//
// PUT /boot-source
func (s *Server) handlePutGuestBootSourceRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestBootSource",
		trace.WithAttributes(otelogen.OperationID("putGuestBootSource")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutGuestBootSourceRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutGuestBootSource(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestBootSourceResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutGuestDriveByIDRequest handles putGuestDriveByID operation.
//
// PUT /drives/{drive_id}
func (s *Server) handlePutGuestDriveByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestDriveByID",
		trace.WithAttributes(otelogen.OperationID("putGuestDriveByID")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePutGuestDriveByIDParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePutGuestDriveByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutGuestNetworkInterfaceByIDRequest handles putGuestNetworkInterfaceByID operation.
//
// PUT /network-interfaces/{iface_id}
func (s *Server) handlePutGuestNetworkInterfaceByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestNetworkInterfaceByID",
		trace.WithAttributes(otelogen.OperationID("putGuestNetworkInterfaceByID")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePutGuestNetworkInterfaceByIDParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePutGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutGuestVsockRequest handles putGuestVsock operation.
//
// PUT /vsock
func (s *Server) handlePutGuestVsockRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestVsock",
		trace.WithAttributes(otelogen.OperationID("putGuestVsock")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutGuestVsockRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutGuestVsock(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestVsockResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutLoggerRequest handles putLogger operation.
//
// PUT /logger
func (s *Server) handlePutLoggerRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutLogger",
		trace.WithAttributes(otelogen.OperationID("putLogger")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutLoggerRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutLogger(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutLoggerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutMachineConfigurationRequest handles putMachineConfiguration operation.
//
// PUT /machine-config
func (s *Server) handlePutMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutMachineConfiguration",
		trace.WithAttributes(otelogen.OperationID("putMachineConfiguration")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutMachineConfigurationRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandlePutMetricsRequest handles putMetrics operation.
//
// PUT /metrics
func (s *Server) handlePutMetricsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutMetrics",
		trace.WithAttributes(otelogen.OperationID("putMetrics")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutMetricsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.PutMetrics(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutMetricsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
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
