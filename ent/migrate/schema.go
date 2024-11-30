// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DecisionsColumns holds the columns for the "decisions" table.
	DecisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "lawyer_decisions", Type: field.TypeInt, Nullable: true},
	}
	// DecisionsTable holds the schema information for the "decisions" table.
	DecisionsTable = &schema.Table{
		Name:       "decisions",
		Columns:    DecisionsColumns,
		PrimaryKey: []*schema.Column{DecisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "decisions_lawyers_decisions",
				Columns:    []*schema.Column{DecisionsColumns[2]},
				RefColumns: []*schema.Column{LawyersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JurisdictionsColumns holds the columns for the "jurisdictions" table.
	JurisdictionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// JurisdictionsTable holds the schema information for the "jurisdictions" table.
	JurisdictionsTable = &schema.Table{
		Name:       "jurisdictions",
		Columns:    JurisdictionsColumns,
		PrimaryKey: []*schema.Column{JurisdictionsColumns[0]},
	}
	// LawyersColumns holds the columns for the "lawyers" table.
	LawyersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "full_name", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "gender", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "practicing_status", Type: field.TypeString},
		{Name: "profile_url", Type: field.TypeString},
		{Name: "average_rating", Type: field.TypeFloat64, Default: 0},
		{Name: "review_count", Type: field.TypeInt, Default: 0},
	}
	// LawyersTable holds the schema information for the "lawyers" table.
	LawyersTable = &schema.Table{
		Name:       "lawyers",
		Columns:    LawyersColumns,
		PrimaryKey: []*schema.Column{LawyersColumns[0]},
	}
	// LawyerJurisdictionsColumns holds the columns for the "lawyer_jurisdictions" table.
	LawyerJurisdictionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "lawyer_id", Type: field.TypeInt},
		{Name: "jurisdiction_id", Type: field.TypeInt},
	}
	// LawyerJurisdictionsTable holds the schema information for the "lawyer_jurisdictions" table.
	LawyerJurisdictionsTable = &schema.Table{
		Name:       "lawyer_jurisdictions",
		Columns:    LawyerJurisdictionsColumns,
		PrimaryKey: []*schema.Column{LawyerJurisdictionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lawyer_jurisdictions_lawyers_lawyer",
				Columns:    []*schema.Column{LawyerJurisdictionsColumns[1]},
				RefColumns: []*schema.Column{LawyersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "lawyer_jurisdictions_jurisdictions_jurisdiction",
				Columns:    []*schema.Column{LawyerJurisdictionsColumns[2]},
				RefColumns: []*schema.Column{JurisdictionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "lawyerjurisdiction_lawyer_id_jurisdiction_id",
				Unique:  true,
				Columns: []*schema.Column{LawyerJurisdictionsColumns[1], LawyerJurisdictionsColumns[2]},
			},
		},
	}
	// MvpMaterialsColumns holds the columns for the "mvp_materials" table.
	MvpMaterialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "mvp_planned_route_materials", Type: field.TypeInt, Nullable: true},
	}
	// MvpMaterialsTable holds the schema information for the "mvp_materials" table.
	MvpMaterialsTable = &schema.Table{
		Name:       "mvp_materials",
		Columns:    MvpMaterialsColumns,
		PrimaryKey: []*schema.Column{MvpMaterialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mvp_materials_mvp_planned_routes_materials",
				Columns:    []*schema.Column{MvpMaterialsColumns[2]},
				RefColumns: []*schema.Column{MvpPlannedRoutesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MvpPlannedRoutesColumns holds the columns for the "mvp_planned_routes" table.
	MvpPlannedRoutesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "route_name", Type: field.TypeString},
		{Name: "status", Type: field.TypeString, Default: "planned"},
		{Name: "mvp_planned_route_driver", Type: field.TypeInt, Nullable: true},
	}
	// MvpPlannedRoutesTable holds the schema information for the "mvp_planned_routes" table.
	MvpPlannedRoutesTable = &schema.Table{
		Name:       "mvp_planned_routes",
		Columns:    MvpPlannedRoutesColumns,
		PrimaryKey: []*schema.Column{MvpPlannedRoutesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mvp_planned_routes_mvp_staffs_driver",
				Columns:    []*schema.Column{MvpPlannedRoutesColumns[4]},
				RefColumns: []*schema.Column{MvpStaffsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MvpRoutesColumns holds the columns for the "mvp_routes" table.
	MvpRoutesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "day_of_week", Type: field.TypeString},
	}
	// MvpRoutesTable holds the schema information for the "mvp_routes" table.
	MvpRoutesTable = &schema.Table{
		Name:       "mvp_routes",
		Columns:    MvpRoutesColumns,
		PrimaryKey: []*schema.Column{MvpRoutesColumns[0]},
	}
	// MvpStaffsColumns holds the columns for the "mvp_staffs" table.
	MvpStaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime, Nullable: true},
		{Name: "mvp_planned_route_loaders", Type: field.TypeInt, Nullable: true},
	}
	// MvpStaffsTable holds the schema information for the "mvp_staffs" table.
	MvpStaffsTable = &schema.Table{
		Name:       "mvp_staffs",
		Columns:    MvpStaffsColumns,
		PrimaryKey: []*schema.Column{MvpStaffsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mvp_staffs_mvp_planned_routes_loaders",
				Columns:    []*schema.Column{MvpStaffsColumns[7]},
				RefColumns: []*schema.Column{MvpPlannedRoutesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MvpTrucksColumns holds the columns for the "mvp_trucks" table.
	MvpTrucksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "mvp_planned_route_trucks", Type: field.TypeInt, Nullable: true},
	}
	// MvpTrucksTable holds the schema information for the "mvp_trucks" table.
	MvpTrucksTable = &schema.Table{
		Name:       "mvp_trucks",
		Columns:    MvpTrucksColumns,
		PrimaryKey: []*schema.Column{MvpTrucksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mvp_trucks_mvp_planned_routes_trucks",
				Columns:    []*schema.Column{MvpTrucksColumns[2]},
				RefColumns: []*schema.Column{MvpPlannedRoutesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PasswordTokensColumns holds the columns for the "password_tokens" table.
	PasswordTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "password_token_user", Type: field.TypeInt},
	}
	// PasswordTokensTable holds the schema information for the "password_tokens" table.
	PasswordTokensTable = &schema.Table{
		Name:       "password_tokens",
		Columns:    PasswordTokensColumns,
		PrimaryKey: []*schema.Column{PasswordTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_tokens_users_user",
				Columns:    []*schema.Column{PasswordTokensColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RmlUsersColumns holds the columns for the "rml_users" table.
	RmlUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "review_count", Type: field.TypeInt, Default: 0},
	}
	// RmlUsersTable holds the schema information for the "rml_users" table.
	RmlUsersTable = &schema.Table{
		Name:       "rml_users",
		Columns:    RmlUsersColumns,
		PrimaryKey: []*schema.Column{RmlUsersColumns[0]},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rating", Type: field.TypeInt},
		{Name: "comment", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "lawyer_reviews", Type: field.TypeInt, Nullable: true},
		{Name: "rml_user_reviews", Type: field.TypeInt, Nullable: true},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_lawyers_reviews",
				Columns:    []*schema.Column{ReviewsColumns[4]},
				RefColumns: []*schema.Column{LawyersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reviews_rml_users_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{RmlUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DecisionsTable,
		JurisdictionsTable,
		LawyersTable,
		LawyerJurisdictionsTable,
		MvpMaterialsTable,
		MvpPlannedRoutesTable,
		MvpRoutesTable,
		MvpStaffsTable,
		MvpTrucksTable,
		PasswordTokensTable,
		RmlUsersTable,
		ReviewsTable,
		UsersTable,
	}
)

func init() {
	DecisionsTable.ForeignKeys[0].RefTable = LawyersTable
	LawyerJurisdictionsTable.ForeignKeys[0].RefTable = LawyersTable
	LawyerJurisdictionsTable.ForeignKeys[1].RefTable = JurisdictionsTable
	MvpMaterialsTable.ForeignKeys[0].RefTable = MvpPlannedRoutesTable
	MvpPlannedRoutesTable.ForeignKeys[0].RefTable = MvpStaffsTable
	MvpStaffsTable.ForeignKeys[0].RefTable = MvpPlannedRoutesTable
	MvpTrucksTable.ForeignKeys[0].RefTable = MvpPlannedRoutesTable
	PasswordTokensTable.ForeignKeys[0].RefTable = UsersTable
	ReviewsTable.ForeignKeys[0].RefTable = LawyersTable
	ReviewsTable.ForeignKeys[1].RefTable = RmlUsersTable
}
