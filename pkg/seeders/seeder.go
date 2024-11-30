package seeders

import (
	"context"

	"github.com/mikestefanello/pagoda/ent"
)

func SeedAll(client *ent.Client) error {
	ctx := context.Background()

	// Seed RML Users
	if err := SeedRMLUsers(ctx, client); err != nil {
		return err
	}

	// Add any other remaining seeders here

	return nil
}
