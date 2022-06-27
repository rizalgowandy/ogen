// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// ObjectsWithConflictingArrayProperty implements objectsWithConflictingArrayProperty operation.
//
// Objects with conflicting array property.
//
// GET /objectsWithConflictingArrayProperty
func (UnimplementedHandler) ObjectsWithConflictingArrayProperty(ctx context.Context, req ObjectsWithConflictingArrayPropertyReq) (r ObjectsWithConflictingArrayPropertyOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ObjectsWithConflictingProperties implements objectsWithConflictingProperties operation.
//
// Objects with conflicting properties.
//
// GET /objectsWithConflictingProperties
func (UnimplementedHandler) ObjectsWithConflictingProperties(ctx context.Context, req ObjectsWithConflictingPropertiesReq) (r ObjectsWithConflictingPropertiesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ReferencedAllof implements referencedAllof operation.
//
// Referenced allOf.
//
// GET /referencedAllof
func (UnimplementedHandler) ReferencedAllof(ctx context.Context, req Robot) (r ReferencedAllofOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SimpleInteger implements simpleInteger operation.
//
// Simple integers with validation.
//
// GET /simpleInteger
func (UnimplementedHandler) SimpleInteger(ctx context.Context, req int) (r SimpleIntegerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SimpleObjects implements simpleObjects operation.
//
// Simple objects.
//
// GET /simpleObjects
func (UnimplementedHandler) SimpleObjects(ctx context.Context, req SimpleObjectsReq) (r SimpleObjectsOK, _ error) {
	return r, ht.ErrNotImplemented
}
