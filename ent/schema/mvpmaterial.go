package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MvpMaterial holds the schema definition for the MvpMaterial entity.
type MvpMaterial struct {
	ent.Schema
}

// Fields of the MvpMaterial.
func (MvpMaterial) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the MvpMaterial.
func (MvpMaterial) Edges() []ent.Edge {
	return nil
}
