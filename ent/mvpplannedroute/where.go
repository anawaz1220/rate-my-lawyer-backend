// Code generated by ent, DO NOT EDIT.

package mvpplannedroute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLTE(FieldID, id))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldDate, v))
}

// RouteName applies equality check predicate on the "route_name" field. It's identical to RouteNameEQ.
func RouteName(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldRouteName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldStatus, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLTE(FieldDate, v))
}

// RouteNameEQ applies the EQ predicate on the "route_name" field.
func RouteNameEQ(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldRouteName, v))
}

// RouteNameNEQ applies the NEQ predicate on the "route_name" field.
func RouteNameNEQ(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNEQ(FieldRouteName, v))
}

// RouteNameIn applies the In predicate on the "route_name" field.
func RouteNameIn(vs ...string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldIn(FieldRouteName, vs...))
}

// RouteNameNotIn applies the NotIn predicate on the "route_name" field.
func RouteNameNotIn(vs ...string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNotIn(FieldRouteName, vs...))
}

// RouteNameGT applies the GT predicate on the "route_name" field.
func RouteNameGT(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGT(FieldRouteName, v))
}

// RouteNameGTE applies the GTE predicate on the "route_name" field.
func RouteNameGTE(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGTE(FieldRouteName, v))
}

// RouteNameLT applies the LT predicate on the "route_name" field.
func RouteNameLT(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLT(FieldRouteName, v))
}

// RouteNameLTE applies the LTE predicate on the "route_name" field.
func RouteNameLTE(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLTE(FieldRouteName, v))
}

// RouteNameContains applies the Contains predicate on the "route_name" field.
func RouteNameContains(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldContains(FieldRouteName, v))
}

// RouteNameHasPrefix applies the HasPrefix predicate on the "route_name" field.
func RouteNameHasPrefix(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldHasPrefix(FieldRouteName, v))
}

// RouteNameHasSuffix applies the HasSuffix predicate on the "route_name" field.
func RouteNameHasSuffix(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldHasSuffix(FieldRouteName, v))
}

// RouteNameEqualFold applies the EqualFold predicate on the "route_name" field.
func RouteNameEqualFold(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEqualFold(FieldRouteName, v))
}

// RouteNameContainsFold applies the ContainsFold predicate on the "route_name" field.
func RouteNameContainsFold(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldContainsFold(FieldRouteName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.FieldContainsFold(FieldStatus, v))
}

// HasTrucks applies the HasEdge predicate on the "trucks" edge.
func HasTrucks() predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TrucksTable, TrucksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTrucksWith applies the HasEdge predicate on the "trucks" edge with a given conditions (other predicates).
func HasTrucksWith(preds ...predicate.MvpTruck) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := newTrucksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDriver applies the HasEdge predicate on the "driver" edge.
func HasDriver() predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DriverTable, DriverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDriverWith applies the HasEdge predicate on the "driver" edge with a given conditions (other predicates).
func HasDriverWith(preds ...predicate.MvpStaff) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := newDriverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoaders applies the HasEdge predicate on the "loaders" edge.
func HasLoaders() predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LoadersTable, LoadersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoadersWith applies the HasEdge predicate on the "loaders" edge with a given conditions (other predicates).
func HasLoadersWith(preds ...predicate.MvpStaff) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := newLoadersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMaterials applies the HasEdge predicate on the "materials" edge.
func HasMaterials() predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MaterialsTable, MaterialsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMaterialsWith applies the HasEdge predicate on the "materials" edge with a given conditions (other predicates).
func HasMaterialsWith(preds ...predicate.MvpMaterial) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(func(s *sql.Selector) {
		step := newMaterialsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MvpPlannedRoute) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MvpPlannedRoute) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MvpPlannedRoute) predicate.MvpPlannedRoute {
	return predicate.MvpPlannedRoute(sql.NotPredicates(p))
}
