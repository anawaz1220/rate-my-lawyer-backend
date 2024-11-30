// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/mvptruck"
)

// MvpTruck is the model entity for the MvpTruck schema.
type MvpTruck struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name                     string `json:"name,omitempty"`
	mvp_planned_route_trucks *int
	selectValues             sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MvpTruck) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mvptruck.FieldID:
			values[i] = new(sql.NullInt64)
		case mvptruck.FieldName:
			values[i] = new(sql.NullString)
		case mvptruck.ForeignKeys[0]: // mvp_planned_route_trucks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MvpTruck fields.
func (mt *MvpTruck) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mvptruck.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mt.ID = int(value.Int64)
		case mvptruck.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mt.Name = value.String
			}
		case mvptruck.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mvp_planned_route_trucks", value)
			} else if value.Valid {
				mt.mvp_planned_route_trucks = new(int)
				*mt.mvp_planned_route_trucks = int(value.Int64)
			}
		default:
			mt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MvpTruck.
// This includes values selected through modifiers, order, etc.
func (mt *MvpTruck) Value(name string) (ent.Value, error) {
	return mt.selectValues.Get(name)
}

// Update returns a builder for updating this MvpTruck.
// Note that you need to call MvpTruck.Unwrap() before calling this method if this MvpTruck
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MvpTruck) Update() *MvpTruckUpdateOne {
	return NewMvpTruckClient(mt.config).UpdateOne(mt)
}

// Unwrap unwraps the MvpTruck entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mt *MvpTruck) Unwrap() *MvpTruck {
	_tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MvpTruck is not a transactional entity")
	}
	mt.config.driver = _tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MvpTruck) String() string {
	var builder strings.Builder
	builder.WriteString("MvpTruck(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mt.ID))
	builder.WriteString("name=")
	builder.WriteString(mt.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MvpTrucks is a parsable slice of MvpTruck.
type MvpTrucks []*MvpTruck