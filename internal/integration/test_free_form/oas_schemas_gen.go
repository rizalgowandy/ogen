// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/jx"
)

type Path1Def map[string]jx.Raw

func (s *Path1Def) init() Path1Def {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

// Path1DefStatusCode wraps Path1Def with StatusCode.
type Path1DefStatusCode struct {
	StatusCode int
	Response   Path1Def
}

// GetStatusCode returns the value of StatusCode.
func (s *Path1DefStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *Path1DefStatusCode) GetResponse() Path1Def {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *Path1DefStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *Path1DefStatusCode) SetResponse(val Path1Def) {
	s.Response = val
}

type Path2Def map[string]jx.Raw

func (s *Path2Def) init() Path2Def {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

// Path2DefStatusCode wraps Path2Def with StatusCode.
type Path2DefStatusCode struct {
	StatusCode int
	Response   Path2Def
}

// GetStatusCode returns the value of StatusCode.
func (s *Path2DefStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *Path2DefStatusCode) GetResponse() Path2Def {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *Path2DefStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *Path2DefStatusCode) SetResponse(val Path2Def) {
	s.Response = val
}
