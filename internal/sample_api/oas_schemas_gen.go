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

// Ref: #/components/schemas/Data
type Data struct {
	Description OptString `json:"description"`
}

// FoobarGetResNotFound is response for FoobarGet operation.
type FoobarGetResNotFound struct{}

func (*FoobarGetResNotFound) foobarGetRes() {}

type FoobarPostDefApplicationJSON struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// FoobarPostDefApplicationJSONStatusCode wraps FoobarPostDefApplicationJSON with StatusCode.
type FoobarPostDefApplicationJSONStatusCode struct {
	StatusCode int
	Response   FoobarPostDefApplicationJSON
}

func (*FoobarPostDefApplicationJSONStatusCode) foobarPostRes() {}

// FoobarPostResNotFound is response for FoobarPost operation.
type FoobarPostResNotFound struct{}

func (*FoobarPostResNotFound) foobarPostRes() {}

// FoobarPutDef is default response for FoobarPut operation.
type FoobarPutDef struct{}

// FoobarPutDefStatusCode wraps FoobarPutDef with StatusCode.
type FoobarPutDefStatusCode struct {
	StatusCode int
	Response   FoobarPutDef
}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// NewOptData returns new OptData with value set to v.
func NewOptData(v Data) OptData {
	return OptData{
		Value: v,
		Set:   true,
	}
}

// OptData is optional Data.
type OptData struct {
	Value Data
	Set   bool
}

// IsSet returns true if OptData was set.
func (o OptData) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptData) Reset() {
	var v Data
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptData) SetTo(v Data) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptData) Get() (v Data, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptDuration returns new OptDuration with value set to v.
func NewOptDuration(v time.Duration) OptDuration {
	return OptDuration{
		Value: v,
		Set:   true,
	}
}

// OptDuration is optional time.Duration.
type OptDuration struct {
	Value time.Duration
	Set   bool
}

// IsSet returns true if OptDuration was set.
func (o OptDuration) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDuration) Reset() {
	var v time.Duration
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDuration) SetTo(v time.Duration) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDuration) Get() (v time.Duration, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptNilString returns new OptNilString with value set to v.
func NewOptNilString(v string) OptNilString {
	return OptNilString{
		Value: v,
		Set:   true,
	}
}

// OptNilString is optional nullable string.
type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilString was set.
func (o OptNilString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilString) SetTo(v string) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilString) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o OptNilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptPetType returns new OptPetType with value set to v.
func NewOptPetType(v PetType) OptPetType {
	return OptPetType{
		Value: v,
		Set:   true,
	}
}

// OptPetType is optional PetType.
type OptPetType struct {
	Value PetType
	Set   bool
}

// IsSet returns true if OptPetType was set.
func (o OptPetType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPetType) Reset() {
	var v PetType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPetType) SetTo(v PetType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPetType) Get() (v PetType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptTime returns new OptTime with value set to v.
func NewOptTime(v time.Time) OptTime {
	return OptTime{
		Value: v,
		Set:   true,
	}
}

// OptTime is optional time.Time.
type OptTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptTime was set.
func (o OptTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// NewOptUUID returns new OptUUID with value set to v.
func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}

// OptUUID is optional uuid.UUID.
type OptUUID struct {
	Value uuid.UUID
	Set   bool
}

// IsSet returns true if OptUUID was set.
func (o OptUUID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Ref: #/components/schemas/Pet
type Pet struct {
	Birthday     time.Time     `json:"birthday"`
	Friends      []Pet         `json:"friends"`
	ID           int64         `json:"id"`
	IP           net.IP        `json:"ip"`
	IPV4         net.IP        `json:"ip_v4"`
	IPV6         net.IP        `json:"ip_v6"`
	Kind         PetKind       `json:"kind"`
	Name         string        `json:"name"`
	Next         OptData       `json:"next"`
	Nickname     NilString     `json:"nickname"`
	NullStr      OptNilString  `json:"nullStr"`
	Primary      *Pet          `json:"primary"`
	Rate         time.Duration `json:"rate"`
	Tag          OptUUID       `json:"tag"`
	TestArray1   [][]string    `json:"testArray1"`
	TestDate     OptTime       `json:"testDate"`
	TestDateTime OptTime       `json:"testDateTime"`
	TestDuration OptDuration   `json:"testDuration"`
	TestFloat1   OptFloat64    `json:"testFloat1"`
	TestInteger1 OptInt        `json:"testInteger1"`
	TestTime     OptTime       `json:"testTime"`
	Type         OptPetType    `json:"type"`
	UniqueID     uuid.UUID     `json:"unique_id"`
	URI          url.URL       `json:"uri"`
}

func (*Pet) foobarGetRes()  {}
func (*Pet) foobarPostRes() {}
func (*Pet) petCreateReq()  {}
func (*Pet) petGetRes()     {}

type PetCreateReqTextPlain struct{}

func (*PetCreateReqTextPlain) petCreateReq() {}

type PetGetDefApplicationJSON struct {
	Message string `json:"message"`
}

// PetGetDefApplicationJSONStatusCode wraps PetGetDefApplicationJSON with StatusCode.
type PetGetDefApplicationJSONStatusCode struct {
	StatusCode int
	Response   PetGetDefApplicationJSON
}

func (*PetGetDefApplicationJSONStatusCode) petGetRes() {}

type PetKind string

const (
	PetKindBig  PetKind = "big"
	PetKindSmol PetKind = "smol"
)

type PetType string

const (
	PetTypeFifa PetType = "fifa"
	PetTypeFofa PetType = "fofa"
)
