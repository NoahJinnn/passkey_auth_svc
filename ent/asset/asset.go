// Code generated by ent, DO NOT EDIT.

package asset

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
)

const (
	// Label holds the string label denoting the asset type in the database.
	Label = "asset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSheet holds the string denoting the sheet field in the database.
	FieldSheet = "sheet"
	// FieldSection holds the string denoting the section field in the database.
	FieldSection = "section"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldProviderName holds the string denoting the provider_name field in the database.
	FieldProviderName = "provider_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the asset in the database.
	Table = "assets"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "assets"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for asset fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSheet,
	FieldSection,
	FieldType,
	FieldProviderName,
	FieldDescription,
	FieldCurrency,
	FieldValue,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSheet holds the default value on creation for the "sheet" field.
	DefaultSheet int32
	// DefaultSection holds the default value on creation for the "section" field.
	DefaultSection int32
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// DefaultProviderName holds the default value on creation for the "provider_name" field.
	DefaultProviderName string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultCurrency holds the default value on creation for the "currency" field.
	DefaultCurrency string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Asset queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySheet orders the results by the sheet field.
func BySheet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSheet, opts...).ToFunc()
}

// BySection orders the results by the section field.
func BySection(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSection, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByProviderName orders the results by the provider_name field.
func ByProviderName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
