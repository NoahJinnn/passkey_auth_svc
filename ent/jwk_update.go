// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NoahJinnn/passkey_auth_svc/ent/jwk"
	"github.com/NoahJinnn/passkey_auth_svc/ent/predicate"
)

// JwkUpdate is the builder for updating Jwk entities.
type JwkUpdate struct {
	config
	hooks    []Hook
	mutation *JwkMutation
}

// Where appends a list predicates to the JwkUpdate builder.
func (ju *JwkUpdate) Where(ps ...predicate.Jwk) *JwkUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetKeyData sets the "key_data" field.
func (ju *JwkUpdate) SetKeyData(s string) *JwkUpdate {
	ju.mutation.SetKeyData(s)
	return ju
}

// Mutation returns the JwkMutation object of the builder.
func (ju *JwkUpdate) Mutation() *JwkMutation {
	return ju.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JwkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ju.sqlSave, ju.mutation, ju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JwkUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JwkUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JwkUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ju *JwkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jwk.Table, jwk.Columns, sqlgraph.NewFieldSpec(jwk.FieldID, field.TypeUint))
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.KeyData(); ok {
		_spec.SetField(jwk.FieldKeyData, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jwk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ju.mutation.done = true
	return n, nil
}

// JwkUpdateOne is the builder for updating a single Jwk entity.
type JwkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JwkMutation
}

// SetKeyData sets the "key_data" field.
func (juo *JwkUpdateOne) SetKeyData(s string) *JwkUpdateOne {
	juo.mutation.SetKeyData(s)
	return juo
}

// Mutation returns the JwkMutation object of the builder.
func (juo *JwkUpdateOne) Mutation() *JwkMutation {
	return juo.mutation
}

// Where appends a list predicates to the JwkUpdate builder.
func (juo *JwkUpdateOne) Where(ps ...predicate.Jwk) *JwkUpdateOne {
	juo.mutation.Where(ps...)
	return juo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JwkUpdateOne) Select(field string, fields ...string) *JwkUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Jwk entity.
func (juo *JwkUpdateOne) Save(ctx context.Context) (*Jwk, error) {
	return withHooks(ctx, juo.sqlSave, juo.mutation, juo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JwkUpdateOne) SaveX(ctx context.Context) *Jwk {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JwkUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JwkUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (juo *JwkUpdateOne) sqlSave(ctx context.Context) (_node *Jwk, err error) {
	_spec := sqlgraph.NewUpdateSpec(jwk.Table, jwk.Columns, sqlgraph.NewFieldSpec(jwk.FieldID, field.TypeUint))
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Jwk.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jwk.FieldID)
		for _, f := range fields {
			if !jwk.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jwk.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := juo.mutation.KeyData(); ok {
		_spec.SetField(jwk.FieldKeyData, field.TypeString, value)
	}
	_node = &Jwk{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jwk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	juo.mutation.done = true
	return _node, nil
}
