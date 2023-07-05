// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/asset"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
)

// AssetCreate is the builder for creating a Asset entity.
type AssetCreate struct {
	config
	mutation *AssetMutation
	hooks    []Hook
}

// SetInstitutionID sets the "institution_id" field.
func (ac *AssetCreate) SetInstitutionID(u uuid.UUID) *AssetCreate {
	ac.mutation.SetInstitutionID(u)
	return ac
}

// SetNillableInstitutionID sets the "institution_id" field if the given value is not nil.
func (ac *AssetCreate) SetNillableInstitutionID(u *uuid.UUID) *AssetCreate {
	if u != nil {
		ac.SetInstitutionID(*u)
	}
	return ac
}

// SetData sets the "data" field.
func (ac *AssetCreate) SetData(s string) *AssetCreate {
	ac.mutation.SetData(s)
	return ac
}

// SetNillableData sets the "data" field if the given value is not nil.
func (ac *AssetCreate) SetNillableData(s *string) *AssetCreate {
	if s != nil {
		ac.SetData(*s)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AssetCreate) SetCreatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableCreatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AssetCreate) SetUpdatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableUpdatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AssetCreate) SetID(u uuid.UUID) *AssetCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AssetCreate) SetNillableID(u *uuid.UUID) *AssetCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (ac *AssetCreate) SetInstitution(i *Institution) *AssetCreate {
	return ac.SetInstitutionID(i.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (ac *AssetCreate) Mutation() *AssetMutation {
	return ac.mutation
}

// Save creates the Asset in the database.
func (ac *AssetCreate) Save(ctx context.Context) (*Asset, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssetCreate) SaveX(ctx context.Context) *Asset {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AssetCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AssetCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AssetCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := asset.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := asset.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := asset.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AssetCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Asset.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Asset.updated_at"`)}
	}
	return nil
}

func (ac *AssetCreate) sqlSave(ctx context.Context) (*Asset, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AssetCreate) createSpec() (*Asset, *sqlgraph.CreateSpec) {
	var (
		_node = &Asset{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(asset.Table, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Data(); ok {
		_spec.SetField(asset.FieldData, field.TypeString, value)
		_node.Data = &value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(asset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(asset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.InstitutionTable,
			Columns: []string{asset.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InstitutionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssetCreateBulk is the builder for creating many Asset entities in bulk.
type AssetCreateBulk struct {
	config
	builders []*AssetCreate
}

// Save creates the Asset entities in the database.
func (acb *AssetCreateBulk) Save(ctx context.Context) ([]*Asset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Asset, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AssetCreateBulk) SaveX(ctx context.Context) []*Asset {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AssetCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AssetCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}