package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RMLUser holds the schema definition for the RMLUser entity.
type RMLUser struct {
	ent.Schema
}

// Fields of the RMLUser.
func (RMLUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("role").Default("user"),
		field.Int("review_count").Default(0),
	}
}

// Edges of the RMLUser.
func (RMLUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reviews", Review.Type),
	}
}
