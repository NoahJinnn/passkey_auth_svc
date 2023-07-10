// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/predicate"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/transaction"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetAccountID sets the "account_id" field.
func (tu *TransactionUpdate) SetAccountID(u uuid.UUID) *TransactionUpdate {
	tu.mutation.SetAccountID(u)
	return tu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAccountID(u *uuid.UUID) *TransactionUpdate {
	if u != nil {
		tu.SetAccountID(*u)
	}
	return tu
}

// ClearAccountID clears the value of the "account_id" field.
func (tu *TransactionUpdate) ClearAccountID() *TransactionUpdate {
	tu.mutation.ClearAccountID()
	return tu
}

// SetData sets the "data" field.
func (tu *TransactionUpdate) SetData(s string) *TransactionUpdate {
	tu.mutation.SetData(s)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TransactionUpdate) SetUpdatedAt(t time.Time) *TransactionUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetAccount sets the "account" edge to the Account entity.
func (tu *TransactionUpdate) SetAccount(a *Account) *TransactionUpdate {
	return tu.SetAccountID(a.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tu *TransactionUpdate) ClearAccount() *TransactionUpdate {
	tu.mutation.ClearAccount()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TransactionUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := transaction.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Data(); ok {
		_spec.SetField(transaction.FieldData, field.TypeString, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(transaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.AccountTable,
			Columns: []string{transaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.AccountTable,
			Columns: []string{transaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetAccountID sets the "account_id" field.
func (tuo *TransactionUpdateOne) SetAccountID(u uuid.UUID) *TransactionUpdateOne {
	tuo.mutation.SetAccountID(u)
	return tuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAccountID(u *uuid.UUID) *TransactionUpdateOne {
	if u != nil {
		tuo.SetAccountID(*u)
	}
	return tuo
}

// ClearAccountID clears the value of the "account_id" field.
func (tuo *TransactionUpdateOne) ClearAccountID() *TransactionUpdateOne {
	tuo.mutation.ClearAccountID()
	return tuo
}

// SetData sets the "data" field.
func (tuo *TransactionUpdateOne) SetData(s string) *TransactionUpdateOne {
	tuo.mutation.SetData(s)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TransactionUpdateOne) SetUpdatedAt(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetAccount sets the "account" edge to the Account entity.
func (tuo *TransactionUpdateOne) SetAccount(a *Account) *TransactionUpdateOne {
	return tuo.SetAccountID(a.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tuo *TransactionUpdateOne) ClearAccount() *TransactionUpdateOne {
	tuo.mutation.ClearAccount()
	return tuo
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TransactionUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := transaction.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Data(); ok {
		_spec.SetField(transaction.FieldData, field.TypeString, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(transaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.AccountTable,
			Columns: []string{transaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.AccountTable,
			Columns: []string{transaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
