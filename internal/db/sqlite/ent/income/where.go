// Code generated by ent, DO NOT EDIT.

package income

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Income {
	return predicate.Income(sql.FieldLTE(FieldID, id))
}

// ProviderName applies equality check predicate on the "provider_name" field. It's identical to ProviderNameEQ.
func ProviderName(v string) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldProviderName, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldData, v))
}

// ProviderNameEQ applies the EQ predicate on the "provider_name" field.
func ProviderNameEQ(v string) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldProviderName, v))
}

// ProviderNameNEQ applies the NEQ predicate on the "provider_name" field.
func ProviderNameNEQ(v string) predicate.Income {
	return predicate.Income(sql.FieldNEQ(FieldProviderName, v))
}

// ProviderNameIn applies the In predicate on the "provider_name" field.
func ProviderNameIn(vs ...string) predicate.Income {
	return predicate.Income(sql.FieldIn(FieldProviderName, vs...))
}

// ProviderNameNotIn applies the NotIn predicate on the "provider_name" field.
func ProviderNameNotIn(vs ...string) predicate.Income {
	return predicate.Income(sql.FieldNotIn(FieldProviderName, vs...))
}

// ProviderNameGT applies the GT predicate on the "provider_name" field.
func ProviderNameGT(v string) predicate.Income {
	return predicate.Income(sql.FieldGT(FieldProviderName, v))
}

// ProviderNameGTE applies the GTE predicate on the "provider_name" field.
func ProviderNameGTE(v string) predicate.Income {
	return predicate.Income(sql.FieldGTE(FieldProviderName, v))
}

// ProviderNameLT applies the LT predicate on the "provider_name" field.
func ProviderNameLT(v string) predicate.Income {
	return predicate.Income(sql.FieldLT(FieldProviderName, v))
}

// ProviderNameLTE applies the LTE predicate on the "provider_name" field.
func ProviderNameLTE(v string) predicate.Income {
	return predicate.Income(sql.FieldLTE(FieldProviderName, v))
}

// ProviderNameContains applies the Contains predicate on the "provider_name" field.
func ProviderNameContains(v string) predicate.Income {
	return predicate.Income(sql.FieldContains(FieldProviderName, v))
}

// ProviderNameHasPrefix applies the HasPrefix predicate on the "provider_name" field.
func ProviderNameHasPrefix(v string) predicate.Income {
	return predicate.Income(sql.FieldHasPrefix(FieldProviderName, v))
}

// ProviderNameHasSuffix applies the HasSuffix predicate on the "provider_name" field.
func ProviderNameHasSuffix(v string) predicate.Income {
	return predicate.Income(sql.FieldHasSuffix(FieldProviderName, v))
}

// ProviderNameEqualFold applies the EqualFold predicate on the "provider_name" field.
func ProviderNameEqualFold(v string) predicate.Income {
	return predicate.Income(sql.FieldEqualFold(FieldProviderName, v))
}

// ProviderNameContainsFold applies the ContainsFold predicate on the "provider_name" field.
func ProviderNameContainsFold(v string) predicate.Income {
	return predicate.Income(sql.FieldContainsFold(FieldProviderName, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Income {
	return predicate.Income(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Income {
	return predicate.Income(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Income {
	return predicate.Income(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Income {
	return predicate.Income(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Income {
	return predicate.Income(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Income {
	return predicate.Income(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Income {
	return predicate.Income(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Income {
	return predicate.Income(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Income {
	return predicate.Income(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Income {
	return predicate.Income(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Income {
	return predicate.Income(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Income {
	return predicate.Income(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Income {
	return predicate.Income(sql.FieldContainsFold(FieldData, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Income) predicate.Income {
	return predicate.Income(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Income) predicate.Income {
	return predicate.Income(func(s *sql.Selector) {
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
func Not(p predicate.Income) predicate.Income {
	return predicate.Income(func(s *sql.Selector) {
		p(s.Not())
	})
}
