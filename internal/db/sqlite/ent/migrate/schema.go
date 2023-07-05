// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "data", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_id", Type: field.TypeUUID, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_institutions_accounts",
				Columns:    []*schema.Column{AccountsColumns[4]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "data", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_id", Type: field.TypeUUID, Nullable: true},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_institutions_assets",
				Columns:    []*schema.Column{AssetsColumns[4]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConnectionsColumns holds the columns for the "connections" table.
	ConnectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "data", Type: field.TypeString},
		{Name: "env", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// ConnectionsTable holds the schema information for the "connections" table.
	ConnectionsTable = &schema.Table{
		Name:       "connections",
		Columns:    ConnectionsColumns,
		PrimaryKey: []*schema.Column{ConnectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "connections_institutions_connection",
				Columns:    []*schema.Column{ConnectionsColumns[5]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:       "institutions",
		Columns:    InstitutionsColumns,
		PrimaryKey: []*schema.Column{InstitutionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AssetsTable,
		ConnectionsTable,
		InstitutionsTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = InstitutionsTable
	AssetsTable.ForeignKeys[0].RefTable = InstitutionsTable
	ConnectionsTable.ForeignKeys[0].RefTable = InstitutionsTable
}