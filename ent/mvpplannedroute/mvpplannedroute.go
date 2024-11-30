// Code generated by ent, DO NOT EDIT.

package mvpplannedroute

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mvpplannedroute type in the database.
	Label = "mvp_planned_route"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldRouteName holds the string denoting the route_name field in the database.
	FieldRouteName = "route_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeTrucks holds the string denoting the trucks edge name in mutations.
	EdgeTrucks = "trucks"
	// EdgeDriver holds the string denoting the driver edge name in mutations.
	EdgeDriver = "driver"
	// EdgeLoaders holds the string denoting the loaders edge name in mutations.
	EdgeLoaders = "loaders"
	// EdgeMaterials holds the string denoting the materials edge name in mutations.
	EdgeMaterials = "materials"
	// Table holds the table name of the mvpplannedroute in the database.
	Table = "mvp_planned_routes"
	// TrucksTable is the table that holds the trucks relation/edge.
	TrucksTable = "mvp_trucks"
	// TrucksInverseTable is the table name for the MvpTruck entity.
	// It exists in this package in order to avoid circular dependency with the "mvptruck" package.
	TrucksInverseTable = "mvp_trucks"
	// TrucksColumn is the table column denoting the trucks relation/edge.
	TrucksColumn = "mvp_planned_route_trucks"
	// DriverTable is the table that holds the driver relation/edge.
	DriverTable = "mvp_planned_routes"
	// DriverInverseTable is the table name for the MvpStaff entity.
	// It exists in this package in order to avoid circular dependency with the "mvpstaff" package.
	DriverInverseTable = "mvp_staffs"
	// DriverColumn is the table column denoting the driver relation/edge.
	DriverColumn = "mvp_planned_route_driver"
	// LoadersTable is the table that holds the loaders relation/edge.
	LoadersTable = "mvp_staffs"
	// LoadersInverseTable is the table name for the MvpStaff entity.
	// It exists in this package in order to avoid circular dependency with the "mvpstaff" package.
	LoadersInverseTable = "mvp_staffs"
	// LoadersColumn is the table column denoting the loaders relation/edge.
	LoadersColumn = "mvp_planned_route_loaders"
	// MaterialsTable is the table that holds the materials relation/edge.
	MaterialsTable = "mvp_materials"
	// MaterialsInverseTable is the table name for the MvpMaterial entity.
	// It exists in this package in order to avoid circular dependency with the "mvpmaterial" package.
	MaterialsInverseTable = "mvp_materials"
	// MaterialsColumn is the table column denoting the materials relation/edge.
	MaterialsColumn = "mvp_planned_route_materials"
)

// Columns holds all SQL columns for mvpplannedroute fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldRouteName,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mvp_planned_routes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"mvp_planned_route_driver",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
)

// OrderOption defines the ordering options for the MvpPlannedRoute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByRouteName orders the results by the route_name field.
func ByRouteName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRouteName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTrucksCount orders the results by trucks count.
func ByTrucksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTrucksStep(), opts...)
	}
}

// ByTrucks orders the results by trucks terms.
func ByTrucks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTrucksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDriverField orders the results by driver field.
func ByDriverField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDriverStep(), sql.OrderByField(field, opts...))
	}
}

// ByLoadersCount orders the results by loaders count.
func ByLoadersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLoadersStep(), opts...)
	}
}

// ByLoaders orders the results by loaders terms.
func ByLoaders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLoadersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMaterialsCount orders the results by materials count.
func ByMaterialsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMaterialsStep(), opts...)
	}
}

// ByMaterials orders the results by materials terms.
func ByMaterials(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMaterialsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTrucksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TrucksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TrucksTable, TrucksColumn),
	)
}
func newDriverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DriverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DriverTable, DriverColumn),
	)
}
func newLoadersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LoadersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LoadersTable, LoadersColumn),
	)
}
func newMaterialsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MaterialsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MaterialsTable, MaterialsColumn),
	)
}