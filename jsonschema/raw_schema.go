package jsonschema

import (
	"reflect"

	"github.com/go-faster/errors"
	"github.com/go-json-experiment/json"

	ogenjson "github.com/ogen-go/ogen/json"
)

// Num represents JSON number.
type Num json.RawValue

// MarshalNextJSON implements json.MarshalerV2.
func (n Num) MarshalNextJSON(opts json.MarshalOptions, e *json.Encoder) error {
	val := json.RawValue(n)
	return opts.MarshalNext(e, val)
}

// UnmarshalNextJSON implements json.UnmarshalerV2.
func (n *Num) UnmarshalNextJSON(opts json.UnmarshalOptions, d *json.Decoder) error {
	offset := d.InputOffset()
	// Check Kind for invalid, next call will return error.
	if kind := d.PeekKind(); kind != '0' && kind != 0 {
		return &json.SemanticError{
			ByteOffset:  offset,
			JSONPointer: d.StackPointer(),
			JSONKind:    kind,
			GoType:      reflect.TypeOf(n),
			Err:         errors.Errorf("unexpected type %s", kind.String()),
		}
	}

	val, err := d.ReadValue()
	if err != nil {
		return err
	}

	*n = Num(val)
	return nil
}

// RawSchema is unparsed JSON Schema.
type RawSchema struct {
	Ref                  string                   `json:"$ref,omitzero"`
	Summary              string                   `json:"summary,omitzero"`
	Description          string                   `json:"description,omitzero"`
	Type                 string                   `json:"type,omitzero"`
	Format               string                   `json:"format,omitzero"`
	Properties           RawProperties            `json:"properties,omitzero"`
	AdditionalProperties *AdditionalProperties    `json:"additionalProperties,omitzero"`
	PatternProperties    RawPatternProperties     `json:"patternProperties,omitzero"`
	Required             []string                 `json:"required,omitzero"`
	Items                *RawSchema               `json:"items,omitzero"`
	Nullable             bool                     `json:"nullable,omitzero"`
	AllOf                []*RawSchema             `json:"allOf,omitzero"`
	OneOf                []*RawSchema             `json:"oneOf,omitzero"`
	AnyOf                []*RawSchema             `json:"anyOf,omitzero"`
	Discriminator        *Discriminator           `json:"discriminator,omitzero"`
	Enum                 Enum                     `json:"enum,omitzero"`
	MultipleOf           Num                      `json:"multipleOf,omitzero"`
	Maximum              Num                      `json:"maximum,omitzero"`
	ExclusiveMaximum     bool                     `json:"exclusiveMaximum,omitzero"`
	Minimum              Num                      `json:"minimum,omitzero"`
	ExclusiveMinimum     bool                     `json:"exclusiveMinimum,omitzero"`
	MaxLength            *uint64                  `json:"maxLength,omitzero"`
	MinLength            *uint64                  `json:"minLength,omitzero"`
	Pattern              string                   `json:"pattern,omitzero"`
	MaxItems             *uint64                  `json:"maxItems,omitzero"`
	MinItems             *uint64                  `json:"minItems,omitzero"`
	UniqueItems          bool                     `json:"uniqueItems,omitzero"`
	MaxProperties        *uint64                  `json:"maxProperties,omitzero"`
	MinProperties        *uint64                  `json:"minProperties,omitzero"`
	Default              json.RawValue            `json:"default,omitzero"`
	Example              json.RawValue            `json:"example,omitzero"`
	Deprecated           bool                     `json:"deprecated,omitzero"`
	ContentEncoding      string                   `json:"contentEncoding,omitzero"`
	ContentMediaType     string                   `json:"contentMediaType,omitzero"`
	XAnnotations         map[string]json.RawValue `json:",inline"`
}

// Enum is JSON Schema enum validator description.
type Enum []json.RawValue

// UnmarshalNextJSON implements json.UnmarshalerV2.
func (n *Enum) UnmarshalNextJSON(opts json.UnmarshalOptions, d *json.Decoder) error {
	offset := d.InputOffset()
	// Check Kind for invalid, next call will return error.
	if kind := d.PeekKind(); kind != '[' && kind != 0 {
		return &json.SemanticError{
			ByteOffset:  offset,
			JSONPointer: d.StackPointer(),
			JSONKind:    kind,
			GoType:      reflect.TypeOf(n),
			Err:         errors.Errorf("unexpected type %s", kind.String()),
		}
	}
	// Read the opening bracket.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	// Keep non-nil value, to distinguish from not set array.
	values := Enum{}
	for {
		if kind := d.PeekKind(); kind == ']' || kind == 0 {
			break
		}
		var (
			val    json.RawValue
			offset = d.InputOffset()
		)
		if err := opts.UnmarshalNext(d, &val); err != nil {
			return err
		}
		for _, val2 := range values {
			if ok, _ := ogenjson.Equal(val, val2); ok {
				return &json.SemanticError{
					ByteOffset:  offset,
					JSONPointer: d.StackPointer(),
					JSONKind:    val.Kind(),
					GoType:      reflect.TypeOf(val),
					Err:         errors.Errorf("duplicate value %s", val.String()),
				}
			}
		}
		values = append(values, val)
	}
	// Read the closing bracket.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	*n = values
	return nil
}

