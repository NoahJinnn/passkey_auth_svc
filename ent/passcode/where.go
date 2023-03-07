// Code generated by ent, DO NOT EDIT.

package passcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldUserID, v))
}

// TTL applies equality check predicate on the "ttl" field. It's identical to TTLEQ.
func TTL(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldTTL, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldCode, v))
}

// TryCount applies equality check predicate on the "try_count" field. It's identical to TryCountEQ.
func TryCount(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldTryCount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldUpdatedAt, v))
}

// EmailID applies equality check predicate on the "email_id" field. It's identical to EmailIDEQ.
func EmailID(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldEmailID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Passcode {
	return predicate.Passcode(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Passcode {
	return predicate.Passcode(sql.FieldNotNull(FieldUserID))
}

// TTLEQ applies the EQ predicate on the "ttl" field.
func TTLEQ(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldTTL, v))
}

// TTLNEQ applies the NEQ predicate on the "ttl" field.
func TTLNEQ(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldTTL, v))
}

// TTLIn applies the In predicate on the "ttl" field.
func TTLIn(vs ...int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldTTL, vs...))
}

// TTLNotIn applies the NotIn predicate on the "ttl" field.
func TTLNotIn(vs ...int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldTTL, vs...))
}

// TTLGT applies the GT predicate on the "ttl" field.
func TTLGT(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldTTL, v))
}

// TTLGTE applies the GTE predicate on the "ttl" field.
func TTLGTE(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldTTL, v))
}

// TTLLT applies the LT predicate on the "ttl" field.
func TTLLT(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldTTL, v))
}

// TTLLTE applies the LTE predicate on the "ttl" field.
func TTLLTE(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldTTL, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Passcode {
	return predicate.Passcode(sql.FieldContainsFold(FieldCode, v))
}

// TryCountEQ applies the EQ predicate on the "try_count" field.
func TryCountEQ(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldTryCount, v))
}

// TryCountNEQ applies the NEQ predicate on the "try_count" field.
func TryCountNEQ(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldTryCount, v))
}

// TryCountIn applies the In predicate on the "try_count" field.
func TryCountIn(vs ...int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldTryCount, vs...))
}

// TryCountNotIn applies the NotIn predicate on the "try_count" field.
func TryCountNotIn(vs ...int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldTryCount, vs...))
}

// TryCountGT applies the GT predicate on the "try_count" field.
func TryCountGT(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldTryCount, v))
}

// TryCountGTE applies the GTE predicate on the "try_count" field.
func TryCountGTE(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldTryCount, v))
}

// TryCountLT applies the LT predicate on the "try_count" field.
func TryCountLT(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldTryCount, v))
}

// TryCountLTE applies the LTE predicate on the "try_count" field.
func TryCountLTE(v int32) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldTryCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Passcode {
	return predicate.Passcode(sql.FieldLTE(FieldUpdatedAt, v))
}

// EmailIDEQ applies the EQ predicate on the "email_id" field.
func EmailIDEQ(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldEQ(FieldEmailID, v))
}

// EmailIDNEQ applies the NEQ predicate on the "email_id" field.
func EmailIDNEQ(v uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNEQ(FieldEmailID, v))
}

// EmailIDIn applies the In predicate on the "email_id" field.
func EmailIDIn(vs ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldIn(FieldEmailID, vs...))
}

// EmailIDNotIn applies the NotIn predicate on the "email_id" field.
func EmailIDNotIn(vs ...uuid.UUID) predicate.Passcode {
	return predicate.Passcode(sql.FieldNotIn(FieldEmailID, vs...))
}

// EmailIDIsNil applies the IsNil predicate on the "email_id" field.
func EmailIDIsNil() predicate.Passcode {
	return predicate.Passcode(sql.FieldIsNull(FieldEmailID))
}

// EmailIDNotNil applies the NotNil predicate on the "email_id" field.
func EmailIDNotNil() predicate.Passcode {
	return predicate.Passcode(sql.FieldNotNull(FieldEmailID))
}

// HasEmail applies the HasEdge predicate on the "email" edge.
func HasEmail() predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmailTable, EmailColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmailWith applies the HasEdge predicate on the "email" edge with a given conditions (other predicates).
func HasEmailWith(preds ...predicate.Email) predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmailInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmailTable, EmailColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Passcode) predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Passcode) predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
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
func Not(p predicate.Passcode) predicate.Passcode {
	return predicate.Passcode(func(s *sql.Selector) {
		p(s.Not())
	})
}
