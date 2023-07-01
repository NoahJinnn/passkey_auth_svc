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
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/provider"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderMutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *ProviderUpdate) SetUserID(u uuid.UUID) *ProviderUpdate {
	pu.mutation.SetUserID(u)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableUserID(u *uuid.UUID) *ProviderUpdate {
	if u != nil {
		pu.SetUserID(*u)
	}
	return pu
}

// ClearUserID clears the value of the "user_id" field.
func (pu *ProviderUpdate) ClearUserID() *ProviderUpdate {
	pu.mutation.ClearUserID()
	return pu
}

// SetVerified sets the "verified" field.
func (pu *ProviderUpdate) SetVerified(b bool) *ProviderUpdate {
	pu.mutation.SetVerified(b)
	return pu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableVerified(b *bool) *ProviderUpdate {
	if b != nil {
		pu.SetVerified(*b)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProviderUpdate) SetUpdatedAt(t time.Time) *ProviderUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// Mutation returns the ProviderMutation object of the builder.
func (pu *ProviderUpdate) Mutation() *ProviderMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProviderUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := provider.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(provider.FieldUserID, field.TypeUUID, value)
	}
	if pu.mutation.UserIDCleared() {
		_spec.ClearField(provider.FieldUserID, field.TypeUUID)
	}
	if value, ok := pu.mutation.Verified(); ok {
		_spec.SetField(provider.FieldVerified, field.TypeBool, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(provider.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderMutation
}

// SetUserID sets the "user_id" field.
func (puo *ProviderUpdateOne) SetUserID(u uuid.UUID) *ProviderUpdateOne {
	puo.mutation.SetUserID(u)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableUserID(u *uuid.UUID) *ProviderUpdateOne {
	if u != nil {
		puo.SetUserID(*u)
	}
	return puo
}

// ClearUserID clears the value of the "user_id" field.
func (puo *ProviderUpdateOne) ClearUserID() *ProviderUpdateOne {
	puo.mutation.ClearUserID()
	return puo
}

// SetVerified sets the "verified" field.
func (puo *ProviderUpdateOne) SetVerified(b bool) *ProviderUpdateOne {
	puo.mutation.SetVerified(b)
	return puo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableVerified(b *bool) *ProviderUpdateOne {
	if b != nil {
		puo.SetVerified(*b)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProviderUpdateOne) SetUpdatedAt(t time.Time) *ProviderUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// Mutation returns the ProviderMutation object of the builder.
func (puo *ProviderUpdateOne) Mutation() *ProviderMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (puo *ProviderUpdateOne) Where(ps ...predicate.Provider) *ProviderUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProviderUpdateOne) Select(field string, fields ...string) *ProviderUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provider entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProviderUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := provider.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (_node *Provider, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Provider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provider.FieldID)
		for _, f := range fields {
			if !provider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(provider.FieldUserID, field.TypeUUID, value)
	}
	if puo.mutation.UserIDCleared() {
		_spec.ClearField(provider.FieldUserID, field.TypeUUID)
	}
	if value, ok := puo.mutation.Verified(); ok {
		_spec.SetField(provider.FieldVerified, field.TypeBool, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(provider.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Provider{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
