// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
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
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ErrorStatusCode struct {
	StatusCode int   `json:"-"`
	Response   Error `json:"-"`
}

func (*ErrorStatusCode) foobarPostResponse() {}

type FoobarPutDefault struct {
	StatusCode int `json:"-"`
}

type NotFound struct{}

func (*NotFound) foobarGetResponse()  {}
func (*NotFound) foobarPostResponse() {}

// Pet describes #/components/schemas/Pet.
type Pet struct {
	Birthday     time.Time         `json:"birthday"`
	Friends      *[]Pet            `json:"friends"`
	ID           int64             `json:"id"`
	Name         string            `json:"name"`
	Nickname     NilString         `json:"nickname"`
	NullStr      OptionalNilString `json:"nullStr"`
	Rate         time.Duration     `json:"rate"`
	Tag          OptionalUUID      `json:"tag"`
	TestArray1   *[][]string       `json:"testArray1"`
	TestDate     OptionalTime      `json:"testDate"`
	TestDateTime OptionalTime      `json:"testDateTime"`
	TestDuration OptionalDuration  `json:"testDuration"`
	TestFloat1   OptionalFloat64   `json:"testFloat1"`
	TestInteger1 OptionalInt       `json:"testInteger1"`
	TestTime     OptionalTime      `json:"testTime"`
	Type         *PetType          `json:"type"`
	UniqueID     uuid.UUID         `json:"unique_id"`
}

func (*Pet) foobarGetResponse()  {}
func (*Pet) foobarPostResponse() {}
func (*Pet) petCreateRequest()   {}
func (*Pet) petGetResponse()     {}

type PetCreateTextPlainRequest string

func (*PetCreateTextPlainRequest) petCreateRequest() {}

type PetGetDefault struct {
	Message string `json:"message"`
}

type PetGetDefaultStatusCode struct {
	StatusCode int           `json:"-"`
	Response   PetGetDefault `json:"-"`
}

func (*PetGetDefaultStatusCode) petGetResponse() {}

type PetType string

const (
	PetTypeFifa PetType = "fifa"
	PetTypeFofa PetType = "fofa"
)