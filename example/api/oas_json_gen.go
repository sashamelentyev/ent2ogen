// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Keyboard) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Keyboard) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("name")
		e.Str(s.Name)
	}
	{

		e.FieldStart("switches")
		s.Switches.Encode(e)
	}
	{

		e.FieldStart("keycaps")
		s.Keycaps.Encode(e)
	}
	{

		e.FieldStart("price")
		e.Int64(s.Price)
	}
	{

		e.FieldStart("discount")
		s.Discount.Encode(e)
	}
}

var jsonFieldsNameOfKeyboard = [6]string{
	0: "id",
	1: "name",
	2: "switches",
	3: "keycaps",
	4: "price",
	5: "discount",
}

// Decode decodes Keyboard from json.
func (s *Keyboard) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Keyboard to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "switches":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Switches.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"switches\"")
			}
		case "keycaps":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				if err := s.Keycaps.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"keycaps\"")
			}
		case "price":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int64()
				s.Price = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		case "discount":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				if err := s.Discount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"discount\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Keyboard")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00111111,
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
				if fieldIdx < len(jsonFieldsNameOfKeyboard) {
					name = jsonFieldsNameOfKeyboard[fieldIdx]
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
func (s *Keyboard) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Keyboard) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Keycaps) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Keycaps) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("name")
		e.Str(s.Name)
	}
	{

		e.FieldStart("profile")
		e.Str(s.Profile)
	}
	{

		e.FieldStart("material")
		s.Material.Encode(e)
	}
}

var jsonFieldsNameOfKeycaps = [4]string{
	0: "id",
	1: "name",
	2: "profile",
	3: "material",
}

// Decode decodes Keycaps from json.
func (s *Keycaps) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Keycaps to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "profile":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Profile = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile\"")
			}
		case "material":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				if err := s.Material.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"material\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Keycaps")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
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
				if fieldIdx < len(jsonFieldsNameOfKeycaps) {
					name = jsonFieldsNameOfKeycaps[fieldIdx]
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
func (s *Keycaps) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Keycaps) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes KeycapsMaterial as json.
func (s KeycapsMaterial) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes KeycapsMaterial from json.
func (s *KeycapsMaterial) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode KeycapsMaterial to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch KeycapsMaterial(v) {
	case KeycapsMaterialABS:
		*s = KeycapsMaterialABS
	case KeycapsMaterialPBT:
		*s = KeycapsMaterialPBT
	default:
		*s = KeycapsMaterial(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s KeycapsMaterial) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *KeycapsMaterial) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o NilInt64) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *NilInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilInt64 to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int64
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s NilInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *NilInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Switches) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Switches) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("name")
		e.Str(s.Name)
	}
	{

		e.FieldStart("switch_type")
		s.SwitchType.Encode(e)
	}
}

var jsonFieldsNameOfSwitches = [3]string{
	0: "id",
	1: "name",
	2: "switch_type",
}

// Decode decodes Switches from json.
func (s *Switches) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Switches to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "switch_type":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.SwitchType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"switch_type\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Switches")
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
				if fieldIdx < len(jsonFieldsNameOfSwitches) {
					name = jsonFieldsNameOfSwitches[fieldIdx]
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
func (s *Switches) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Switches) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SwitchesSwitchType as json.
func (s SwitchesSwitchType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes SwitchesSwitchType from json.
func (s *SwitchesSwitchType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SwitchesSwitchType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch SwitchesSwitchType(v) {
	case SwitchesSwitchTypeMechanical:
		*s = SwitchesSwitchTypeMechanical
	case SwitchesSwitchTypeOptical:
		*s = SwitchesSwitchTypeOptical
	case SwitchesSwitchTypeElectrocapacitive:
		*s = SwitchesSwitchTypeElectrocapacitive
	default:
		*s = SwitchesSwitchType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SwitchesSwitchType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SwitchesSwitchType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
