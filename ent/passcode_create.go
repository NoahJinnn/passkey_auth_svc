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
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/user"
)

// PasscodeCreate is the builder for creating a Passcode entity.
type PasscodeCreate struct {
	config
	mutation *PasscodeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pc *PasscodeCreate) SetUserID(u uuid.UUID) *PasscodeCreate {
	pc.mutation.SetUserID(u)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *PasscodeCreate) SetNillableUserID(u *uuid.UUID) *PasscodeCreate {
	if u != nil {
		pc.SetUserID(*u)
	}
	return pc
}

// SetTTL sets the "ttl" field.
func (pc *PasscodeCreate) SetTTL(i int32) *PasscodeCreate {
	pc.mutation.SetTTL(i)
	return pc
}

// SetCode sets the "code" field.
func (pc *PasscodeCreate) SetCode(s string) *PasscodeCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetTryCount sets the "try_count" field.
func (pc *PasscodeCreate) SetTryCount(i int32) *PasscodeCreate {
	pc.mutation.SetTryCount(i)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PasscodeCreate) SetCreatedAt(t time.Time) *PasscodeCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PasscodeCreate) SetNillableCreatedAt(t *time.Time) *PasscodeCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PasscodeCreate) SetUpdatedAt(t time.Time) *PasscodeCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PasscodeCreate) SetNillableUpdatedAt(t *time.Time) *PasscodeCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetEmailID sets the "email_id" field.
func (pc *PasscodeCreate) SetEmailID(u uuid.UUID) *PasscodeCreate {
	pc.mutation.SetEmailID(u)
	return pc
}

// SetNillableEmailID sets the "email_id" field if the given value is not nil.
func (pc *PasscodeCreate) SetNillableEmailID(u *uuid.UUID) *PasscodeCreate {
	if u != nil {
		pc.SetEmailID(*u)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PasscodeCreate) SetID(u uuid.UUID) *PasscodeCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PasscodeCreate) SetNillableID(u *uuid.UUID) *PasscodeCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetEmail sets the "email" edge to the Email entity.
func (pc *PasscodeCreate) SetEmail(e *Email) *PasscodeCreate {
	return pc.SetEmailID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pc *PasscodeCreate) SetUser(u *User) *PasscodeCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PasscodeMutation object of the builder.
func (pc *PasscodeCreate) Mutation() *PasscodeMutation {
	return pc.mutation
}

// Save creates the Passcode in the database.
func (pc *PasscodeCreate) Save(ctx context.Context) (*Passcode, error) {
	pc.defaults()
	return withHooks[*Passcode, PasscodeMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PasscodeCreate) SaveX(ctx context.Context) *Passcode {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PasscodeCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PasscodeCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PasscodeCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := passcode.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := passcode.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := passcode.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PasscodeCreate) check() error {
	if _, ok := pc.mutation.TTL(); !ok {
		return &ValidationError{Name: "ttl", err: errors.New(`ent: missing required field "Passcode.ttl"`)}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Passcode.code"`)}
	}
	if _, ok := pc.mutation.TryCount(); !ok {
		return &ValidationError{Name: "try_count", err: errors.New(`ent: missing required field "Passcode.try_count"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Passcode.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Passcode.updated_at"`)}
	}
	return nil
}

func (pc *PasscodeCreate) sqlSave(ctx context.Context) (*Passcode, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PasscodeCreate) createSpec() (*Passcode, *sqlgraph.CreateSpec) {
	var (
		_node = &Passcode{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(passcode.Table, sqlgraph.NewFieldSpec(passcode.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.TTL(); ok {
		_spec.SetField(passcode.FieldTTL, field.TypeInt32, value)
		_node.TTL = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(passcode.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.TryCount(); ok {
		_spec.SetField(passcode.FieldTryCount, field.TypeInt32, value)
		_node.TryCount = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(passcode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(passcode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.EmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.EmailTable,
			Columns: []string{passcode.EmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: email.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EmailID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.UserTable,
			Columns: []string{passcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
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

// PasscodeCreateBulk is the builder for creating many Passcode entities in bulk.
type PasscodeCreateBulk struct {
	config
	builders []*PasscodeCreate
}

// Save creates the Passcode entities in the database.
func (pcb *PasscodeCreateBulk) Save(ctx context.Context) ([]*Passcode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Passcode, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasscodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PasscodeCreateBulk) SaveX(ctx context.Context) []*Passcode {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PasscodeCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PasscodeCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
