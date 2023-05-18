// Code generated by ent, DO NOT EDIT.

package webauthnsessiondata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
)

const (
	// Label holds the string label denoting the webauthnsessiondata type in the database.
	Label = "webauthn_session_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChallenge holds the string denoting the challenge field in the database.
	FieldChallenge = "challenge"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserVerification holds the string denoting the user_verification field in the database.
	FieldUserVerification = "user_verification"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWebauthnSessionDataAllowedCredentials holds the string denoting the webauthn_session_data_allowed_credentials edge name in mutations.
	EdgeWebauthnSessionDataAllowedCredentials = "webauthn_session_data_allowed_credentials"
	// Table holds the table name of the webauthnsessiondata in the database.
	Table = "webauthn_session_data"
	// WebauthnSessionDataAllowedCredentialsTable is the table that holds the webauthn_session_data_allowed_credentials relation/edge.
	WebauthnSessionDataAllowedCredentialsTable = "webauthn_session_data_allowed_credentials"
	// WebauthnSessionDataAllowedCredentialsInverseTable is the table name for the WebauthnSessionDataAllowedCredential entity.
	// It exists in this package in order to avoid circular dependency with the "webauthnsessiondataallowedcredential" package.
	WebauthnSessionDataAllowedCredentialsInverseTable = "webauthn_session_data_allowed_credentials"
	// WebauthnSessionDataAllowedCredentialsColumn is the table column denoting the webauthn_session_data_allowed_credentials relation/edge.
	WebauthnSessionDataAllowedCredentialsColumn = "webauthn_session_data_id"
)

// Columns holds all SQL columns for webauthnsessiondata fields.
var Columns = []string{
	FieldID,
	FieldChallenge,
	FieldUserID,
	FieldUserVerification,
	FieldOperation,
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

// OrderOption defines the ordering options for the WebauthnSessionData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChallenge orders the results by the challenge field.
func ByChallenge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChallenge, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserVerification orders the results by the user_verification field.
func ByUserVerification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserVerification, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWebauthnSessionDataAllowedCredentialsCount orders the results by webauthn_session_data_allowed_credentials count.
func ByWebauthnSessionDataAllowedCredentialsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWebauthnSessionDataAllowedCredentialsStep(), opts...)
	}
}

// ByWebauthnSessionDataAllowedCredentials orders the results by webauthn_session_data_allowed_credentials terms.
func ByWebauthnSessionDataAllowedCredentials(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWebauthnSessionDataAllowedCredentialsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWebauthnSessionDataAllowedCredentialsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WebauthnSessionDataAllowedCredentialsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WebauthnSessionDataAllowedCredentialsTable, WebauthnSessionDataAllowedCredentialsColumn),
	)
}
