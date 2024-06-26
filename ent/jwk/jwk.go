// Code generated by ent, DO NOT EDIT.

package jwk

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the jwk type in the database.
	Label = "jwk"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKeyData holds the string denoting the key_data field in the database.
	FieldKeyData = "key_data"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the jwk in the database.
	Table = "jwks"
)

// Columns holds all SQL columns for jwk fields.
var Columns = []string{
	FieldID,
	FieldKeyData,
	FieldCreatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Jwk queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKeyData orders the results by the key_data field.
func ByKeyData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyData, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
