// Code generated by ent, DO NOT EDIT.

package webauthncredential

import (
	"time"
)

const (
	// Label holds the string label denoting the webauthncredential type in the database.
	Label = "webauthn_credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// FieldAttestationType holds the string denoting the attestation_type field in the database.
	FieldAttestationType = "attestation_type"
	// FieldAaguid holds the string denoting the aaguid field in the database.
	FieldAaguid = "aaguid"
	// FieldSignCount holds the string denoting the sign_count field in the database.
	FieldSignCount = "sign_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBackupEligible holds the string denoting the backup_eligible field in the database.
	FieldBackupEligible = "backup_eligible"
	// FieldBackupState holds the string denoting the backup_state field in the database.
	FieldBackupState = "backup_state"
	// FieldLastUsedAt holds the string denoting the last_used_at field in the database.
	FieldLastUsedAt = "last_used_at"
	// EdgeWebauthnCredentialTransports holds the string denoting the webauthn_credential_transports edge name in mutations.
	EdgeWebauthnCredentialTransports = "webauthn_credential_transports"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the webauthncredential in the database.
	Table = "webauthn_credentials"
	// WebauthnCredentialTransportsTable is the table that holds the webauthn_credential_transports relation/edge.
	WebauthnCredentialTransportsTable = "webauthn_credential_transports"
	// WebauthnCredentialTransportsInverseTable is the table name for the WebauthnCredentialTransport entity.
	// It exists in this package in order to avoid circular dependency with the "webauthncredentialtransport" package.
	WebauthnCredentialTransportsInverseTable = "webauthn_credential_transports"
	// WebauthnCredentialTransportsColumn is the table column denoting the webauthn_credential_transports relation/edge.
	WebauthnCredentialTransportsColumn = "webauthn_credential_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "webauthn_credentials"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for webauthncredential fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldPublicKey,
	FieldAttestationType,
	FieldAaguid,
	FieldSignCount,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldBackupEligible,
	FieldBackupState,
	FieldLastUsedAt,
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