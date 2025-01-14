// Code generated by ent, DO NOT EDIT.

package mvpstaff

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the mvpstaff type in the database.
	Label = "mvp_staff"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// Table holds the table name of the mvpstaff in the database.
	Table = "mvp_staffs"
)

// Columns holds all SQL columns for mvpstaff fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRole,
	FieldEmail,
	FieldPhone,
	FieldLastName,
	FieldBirthday,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mvp_staffs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"mvp_planned_route_loaders",
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
	// RoleValidator is a validator for the "role" field. It is called by the builders before save.
	RoleValidator func(string) error
)

// OrderOption defines the ordering options for the MvpStaff queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}
