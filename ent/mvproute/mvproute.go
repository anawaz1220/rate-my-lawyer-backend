// Code generated by ent, DO NOT EDIT.

package mvproute

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the mvproute type in the database.
	Label = "mvp_route"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDayOfWeek holds the string denoting the day_of_week field in the database.
	FieldDayOfWeek = "day_of_week"
	// Table holds the table name of the mvproute in the database.
	Table = "mvp_routes"
)

// Columns holds all SQL columns for mvproute fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDayOfWeek,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the MvpRoute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDayOfWeek orders the results by the day_of_week field.
func ByDayOfWeek(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDayOfWeek, opts...).ToFunc()
}
