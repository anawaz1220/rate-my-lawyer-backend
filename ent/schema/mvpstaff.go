package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MvpStaff holds the schema definition for the MvpStaff entity.
type MvpStaff struct {
	ent.Schema
}

// Fields of the MvpStaff.
func (MvpStaff) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("role").
			Validate(func(s string) error {
				if s != "DRIVER" && s != "LOADER" {
					return fmt.Errorf("invalid role: %s", s)
				}
				return nil
			}),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("last_name"),
		field.Time("birthday").Optional(),
	}
}

// Edges of the MvpStaff.
func (MvpStaff) Edges() []ent.Edge {
	return nil
}
