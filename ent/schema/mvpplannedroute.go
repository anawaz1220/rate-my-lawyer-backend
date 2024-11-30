package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MvpPlannedRoute holds the schema definition for the MvpPlannedRoute entity.
type MvpPlannedRoute struct {
	ent.Schema
}

// Fields of the MvpPlannedRoute.
func (MvpPlannedRoute) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.String("route_name"),
		field.String("status").
			Default("planned").
			Validate(func(s string) error {
				if s != "planned" && s != "edited" {
					return fmt.Errorf("invalid status: %s", s)
				}
				return nil
			}),
	}
}

// Edges of the MvpPlannedRoute.
func (MvpPlannedRoute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("trucks", MvpTruck.Type),
		edge.To("driver", MvpStaff.Type).
			Unique(),
		edge.To("loaders", MvpStaff.Type),
		edge.To("materials", MvpMaterial.Type),
	}
}
