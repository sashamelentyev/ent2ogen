// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"github.com/go-faster/errors"
)

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

// NewOptNilBool returns new OptNilBool with value set to v.
func NewOptNilBool(v bool) OptNilBool {
	return OptNilBool{
		Value: v,
		Set:   true,
	}
}

// OptNilBool is optional nullable bool.
type OptNilBool struct {
	Value bool
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilBool was set.
func (o OptNilBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilBool) SetTo(v bool) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilBool) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o OptNilBool) Get() (v bool, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilSchemaAOptionalNullableEnum returns new OptNilSchemaAOptionalNullableEnum with value set to v.
func NewOptNilSchemaAOptionalNullableEnum(v SchemaAOptionalNullableEnum) OptNilSchemaAOptionalNullableEnum {
	return OptNilSchemaAOptionalNullableEnum{
		Value: v,
		Set:   true,
	}
}

// OptNilSchemaAOptionalNullableEnum is optional nullable SchemaAOptionalNullableEnum.
type OptNilSchemaAOptionalNullableEnum struct {
	Value SchemaAOptionalNullableEnum
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilSchemaAOptionalNullableEnum was set.
func (o OptNilSchemaAOptionalNullableEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilSchemaAOptionalNullableEnum) Reset() {
	var v SchemaAOptionalNullableEnum
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilSchemaAOptionalNullableEnum) SetTo(v SchemaAOptionalNullableEnum) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilSchemaAOptionalNullableEnum) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o OptNilSchemaAOptionalNullableEnum) Get() (v SchemaAOptionalNullableEnum, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilSchemaAOptionalNullableEnum) Or(d SchemaAOptionalNullableEnum) SchemaAOptionalNullableEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSchemaB returns new OptSchemaB with value set to v.
func NewOptSchemaB(v SchemaB) OptSchemaB {
	return OptSchemaB{
		Value: v,
		Set:   true,
	}
}

// OptSchemaB is optional SchemaB.
type OptSchemaB struct {
	Value SchemaB
	Set   bool
}

