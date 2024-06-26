// Code generated by ent, DO NOT EDIT.

package webauthnsessiondataallowedcredential

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
)

const (
	// Label holds the string label denoting the webauthnsessiondataallowedcredential type in the database.
	Label = "webauthn_session_data_allowed_credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCredentialID holds the string denoting the credential_id field in the database.
	FieldCredentialID = "credential_id"
	// FieldWebauthnSessionDataID holds the string denoting the webauthn_session_data_id field in the database.
	FieldWebauthnSessionDataID = "webauthn_session_data_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWebauthnSessionData holds the string denoting the webauthn_session_data edge name in mutations.
	EdgeWebauthnSessionData = "webauthn_session_data"
	// Table holds the table name of the webauthnsessiondataallowedcredential in the database.
	Table = "webauthn_session_data_allowed_credentials"
	// WebauthnSessionDataTable is the table that holds the webauthn_session_data relation/edge.
	WebauthnSessionDataTable = "webauthn_session_data_allowed_credentials"
	// WebauthnSessionDataInverseTable is the table name for the WebauthnSessionData entity.
	// It exists in this package in order to avoid circular dependency with the "webauthnsessiondata" package.
	WebauthnSessionDataInverseTable = "webauthn_session_data"
	// WebauthnSessionDataColumn is the table column denoting the webauthn_session_data relation/edge.
	WebauthnSessionDataColumn = "webauthn_session_data_id"
)

// Columns holds all SQL columns for webauthnsessiondataallowedcredential fields.
var Columns = []string{
	FieldID,
	FieldCredentialID,
	FieldWebauthnSessionDataID,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the WebauthnSessionDataAllowedCredential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCredentialID orders the results by the credential_id field.
func ByCredentialID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredentialID, opts...).ToFunc()
}

// ByWebauthnSessionDataID orders the results by the webauthn_session_data_id field.
func ByWebauthnSessionDataID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebauthnSessionDataID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWebauthnSessionDataField orders the results by webauthn_session_data field.
func ByWebauthnSessionDataField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWebauthnSessionDataStep(), sql.OrderByField(field, opts...))
	}
}
func newWebauthnSessionDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WebauthnSessionDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WebauthnSessionDataTable, WebauthnSessionDataColumn),
	)
}
