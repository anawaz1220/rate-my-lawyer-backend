package seeders

import (
	"context"
	"log"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/rmluser"
)

func SeedRMLUsers(ctx context.Context, client *ent.Client) error {
	defaultAdmin := &ent.RMLUser{
		Name:     "Super Admin",
		Email:    "admin@ratemylawyer.ca",
		Password: "rml@123", // Note: In a real application, you should hash the password
		Role:     "Administrator",
	}

	exists, err := client.RMLUser.Query().
		Where(rmluser.EmailEQ(defaultAdmin.Email)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if !exists {
		_, err := client.RMLUser.Create().
			SetName(defaultAdmin.Name).
			SetEmail(defaultAdmin.Email).
			SetPassword(defaultAdmin.Password).
			SetRole(defaultAdmin.Role).
			Save(ctx)
		if err != nil {
			return err
		}
		log.Printf("Created default admin user: %s", defaultAdmin.Email)
	} else {
		log.Printf("Default admin user already exists: %s", defaultAdmin.Email)
	}

	return nil
}