// IsSet returns true if OptSchemaB was set.
func (o OptSchemaB) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSchemaB) Reset() {
	var v SchemaB
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSchemaB) SetTo(v SchemaB) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSchemaB) Get() (v SchemaB, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSchemaB) Or(d SchemaB) SchemaB {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SchemaA
type SchemaA struct {
	Int64                           int64                             `json:"int64"`
	StringFoobarBind                string                            `json:"string_foobar_bind"`
	StringOptionalNullable          NilString                         `json:"string_optional_nullable"`
	OptionalNullableBool            OptNilBool                        `json:"optional_nullable_bool"`
	JsontypeStrings                 []string                          `json:"jsontype_strings"`
	JsontypeStringsOptional         []string                          `json:"jsontype_strings_optional"`
	JsontypeInts                    []int                             `json:"jsontype_ints"`
	JsontypeIntsOptional            []int                             `json:"jsontype_ints_optional"`
	RequiredEnum                    SchemaARequiredEnum               `json:"required_enum"`
	OptionalNullableEnum            OptNilSchemaAOptionalNullableEnum `json:"optional_nullable_enum"`
	EdgeSchemabUniqueRequired       SchemaB                           `json:"edge_schemab_unique_required"`
	EdgeSchemabUniqueRequiredBsBind SchemaB                           `json:"edge_schemab_unique_required_bs_bind"`
	EdgeSchemabUniqueOptional       OptSchemaB                        `json:"edge_schemab_unique_optional"`
	EdgeSchemab                     []SchemaB                         `json:"edge_schemab"`
	EdgeSchemaaRecursive            []SchemaA                         `json:"edge_schemaa_recursive"`
}

// GetInt64 returns the value of Int64.
func (s *SchemaA) GetInt64() int64 {
	return s.Int64
}

// GetStringFoobarBind returns the value of StringFoobarBind.
func (s *SchemaA) GetStringFoobarBind() string {
	return s.StringFoobarBind
}

// GetStringOptionalNullable returns the value of StringOptionalNullable.
func (s *SchemaA) GetStringOptionalNullable() NilString {
	return s.StringOptionalNullable
}

// GetOptionalNullableBool returns the value of OptionalNullableBool.
func (s *SchemaA) GetOptionalNullableBool() OptNilBool {
	return s.OptionalNullableBool
}

// GetJsontypeStrings returns the value of JsontypeStrings.
func (s *SchemaA) GetJsontypeStrings() []string {
	return s.JsontypeStrings
}

// GetJsontypeStringsOptional returns the value of JsontypeStringsOptional.
func (s *SchemaA) GetJsontypeStringsOptional() []string {
	return s.JsontypeStringsOptional
}

// GetJsontypeInts returns the value of JsontypeInts.
func (s *SchemaA) GetJsontypeInts() []int {
	return s.JsontypeInts
}

// GetJsontypeIntsOptional returns the value of JsontypeIntsOptional.
func (s *SchemaA) GetJsontypeIntsOptional() []int {
	return s.JsontypeIntsOptional
}

// GetRequiredEnum returns the value of RequiredEnum.
func (s *SchemaA) GetRequiredEnum() SchemaARequiredEnum {
	return s.RequiredEnum
}

// GetOptionalNullableEnum returns the value of OptionalNullableEnum.
func (s *SchemaA) GetOptionalNullableEnum() OptNilSchemaAOptionalNullableEnum {
	return s.OptionalNullableEnum
}

// GetEdgeSchemabUniqueRequired returns the value of EdgeSchemabUniqueRequired.
func (s *SchemaA) GetEdgeSchemabUniqueRequired() SchemaB {
	return s.EdgeSchemabUniqueRequired
}

// GetEdgeSchemabUniqueRequiredBsBind returns the value of EdgeSchemabUniqueRequiredBsBind.
func (s *SchemaA) GetEdgeSchemabUniqueRequiredBsBind() SchemaB {
	return s.EdgeSchemabUniqueRequiredBsBind
}

// GetEdgeSchemabUniqueOptional returns the value of EdgeSchemabUniqueOptional.
func (s *SchemaA) GetEdgeSchemabUniqueOptional() OptSchemaB {
	return s.EdgeSchemabUniqueOptional
}

// GetEdgeSchemab returns the value of EdgeSchemab.
func (s *SchemaA) GetEdgeSchemab() []SchemaB {
	return s.EdgeSchemab
}

// GetEdgeSchemaaRecursive returns the value of EdgeSchemaaRecursive.
func (s *SchemaA) GetEdgeSchemaaRecursive() []SchemaA {
	return s.EdgeSchemaaRecursive
}

// SetInt64 sets the value of Int64.
func (s *SchemaA) SetInt64(val int64) {
	s.Int64 = val
}

// SetStringFoobarBind sets the value of StringFoobarBind.
func (s *SchemaA) SetStringFoobarBind(val string) {
	s.StringFoobarBind = val
}

// SetStringOptionalNullable sets the value of StringOptionalNullable.
func (s *SchemaA) SetStringOptionalNullable(val NilString) {
	s.StringOptionalNullable = val
}

// SetOptionalNullableBool sets the value of OptionalNullableBool.
func (s *SchemaA) SetOptionalNullableBool(val OptNilBool) {
	s.OptionalNullableBool = val
}

// SetJsontypeStrings sets the value of JsontypeStrings.
func (s *SchemaA) SetJsontypeStrings(val []string) {
	s.JsontypeStrings = val
}

// SetJsontypeStringsOptional sets the value of JsontypeStringsOptional.
func (s *SchemaA) SetJsontypeStringsOptional(val []string) {
	s.JsontypeStringsOptional = val
}

// SetJsontypeInts sets the value of JsontypeInts.
func (s *SchemaA) SetJsontypeInts(val []int) {
	s.JsontypeInts = val
}

// SetJsontypeIntsOptional sets the value of JsontypeIntsOptional.
func (s *SchemaA) SetJsontypeIntsOptional(val []int) {
	s.JsontypeIntsOptional = val
}

// SetRequiredEnum sets the value of RequiredEnum.
func (s *SchemaA) SetRequiredEnum(val SchemaARequiredEnum) {
	s.RequiredEnum = val
}

// SetOptionalNullableEnum sets the value of OptionalNullableEnum.
func (s *SchemaA) SetOptionalNullableEnum(val OptNilSchemaAOptionalNullableEnum) {
	s.OptionalNullableEnum = val
}

// SetEdgeSchemabUniqueRequired sets the value of EdgeSchemabUniqueRequired.
func (s *SchemaA) SetEdgeSchemabUniqueRequired(val SchemaB) {
	s.EdgeSchemabUniqueRequired = val
}

// SetEdgeSchemabUniqueRequiredBsBind sets the value of EdgeSchemabUniqueRequiredBsBind.
func (s *SchemaA) SetEdgeSchemabUniqueRequiredBsBind(val SchemaB) {
	s.EdgeSchemabUniqueRequiredBsBind = val
}

// SetEdgeSchemabUniqueOptional sets the value of EdgeSchemabUniqueOptional.
func (s *SchemaA) SetEdgeSchemabUniqueOptional(val OptSchemaB) {
	s.EdgeSchemabUniqueOptional = val
}

// SetEdgeSchemab sets the value of EdgeSchemab.
func (s *SchemaA) SetEdgeSchemab(val []SchemaB) {
	s.EdgeSchemab = val
}

// SetEdgeSchemaaRecursive sets the value of EdgeSchemaaRecursive.
func (s *SchemaA) SetEdgeSchemaaRecursive(val []SchemaA) {
	s.EdgeSchemaaRecursive = val
}

type SchemaAOptionalNullableEnum string

const (
	SchemaAOptionalNullableEnumC SchemaAOptionalNullableEnum = "c"
	SchemaAOptionalNullableEnumD SchemaAOptionalNullableEnum = "d"
)

// MarshalText implements encoding.TextMarshaler.
func (s SchemaAOptionalNullableEnum) MarshalText() ([]byte, error) {
	switch s {
	case SchemaAOptionalNullableEnumC:
		return []byte(s), nil
	case SchemaAOptionalNullableEnumD:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *SchemaAOptionalNullableEnum) UnmarshalText(data []byte) error {
	switch SchemaAOptionalNullableEnum(data) {
	case SchemaAOptionalNullableEnumC:
		*s = SchemaAOptionalNullableEnumC
		return nil
	case SchemaAOptionalNullableEnumD:
		*s = SchemaAOptionalNullableEnumD
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type SchemaARequiredEnum string

const (
	SchemaARequiredEnumA SchemaARequiredEnum = "a"
	SchemaARequiredEnumB SchemaARequiredEnum = "b"
)

// MarshalText implements encoding.TextMarshaler.
func (s SchemaARequiredEnum) MarshalText() ([]byte, error) {
	switch s {
	case SchemaARequiredEnumA:
		return []byte(s), nil
	case SchemaARequiredEnumB:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *SchemaARequiredEnum) UnmarshalText(data []byte) error {
	switch SchemaARequiredEnum(data) {
	case SchemaARequiredEnumA:
		*s = SchemaARequiredEnumA
		return nil
	case SchemaARequiredEnumB:
		*s = SchemaARequiredEnumB
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/SchemaB
type SchemaB struct {
	ID int64 `json:"id"`
}

// GetID returns the value of ID.
func (s *SchemaB) GetID() int64 {
	return s.ID
}

// SetID sets the value of ID.
func (s *SchemaB) SetID(val int64) {
	s.ID = val
}
