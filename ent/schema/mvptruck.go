package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MvpTruck holds the schema definition for the MvpTruck entity.
type MvpTruck struct {
	ent.Schema
}

// Fields of the MvpTruck.
func (MvpTruck) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the MvpTruck.
func (MvpTruck) Edges() []ent.Edge {
	return nil
}
