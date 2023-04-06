// Code generated by ent, DO NOT EDIT.

package jwk

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.Jwk {
	return predicate.Jwk(sql.FieldLTE(FieldID, id))
}

// KeyData applies equality check predicate on the "key_data" field. It's identical to KeyDataEQ.
func KeyData(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldKeyData, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldCreatedAt, v))
}

// KeyDataEQ applies the EQ predicate on the "key_data" field.
func KeyDataEQ(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldKeyData, v))
}

// KeyDataNEQ applies the NEQ predicate on the "key_data" field.
func KeyDataNEQ(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldNEQ(FieldKeyData, v))
}

// KeyDataIn applies the In predicate on the "key_data" field.
func KeyDataIn(vs ...string) predicate.Jwk {
	return predicate.Jwk(sql.FieldIn(FieldKeyData, vs...))
}

// KeyDataNotIn applies the NotIn predicate on the "key_data" field.
func KeyDataNotIn(vs ...string) predicate.Jwk {
	return predicate.Jwk(sql.FieldNotIn(FieldKeyData, vs...))
}

// KeyDataGT applies the GT predicate on the "key_data" field.
func KeyDataGT(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldGT(FieldKeyData, v))
}

// KeyDataGTE applies the GTE predicate on the "key_data" field.
func KeyDataGTE(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldGTE(FieldKeyData, v))
}

// KeyDataLT applies the LT predicate on the "key_data" field.
func KeyDataLT(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldLT(FieldKeyData, v))
}

// KeyDataLTE applies the LTE predicate on the "key_data" field.
func KeyDataLTE(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldLTE(FieldKeyData, v))
}

// KeyDataContains applies the Contains predicate on the "key_data" field.
func KeyDataContains(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldContains(FieldKeyData, v))
}

// KeyDataHasPrefix applies the HasPrefix predicate on the "key_data" field.
func KeyDataHasPrefix(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldHasPrefix(FieldKeyData, v))
}

// KeyDataHasSuffix applies the HasSuffix predicate on the "key_data" field.
func KeyDataHasSuffix(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldHasSuffix(FieldKeyData, v))
}

// KeyDataEqualFold applies the EqualFold predicate on the "key_data" field.
func KeyDataEqualFold(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldEqualFold(FieldKeyData, v))
}

// KeyDataContainsFold applies the ContainsFold predicate on the "key_data" field.
func KeyDataContainsFold(v string) predicate.Jwk {
	return predicate.Jwk(sql.FieldContainsFold(FieldKeyData, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Jwk {
	return predicate.Jwk(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Jwk) predicate.Jwk {
	return predicate.Jwk(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Jwk) predicate.Jwk {
	return predicate.Jwk(func(s *sql.Selector) {
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
func Not(p predicate.Jwk) predicate.Jwk {
	return predicate.Jwk(func(s *sql.Selector) {
		p(s.Not())
	})
}