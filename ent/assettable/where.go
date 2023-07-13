// Code generated by ent, DO NOT EDIT.

package assettable

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldUserID, v))
}

// Sheet applies equality check predicate on the "sheet" field. It's identical to SheetEQ.
func Sheet(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldSheet, v))
}

// Section applies equality check predicate on the "section" field. It's identical to SectionEQ.
func Section(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldSection, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotNull(FieldUserID))
}

// SheetEQ applies the EQ predicate on the "sheet" field.
func SheetEQ(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldSheet, v))
}

// SheetNEQ applies the NEQ predicate on the "sheet" field.
func SheetNEQ(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldSheet, v))
}

// SheetIn applies the In predicate on the "sheet" field.
func SheetIn(vs ...int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldSheet, vs...))
}

// SheetNotIn applies the NotIn predicate on the "sheet" field.
func SheetNotIn(vs ...int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldSheet, vs...))
}

// SheetGT applies the GT predicate on the "sheet" field.
func SheetGT(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldSheet, v))
}

// SheetGTE applies the GTE predicate on the "sheet" field.
func SheetGTE(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldSheet, v))
}

// SheetLT applies the LT predicate on the "sheet" field.
func SheetLT(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldSheet, v))
}

// SheetLTE applies the LTE predicate on the "sheet" field.
func SheetLTE(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldSheet, v))
}

// SheetIsNil applies the IsNil predicate on the "sheet" field.
func SheetIsNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIsNull(FieldSheet))
}

// SheetNotNil applies the NotNil predicate on the "sheet" field.
func SheetNotNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotNull(FieldSheet))
}

// SectionEQ applies the EQ predicate on the "section" field.
func SectionEQ(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldSection, v))
}

// SectionNEQ applies the NEQ predicate on the "section" field.
func SectionNEQ(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldSection, v))
}

// SectionIn applies the In predicate on the "section" field.
func SectionIn(vs ...int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldSection, vs...))
}

// SectionNotIn applies the NotIn predicate on the "section" field.
func SectionNotIn(vs ...int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldSection, vs...))
}

// SectionGT applies the GT predicate on the "section" field.
func SectionGT(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldSection, v))
}

// SectionGTE applies the GTE predicate on the "section" field.
func SectionGTE(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldSection, v))
}

// SectionLT applies the LT predicate on the "section" field.
func SectionLT(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldSection, v))
}

// SectionLTE applies the LTE predicate on the "section" field.
func SectionLTE(v int32) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldSection, v))
}

// SectionIsNil applies the IsNil predicate on the "section" field.
func SectionIsNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIsNull(FieldSection))
}

// SectionNotNil applies the NotNil predicate on the "section" field.
func SectionNotNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotNull(FieldSection))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AssetTable {
	return predicate.AssetTable(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.AssetTable {
	return predicate.AssetTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.AssetTable {
	return predicate.AssetTable(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AssetTable) predicate.AssetTable {
	return predicate.AssetTable(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AssetTable) predicate.AssetTable {
	return predicate.AssetTable(func(s *sql.Selector) {
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
func Not(p predicate.AssetTable) predicate.AssetTable {
	return predicate.AssetTable(func(s *sql.Selector) {
		p(s.Not())
	})
}