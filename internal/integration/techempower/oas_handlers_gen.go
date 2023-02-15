// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleCachingRequest handles Caching operation.
//
// Test #7. The Caching test exercises the preferred in-memory or separate-process caching technology
// for the platform or framework. For implementation simplicity, the requirements are very similar to
// the multiple database-query test Test #3, but use a separate database table. The requirements are
// quite generous, affording each framework fairly broad freedom to meet the requirements in the
// manner that best represents the canonical non-distributed caching approach for the framework.
// (Note: a distributed caching test type could be added later.).
//
// GET /cached-worlds
func (s *Server) handleCachingRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Caching"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/cached-worlds"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Caching",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Caching",
			ID:   "Caching",
		}
	)
	params, err := decodeCachingParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response WorldObjects
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Caching",
			OperationID:   "Caching",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "count",
					In:   "query",
				}: params.Count,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = CachingParams
			Response = WorldObjects
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCachingParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Caching(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Caching(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCachingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleDBRequest handles DB operation.
//
// Test #2. The Single Database Query test exercises the framework's object-relational mapper (ORM),
// random number generator, database driver, and database connection pool.
//
// GET /db
func (s *Server) handleDBRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("DB"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/db"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DB",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err error
	)

	var response *WorldObject
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "DB",
			OperationID:   "DB",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *WorldObject
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DB(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.DB(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDBResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleJSONRequest handles json operation.
//
// Test #1. The JSON Serialization test exercises the framework fundamentals including keep-alive
// support, request routing, request header parsing, object instantiation, JSON serialization,
// response header generation, and request count throughput.
//
// GET /json
func (s *Server) handleJSONRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("json"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/json"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "JSON",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err error
	)

	var response *HelloWorld
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "JSON",
			OperationID:   "json",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *HelloWorld
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.JSON(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.JSON(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeJSONResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleQueriesRequest handles Queries operation.
//
// Test #3. The Multiple Database Queries test is a variation of Test #2 and also uses the World
// table. Multiple rows are fetched to more dramatically punish the database driver and connection
// pool. At the highest queries-per-request tested (20), this test demonstrates all frameworks'
// convergence toward zero requests-per-second as database activity increases.
//
// GET /queries
func (s *Server) handleQueriesRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Queries"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/queries"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Queries",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Queries",
			ID:   "Queries",
		}
	)
	params, err := decodeQueriesParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response WorldObjects
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Queries",
			OperationID:   "Queries",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "queries",
					In:   "query",
				}: params.Queries,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = QueriesParams
			Response = WorldObjects
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackQueriesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Queries(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Queries(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeQueriesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleUpdatesRequest handles Updates operation.
//
// Test #5. The Database Updates test is a variation of Test #3 that exercises the ORM's persistence
// of objects and the database driver's performance at running UPDATE statements or similar. The
// spirit of this test is to exercise a variable number of read-then-write style database operations.
//
// GET /updates
func (s *Server) handleUpdatesRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Updates"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/updates"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Updates",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Updates",
			ID:   "Updates",
		}
	)
	params, err := decodeUpdatesParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response WorldObjects
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Updates",
			OperationID:   "Updates",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "queries",
					In:   "query",
				}: params.Queries,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UpdatesParams
			Response = WorldObjects
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdatesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Updates(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Updates(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdatesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
