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

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// AnswerCallbackQueryPost implements  operation.
	AnswerCallbackQueryPost(ctx context.Context, req AnswerCallbackQueryPostReqApplicationJSON) (AnswerCallbackQueryPostRes, error)
	// AnswerPreCheckoutQueryPost implements  operation.
	AnswerPreCheckoutQueryPost(ctx context.Context, req AnswerPreCheckoutQueryPostReqApplicationJSON) (AnswerPreCheckoutQueryPostRes, error)
	// AnswerShippingQueryPost implements  operation.
	AnswerShippingQueryPost(ctx context.Context, req AnswerShippingQueryPostReqApplicationJSON) (AnswerShippingQueryPostRes, error)
	// ClosePost implements  operation.
	ClosePost(ctx context.Context) (ClosePostRes, error)
	// DeleteStickerFromSetPost implements  operation.
	DeleteStickerFromSetPost(ctx context.Context, req DeleteStickerFromSetPostReqApplicationJSON) (DeleteStickerFromSetPostRes, error)
	// DeleteWebhookPost implements  operation.
	DeleteWebhookPost(ctx context.Context, req DeleteWebhookPostReqApplicationJSON) (DeleteWebhookPostRes, error)
	// GetFilePost implements  operation.
	GetFilePost(ctx context.Context, req GetFilePostReqApplicationJSON) (GetFilePostRes, error)
	// GetGameHighScoresPost implements  operation.
	GetGameHighScoresPost(ctx context.Context, req GetGameHighScoresPostReqApplicationJSON) (GetGameHighScoresPostRes, error)
	// GetMePost implements  operation.
	GetMePost(ctx context.Context) (GetMePostRes, error)
	// GetMyCommandsPost implements  operation.
	GetMyCommandsPost(ctx context.Context) (GetMyCommandsPostRes, error)
	// GetStickerSetPost implements  operation.
	GetStickerSetPost(ctx context.Context, req GetStickerSetPostReqApplicationJSON) (GetStickerSetPostRes, error)
	// GetUpdatesPost implements  operation.
	GetUpdatesPost(ctx context.Context, req GetUpdatesPostReqApplicationJSON) (GetUpdatesPostRes, error)
	// GetUserProfilePhotosPost implements  operation.
	GetUserProfilePhotosPost(ctx context.Context, req GetUserProfilePhotosPostReqApplicationJSON) (GetUserProfilePhotosPostRes, error)
	// GetWebhookInfoPost implements  operation.
	GetWebhookInfoPost(ctx context.Context) (GetWebhookInfoPostRes, error)
	// LogOutPost implements  operation.
	LogOutPost(ctx context.Context) (LogOutPostRes, error)
	// SendGamePost implements  operation.
	SendGamePost(ctx context.Context, req SendGamePostReqApplicationJSON) (SendGamePostRes, error)
	// SendInvoicePost implements  operation.
	SendInvoicePost(ctx context.Context, req SendInvoicePostReqApplicationJSON) (SendInvoicePostRes, error)
	// SetMyCommandsPost implements  operation.
	SetMyCommandsPost(ctx context.Context, req SetMyCommandsPostReqApplicationJSON) (SetMyCommandsPostRes, error)
	// SetStickerPositionInSetPost implements  operation.
	SetStickerPositionInSetPost(ctx context.Context, req SetStickerPositionInSetPostReqApplicationJSON) (SetStickerPositionInSetPostRes, error)
}
