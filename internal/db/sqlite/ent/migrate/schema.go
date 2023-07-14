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
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_accounts", Type: field.TypeUUID, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_institutions_accounts",
				Columns:    []*schema.Column{AccountsColumns[5]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConnectionsColumns holds the columns for the "connections" table.
	ConnectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ConnectionsTable holds the schema information for the "connections" table.
	ConnectionsTable = &schema.Table{
		Name:       "connections",
		Columns:    ConnectionsColumns,
		PrimaryKey: []*schema.Column{ConnectionsColumns[0]},
	}
	// IncomesColumns holds the columns for the "incomes" table.
	IncomesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_incomes", Type: field.TypeUUID, Nullable: true},
	}
	// IncomesTable holds the schema information for the "incomes" table.
	IncomesTable = &schema.Table{
		Name:       "incomes",
		Columns:    IncomesColumns,
		PrimaryKey: []*schema.Column{IncomesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "incomes_institutions_incomes",
				Columns:    []*schema.Column{IncomesColumns[5]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "institution_connection", Type: field.TypeUUID, Nullable: true},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:       "institutions",
		Columns:    InstitutionsColumns,
		PrimaryKey: []*schema.Column{InstitutionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "institutions_connections_connection",
				Columns:    []*schema.Column{InstitutionsColumns[5]},
				RefColumns: []*schema.Column{ConnectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ManualItemsColumns holds the columns for the "manual_items" table.
	ManualItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "item_table_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ManualItemsTable holds the schema information for the "manual_items" table.
	ManualItemsTable = &schema.Table{
		Name:       "manual_items",
		Columns:    ManualItemsColumns,
		PrimaryKey: []*schema.Column{ManualItemsColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		ConnectionsTable,
		IncomesTable,
		InstitutionsTable,
		ManualItemsTable,
		TransactionsTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = InstitutionsTable
	IncomesTable.ForeignKeys[0].RefTable = InstitutionsTable
	InstitutionsTable.ForeignKeys[0].RefTable = ConnectionsTable
}
