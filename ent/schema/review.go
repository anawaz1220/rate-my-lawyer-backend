package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int("rating").Min(1).Max(5),
		field.String("comment"),
		field.Time("created_at").Default(time.Now),
		field.Int("lawyer_reviews").Optional().Nillable(),
		field.Int("rml_user_reviews").Optional().Nillable(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lawyer", Lawyer.Type).Ref("reviews").Unique().Field("lawyer_reviews"),
		edge.From("user", RMLUser.Type).Ref("reviews").Unique().Field("rml_user_reviews"),
	}
}
