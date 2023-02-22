// Code generated by ent, DO NOT EDIT.

package loan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldUserID, v))
}

// LenderName applies equality check predicate on the "lender_name" field. It's identical to LenderNameEQ.
func LenderName(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldLenderName, v))
}

// LoanType applies equality check predicate on the "loan_type" field. It's identical to LoanTypeEQ.
func LoanType(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldLoanType, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldBalance, v))
}

// InterestRate applies equality check predicate on the "interest_rate" field. It's identical to InterestRateEQ.
func InterestRate(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldInterestRate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Loan {
	return predicate.Loan(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Loan {
	return predicate.Loan(sql.FieldNotNull(FieldUserID))
}

// LenderNameEQ applies the EQ predicate on the "lender_name" field.
func LenderNameEQ(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldLenderName, v))
}

// LenderNameNEQ applies the NEQ predicate on the "lender_name" field.
func LenderNameNEQ(v string) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldLenderName, v))
}

// LenderNameIn applies the In predicate on the "lender_name" field.
func LenderNameIn(vs ...string) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldLenderName, vs...))
}

// LenderNameNotIn applies the NotIn predicate on the "lender_name" field.
func LenderNameNotIn(vs ...string) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldLenderName, vs...))
}

// LenderNameGT applies the GT predicate on the "lender_name" field.
func LenderNameGT(v string) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldLenderName, v))
}

// LenderNameGTE applies the GTE predicate on the "lender_name" field.
func LenderNameGTE(v string) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldLenderName, v))
}

// LenderNameLT applies the LT predicate on the "lender_name" field.
func LenderNameLT(v string) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldLenderName, v))
}

// LenderNameLTE applies the LTE predicate on the "lender_name" field.
func LenderNameLTE(v string) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldLenderName, v))
}

// LenderNameContains applies the Contains predicate on the "lender_name" field.
func LenderNameContains(v string) predicate.Loan {
	return predicate.Loan(sql.FieldContains(FieldLenderName, v))
}

// LenderNameHasPrefix applies the HasPrefix predicate on the "lender_name" field.
func LenderNameHasPrefix(v string) predicate.Loan {
	return predicate.Loan(sql.FieldHasPrefix(FieldLenderName, v))
}

// LenderNameHasSuffix applies the HasSuffix predicate on the "lender_name" field.
func LenderNameHasSuffix(v string) predicate.Loan {
	return predicate.Loan(sql.FieldHasSuffix(FieldLenderName, v))
}

// LenderNameEqualFold applies the EqualFold predicate on the "lender_name" field.
func LenderNameEqualFold(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEqualFold(FieldLenderName, v))
}

// LenderNameContainsFold applies the ContainsFold predicate on the "lender_name" field.
func LenderNameContainsFold(v string) predicate.Loan {
	return predicate.Loan(sql.FieldContainsFold(FieldLenderName, v))
}

// LoanTypeEQ applies the EQ predicate on the "loan_type" field.
func LoanTypeEQ(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldLoanType, v))
}

// LoanTypeNEQ applies the NEQ predicate on the "loan_type" field.
func LoanTypeNEQ(v string) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldLoanType, v))
}

// LoanTypeIn applies the In predicate on the "loan_type" field.
func LoanTypeIn(vs ...string) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldLoanType, vs...))
}

// LoanTypeNotIn applies the NotIn predicate on the "loan_type" field.
func LoanTypeNotIn(vs ...string) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldLoanType, vs...))
}

// LoanTypeGT applies the GT predicate on the "loan_type" field.
func LoanTypeGT(v string) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldLoanType, v))
}

// LoanTypeGTE applies the GTE predicate on the "loan_type" field.
func LoanTypeGTE(v string) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldLoanType, v))
}

// LoanTypeLT applies the LT predicate on the "loan_type" field.
func LoanTypeLT(v string) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldLoanType, v))
}

// LoanTypeLTE applies the LTE predicate on the "loan_type" field.
func LoanTypeLTE(v string) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldLoanType, v))
}

// LoanTypeContains applies the Contains predicate on the "loan_type" field.
func LoanTypeContains(v string) predicate.Loan {
	return predicate.Loan(sql.FieldContains(FieldLoanType, v))
}

// LoanTypeHasPrefix applies the HasPrefix predicate on the "loan_type" field.
func LoanTypeHasPrefix(v string) predicate.Loan {
	return predicate.Loan(sql.FieldHasPrefix(FieldLoanType, v))
}

// LoanTypeHasSuffix applies the HasSuffix predicate on the "loan_type" field.
func LoanTypeHasSuffix(v string) predicate.Loan {
	return predicate.Loan(sql.FieldHasSuffix(FieldLoanType, v))
}

// LoanTypeEqualFold applies the EqualFold predicate on the "loan_type" field.
func LoanTypeEqualFold(v string) predicate.Loan {
	return predicate.Loan(sql.FieldEqualFold(FieldLoanType, v))
}

// LoanTypeContainsFold applies the ContainsFold predicate on the "loan_type" field.
func LoanTypeContainsFold(v string) predicate.Loan {
	return predicate.Loan(sql.FieldContainsFold(FieldLoanType, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldBalance, v))
}

// InterestRateEQ applies the EQ predicate on the "interest_rate" field.
func InterestRateEQ(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldInterestRate, v))
}

// InterestRateNEQ applies the NEQ predicate on the "interest_rate" field.
func InterestRateNEQ(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldInterestRate, v))
}

// InterestRateIn applies the In predicate on the "interest_rate" field.
func InterestRateIn(vs ...float64) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldInterestRate, vs...))
}

// InterestRateNotIn applies the NotIn predicate on the "interest_rate" field.
func InterestRateNotIn(vs ...float64) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldInterestRate, vs...))
}

// InterestRateGT applies the GT predicate on the "interest_rate" field.
func InterestRateGT(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldInterestRate, v))
}

// InterestRateGTE applies the GTE predicate on the "interest_rate" field.
func InterestRateGTE(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldInterestRate, v))
}

// InterestRateLT applies the LT predicate on the "interest_rate" field.
func InterestRateLT(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldInterestRate, v))
}

// InterestRateLTE applies the LTE predicate on the "interest_rate" field.
func InterestRateLTE(v float64) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldInterestRate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Loan {
	return predicate.Loan(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
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
func And(predicates ...predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
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
func Not(p predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		p(s.Not())
	})
}
