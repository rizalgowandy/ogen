// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
)

type AnyContentTypeBinaryStringSchemaDefaultDef struct {
	Data io.Reader
}

func (s AnyContentTypeBinaryStringSchemaDefaultDef) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// AnyContentTypeBinaryStringSchemaDefaultDefStatusCode wraps AnyContentTypeBinaryStringSchemaDefaultDef with StatusCode.
type AnyContentTypeBinaryStringSchemaDefaultDefStatusCode struct {
	StatusCode int
	Response   AnyContentTypeBinaryStringSchemaDefaultDef
}

type AnyContentTypeBinaryStringSchemaOK struct {
	Data io.Reader
}

func (s AnyContentTypeBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// Headers200OK is response for Headers200 operation.
type Headers200OK struct {
	TestHeader string
}

// HeadersCombinedDef is default response for HeadersCombined operation.
type HeadersCombinedDef struct {
	TestHeader string
	StatusCode int
}

func (*HeadersCombinedDef) headersCombinedRes() {}

// HeadersCombinedOK is response for HeadersCombined operation.
type HeadersCombinedOK struct {
	TestHeader string
}

func (*HeadersCombinedOK) headersCombinedRes() {}

// HeadersDefaultDef is default response for HeadersDefault operation.
type HeadersDefaultDef struct {
	TestHeader string
	StatusCode int
}

// NewNilInt returns new NilInt with value set to v.
func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}

// NilInt is nullable int.
type NilInt struct {
	Value int
	Null  bool
}

// SetTo sets value to v.
func (o *NilInt) SetTo(v int) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilInt) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o NilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilInt) multipleGenericResponsesRes() {}

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

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilString) multipleGenericResponsesRes() {}

type OctetStreamBinaryStringSchemaOK struct {
	Data io.Reader
}

func (s OctetStreamBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

type OctetStreamEmptySchemaOK struct {
	Data io.Reader
}

func (s OctetStreamEmptySchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}
