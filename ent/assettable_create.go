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
	"github.com/hellohq/hqservice/ent/assettable"
	"github.com/hellohq/hqservice/ent/user"
)

// AssetTableCreate is the builder for creating a AssetTable entity.
type AssetTableCreate struct {
	config
	mutation *AssetTableMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (atc *AssetTableCreate) SetUserID(u uuid.UUID) *AssetTableCreate {
	atc.mutation.SetUserID(u)
	return atc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableUserID(u *uuid.UUID) *AssetTableCreate {
	if u != nil {
		atc.SetUserID(*u)
	}
	return atc
}

// SetSheet sets the "sheet" field.
func (atc *AssetTableCreate) SetSheet(i int32) *AssetTableCreate {
	atc.mutation.SetSheet(i)
	return atc
}

// SetNillableSheet sets the "sheet" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableSheet(i *int32) *AssetTableCreate {
	if i != nil {
		atc.SetSheet(*i)
	}
	return atc
}

// SetSection sets the "section" field.
func (atc *AssetTableCreate) SetSection(i int32) *AssetTableCreate {
	atc.mutation.SetSection(i)
	return atc
}

// SetNillableSection sets the "section" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableSection(i *int32) *AssetTableCreate {
	if i != nil {
		atc.SetSection(*i)
	}
	return atc
}

// SetDescription sets the "description" field.
func (atc *AssetTableCreate) SetDescription(s string) *AssetTableCreate {
	atc.mutation.SetDescription(s)
	return atc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableDescription(s *string) *AssetTableCreate {
	if s != nil {
		atc.SetDescription(*s)
	}
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *AssetTableCreate) SetCreatedAt(t time.Time) *AssetTableCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableCreatedAt(t *time.Time) *AssetTableCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *AssetTableCreate) SetUpdatedAt(t time.Time) *AssetTableCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableUpdatedAt(t *time.Time) *AssetTableCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *AssetTableCreate) SetID(u uuid.UUID) *AssetTableCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *AssetTableCreate) SetNillableID(u *uuid.UUID) *AssetTableCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *AssetTableCreate) SetUser(u *User) *AssetTableCreate {
	return atc.SetUserID(u.ID)
}

// Mutation returns the AssetTableMutation object of the builder.
func (atc *AssetTableCreate) Mutation() *AssetTableMutation {
	return atc.mutation
}

// Save creates the AssetTable in the database.
func (atc *AssetTableCreate) Save(ctx context.Context) (*AssetTable, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AssetTableCreate) SaveX(ctx context.Context) *AssetTable {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AssetTableCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AssetTableCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AssetTableCreate) defaults() {
	if _, ok := atc.mutation.Sheet(); !ok {
		v := assettable.DefaultSheet
		atc.mutation.SetSheet(v)
	}
	if _, ok := atc.mutation.Section(); !ok {
		v := assettable.DefaultSection
		atc.mutation.SetSection(v)
	}
	if _, ok := atc.mutation.Description(); !ok {
		v := assettable.DefaultDescription
		atc.mutation.SetDescription(v)
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := assettable.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := assettable.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := assettable.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AssetTableCreate) check() error {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AssetTable.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AssetTable.updated_at"`)}
	}
	return nil
}

func (atc *AssetTableCreate) sqlSave(ctx context.Context) (*AssetTable, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
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
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AssetTableCreate) createSpec() (*AssetTable, *sqlgraph.CreateSpec) {
	var (
		_node = &AssetTable{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(assettable.Table, sqlgraph.NewFieldSpec(assettable.FieldID, field.TypeUUID))
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.Sheet(); ok {
		_spec.SetField(assettable.FieldSheet, field.TypeInt32, value)
		_node.Sheet = value
	}
	if value, ok := atc.mutation.Section(); ok {
		_spec.SetField(assettable.FieldSection, field.TypeInt32, value)
		_node.Section = value
	}
	if value, ok := atc.mutation.Description(); ok {
		_spec.SetField(assettable.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(assettable.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.SetField(assettable.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assettable.UserTable,
			Columns: []string{assettable.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssetTableCreateBulk is the builder for creating many AssetTable entities in bulk.
type AssetTableCreateBulk struct {
	config
	builders []*AssetTableCreate
}

// Save creates the AssetTable entities in the database.
func (atcb *AssetTableCreateBulk) Save(ctx context.Context) ([]*AssetTable, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AssetTable, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetTableMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AssetTableCreateBulk) SaveX(ctx context.Context) []*AssetTable {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AssetTableCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AssetTableCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}