// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/lawyerjurisdiction"
)

// LawyerJurisdiction is the model entity for the LawyerJurisdiction schema.
type LawyerJurisdiction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LawyerID holds the value of the "lawyer_id" field.
	LawyerID int `json:"lawyer_id,omitempty"`
	// JurisdictionID holds the value of the "jurisdiction_id" field.
	JurisdictionID int `json:"jurisdiction_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LawyerJurisdictionQuery when eager-loading is set.
	Edges        LawyerJurisdictionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LawyerJurisdictionEdges holds the relations/edges for other nodes in the graph.
type LawyerJurisdictionEdges struct {
	// Lawyer holds the value of the lawyer edge.
	Lawyer *Lawyer `json:"lawyer,omitempty"`
	// Jurisdiction holds the value of the jurisdiction edge.
	Jurisdiction *Jurisdiction `json:"jurisdiction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LawyerOrErr returns the Lawyer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LawyerJurisdictionEdges) LawyerOrErr() (*Lawyer, error) {
	if e.Lawyer != nil {
		return e.Lawyer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: lawyer.Label}
	}
	return nil, &NotLoadedError{edge: "lawyer"}
}

// JurisdictionOrErr returns the Jurisdiction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LawyerJurisdictionEdges) JurisdictionOrErr() (*Jurisdiction, error) {
	if e.Jurisdiction != nil {
		return e.Jurisdiction, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: jurisdiction.Label}
	}
	return nil, &NotLoadedError{edge: "jurisdiction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LawyerJurisdiction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lawyerjurisdiction.FieldID, lawyerjurisdiction.FieldLawyerID, lawyerjurisdiction.FieldJurisdictionID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LawyerJurisdiction fields.
func (lj *LawyerJurisdiction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lawyerjurisdiction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lj.ID = int(value.Int64)
		case lawyerjurisdiction.FieldLawyerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lawyer_id", values[i])
			} else if value.Valid {
				lj.LawyerID = int(value.Int64)
			}
		case lawyerjurisdiction.FieldJurisdictionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field jurisdiction_id", values[i])
			} else if value.Valid {
				lj.JurisdictionID = int(value.Int64)
			}
		default:
			lj.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LawyerJurisdiction.
// This includes values selected through modifiers, order, etc.
func (lj *LawyerJurisdiction) Value(name string) (ent.Value, error) {
	return lj.selectValues.Get(name)
}

// QueryLawyer queries the "lawyer" edge of the LawyerJurisdiction entity.
func (lj *LawyerJurisdiction) QueryLawyer() *LawyerQuery {
	return NewLawyerJurisdictionClient(lj.config).QueryLawyer(lj)
}

// QueryJurisdiction queries the "jurisdiction" edge of the LawyerJurisdiction entity.
func (lj *LawyerJurisdiction) QueryJurisdiction() *JurisdictionQuery {
	return NewLawyerJurisdictionClient(lj.config).QueryJurisdiction(lj)
}

// Update returns a builder for updating this LawyerJurisdiction.
// Note that you need to call LawyerJurisdiction.Unwrap() before calling this method if this LawyerJurisdiction
// was returned from a transaction, and the transaction was committed or rolled back.
func (lj *LawyerJurisdiction) Update() *LawyerJurisdictionUpdateOne {
	return NewLawyerJurisdictionClient(lj.config).UpdateOne(lj)
}

// Unwrap unwraps the LawyerJurisdiction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lj *LawyerJurisdiction) Unwrap() *LawyerJurisdiction {
	_tx, ok := lj.config.driver.(*txDriver)
	if !ok {
		panic("ent: LawyerJurisdiction is not a transactional entity")
	}
	lj.config.driver = _tx.drv
	return lj
}

// String implements the fmt.Stringer.
func (lj *LawyerJurisdiction) String() string {
	var builder strings.Builder
	builder.WriteString("LawyerJurisdiction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lj.ID))
	builder.WriteString("lawyer_id=")
	builder.WriteString(fmt.Sprintf("%v", lj.LawyerID))
	builder.WriteString(", ")
	builder.WriteString("jurisdiction_id=")
	builder.WriteString(fmt.Sprintf("%v", lj.JurisdictionID))
	builder.WriteByte(')')
	return builder.String()
}

// LawyerJurisdictions is a parsable slice of LawyerJurisdiction.
type LawyerJurisdictions []*LawyerJurisdiction