// RawProperty is item of RawProperties.
type RawProperty struct {
	Name   string
	Schema *RawSchema
}

// RawProperties is unparsed JSON Schema properties validator description.
type RawProperties []RawProperty

// MarshalNextJSON implements json.MarshalerV2.
func (p RawProperties) MarshalNextJSON(opts json.MarshalOptions, e *json.Encoder) error {
	if err := e.WriteToken(json.ObjectStart); err != nil {
		return err
	}
	for _, member := range p {
		if err := opts.MarshalNext(e, member.Name); err != nil {
			return err
		}
		if err := opts.MarshalNext(e, member.Schema); err != nil {
			return err
		}
	}
	if err := e.WriteToken(json.ObjectEnd); err != nil {
		return err
	}
	return nil
}

// UnmarshalNextJSON implements json.UnmarshalerV2.
func (p *RawProperties) UnmarshalNextJSON(opts json.UnmarshalOptions, d *json.Decoder) error {
	offset := d.InputOffset()
	// Check Kind for invalid, next call will return error.
	if kind := d.PeekKind(); kind != '{' && kind != 0 {
		return &json.SemanticError{
			ByteOffset:  offset,
			JSONPointer: d.StackPointer(),
			JSONKind:    kind,
			GoType:      reflect.TypeOf(p),
			Err:         errors.Errorf("unexpected type %s", kind.String()),
		}
	}
	// Read the opening brace.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	// Keep non-nil value, to distinguish from not set object.
	properties := RawProperties{}
	for d.PeekKind() == '"' {
		var (
			name   string
			schema *RawSchema
		)
		if err := opts.UnmarshalNext(d, &name); err != nil {
			return err
		}
		if err := opts.UnmarshalNext(d, &schema); err != nil {
			return err
		}
		properties = append(properties, RawProperty{Name: name, Schema: schema})
	}
	// Read the closing brace.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	*p = properties
	return nil
}

// AdditionalProperties is JSON Schema additionalProperties validator description.
type AdditionalProperties struct {
	Bool   *bool
	Schema RawSchema
}

// MarshalNextJSON implements json.MarshalerV2.
func (p AdditionalProperties) MarshalNextJSON(opts json.MarshalOptions, e *json.Encoder) error {
	if p.Bool != nil {
		return opts.MarshalNext(e, p.Bool)
	}
	return opts.MarshalNext(e, p.Schema)
}

// UnmarshalNextJSON implements json.UnmarshalerV2.
func (p *AdditionalProperties) UnmarshalNextJSON(opts json.UnmarshalOptions, d *json.Decoder) error {
	offset := d.InputOffset()
	switch kind := d.PeekKind(); kind {
	case 't', 'f':
		return opts.UnmarshalNext(d, &p.Bool)
	case '{':
		return opts.UnmarshalNext(d, &p.Schema)
	default:
		return &json.SemanticError{
			ByteOffset:  offset,
			JSONPointer: d.StackPointer(),
			JSONKind:    kind,
			GoType:      reflect.TypeOf(p),
			Err:         errors.Errorf("unexpected type %s", kind.String()),
		}
	}
}

// RawPatternProperty is item of RawPatternProperties.
type RawPatternProperty struct {
	Pattern string
	Schema  *RawSchema
}

// RawPatternProperties is unparsed JSON Schema patternProperties validator description.
type RawPatternProperties []RawPatternProperty

// MarshalNextJSON implements json.MarshalerV2.
func (r RawPatternProperties) MarshalNextJSON(opts json.MarshalOptions, e *json.Encoder) error {
	if err := e.WriteToken(json.ObjectStart); err != nil {
		return err
	}
	for _, member := range r {
		if err := opts.MarshalNext(e, member.Pattern); err != nil {
			return err
		}
		if err := opts.MarshalNext(e, member.Schema); err != nil {
			return err
		}
	}
	if err := e.WriteToken(json.ObjectEnd); err != nil {
		return err
	}
	return nil
}

// UnmarshalNextJSON implements json.UnmarshalerV2.
func (r *RawPatternProperties) UnmarshalNextJSON(opts json.UnmarshalOptions, d *json.Decoder) error {
	offset := d.InputOffset()
	// Check Kind for invalid, next call will return error.
	if kind := d.PeekKind(); kind != '{' && kind != 0 {
		return &json.SemanticError{
			ByteOffset:  offset,
			JSONPointer: d.StackPointer(),
			JSONKind:    kind,
			GoType:      reflect.TypeOf(r),
			Err:         errors.Errorf("unexpected type %s", kind.String()),
		}
	}
	// Read the opening brace.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	// Keep non-nil value, to distinguish from not set object.
	patternProperties := RawPatternProperties{}
	for d.PeekKind() == '"' {
		var (
			pattern string
			schema  *RawSchema
		)
		if err := opts.UnmarshalNext(d, &pattern); err != nil {
			return err
		}
		if err := opts.UnmarshalNext(d, &schema); err != nil {
			return err
		}
		patternProperties = append(patternProperties, RawPatternProperty{Pattern: pattern, Schema: schema})
	}
	// Read the closing brace.
	if _, err := d.ReadToken(); err != nil {
		return err
	}

	*r = patternProperties
	return nil
}

// Discriminator is JSON Schema discriminator description.
type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitzero"`
}
