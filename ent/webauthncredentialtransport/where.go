// Code generated by ent, DO NOT EDIT.

package webauthncredentialtransport

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldName, v))
}

// WebauthnCredentialID applies equality check predicate on the "webauthn_credential_id" field. It's identical to WebauthnCredentialIDEQ.
func WebauthnCredentialID(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldWebauthnCredentialID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldContainsFold(FieldName, v))
}

// WebauthnCredentialIDEQ applies the EQ predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDEQ(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEQ(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDNEQ applies the NEQ predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDNEQ(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNEQ(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDIn applies the In predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDIn(vs ...string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldIn(FieldWebauthnCredentialID, vs...))
}

// WebauthnCredentialIDNotIn applies the NotIn predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDNotIn(vs ...string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNotIn(FieldWebauthnCredentialID, vs...))
}

// WebauthnCredentialIDGT applies the GT predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDGT(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGT(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDGTE applies the GTE predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDGTE(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldGTE(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDLT applies the LT predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDLT(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLT(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDLTE applies the LTE predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDLTE(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldLTE(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDContains applies the Contains predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDContains(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldContains(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDHasPrefix applies the HasPrefix predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDHasPrefix(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldHasPrefix(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDHasSuffix applies the HasSuffix predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDHasSuffix(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldHasSuffix(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDIsNil applies the IsNil predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDIsNil() predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldIsNull(FieldWebauthnCredentialID))
}

// WebauthnCredentialIDNotNil applies the NotNil predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDNotNil() predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldNotNull(FieldWebauthnCredentialID))
}

// WebauthnCredentialIDEqualFold applies the EqualFold predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDEqualFold(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldEqualFold(FieldWebauthnCredentialID, v))
}

// WebauthnCredentialIDContainsFold applies the ContainsFold predicate on the "webauthn_credential_id" field.
func WebauthnCredentialIDContainsFold(v string) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(sql.FieldContainsFold(FieldWebauthnCredentialID, v))
}

// HasWebauthnCredential applies the HasEdge predicate on the "webauthn_credential" edge.
func HasWebauthnCredential() predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WebauthnCredentialTable, WebauthnCredentialColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebauthnCredentialWith applies the HasEdge predicate on the "webauthn_credential" edge with a given conditions (other predicates).
func HasWebauthnCredentialWith(preds ...predicate.WebauthnCredential) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
		step := newWebauthnCredentialStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WebauthnCredentialTransport) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WebauthnCredentialTransport) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
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
func Not(p predicate.WebauthnCredentialTransport) predicate.WebauthnCredentialTransport {
	return predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
		p(s.Not())
	})
}
