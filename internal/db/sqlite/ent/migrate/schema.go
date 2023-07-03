// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProvidersTable,
	}
)

func init() {
}
