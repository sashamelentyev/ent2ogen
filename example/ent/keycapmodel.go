// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
)

// KeycapModel is the model entity for the KeycapModel schema.
type KeycapModel struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Profile holds the value of the "profile" field.
	Profile string `json:"profile,omitempty"`
	// Material holds the value of the "material" field.
	Material     keycapmodel.Material `json:"material,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KeycapModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keycapmodel.FieldID:
			values[i] = new(sql.NullInt64)
		case keycapmodel.FieldName, keycapmodel.FieldProfile, keycapmodel.FieldMaterial:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KeycapModel fields.
func (km *KeycapModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keycapmodel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			km.ID = int64(value.Int64)
		case keycapmodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				km.Name = value.String
			}
		case keycapmodel.FieldProfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile", values[i])
			} else if value.Valid {
				km.Profile = value.String
			}
		case keycapmodel.FieldMaterial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field material", values[i])
			} else if value.Valid {
				km.Material = keycapmodel.Material(value.String)
			}
		default:
			km.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the KeycapModel.
// This includes values selected through modifiers, order, etc.
func (km *KeycapModel) Value(name string) (ent.Value, error) {
	return km.selectValues.Get(name)
}

// Update returns a builder for updating this KeycapModel.
// Note that you need to call KeycapModel.Unwrap() before calling this method if this KeycapModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (km *KeycapModel) Update() *KeycapModelUpdateOne {
	return NewKeycapModelClient(km.config).UpdateOne(km)
}

// Unwrap unwraps the KeycapModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (km *KeycapModel) Unwrap() *KeycapModel {
	_tx, ok := km.config.driver.(*txDriver)
	if !ok {
		panic("ent: KeycapModel is not a transactional entity")
	}
	km.config.driver = _tx.drv
	return km
}

// String implements the fmt.Stringer.
func (km *KeycapModel) String() string {
	var builder strings.Builder
	builder.WriteString("KeycapModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", km.ID))
	builder.WriteString("name=")
	builder.WriteString(km.Name)
	builder.WriteString(", ")
	builder.WriteString("profile=")
	builder.WriteString(km.Profile)
	builder.WriteString(", ")
	builder.WriteString("material=")
	builder.WriteString(fmt.Sprintf("%v", km.Material))
	builder.WriteByte(')')
	return builder.String()
}

// KeycapModels is a parsable slice of KeycapModel.
type KeycapModels []*KeycapModel
