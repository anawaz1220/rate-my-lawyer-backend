package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LawyerJurisdiction holds the schema definition for the LawyerJurisdiction entity.
type LawyerJurisdiction struct {
	ent.Schema
}

// Fields of the LawyerJurisdiction.
func (LawyerJurisdiction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("lawyer_id"),
		field.Int("jurisdiction_id"),
	}
}

// Edges of the LawyerJurisdiction.
func (LawyerJurisdiction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lawyer", Lawyer.Type).Unique().Required().Field("lawyer_id"),
		edge.To("jurisdiction", Jurisdiction.Type).Unique().Required().Field("jurisdiction_id"),
	}
}
