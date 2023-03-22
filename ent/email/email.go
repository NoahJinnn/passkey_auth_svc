// Code generated by ent, DO NOT EDIT.

package email

import (
	"time"
)

const (
	// Label holds the string label denoting the email type in the database.
	Label = "email"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeIdentities holds the string denoting the identities edge name in mutations.
	EdgeIdentities = "identities"
	// EdgePasscodes holds the string denoting the passcodes edge name in mutations.
	EdgePasscodes = "passcodes"
	// EdgePrimaryEmail holds the string denoting the primary_email edge name in mutations.
	EdgePrimaryEmail = "primary_email"
	// Table holds the table name of the email in the database.
	Table = "emails"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "emails"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// IdentitiesTable is the table that holds the identities relation/edge.
	IdentitiesTable = "identities"
	// IdentitiesInverseTable is the table name for the Identity entity.
	// It exists in this package in order to avoid circular dependency with the "identity" package.
	IdentitiesInverseTable = "identities"
	// IdentitiesColumn is the table column denoting the identities relation/edge.
	IdentitiesColumn = "email_id"
	// PasscodesTable is the table that holds the passcodes relation/edge.
	PasscodesTable = "passcodes"
	// PasscodesInverseTable is the table name for the Passcode entity.
	// It exists in this package in order to avoid circular dependency with the "passcode" package.
	PasscodesInverseTable = "passcodes"
	// PasscodesColumn is the table column denoting the passcodes relation/edge.
	PasscodesColumn = "email_id"
	// PrimaryEmailTable is the table that holds the primary_email relation/edge.
	PrimaryEmailTable = "primary_emails"
	// PrimaryEmailInverseTable is the table name for the PrimaryEmail entity.
	// It exists in this package in order to avoid circular dependency with the "primaryemail" package.
	PrimaryEmailInverseTable = "primary_emails"
	// PrimaryEmailColumn is the table column denoting the primary_email relation/edge.
	PrimaryEmailColumn = "email_id"
)

// Columns holds all SQL columns for email fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAddress,
	FieldVerified,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
