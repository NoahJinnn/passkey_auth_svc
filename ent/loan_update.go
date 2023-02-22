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
	"github.com/hellohq/hqservice/ent/loan"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/user"
)

// LoanUpdate is the builder for updating Loan entities.
type LoanUpdate struct {
	config
	hooks    []Hook
	mutation *LoanMutation
}

// Where appends a list predicates to the LoanUpdate builder.
func (lu *LoanUpdate) Where(ps ...predicate.Loan) *LoanUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *LoanUpdate) SetUserID(u uint) *LoanUpdate {
	lu.mutation.SetUserID(u)
	return lu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lu *LoanUpdate) SetNillableUserID(u *uint) *LoanUpdate {
	if u != nil {
		lu.SetUserID(*u)
	}
	return lu
}

// ClearUserID clears the value of the "user_id" field.
func (lu *LoanUpdate) ClearUserID() *LoanUpdate {
	lu.mutation.ClearUserID()
	return lu
}

// SetAssetInfoID sets the "asset_info_id" field.
func (lu *LoanUpdate) SetAssetInfoID(u uint) *LoanUpdate {
	lu.mutation.ResetAssetInfoID()
	lu.mutation.SetAssetInfoID(u)
	return lu
}

// AddAssetInfoID adds u to the "asset_info_id" field.
func (lu *LoanUpdate) AddAssetInfoID(u int) *LoanUpdate {
	lu.mutation.AddAssetInfoID(u)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LoanUpdate) SetCreatedAt(t time.Time) *LoanUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LoanUpdate) SetUpdatedAt(t time.Time) *LoanUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetUser sets the "user" edge to the User entity.
func (lu *LoanUpdate) SetUser(u *User) *LoanUpdate {
	return lu.SetUserID(u.ID)
}

// Mutation returns the LoanMutation object of the builder.
func (lu *LoanUpdate) Mutation() *LoanMutation {
	return lu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LoanUpdate) ClearUser() *LoanUpdate {
	lu.mutation.ClearUser()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LoanUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LoanMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LoanUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LoanUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LoanUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LoanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(loan.Table, loan.Columns, sqlgraph.NewFieldSpec(loan.FieldID, field.TypeUint))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.AssetInfoID(); ok {
		_spec.SetField(loan.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := lu.mutation.AddedAssetInfoID(); ok {
		_spec.AddField(loan.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(loan.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(loan.FieldUpdatedAt, field.TypeTime, value)
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LoanUpdateOne is the builder for updating a single Loan entity.
type LoanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoanMutation
}

// SetUserID sets the "user_id" field.
func (luo *LoanUpdateOne) SetUserID(u uint) *LoanUpdateOne {
	luo.mutation.SetUserID(u)
	return luo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (luo *LoanUpdateOne) SetNillableUserID(u *uint) *LoanUpdateOne {
	if u != nil {
		luo.SetUserID(*u)
	}
	return luo
}

// ClearUserID clears the value of the "user_id" field.
func (luo *LoanUpdateOne) ClearUserID() *LoanUpdateOne {
	luo.mutation.ClearUserID()
	return luo
}

// SetAssetInfoID sets the "asset_info_id" field.
func (luo *LoanUpdateOne) SetAssetInfoID(u uint) *LoanUpdateOne {
	luo.mutation.ResetAssetInfoID()
	luo.mutation.SetAssetInfoID(u)
	return luo
}

// AddAssetInfoID adds u to the "asset_info_id" field.
func (luo *LoanUpdateOne) AddAssetInfoID(u int) *LoanUpdateOne {
	luo.mutation.AddAssetInfoID(u)
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LoanUpdateOne) SetCreatedAt(t time.Time) *LoanUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LoanUpdateOne) SetUpdatedAt(t time.Time) *LoanUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetUser sets the "user" edge to the User entity.
func (luo *LoanUpdateOne) SetUser(u *User) *LoanUpdateOne {
	return luo.SetUserID(u.ID)
}

// Mutation returns the LoanMutation object of the builder.
func (luo *LoanUpdateOne) Mutation() *LoanMutation {
	return luo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LoanUpdateOne) ClearUser() *LoanUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// Where appends a list predicates to the LoanUpdate builder.
func (luo *LoanUpdateOne) Where(ps ...predicate.Loan) *LoanUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LoanUpdateOne) Select(field string, fields ...string) *LoanUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Loan entity.
func (luo *LoanUpdateOne) Save(ctx context.Context) (*Loan, error) {
	return withHooks[*Loan, LoanMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LoanUpdateOne) SaveX(ctx context.Context) *Loan {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LoanUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LoanUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LoanUpdateOne) sqlSave(ctx context.Context) (_node *Loan, err error) {
	_spec := sqlgraph.NewUpdateSpec(loan.Table, loan.Columns, sqlgraph.NewFieldSpec(loan.FieldID, field.TypeUint))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Loan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loan.FieldID)
		for _, f := range fields {
			if !loan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != loan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.AssetInfoID(); ok {
		_spec.SetField(loan.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := luo.mutation.AddedAssetInfoID(); ok {
		_spec.AddField(loan.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(loan.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(loan.FieldUpdatedAt, field.TypeTime, value)
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Loan{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
