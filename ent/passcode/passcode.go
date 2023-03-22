// Code generated by ent, DO NOT EDIT.

package passcode

import (
	"time"
)

const (
	// Label holds the string label denoting the passcode type in the database.
	Label = "passcode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTTL holds the string denoting the ttl field in the database.
	FieldTTL = "ttl"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldTryCount holds the string denoting the try_count field in the database.
	FieldTryCount = "try_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmailID holds the string denoting the email_id field in the database.
	FieldEmailID = "email_id"
	// EdgeEmail holds the string denoting the email edge name in mutations.
	EdgeEmail = "email"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the passcode in the database.
	Table = "passcodes"
	// EmailTable is the table that holds the email relation/edge.
	EmailTable = "passcodes"
	// EmailInverseTable is the table name for the Email entity.
	// It exists in this package in order to avoid circular dependency with the "email" package.
	EmailInverseTable = "emails"
	// EmailColumn is the table column denoting the email relation/edge.
	EmailColumn = "email_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "passcodes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for passcode fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTTL,
	FieldCode,
	FieldTryCount,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmailID,
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
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
