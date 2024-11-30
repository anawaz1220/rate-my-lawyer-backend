package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Jurisdiction holds the schema definition for the Jurisdiction entity.
type Jurisdiction struct {
	ent.Schema
}

// Fields of the Jurisdiction.
func (Jurisdiction) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Jurisdiction.
func (Jurisdiction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lawyers", Lawyer.Type).Ref("jurisdictions").Through("lawyer_jurisdictions", LawyerJurisdiction.Type),
	}
}
