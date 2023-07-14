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
	"github.com/hellohq/hqservice/ent/finitemtable"
	"github.com/hellohq/hqservice/ent/fvsession"
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddEmailIDs adds the "emails" edge to the Email entity by IDs.
func (uc *UserCreate) AddEmailIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddEmailIDs(ids...)
	return uc
}

// AddEmails adds the "emails" edges to the Email entity.
func (uc *UserCreate) AddEmails(e ...*Email) *UserCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEmailIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (uc *UserCreate) AddPasscodeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddPasscodeIDs(ids...)
	return uc
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (uc *UserCreate) AddPasscodes(p ...*Passcode) *UserCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPasscodeIDs(ids...)
}

// AddWebauthnCredentialIDs adds the "webauthn_credentials" edge to the WebauthnCredential entity by IDs.
func (uc *UserCreate) AddWebauthnCredentialIDs(ids ...string) *UserCreate {
	uc.mutation.AddWebauthnCredentialIDs(ids...)
	return uc
}

// AddWebauthnCredentials adds the "webauthn_credentials" edges to the WebauthnCredential entity.
func (uc *UserCreate) AddWebauthnCredentials(w ...*WebauthnCredential) *UserCreate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddWebauthnCredentialIDs(ids...)
}

// AddFinItemTableIDs adds the "fin_item_tables" edge to the FinItemTable entity by IDs.
func (uc *UserCreate) AddFinItemTableIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddFinItemTableIDs(ids...)
	return uc
}

// AddFinItemTables adds the "fin_item_tables" edges to the FinItemTable entity.
func (uc *UserCreate) AddFinItemTables(f ...*FinItemTable) *UserCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddFinItemTableIDs(ids...)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (uc *UserCreate) SetPrimaryEmailID(id uuid.UUID) *UserCreate {
	uc.mutation.SetPrimaryEmailID(id)
	return uc
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillablePrimaryEmailID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetPrimaryEmailID(*id)
	}
	return uc
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (uc *UserCreate) SetPrimaryEmail(p *PrimaryEmail) *UserCreate {
	return uc.SetPrimaryEmailID(p.ID)
}

// SetFvSessionID sets the "fv_session" edge to the FvSession entity by ID.
func (uc *UserCreate) SetFvSessionID(id uuid.UUID) *UserCreate {
	uc.mutation.SetFvSessionID(id)
	return uc
}

// SetNillableFvSessionID sets the "fv_session" edge to the FvSession entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableFvSessionID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetFvSessionID(*id)
	}
	return uc
}

// SetFvSession sets the "fv_session" edge to the FvSession entity.
func (uc *UserCreate) SetFvSession(f *FvSession) *UserCreate {
	return uc.SetFvSessionID(f.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.EmailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PasscodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.WebauthnCredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FinItemTablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FinItemTablesTable,
			Columns: []string{user.FinItemTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(finitemtable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PrimaryEmailTable,
			Columns: []string{user.PrimaryEmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(primaryemail.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FvSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FvSessionTable,
			Columns: []string{user.FvSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fvsession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
