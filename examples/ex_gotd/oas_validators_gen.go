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
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

func (s *Chat) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ChatLocation) Validate() error {
	var failures []validate.FieldError
	{
		// Validate "address" property.
		validator := validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    64,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Address)); err != nil {
			failures = append(failures, validate.FieldError{Name: "address", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ForwardMessageResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Game) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Message) Validate() error {
	var failures []validate.FieldError
	if s.Chat == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "chat",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Poll) Validate() error {
	var failures []validate.FieldError
	{
		// Validate "question" property.
		validator := validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    300,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Question)); err != nil {
			failures = append(failures, validate.FieldError{Name: "question", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendAnimationResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendContactResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendDiceResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendDocumentResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendGameResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoiceResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendLocationResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendMessageResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendPhotoResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendPollResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendStickerResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVenueResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVideoNoteResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVideoResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVoiceResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
