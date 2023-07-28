// Code generated by ent, DO NOT EDIT.

package connection

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
)

const (
	// Label holds the string label denoting the connection type in the database.
	Label = "connection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProviderName holds the string denoting the provider_name field in the database.
	FieldProviderName = "provider_name"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// Table holds the table name of the connection in the database.
	Table = "connections"
)

// Columns holds all SQL columns for connection fields.
var Columns = []string{
	FieldID,
	FieldProviderName,
	FieldData,
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
	// DefaultProviderName holds the default value on creation for the "provider_name" field.
	DefaultProviderName string
	// DefaultData holds the default value on creation for the "data" field.
	DefaultData string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Connection queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProviderName orders the results by the provider_name field.
func ByProviderName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderName, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}
