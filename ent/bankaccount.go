// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hellohq/hqservice/ent/bankaccount"
	"github.com/hellohq/hqservice/ent/user"
)

// BankAccount is the model entity for the BankAccount schema.
type BankAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint `json:"user_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID string `json:"account_id,omitempty"`
	// InstitutionInfo holds the value of the "institution_info" field.
	InstitutionInfo struct{} `json:"institution_info,omitempty"`
	// AccountInfo holds the value of the "account_info" field.
	AccountInfo struct{} `json:"account_info,omitempty"`
	// SensibleData holds the value of the "sensible_data" field.
	SensibleData string `json:"sensible_data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BankAccountQuery when eager-loading is set.
	Edges BankAccountEdges `json:"edges"`
}

// BankAccountEdges holds the relations/edges for other nodes in the graph.
type BankAccountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BankAccountEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BankAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bankaccount.FieldInstitutionInfo, bankaccount.FieldAccountInfo:
			values[i] = new([]byte)
		case bankaccount.FieldID, bankaccount.FieldUserID:
			values[i] = new(sql.NullInt64)
		case bankaccount.FieldAccountID, bankaccount.FieldSensibleData:
			values[i] = new(sql.NullString)
		case bankaccount.FieldCreatedAt, bankaccount.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BankAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BankAccount fields.
func (ba *BankAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bankaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ba.ID = uint(value.Int64)
		case bankaccount.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ba.UserID = uint(value.Int64)
			}
		case bankaccount.FieldAccountID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				ba.AccountID = value.String
			}
		case bankaccount.FieldInstitutionInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field institution_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ba.InstitutionInfo); err != nil {
					return fmt.Errorf("unmarshal field institution_info: %w", err)
				}
			}
		case bankaccount.FieldAccountInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field account_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ba.AccountInfo); err != nil {
					return fmt.Errorf("unmarshal field account_info: %w", err)
				}
			}
		case bankaccount.FieldSensibleData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sensible_data", values[i])
			} else if value.Valid {
				ba.SensibleData = value.String
			}
		case bankaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ba.CreatedAt = value.Time
			}
		case bankaccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ba.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the BankAccount entity.
func (ba *BankAccount) QueryUser() *UserQuery {
	return NewBankAccountClient(ba.config).QueryUser(ba)
}

// Update returns a builder for updating this BankAccount.
// Note that you need to call BankAccount.Unwrap() before calling this method if this BankAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ba *BankAccount) Update() *BankAccountUpdateOne {
	return NewBankAccountClient(ba.config).UpdateOne(ba)
}

// Unwrap unwraps the BankAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ba *BankAccount) Unwrap() *BankAccount {
	_tx, ok := ba.config.driver.(*txDriver)
	if !ok {
		panic("ent: BankAccount is not a transactional entity")
	}
	ba.config.driver = _tx.drv
	return ba
}

// String implements the fmt.Stringer.
func (ba *BankAccount) String() string {
	var builder strings.Builder
	builder.WriteString("BankAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ba.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ba.UserID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(ba.AccountID)
	builder.WriteString(", ")
	builder.WriteString("institution_info=")
	builder.WriteString(fmt.Sprintf("%v", ba.InstitutionInfo))
	builder.WriteString(", ")
	builder.WriteString("account_info=")
	builder.WriteString(fmt.Sprintf("%v", ba.AccountInfo))
	builder.WriteString(", ")
	builder.WriteString("sensible_data=")
	builder.WriteString(ba.SensibleData)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ba.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ba.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BankAccounts is a parsable slice of BankAccount.
type BankAccounts []*BankAccount
