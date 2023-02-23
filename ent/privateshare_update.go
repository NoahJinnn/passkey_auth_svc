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
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/privateshare"
	"github.com/hellohq/hqservice/ent/user"
)

// PrivateShareUpdate is the builder for updating PrivateShare entities.
type PrivateShareUpdate struct {
	config
	hooks    []Hook
	mutation *PrivateShareMutation
}

// Where appends a list predicates to the PrivateShareUpdate builder.
func (psu *PrivateShareUpdate) Where(ps ...predicate.PrivateShare) *PrivateShareUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetUserID sets the "user_id" field.
func (psu *PrivateShareUpdate) SetUserID(u uint) *PrivateShareUpdate {
	psu.mutation.SetUserID(u)
	return psu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (psu *PrivateShareUpdate) SetNillableUserID(u *uint) *PrivateShareUpdate {
	if u != nil {
		psu.SetUserID(*u)
	}
	return psu
}

// ClearUserID clears the value of the "user_id" field.
func (psu *PrivateShareUpdate) ClearUserID() *PrivateShareUpdate {
	psu.mutation.ClearUserID()
	return psu
}

// SetAssetInfoID sets the "asset_info_id" field.
func (psu *PrivateShareUpdate) SetAssetInfoID(u uint) *PrivateShareUpdate {
	psu.mutation.ResetAssetInfoID()
	psu.mutation.SetAssetInfoID(u)
	return psu
}

// AddAssetInfoID adds u to the "asset_info_id" field.
func (psu *PrivateShareUpdate) AddAssetInfoID(u int) *PrivateShareUpdate {
	psu.mutation.AddAssetInfoID(u)
	return psu
}

// SetCreatedAt sets the "created_at" field.
func (psu *PrivateShareUpdate) SetCreatedAt(t time.Time) *PrivateShareUpdate {
	psu.mutation.SetCreatedAt(t)
	return psu
}

// SetUpdatedAt sets the "updated_at" field.
func (psu *PrivateShareUpdate) SetUpdatedAt(t time.Time) *PrivateShareUpdate {
	psu.mutation.SetUpdatedAt(t)
	return psu
}

// SetUser sets the "user" edge to the User entity.
func (psu *PrivateShareUpdate) SetUser(u *User) *PrivateShareUpdate {
	return psu.SetUserID(u.ID)
}

// Mutation returns the PrivateShareMutation object of the builder.
func (psu *PrivateShareUpdate) Mutation() *PrivateShareMutation {
	return psu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (psu *PrivateShareUpdate) ClearUser() *PrivateShareUpdate {
	psu.mutation.ClearUser()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PrivateShareUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PrivateShareMutation](ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PrivateShareUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PrivateShareUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PrivateShareUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psu *PrivateShareUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(privateshare.Table, privateshare.Columns, sqlgraph.NewFieldSpec(privateshare.FieldID, field.TypeUint))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.AssetInfoID(); ok {
		_spec.SetField(privateshare.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := psu.mutation.AddedAssetInfoID(); ok {
		_spec.AddField(privateshare.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := psu.mutation.CreatedAt(); ok {
		_spec.SetField(privateshare.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := psu.mutation.UpdatedAt(); ok {
		_spec.SetField(privateshare.FieldUpdatedAt, field.TypeTime, value)
	}
	if psu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privateshare.UserTable,
			Columns: []string{privateshare.UserColumn},
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
	if nodes := psu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privateshare.UserTable,
			Columns: []string{privateshare.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{privateshare.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PrivateShareUpdateOne is the builder for updating a single PrivateShare entity.
type PrivateShareUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrivateShareMutation
}

// SetUserID sets the "user_id" field.
func (psuo *PrivateShareUpdateOne) SetUserID(u uint) *PrivateShareUpdateOne {
	psuo.mutation.SetUserID(u)
	return psuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (psuo *PrivateShareUpdateOne) SetNillableUserID(u *uint) *PrivateShareUpdateOne {
	if u != nil {
		psuo.SetUserID(*u)
	}
	return psuo
}

// ClearUserID clears the value of the "user_id" field.
func (psuo *PrivateShareUpdateOne) ClearUserID() *PrivateShareUpdateOne {
	psuo.mutation.ClearUserID()
	return psuo
}

// SetAssetInfoID sets the "asset_info_id" field.
func (psuo *PrivateShareUpdateOne) SetAssetInfoID(u uint) *PrivateShareUpdateOne {
	psuo.mutation.ResetAssetInfoID()
	psuo.mutation.SetAssetInfoID(u)
	return psuo
}

// AddAssetInfoID adds u to the "asset_info_id" field.
func (psuo *PrivateShareUpdateOne) AddAssetInfoID(u int) *PrivateShareUpdateOne {
	psuo.mutation.AddAssetInfoID(u)
	return psuo
}

// SetCreatedAt sets the "created_at" field.
func (psuo *PrivateShareUpdateOne) SetCreatedAt(t time.Time) *PrivateShareUpdateOne {
	psuo.mutation.SetCreatedAt(t)
	return psuo
}

// SetUpdatedAt sets the "updated_at" field.
func (psuo *PrivateShareUpdateOne) SetUpdatedAt(t time.Time) *PrivateShareUpdateOne {
	psuo.mutation.SetUpdatedAt(t)
	return psuo
}

// SetUser sets the "user" edge to the User entity.
func (psuo *PrivateShareUpdateOne) SetUser(u *User) *PrivateShareUpdateOne {
	return psuo.SetUserID(u.ID)
}

// Mutation returns the PrivateShareMutation object of the builder.
func (psuo *PrivateShareUpdateOne) Mutation() *PrivateShareMutation {
	return psuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (psuo *PrivateShareUpdateOne) ClearUser() *PrivateShareUpdateOne {
	psuo.mutation.ClearUser()
	return psuo
}

// Where appends a list predicates to the PrivateShareUpdate builder.
func (psuo *PrivateShareUpdateOne) Where(ps ...predicate.PrivateShare) *PrivateShareUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PrivateShareUpdateOne) Select(field string, fields ...string) *PrivateShareUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PrivateShare entity.
func (psuo *PrivateShareUpdateOne) Save(ctx context.Context) (*PrivateShare, error) {
	return withHooks[*PrivateShare, PrivateShareMutation](ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PrivateShareUpdateOne) SaveX(ctx context.Context) *PrivateShare {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PrivateShareUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PrivateShareUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psuo *PrivateShareUpdateOne) sqlSave(ctx context.Context) (_node *PrivateShare, err error) {
	_spec := sqlgraph.NewUpdateSpec(privateshare.Table, privateshare.Columns, sqlgraph.NewFieldSpec(privateshare.FieldID, field.TypeUint))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PrivateShare.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, privateshare.FieldID)
		for _, f := range fields {
			if !privateshare.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != privateshare.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.AssetInfoID(); ok {
		_spec.SetField(privateshare.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := psuo.mutation.AddedAssetInfoID(); ok {
		_spec.AddField(privateshare.FieldAssetInfoID, field.TypeUint, value)
	}
	if value, ok := psuo.mutation.CreatedAt(); ok {
		_spec.SetField(privateshare.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := psuo.mutation.UpdatedAt(); ok {
		_spec.SetField(privateshare.FieldUpdatedAt, field.TypeTime, value)
	}
	if psuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privateshare.UserTable,
			Columns: []string{privateshare.UserColumn},
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
	if nodes := psuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privateshare.UserTable,
			Columns: []string{privateshare.UserColumn},
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
	_node = &PrivateShare{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{privateshare.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}
