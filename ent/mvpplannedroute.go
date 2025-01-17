// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/mvpplannedroute"
	"github.com/mikestefanello/pagoda/ent/mvpstaff"
)

// MvpPlannedRoute is the model entity for the MvpPlannedRoute schema.
type MvpPlannedRoute struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// RouteName holds the value of the "route_name" field.
	RouteName string `json:"route_name,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MvpPlannedRouteQuery when eager-loading is set.
	Edges                    MvpPlannedRouteEdges `json:"edges"`
	mvp_planned_route_driver *int
	selectValues             sql.SelectValues
}

// MvpPlannedRouteEdges holds the relations/edges for other nodes in the graph.
type MvpPlannedRouteEdges struct {
	// Trucks holds the value of the trucks edge.
	Trucks []*MvpTruck `json:"trucks,omitempty"`
	// Driver holds the value of the driver edge.
	Driver *MvpStaff `json:"driver,omitempty"`
	// Loaders holds the value of the loaders edge.
	Loaders []*MvpStaff `json:"loaders,omitempty"`
	// Materials holds the value of the materials edge.
	Materials []*MvpMaterial `json:"materials,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TrucksOrErr returns the Trucks value or an error if the edge
// was not loaded in eager-loading.
func (e MvpPlannedRouteEdges) TrucksOrErr() ([]*MvpTruck, error) {
	if e.loadedTypes[0] {
		return e.Trucks, nil
	}
	return nil, &NotLoadedError{edge: "trucks"}
}

// DriverOrErr returns the Driver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MvpPlannedRouteEdges) DriverOrErr() (*MvpStaff, error) {
	if e.Driver != nil {
		return e.Driver, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: mvpstaff.Label}
	}
	return nil, &NotLoadedError{edge: "driver"}
}

// LoadersOrErr returns the Loaders value or an error if the edge
// was not loaded in eager-loading.
func (e MvpPlannedRouteEdges) LoadersOrErr() ([]*MvpStaff, error) {
	if e.loadedTypes[2] {
		return e.Loaders, nil
	}
	return nil, &NotLoadedError{edge: "loaders"}
}

// MaterialsOrErr returns the Materials value or an error if the edge
// was not loaded in eager-loading.
func (e MvpPlannedRouteEdges) MaterialsOrErr() ([]*MvpMaterial, error) {
	if e.loadedTypes[3] {
		return e.Materials, nil
	}
	return nil, &NotLoadedError{edge: "materials"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MvpPlannedRoute) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mvpplannedroute.FieldID:
			values[i] = new(sql.NullInt64)
		case mvpplannedroute.FieldRouteName, mvpplannedroute.FieldStatus:
			values[i] = new(sql.NullString)
		case mvpplannedroute.FieldDate:
			values[i] = new(sql.NullTime)
		case mvpplannedroute.ForeignKeys[0]: // mvp_planned_route_driver
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MvpPlannedRoute fields.
func (mpr *MvpPlannedRoute) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mvpplannedroute.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mpr.ID = int(value.Int64)
		case mvpplannedroute.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				mpr.Date = value.Time
			}
		case mvpplannedroute.FieldRouteName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field route_name", values[i])
			} else if value.Valid {
				mpr.RouteName = value.String
			}
		case mvpplannedroute.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mpr.Status = value.String
			}
		case mvpplannedroute.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mvp_planned_route_driver", value)
			} else if value.Valid {
				mpr.mvp_planned_route_driver = new(int)
				*mpr.mvp_planned_route_driver = int(value.Int64)
			}
		default:
			mpr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MvpPlannedRoute.
// This includes values selected through modifiers, order, etc.
func (mpr *MvpPlannedRoute) Value(name string) (ent.Value, error) {
	return mpr.selectValues.Get(name)
}

// QueryTrucks queries the "trucks" edge of the MvpPlannedRoute entity.
func (mpr *MvpPlannedRoute) QueryTrucks() *MvpTruckQuery {
	return NewMvpPlannedRouteClient(mpr.config).QueryTrucks(mpr)
}

// QueryDriver queries the "driver" edge of the MvpPlannedRoute entity.
func (mpr *MvpPlannedRoute) QueryDriver() *MvpStaffQuery {
	return NewMvpPlannedRouteClient(mpr.config).QueryDriver(mpr)
}

// QueryLoaders queries the "loaders" edge of the MvpPlannedRoute entity.
func (mpr *MvpPlannedRoute) QueryLoaders() *MvpStaffQuery {
	return NewMvpPlannedRouteClient(mpr.config).QueryLoaders(mpr)
}

// QueryMaterials queries the "materials" edge of the MvpPlannedRoute entity.
func (mpr *MvpPlannedRoute) QueryMaterials() *MvpMaterialQuery {
	return NewMvpPlannedRouteClient(mpr.config).QueryMaterials(mpr)
}

// Update returns a builder for updating this MvpPlannedRoute.
// Note that you need to call MvpPlannedRoute.Unwrap() before calling this method if this MvpPlannedRoute
// was returned from a transaction, and the transaction was committed or rolled back.
func (mpr *MvpPlannedRoute) Update() *MvpPlannedRouteUpdateOne {
	return NewMvpPlannedRouteClient(mpr.config).UpdateOne(mpr)
}

// Unwrap unwraps the MvpPlannedRoute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mpr *MvpPlannedRoute) Unwrap() *MvpPlannedRoute {
	_tx, ok := mpr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MvpPlannedRoute is not a transactional entity")
	}
	mpr.config.driver = _tx.drv
	return mpr
}

// String implements the fmt.Stringer.
func (mpr *MvpPlannedRoute) String() string {
	var builder strings.Builder
	builder.WriteString("MvpPlannedRoute(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mpr.ID))
	builder.WriteString("date=")
	builder.WriteString(mpr.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("route_name=")
	builder.WriteString(mpr.RouteName)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(mpr.Status)
	builder.WriteByte(')')
	return builder.String()
}

// MvpPlannedRoutes is a parsable slice of MvpPlannedRoute.
type MvpPlannedRoutes []*MvpPlannedRoute
