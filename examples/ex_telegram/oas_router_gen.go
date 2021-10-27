// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
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
)

func Register(r chi.Router, s Server, opts ...Option) {
	r.MethodFunc("POST", "/answerCallbackQuery", NewAnswerCallbackQueryPostHandler(s, opts...))
	r.MethodFunc("POST", "/answerPreCheckoutQuery", NewAnswerPreCheckoutQueryPostHandler(s, opts...))
	r.MethodFunc("POST", "/answerShippingQuery", NewAnswerShippingQueryPostHandler(s, opts...))
	r.MethodFunc("POST", "/close", NewClosePostHandler(s, opts...))
	r.MethodFunc("POST", "/deleteStickerFromSet", NewDeleteStickerFromSetPostHandler(s, opts...))
	r.MethodFunc("POST", "/deleteWebhook", NewDeleteWebhookPostHandler(s, opts...))
	r.MethodFunc("POST", "/getFile", NewGetFilePostHandler(s, opts...))
	r.MethodFunc("POST", "/getGameHighScores", NewGetGameHighScoresPostHandler(s, opts...))
	r.MethodFunc("POST", "/getMe", NewGetMePostHandler(s, opts...))
	r.MethodFunc("POST", "/getMyCommands", NewGetMyCommandsPostHandler(s, opts...))
	r.MethodFunc("POST", "/getStickerSet", NewGetStickerSetPostHandler(s, opts...))
	r.MethodFunc("POST", "/getUpdates", NewGetUpdatesPostHandler(s, opts...))
	r.MethodFunc("POST", "/getUserProfilePhotos", NewGetUserProfilePhotosPostHandler(s, opts...))
	r.MethodFunc("POST", "/getWebhookInfo", NewGetWebhookInfoPostHandler(s, opts...))
	r.MethodFunc("POST", "/logOut", NewLogOutPostHandler(s, opts...))
	r.MethodFunc("POST", "/sendGame", NewSendGamePostHandler(s, opts...))
	r.MethodFunc("POST", "/sendInvoice", NewSendInvoicePostHandler(s, opts...))
	r.MethodFunc("POST", "/setMyCommands", NewSetMyCommandsPostHandler(s, opts...))
	r.MethodFunc("POST", "/setStickerPositionInSet", NewSetStickerPositionInSetPostHandler(s, opts...))
	r.MethodFunc("POST", "/setWebhook", NewSetWebhookPostHandler(s, opts...))
	r.MethodFunc("POST", "/uploadStickerFile", NewUploadStickerFilePostHandler(s, opts...))
}
