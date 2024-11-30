package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MvpRoute holds the schema definition for the MvpRoute entity.
type MvpRoute struct {
	ent.Schema
}

// Fields of the MvpRoute.
func (MvpRoute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("day_of_week"),
	}
}
