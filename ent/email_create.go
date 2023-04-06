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
	"github.com/hellohq/hqservice/ent/identity"
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/user"
)

// EmailCreate is the builder for creating a Email entity.
type EmailCreate struct {
	config
	mutation *EmailMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ec *EmailCreate) SetUserID(u uuid.UUID) *EmailCreate {
	ec.mutation.SetUserID(u)
	return ec
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ec *EmailCreate) SetNillableUserID(u *uuid.UUID) *EmailCreate {
	if u != nil {
		ec.SetUserID(*u)
	}
	return ec
}

// SetAddress sets the "address" field.
func (ec *EmailCreate) SetAddress(s string) *EmailCreate {
	ec.mutation.SetAddress(s)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EmailCreate) SetCreatedAt(t time.Time) *EmailCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EmailCreate) SetNillableCreatedAt(t *time.Time) *EmailCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EmailCreate) SetUpdatedAt(t time.Time) *EmailCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EmailCreate) SetNillableUpdatedAt(t *time.Time) *EmailCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EmailCreate) SetID(u uuid.UUID) *EmailCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EmailCreate) SetNillableID(u *uuid.UUID) *EmailCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// SetUser sets the "user" edge to the User entity.
func (ec *EmailCreate) SetUser(u *User) *EmailCreate {
	return ec.SetUserID(u.ID)
}

// AddIdentityIDs adds the "identities" edge to the Identity entity by IDs.
func (ec *EmailCreate) AddIdentityIDs(ids ...uuid.UUID) *EmailCreate {
	ec.mutation.AddIdentityIDs(ids...)
	return ec
}

// AddIdentities adds the "identities" edges to the Identity entity.
func (ec *EmailCreate) AddIdentities(i ...*Identity) *EmailCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ec.AddIdentityIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (ec *EmailCreate) AddPasscodeIDs(ids ...uuid.UUID) *EmailCreate {
	ec.mutation.AddPasscodeIDs(ids...)
	return ec
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (ec *EmailCreate) AddPasscodes(p ...*Passcode) *EmailCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddPasscodeIDs(ids...)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (ec *EmailCreate) SetPrimaryEmailID(id uuid.UUID) *EmailCreate {
	ec.mutation.SetPrimaryEmailID(id)
	return ec
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (ec *EmailCreate) SetNillablePrimaryEmailID(id *uuid.UUID) *EmailCreate {
	if id != nil {
		ec = ec.SetPrimaryEmailID(*id)
	}
	return ec
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (ec *EmailCreate) SetPrimaryEmail(p *PrimaryEmail) *EmailCreate {
	return ec.SetPrimaryEmailID(p.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (ec *EmailCreate) Mutation() *EmailMutation {
	return ec.mutation
}

// Save creates the Email in the database.
func (ec *EmailCreate) Save(ctx context.Context) (*Email, error) {
	ec.defaults()
	return withHooks[*Email, EmailMutation](ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmailCreate) SaveX(ctx context.Context) *Email {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EmailCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EmailCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EmailCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := email.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := email.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := email.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EmailCreate) check() error {
	if _, ok := ec.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Email.address"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Email.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Email.updated_at"`)}
	}
	return nil
}

func (ec *EmailCreate) sqlSave(ctx context.Context) (*Email, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EmailCreate) createSpec() (*Email, *sqlgraph.CreateSpec) {
	var (
		_node = &Email{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(email.Table, sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Address(); ok {
		_spec.SetField(email.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(email.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(email.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.UserTable,
			Columns: []string{email.UserColumn},
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
	if nodes := ec.mutation.IdentitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   email.IdentitiesTable,
			Columns: []string{email.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: identity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PasscodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   email.PasscodesTable,
			Columns: []string{email.PasscodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passcode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   email.PrimaryEmailTable,
			Columns: []string{email.PrimaryEmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: primaryemail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmailCreateBulk is the builder for creating many Email entities in bulk.
type EmailCreateBulk struct {
	config
	builders []*EmailCreate
}

// Save creates the Email entities in the database.
func (ecb *EmailCreateBulk) Save(ctx context.Context) ([]*Email, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Email, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EmailCreateBulk) SaveX(ctx context.Context) []*Email {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EmailCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EmailCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}