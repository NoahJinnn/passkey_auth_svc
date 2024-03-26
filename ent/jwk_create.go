// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NoahJinnn/passkey_auth_svc/ent/jwk"
)

// JwkCreate is the builder for creating a Jwk entity.
type JwkCreate struct {
	config
	mutation *JwkMutation
	hooks    []Hook
}

// SetKeyData sets the "key_data" field.
func (jc *JwkCreate) SetKeyData(s string) *JwkCreate {
	jc.mutation.SetKeyData(s)
	return jc
}

// SetCreatedAt sets the "created_at" field.
func (jc *JwkCreate) SetCreatedAt(t time.Time) *JwkCreate {
	jc.mutation.SetCreatedAt(t)
	return jc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jc *JwkCreate) SetNillableCreatedAt(t *time.Time) *JwkCreate {
	if t != nil {
		jc.SetCreatedAt(*t)
	}
	return jc
}

// SetID sets the "id" field.
func (jc *JwkCreate) SetID(u uint) *JwkCreate {
	jc.mutation.SetID(u)
	return jc
}

// Mutation returns the JwkMutation object of the builder.
func (jc *JwkCreate) Mutation() *JwkMutation {
	return jc.mutation
}

// Save creates the Jwk in the database.
func (jc *JwkCreate) Save(ctx context.Context) (*Jwk, error) {
	jc.defaults()
	return withHooks(ctx, jc.sqlSave, jc.mutation, jc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JwkCreate) SaveX(ctx context.Context) *Jwk {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JwkCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JwkCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jc *JwkCreate) defaults() {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		v := jwk.DefaultCreatedAt()
		jc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JwkCreate) check() error {
	if _, ok := jc.mutation.KeyData(); !ok {
		return &ValidationError{Name: "key_data", err: errors.New(`ent: missing required field "Jwk.key_data"`)}
	}
	if _, ok := jc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Jwk.created_at"`)}
	}
	return nil
}

func (jc *JwkCreate) sqlSave(ctx context.Context) (*Jwk, error) {
	if err := jc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	jc.mutation.id = &_node.ID
	jc.mutation.done = true
	return _node, nil
}

func (jc *JwkCreate) createSpec() (*Jwk, *sqlgraph.CreateSpec) {
	var (
		_node = &Jwk{config: jc.config}
		_spec = sqlgraph.NewCreateSpec(jwk.Table, sqlgraph.NewFieldSpec(jwk.FieldID, field.TypeUint))
	)
	if id, ok := jc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jc.mutation.KeyData(); ok {
		_spec.SetField(jwk.FieldKeyData, field.TypeString, value)
		_node.KeyData = value
	}
	if value, ok := jc.mutation.CreatedAt(); ok {
		_spec.SetField(jwk.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// JwkCreateBulk is the builder for creating many Jwk entities in bulk.
type JwkCreateBulk struct {
	config
	builders []*JwkCreate
}

// Save creates the Jwk entities in the database.
func (jcb *JwkCreateBulk) Save(ctx context.Context) ([]*Jwk, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Jwk, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JwkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JwkCreateBulk) SaveX(ctx context.Context) []*Jwk {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JwkCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JwkCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}
