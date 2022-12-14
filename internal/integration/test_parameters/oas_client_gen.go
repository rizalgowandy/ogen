// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

var _ Handler = struct {
	*Client
}{}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// ComplicatedParameterNameGet invokes GET /complicatedParameterName operation.
//
// GET /complicatedParameterName
func (c *Client) ComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) error {
	res, err := c.sendComplicatedParameterNameGet(ctx, params)
	_ = res
	return err
}

func (c *Client) sendComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) (res *ComplicatedParameterNameGetOK, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ComplicatedParameterNameGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/complicatedParameterName"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "=" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "=",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Eq))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "+" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "+",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Plus))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "question?" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "question?",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Question))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "and&" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "and&",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.And))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "percent%" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "percent%",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Percent))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeComplicatedParameterNameGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ContentQueryParameter invokes contentQueryParameter operation.
//
// GET /contentQueryParameter
func (c *Client) ContentQueryParameter(ctx context.Context, params ContentQueryParameterParams) (string, error) {
	res, err := c.sendContentQueryParameter(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendContentQueryParameter(ctx context.Context, params ContentQueryParameterParams) (res string, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("contentQueryParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ContentQueryParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/contentQueryParameter"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "param" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "param",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {
				if params.Param.Set {
					params.Param.Encode(e)
				}
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeContentQueryParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// HeaderParameter invokes headerParameter operation.
//
// Test for header param.
//
// GET /headerParameter
func (c *Client) HeaderParameter(ctx context.Context, params HeaderParameterParams) (*Hash, error) {
	res, err := c.sendHeaderParameter(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendHeaderParameter(ctx context.Context, params HeaderParameterParams) (res *Hash, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("headerParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "HeaderParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/headerParameter"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "x-auth-token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAuthToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeHeaderParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectQueryParameter invokes objectQueryParameter operation.
//
// GET /objectQueryParameter
func (c *Client) ObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (*ObjectQueryParameterOK, error) {
	res, err := c.sendObjectQueryParameter(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (res *ObjectQueryParameterOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectQueryParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ObjectQueryParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/objectQueryParameter"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "formObject" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "formObject",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FormObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "deepObject" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.DeepObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeObjectQueryParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PathObjectParameter invokes pathObjectParameter operation.
//
// GET /pathObjectParameter/{param}
func (c *Client) PathObjectParameter(ctx context.Context, params PathObjectParameterParams) (*User, error) {
	res, err := c.sendPathObjectParameter(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendPathObjectParameter(ctx context.Context, params PathObjectParameterParams) (res *User, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pathObjectParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PathObjectParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/pathObjectParameter/"
	{
		// Encode "param" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "param",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {

				params.Param.Encode(e)
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodePathObjectParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SameName invokes sameName operation.
//
// Parameter with different location, but the same name.
//
// GET /same_name/{path}
func (c *Client) SameName(ctx context.Context, params SameNameParams) error {
	res, err := c.sendSameName(ctx, params)
	_ = res
	return err
}

func (c *Client) sendSameName(ctx context.Context, params SameNameParams) (res *SameNameOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sameName"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "SameName",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/same_name/"
	{
		// Encode "path" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "path",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.pathPath))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "path" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "path",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.queryPath))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSameNameResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
