// Code generated by ent, DO NOT EDIT.

package itemtable

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldUserID, v))
}

// Sheet applies equality check predicate on the "sheet" field. It's identical to SheetEQ.
func Sheet(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldSheet, v))
}

// Section applies equality check predicate on the "section" field. It's identical to SectionEQ.
func Section(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldSection, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldCategory, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotNull(FieldUserID))
}

// SheetEQ applies the EQ predicate on the "sheet" field.
func SheetEQ(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldSheet, v))
}

// SheetNEQ applies the NEQ predicate on the "sheet" field.
func SheetNEQ(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldSheet, v))
}

// SheetIn applies the In predicate on the "sheet" field.
func SheetIn(vs ...int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldSheet, vs...))
}

// SheetNotIn applies the NotIn predicate on the "sheet" field.
func SheetNotIn(vs ...int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldSheet, vs...))
}

// SheetGT applies the GT predicate on the "sheet" field.
func SheetGT(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldSheet, v))
}

// SheetGTE applies the GTE predicate on the "sheet" field.
func SheetGTE(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldSheet, v))
}

// SheetLT applies the LT predicate on the "sheet" field.
func SheetLT(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldSheet, v))
}

// SheetLTE applies the LTE predicate on the "sheet" field.
func SheetLTE(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldSheet, v))
}

// SectionEQ applies the EQ predicate on the "section" field.
func SectionEQ(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldSection, v))
}

// SectionNEQ applies the NEQ predicate on the "section" field.
func SectionNEQ(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldSection, v))
}

// SectionIn applies the In predicate on the "section" field.
func SectionIn(vs ...int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldSection, vs...))
}

// SectionNotIn applies the NotIn predicate on the "section" field.
func SectionNotIn(vs ...int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldSection, vs...))
}

// SectionGT applies the GT predicate on the "section" field.
func SectionGT(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldSection, v))
}

// SectionGTE applies the GTE predicate on the "section" field.
func SectionGTE(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldSection, v))
}

// SectionLT applies the LT predicate on the "section" field.
func SectionLT(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldSection, v))
}

// SectionLTE applies the LTE predicate on the "section" field.
func SectionLTE(v int32) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldSection, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldContainsFold(FieldCategory, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ItemTable {
	return predicate.ItemTable(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ItemTable {
	return predicate.ItemTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ItemTable {
	return predicate.ItemTable(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ItemTable) predicate.ItemTable {
	return predicate.ItemTable(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ItemTable) predicate.ItemTable {
	return predicate.ItemTable(func(s *sql.Selector) {
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
func Not(p predicate.ItemTable) predicate.ItemTable {
	return predicate.ItemTable(func(s *sql.Selector) {
		p(s.Not())
	})
}
