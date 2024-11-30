package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lawyer holds the schema definition for the Lawyer entity.
type Lawyer struct {
	ent.Schema
}

// Fields of the Lawyer.
func (Lawyer) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").NotEmpty(),
		field.String("first_name"),
		field.String("middle_name"),
		field.String("last_name"),
		field.String("gender"),
		field.String("address"),
		field.String("phone"),
		field.String("practicing_status"),
		field.String("profile_url"),
		field.Float("average_rating").Default(0),
		field.Int("review_count").Default(0),
	}
}

// Edges of the Lawyer.
func (Lawyer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("jurisdictions", Jurisdiction.Type).Through("lawyer_jurisdictions", LawyerJurisdiction.Type),
		edge.To("decisions", Decision.Type),
		edge.To("reviews", Review.Type),
	}
}
