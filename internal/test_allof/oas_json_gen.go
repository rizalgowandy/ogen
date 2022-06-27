// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s Location) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Location) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("lat")
		e.Float64(s.Lat)
	}
	{

		e.FieldStart("lon")
		e.Float64(s.Lon)
	}
}

var jsonFieldsNameOfLocation = [2]string{
	0: "lat",
	1: "lon",
}

// Decode decodes Location from json.
func (s *Location) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Location to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "lat":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Float64()
				s.Lat = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"lat\"")
			}
		case "lon":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Float64()
				s.Lon = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"lon\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Location")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLocation) {
					name = jsonFieldsNameOfLocation[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Location) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Location) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ObjectsWithConflictingArrayPropertyReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s ObjectsWithConflictingArrayPropertyReq) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("bar")
		e.Int(s.Bar)
	}
	{

		e.FieldStart("foo")
		e.ArrStart()
		for _, elem := range s.Foo {
			e.Int(elem)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfObjectsWithConflictingArrayPropertyReq = [2]string{
	0: "bar",
	1: "foo",
}

// Decode decodes ObjectsWithConflictingArrayPropertyReq from json.
func (s *ObjectsWithConflictingArrayPropertyReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ObjectsWithConflictingArrayPropertyReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "bar":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Bar = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bar\"")
			}
		case "foo":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.Foo = make([]int, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem int
					v, err := d.Int()
					elem = int(v)
					if err != nil {
						return err
					}
					s.Foo = append(s.Foo, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"foo\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ObjectsWithConflictingArrayPropertyReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfObjectsWithConflictingArrayPropertyReq) {
					name = jsonFieldsNameOfObjectsWithConflictingArrayPropertyReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ObjectsWithConflictingArrayPropertyReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ObjectsWithConflictingArrayPropertyReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ObjectsWithConflictingPropertiesReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s ObjectsWithConflictingPropertiesReq) encodeFields(e *jx.Encoder) {
	{
		if s.Bar.Set {
			e.FieldStart("bar")
			s.Bar.Encode(e)
		}
	}
	{

		e.FieldStart("foo")
		e.Str(s.Foo)
	}
}

var jsonFieldsNameOfObjectsWithConflictingPropertiesReq = [2]string{
	0: "bar",
	1: "foo",
}

// Decode decodes ObjectsWithConflictingPropertiesReq from json.
func (s *ObjectsWithConflictingPropertiesReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ObjectsWithConflictingPropertiesReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "bar":
			if err := func() error {
				s.Bar.Reset()
				if err := s.Bar.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bar\"")
			}
		case "foo":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Foo = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"foo\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ObjectsWithConflictingPropertiesReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000010,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfObjectsWithConflictingPropertiesReq) {
					name = jsonFieldsNameOfObjectsWithConflictingPropertiesReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ObjectsWithConflictingPropertiesReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ObjectsWithConflictingPropertiesReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Robot) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Robot) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("state")
		s.State.Encode(e)
	}
	{

		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{

		e.FieldStart("location")
		s.Location.Encode(e)
	}
}

var jsonFieldsNameOfRobot = [3]string{
	0: "state",
	1: "id",
	2: "location",
}

// Decode decodes Robot from json.
func (s *Robot) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Robot to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "state":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.State.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"state\"")
			}
		case "id":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "location":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Location.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"location\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Robot")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfRobot) {
					name = jsonFieldsNameOfRobot[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Robot) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Robot) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RobotState as json.
func (s RobotState) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes RobotState from json.
func (s *RobotState) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RobotState to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch RobotState(v) {
	case RobotStateOn:
		*s = RobotStateOn
	case RobotStateOff:
		*s = RobotStateOff
	default:
		*s = RobotState(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RobotState) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RobotState) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s SimpleObjectsReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s SimpleObjectsReq) encodeFields(e *jx.Encoder) {
	{
		if s.Foo.Set {
			e.FieldStart("foo")
			s.Foo.Encode(e)
		}
	}
	{
		if s.Bar.Set {
			e.FieldStart("bar")
			s.Bar.Encode(e)
		}
	}
}

var jsonFieldsNameOfSimpleObjectsReq = [2]string{
	0: "foo",
	1: "bar",
}

// Decode decodes SimpleObjectsReq from json.
func (s *SimpleObjectsReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SimpleObjectsReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "foo":
			if err := func() error {
				s.Foo.Reset()
				if err := s.Foo.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"foo\"")
			}
		case "bar":
			if err := func() error {
				s.Bar.Reset()
				if err := s.Bar.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bar\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SimpleObjectsReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SimpleObjectsReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SimpleObjectsReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
