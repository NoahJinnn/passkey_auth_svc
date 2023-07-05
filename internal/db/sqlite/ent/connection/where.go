// Code generated by ent, DO NOT EDIT.

package connection

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldID, id))
}

// InstitutionID applies equality check predicate on the "institution_id" field. It's identical to InstitutionIDEQ.
func InstitutionID(v uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldInstitutionID, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldData, v))
}

// Env applies equality check predicate on the "env" field. It's identical to EnvEQ.
func Env(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldEnv, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldUpdatedAt, v))
}

// InstitutionIDEQ applies the EQ predicate on the "institution_id" field.
func InstitutionIDEQ(v uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldInstitutionID, v))
}

// InstitutionIDNEQ applies the NEQ predicate on the "institution_id" field.
func InstitutionIDNEQ(v uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldInstitutionID, v))
}

// InstitutionIDIn applies the In predicate on the "institution_id" field.
func InstitutionIDIn(vs ...uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldInstitutionID, vs...))
}

// InstitutionIDNotIn applies the NotIn predicate on the "institution_id" field.
func InstitutionIDNotIn(vs ...uuid.UUID) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldInstitutionID, vs...))
}

// InstitutionIDIsNil applies the IsNil predicate on the "institution_id" field.
func InstitutionIDIsNil() predicate.Connection {
	return predicate.Connection(sql.FieldIsNull(FieldInstitutionID))
}

// InstitutionIDNotNil applies the NotNil predicate on the "institution_id" field.
func InstitutionIDNotNil() predicate.Connection {
	return predicate.Connection(sql.FieldNotNull(FieldInstitutionID))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldData, v))
}

// EnvEQ applies the EQ predicate on the "env" field.
func EnvEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldEnv, v))
}

// EnvNEQ applies the NEQ predicate on the "env" field.
func EnvNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldEnv, v))
}

// EnvIn applies the In predicate on the "env" field.
func EnvIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldEnv, vs...))
}

// EnvNotIn applies the NotIn predicate on the "env" field.
func EnvNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldEnv, vs...))
}

// EnvGT applies the GT predicate on the "env" field.
func EnvGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldEnv, v))
}

// EnvGTE applies the GTE predicate on the "env" field.
func EnvGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldEnv, v))
}

// EnvLT applies the LT predicate on the "env" field.
func EnvLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldEnv, v))
}

// EnvLTE applies the LTE predicate on the "env" field.
func EnvLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldEnv, v))
}

// EnvContains applies the Contains predicate on the "env" field.
func EnvContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldEnv, v))
}

// EnvHasPrefix applies the HasPrefix predicate on the "env" field.
func EnvHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldEnv, v))
}

// EnvHasSuffix applies the HasSuffix predicate on the "env" field.
func EnvHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldEnv, v))
}

// EnvEqualFold applies the EqualFold predicate on the "env" field.
func EnvEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldEnv, v))
}

// EnvContainsFold applies the ContainsFold predicate on the "env" field.
func EnvContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldEnv, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasInstitution applies the HasEdge predicate on the "institution" edge.
func HasInstitution() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, InstitutionTable, InstitutionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstitutionWith applies the HasEdge predicate on the "institution" edge with a given conditions (other predicates).
func HasInstitutionWith(preds ...predicate.Institution) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := newInstitutionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
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
func Not(p predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		p(s.Not())
	})
}
