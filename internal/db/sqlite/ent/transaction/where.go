// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAccountID, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldData, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDIsNil applies the IsNil predicate on the "account_id" field.
func AccountIDIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldAccountID))
}

// AccountIDNotNil applies the NotNil predicate on the "account_id" field.
func AccountIDNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldAccountID))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldData, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		p(s.Not())
	})
}
