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
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/asset"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
)

// InstitutionCreate is the builder for creating a Institution entity.
type InstitutionCreate struct {
	config
	mutation *InstitutionMutation
	hooks    []Hook
}

// SetProviderName sets the "provider_name" field.
func (ic *InstitutionCreate) SetProviderName(s string) *InstitutionCreate {
	ic.mutation.SetProviderName(s)
	return ic
}

// SetData sets the "data" field.
func (ic *InstitutionCreate) SetData(s string) *InstitutionCreate {
	ic.mutation.SetData(s)
	return ic
}

// SetNillableData sets the "data" field if the given value is not nil.
func (ic *InstitutionCreate) SetNillableData(s *string) *InstitutionCreate {
	if s != nil {
		ic.SetData(*s)
	}
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *InstitutionCreate) SetCreatedAt(t time.Time) *InstitutionCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InstitutionCreate) SetNillableCreatedAt(t *time.Time) *InstitutionCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InstitutionCreate) SetUpdatedAt(t time.Time) *InstitutionCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InstitutionCreate) SetNillableUpdatedAt(t *time.Time) *InstitutionCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InstitutionCreate) SetID(u uuid.UUID) *InstitutionCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InstitutionCreate) SetNillableID(u *uuid.UUID) *InstitutionCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetConnectionID sets the "connection" edge to the Connection entity by ID.
func (ic *InstitutionCreate) SetConnectionID(id uuid.UUID) *InstitutionCreate {
	ic.mutation.SetConnectionID(id)
	return ic
}

// SetNillableConnectionID sets the "connection" edge to the Connection entity by ID if the given value is not nil.
func (ic *InstitutionCreate) SetNillableConnectionID(id *uuid.UUID) *InstitutionCreate {
	if id != nil {
		ic = ic.SetConnectionID(*id)
	}
	return ic
}

// SetConnection sets the "connection" edge to the Connection entity.
func (ic *InstitutionCreate) SetConnection(c *Connection) *InstitutionCreate {
	return ic.SetConnectionID(c.ID)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (ic *InstitutionCreate) AddAccountIDs(ids ...uuid.UUID) *InstitutionCreate {
	ic.mutation.AddAccountIDs(ids...)
	return ic
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (ic *InstitutionCreate) AddAccounts(a ...*Account) *InstitutionCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAccountIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (ic *InstitutionCreate) AddAssetIDs(ids ...uuid.UUID) *InstitutionCreate {
	ic.mutation.AddAssetIDs(ids...)
	return ic
}

// AddAssets adds the "assets" edges to the Asset entity.
func (ic *InstitutionCreate) AddAssets(a ...*Asset) *InstitutionCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAssetIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (ic *InstitutionCreate) Mutation() *InstitutionMutation {
	return ic.mutation
}

// Save creates the Institution in the database.
func (ic *InstitutionCreate) Save(ctx context.Context) (*Institution, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstitutionCreate) SaveX(ctx context.Context) *Institution {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InstitutionCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InstitutionCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InstitutionCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := institution.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := institution.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := institution.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InstitutionCreate) check() error {
	if _, ok := ic.mutation.ProviderName(); !ok {
		return &ValidationError{Name: "provider_name", err: errors.New(`ent: missing required field "Institution.provider_name"`)}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Institution.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Institution.updated_at"`)}
	}
	return nil
}

func (ic *InstitutionCreate) sqlSave(ctx context.Context) (*Institution, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InstitutionCreate) createSpec() (*Institution, *sqlgraph.CreateSpec) {
	var (
		_node = &Institution{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(institution.Table, sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.ProviderName(); ok {
		_spec.SetField(institution.FieldProviderName, field.TypeString, value)
		_node.ProviderName = value
	}
	if value, ok := ic.mutation.Data(); ok {
		_spec.SetField(institution.FieldData, field.TypeString, value)
		_node.Data = &value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(institution.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(institution.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ic.mutation.ConnectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   institution.ConnectionTable,
			Columns: []string{institution.ConnectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(connection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccountsTable,
			Columns: []string{institution.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AssetsTable,
			Columns: []string{institution.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstitutionCreateBulk is the builder for creating many Institution entities in bulk.
type InstitutionCreateBulk struct {
	config
	builders []*InstitutionCreate
}

// Save creates the Institution entities in the database.
func (icb *InstitutionCreateBulk) Save(ctx context.Context) ([]*Institution, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Institution, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstitutionMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InstitutionCreateBulk) SaveX(ctx context.Context) []*Institution {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InstitutionCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InstitutionCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}