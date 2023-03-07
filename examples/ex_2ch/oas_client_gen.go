// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

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

// APICaptcha2chcaptchaIDGet invokes GET /api/captcha/2chcaptcha/id operation.
//
// Получение ид для использования 2chcaptcha.
//
// GET /api/captcha/2chcaptcha/id
func (c *Client) APICaptcha2chcaptchaIDGet(ctx context.Context, params APICaptcha2chcaptchaIDGetParams) (*Captcha, error) {
	res, err := c.sendAPICaptcha2chcaptchaIDGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPICaptcha2chcaptchaIDGet(ctx context.Context, params APICaptcha2chcaptchaIDGetParams) (res *Captcha, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptcha2chcaptchaIDGet",
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
	u.Path += "/api/captcha/2chcaptcha/id"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Board.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
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
	result, err := decodeAPICaptcha2chcaptchaIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptcha2chcaptchaShowGet invokes GET /api/captcha/2chcaptcha/show operation.
//
// Отображение 2chcaptcha по id.
//
// GET /api/captcha/2chcaptcha/show
func (c *Client) APICaptcha2chcaptchaShowGet(ctx context.Context, params APICaptcha2chcaptchaShowGetParams) (APICaptcha2chcaptchaShowGetRes, error) {
	res, err := c.sendAPICaptcha2chcaptchaShowGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPICaptcha2chcaptchaShowGet(ctx context.Context, params APICaptcha2chcaptchaShowGetParams) (res APICaptcha2chcaptchaShowGetRes, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptcha2chcaptchaShowGet",
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
	u.Path += "/api/captcha/2chcaptcha/show"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.ID))
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
	result, err := decodeAPICaptcha2chcaptchaShowGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptchaAppIDPublicKeyGet invokes GET /api/captcha/app/id/{public_key} operation.
//
// Полученный id вам нужно отправить вместе с постом как
// app_response_id.
// При этом нужно отправить app_response = sha256(app_response_id + '|' +
// private key).
// Срок жизни id: 180 секунд.
//
// GET /api/captcha/app/id/{public_key}
func (c *Client) APICaptchaAppIDPublicKeyGet(ctx context.Context, params APICaptchaAppIDPublicKeyGetParams) (*Captcha, error) {
	res, err := c.sendAPICaptchaAppIDPublicKeyGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPICaptchaAppIDPublicKeyGet(ctx context.Context, params APICaptchaAppIDPublicKeyGetParams) (res *Captcha, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptchaAppIDPublicKeyGet",
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
	u.Path += "/api/captcha/app/id/"
	{
		// Encode "public_key" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "public_key",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.PublicKey))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
	}

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Board.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
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
	result, err := decodeAPICaptchaAppIDPublicKeyGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptchaInvisibleRecaptchaIDGet invokes GET /api/captcha/invisible_recaptcha/id operation.
//
// Получение публичного ключа invisible recaptcha.
//
// GET /api/captcha/invisible_recaptcha/id
func (c *Client) APICaptchaInvisibleRecaptchaIDGet(ctx context.Context, params APICaptchaInvisibleRecaptchaIDGetParams) (*Captcha, error) {
	res, err := c.sendAPICaptchaInvisibleRecaptchaIDGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPICaptchaInvisibleRecaptchaIDGet(ctx context.Context, params APICaptchaInvisibleRecaptchaIDGetParams) (res *Captcha, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptchaInvisibleRecaptchaIDGet",
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
	u.Path += "/api/captcha/invisible_recaptcha/id"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Board.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
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
	result, err := decodeAPICaptchaInvisibleRecaptchaIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptchaInvisibleRecaptchaMobileGet invokes GET /api/captcha/invisible_recaptcha/mobile operation.
//
// Получение html страницы для решения капчи, CORS отключён.
//
// GET /api/captcha/invisible_recaptcha/mobile
func (c *Client) APICaptchaInvisibleRecaptchaMobileGet(ctx context.Context) error {
	res, err := c.sendAPICaptchaInvisibleRecaptchaMobileGet(ctx)
	_ = res
	return err
}

func (c *Client) sendAPICaptchaInvisibleRecaptchaMobileGet(ctx context.Context) (res *APICaptchaInvisibleRecaptchaMobileGetOK, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptchaInvisibleRecaptchaMobileGet",
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
	u.Path += "/api/captcha/invisible_recaptcha/mobile"

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
	result, err := decodeAPICaptchaInvisibleRecaptchaMobileGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptchaRecaptchaIDGet invokes GET /api/captcha/recaptcha/id operation.
//
// Получение публичного ключа recaptcha v2.
//
// GET /api/captcha/recaptcha/id
func (c *Client) APICaptchaRecaptchaIDGet(ctx context.Context, params APICaptchaRecaptchaIDGetParams) (*Captcha, error) {
	res, err := c.sendAPICaptchaRecaptchaIDGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPICaptchaRecaptchaIDGet(ctx context.Context, params APICaptchaRecaptchaIDGetParams) (res *Captcha, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptchaRecaptchaIDGet",
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
	u.Path += "/api/captcha/recaptcha/id"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Board.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
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
	result, err := decodeAPICaptchaRecaptchaIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APICaptchaRecaptchaMobileGet invokes GET /api/captcha/recaptcha/mobile operation.
//
// Получение html страницы для решения капчи, CORS отключён.
//
// GET /api/captcha/recaptcha/mobile
func (c *Client) APICaptchaRecaptchaMobileGet(ctx context.Context) error {
	res, err := c.sendAPICaptchaRecaptchaMobileGet(ctx)
	_ = res
	return err
}

func (c *Client) sendAPICaptchaRecaptchaMobileGet(ctx context.Context) (res *APICaptchaRecaptchaMobileGetOK, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APICaptchaRecaptchaMobileGet",
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
	u.Path += "/api/captcha/recaptcha/mobile"

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
	result, err := decodeAPICaptchaRecaptchaMobileGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIDislikeGet invokes GET /api/dislike operation.
//
// Добавление дизлайка на пост.
//
// GET /api/dislike
func (c *Client) APIDislikeGet(ctx context.Context, params APIDislikeGetParams) (*Like, error) {
	res, err := c.sendAPIDislikeGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPIDislikeGet(ctx context.Context, params APIDislikeGetParams) (res *Like, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APIDislikeGet",
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
	u.Path += "/api/dislike"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Board))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "num" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.Num))
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
	result, err := decodeAPIDislikeGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APILikeGet invokes GET /api/like operation.
//
// Добавление лайка на пост.
//
// GET /api/like
func (c *Client) APILikeGet(ctx context.Context, params APILikeGetParams) (*Like, error) {
	res, err := c.sendAPILikeGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPILikeGet(ctx context.Context, params APILikeGetParams) (res *Like, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APILikeGet",
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
	u.Path += "/api/like"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "board" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Board))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "num" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.Num))
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
	result, err := decodeAPILikeGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIMobileV2AfterBoardThreadNumGet invokes GET /api/mobile/v2/after/{board}/{thread}/{num} operation.
//
// Получение постов в треде >= указанного. Не
// рекомендуется использовать для получения треда
// целиком, только для проверки новых постов.
//
// GET /api/mobile/v2/after/{board}/{thread}/{num}
func (c *Client) APIMobileV2AfterBoardThreadNumGet(ctx context.Context, params APIMobileV2AfterBoardThreadNumGetParams) (*MobileThreadPostsAfter, error) {
	res, err := c.sendAPIMobileV2AfterBoardThreadNumGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPIMobileV2AfterBoardThreadNumGet(ctx context.Context, params APIMobileV2AfterBoardThreadNumGetParams) (res *MobileThreadPostsAfter, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APIMobileV2AfterBoardThreadNumGet",
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
	u.Path += "/api/mobile/v2/after/"
	{
		// Encode "board" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "board",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Board))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
	}
	u.Path += "/"
	{
		// Encode "thread" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "thread",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.Thread))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
	}
	u.Path += "/"
	{
		// Encode "num" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "num",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.Num))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
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
	result, err := decodeAPIMobileV2AfterBoardThreadNumGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIMobileV2BoardsGet invokes GET /api/mobile/v2/boards operation.
//
// Получение списка досок и их настроек.
//
// GET /api/mobile/v2/boards
func (c *Client) APIMobileV2BoardsGet(ctx context.Context) (Boards, error) {
	res, err := c.sendAPIMobileV2BoardsGet(ctx)
	_ = res
	return res, err
}

func (c *Client) sendAPIMobileV2BoardsGet(ctx context.Context) (res Boards, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APIMobileV2BoardsGet",
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
	u.Path += "/api/mobile/v2/boards"

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
	result, err := decodeAPIMobileV2BoardsGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIMobileV2InfoBoardThreadGet invokes GET /api/mobile/v2/info/{board}/{thread} operation.
//
// Получение информации о треде.
//
// GET /api/mobile/v2/info/{board}/{thread}
func (c *Client) APIMobileV2InfoBoardThreadGet(ctx context.Context, params APIMobileV2InfoBoardThreadGetParams) (*MobileThreadLastInfo, error) {
	res, err := c.sendAPIMobileV2InfoBoardThreadGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPIMobileV2InfoBoardThreadGet(ctx context.Context, params APIMobileV2InfoBoardThreadGetParams) (res *MobileThreadLastInfo, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APIMobileV2InfoBoardThreadGet",
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
	u.Path += "/api/mobile/v2/info/"
	{
		// Encode "board" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "board",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Board))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
	}
	u.Path += "/"
	{
		// Encode "thread" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "thread",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.Thread))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
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
	result, err := decodeAPIMobileV2InfoBoardThreadGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIMobileV2PostBoardNumGet invokes GET /api/mobile/v2/post/{board}/{num} operation.
//
// Получение информации о посте.
//
// GET /api/mobile/v2/post/{board}/{num}
func (c *Client) APIMobileV2PostBoardNumGet(ctx context.Context, params APIMobileV2PostBoardNumGetParams) (*MobilePost, error) {
	res, err := c.sendAPIMobileV2PostBoardNumGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAPIMobileV2PostBoardNumGet(ctx context.Context, params APIMobileV2PostBoardNumGetParams) (res *MobilePost, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "APIMobileV2PostBoardNumGet",
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
	u.Path += "/api/mobile/v2/post/"
	{
		// Encode "board" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "board",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Board))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
	}
	u.Path += "/"
	{
		// Encode "num" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "num",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.Num))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += encoded
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
	result, err := decodeAPIMobileV2PostBoardNumGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserPassloginPost invokes POST /user/passlogin operation.
//
// Авторизация пасскода.
//
// POST /user/passlogin
func (c *Client) UserPassloginPost(ctx context.Context, request OptUserPassloginPostReq, params UserPassloginPostParams) (*Passcode, error) {
	res, err := c.sendUserPassloginPost(ctx, request, params)
	_ = res
	return res, err
}

func (c *Client) sendUserPassloginPost(ctx context.Context, request OptUserPassloginPostReq, params UserPassloginPostParams) (res *Passcode, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "UserPassloginPost",
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
	u.Path += "/user/passlogin"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "json" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "json",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.JSON))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUserPassloginPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUserPassloginPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserPostingPost invokes POST /user/posting operation.
//
// Создание нового поста или треда.
//
// POST /user/posting
func (c *Client) UserPostingPost(ctx context.Context, request OptUserPostingPostReqForm) (UserPostingPostOK, error) {
	res, err := c.sendUserPostingPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendUserPostingPost(ctx context.Context, request OptUserPostingPostReqForm) (res UserPostingPostOK, err error) {
	var otelAttrs []attribute.KeyValue
	// Validate request before sending.
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := request.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
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
	ctx, span := c.cfg.Tracer.Start(ctx, "UserPostingPost",
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
	u.Path += "/user/posting"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUserPostingPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUserPostingPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserReportPost invokes POST /user/report operation.
//
// Отправка жалобы.
//
// POST /user/report
func (c *Client) UserReportPost(ctx context.Context, request OptUserReportPostReq) (*Report, error) {
	res, err := c.sendUserReportPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendUserReportPost(ctx context.Context, request OptUserReportPostReq) (res *Report, err error) {
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
	ctx, span := c.cfg.Tracer.Start(ctx, "UserReportPost",
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
	u.Path += "/user/report"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUserReportPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUserReportPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
