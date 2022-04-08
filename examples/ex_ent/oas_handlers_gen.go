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
	"net/netip"
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

// HandleCreatePetRequest handles createPet operation.
//
// POST /pets
func (s *Server) handleCreatePetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeCreatePetRequest(r, span)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreatePet(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleCreatePetCategoriesRequest handles createPetCategories operation.
//
// POST /pets/{id}/categories
func (s *Server) handleCreatePetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPetCategories"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetCategories",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeCreatePetCategoriesParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodeCreatePetCategoriesRequest(r, span)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreatePetCategories(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleCreatePetFriendsRequest handles createPetFriends operation.
//
// POST /pets/{id}/friends
func (s *Server) handleCreatePetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPetFriends"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetFriends",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeCreatePetFriendsParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodeCreatePetFriendsRequest(r, span)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreatePetFriends(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleCreatePetOwnerRequest handles createPetOwner operation.
//
// POST /pets/{id}/owner
func (s *Server) handleCreatePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPetOwner"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetOwner",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeCreatePetOwnerParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodeCreatePetOwnerRequest(r, span)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreatePetOwner(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDeletePetRequest handles deletePet operation.
//
// DELETE /pets/{id}
func (s *Server) handleDeletePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deletePet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeletePet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeDeletePetParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.DeletePet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDeletePetOwnerRequest handles deletePetOwner operation.
//
// DELETE /pets/{id}/owner
func (s *Server) handleDeletePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deletePetOwner"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeletePetOwner",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeDeletePetOwnerParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.DeletePetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListPetRequest handles listPet operation.
//
// GET /pets
func (s *Server) handleListPetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listPet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeListPetParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ListPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListPetCategoriesRequest handles listPetCategories operation.
//
// GET /pets/{id}/categories
func (s *Server) handleListPetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listPetCategories"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetCategories",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeListPetCategoriesParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ListPetCategories(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListPetFriendsRequest handles listPetFriends operation.
//
// GET /pets/{id}/friends
func (s *Server) handleListPetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listPetFriends"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetFriends",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeListPetFriendsParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ListPetFriends(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleReadPetRequest handles readPet operation.
//
// GET /pets/{id}
func (s *Server) handleReadPetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readPet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeReadPetParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ReadPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleReadPetOwnerRequest handles readPetOwner operation.
//
// GET /pets/{id}/owner
func (s *Server) handleReadPetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readPetOwner"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPetOwner",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeReadPetOwnerParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ReadPetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleUpdatePetRequest handles updatePet operation.
//
// PATCH /pets/{id}
func (s *Server) handleUpdatePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updatePet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdatePet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeUpdatePetParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodeUpdatePetRequest(r, span)
	if err != nil {
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.UpdatePet(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	span trace.Span,
	otelAttrs []attribute.KeyValue,
	err error,
) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	s.cfg.ErrorHandler(ctx, w, r, http.StatusBadRequest, err)
}
