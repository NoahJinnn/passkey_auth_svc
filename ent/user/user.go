// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeEmails holds the string denoting the emails edge name in mutations.
	EdgeEmails = "emails"
	// EdgePasscodes holds the string denoting the passcodes edge name in mutations.
	EdgePasscodes = "passcodes"
	// EdgePasswordCredential holds the string denoting the password_credential edge name in mutations.
	EdgePasswordCredential = "password_credential"
	// EdgePrimaryEmail holds the string denoting the primary_email edge name in mutations.
	EdgePrimaryEmail = "primary_email"
	// EdgeWebauthnCredentials holds the string denoting the webauthn_credentials edge name in mutations.
	EdgeWebauthnCredentials = "webauthn_credentials"
	// Table holds the table name of the user in the database.
	Table = "users"
	// EmailsTable is the table that holds the emails relation/edge.
	EmailsTable = "emails"
	// EmailsInverseTable is the table name for the Email entity.
	// It exists in this package in order to avoid circular dependency with the "email" package.
	EmailsInverseTable = "emails"
	// EmailsColumn is the table column denoting the emails relation/edge.
	EmailsColumn = "user_id"
	// PasscodesTable is the table that holds the passcodes relation/edge.
	PasscodesTable = "passcodes"
	// PasscodesInverseTable is the table name for the Passcode entity.
	// It exists in this package in order to avoid circular dependency with the "passcode" package.
	PasscodesInverseTable = "passcodes"
	// PasscodesColumn is the table column denoting the passcodes relation/edge.
	PasscodesColumn = "user_id"
	// PasswordCredentialTable is the table that holds the password_credential relation/edge.
	PasswordCredentialTable = "password_credentials"
	// PasswordCredentialInverseTable is the table name for the PasswordCredential entity.
	// It exists in this package in order to avoid circular dependency with the "passwordcredential" package.
	PasswordCredentialInverseTable = "password_credentials"
	// PasswordCredentialColumn is the table column denoting the password_credential relation/edge.
	PasswordCredentialColumn = "user_id"
	// PrimaryEmailTable is the table that holds the primary_email relation/edge.
	PrimaryEmailTable = "primary_emails"
	// PrimaryEmailInverseTable is the table name for the PrimaryEmail entity.
	// It exists in this package in order to avoid circular dependency with the "primaryemail" package.
	PrimaryEmailInverseTable = "primary_emails"
	// PrimaryEmailColumn is the table column denoting the primary_email relation/edge.
	PrimaryEmailColumn = "user_id"
	// WebauthnCredentialsTable is the table that holds the webauthn_credentials relation/edge.
	WebauthnCredentialsTable = "webauthn_credentials"
	// WebauthnCredentialsInverseTable is the table name for the WebauthnCredential entity.
	// It exists in this package in order to avoid circular dependency with the "webauthncredential" package.
	WebauthnCredentialsInverseTable = "webauthn_credentials"
	// WebauthnCredentialsColumn is the table column denoting the webauthn_credentials relation/edge.
	WebauthnCredentialsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
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
