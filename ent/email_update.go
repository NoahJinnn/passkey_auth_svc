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
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ent/identity"
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/user"
)

// EmailUpdate is the builder for updating Email entities.
type EmailUpdate struct {
	config
	hooks    []Hook
	mutation *EmailMutation
}

// Where appends a list predicates to the EmailUpdate builder.
func (eu *EmailUpdate) Where(ps ...predicate.Email) *EmailUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUserID sets the "user_id" field.
func (eu *EmailUpdate) SetUserID(u uuid.UUID) *EmailUpdate {
	eu.mutation.SetUserID(u)
	return eu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (eu *EmailUpdate) SetNillableUserID(u *uuid.UUID) *EmailUpdate {
	if u != nil {
		eu.SetUserID(*u)
	}
	return eu
}

// ClearUserID clears the value of the "user_id" field.
func (eu *EmailUpdate) ClearUserID() *EmailUpdate {
	eu.mutation.ClearUserID()
	return eu
}

// SetAddress sets the "address" field.
func (eu *EmailUpdate) SetAddress(s string) *EmailUpdate {
	eu.mutation.SetAddress(s)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EmailUpdate) SetUpdatedAt(t time.Time) *EmailUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetUser sets the "user" edge to the User entity.
func (eu *EmailUpdate) SetUser(u *User) *EmailUpdate {
	return eu.SetUserID(u.ID)
}

// AddIdentityIDs adds the "identities" edge to the Identity entity by IDs.
func (eu *EmailUpdate) AddIdentityIDs(ids ...uuid.UUID) *EmailUpdate {
	eu.mutation.AddIdentityIDs(ids...)
	return eu
}

// AddIdentities adds the "identities" edges to the Identity entity.
func (eu *EmailUpdate) AddIdentities(i ...*Identity) *EmailUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return eu.AddIdentityIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (eu *EmailUpdate) AddPasscodeIDs(ids ...uuid.UUID) *EmailUpdate {
	eu.mutation.AddPasscodeIDs(ids...)
	return eu
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (eu *EmailUpdate) AddPasscodes(p ...*Passcode) *EmailUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddPasscodeIDs(ids...)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (eu *EmailUpdate) SetPrimaryEmailID(id uuid.UUID) *EmailUpdate {
	eu.mutation.SetPrimaryEmailID(id)
	return eu
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (eu *EmailUpdate) SetNillablePrimaryEmailID(id *uuid.UUID) *EmailUpdate {
	if id != nil {
		eu = eu.SetPrimaryEmailID(*id)
	}
	return eu
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (eu *EmailUpdate) SetPrimaryEmail(p *PrimaryEmail) *EmailUpdate {
	return eu.SetPrimaryEmailID(p.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (eu *EmailUpdate) Mutation() *EmailMutation {
	return eu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (eu *EmailUpdate) ClearUser() *EmailUpdate {
	eu.mutation.ClearUser()
	return eu
}

// ClearIdentities clears all "identities" edges to the Identity entity.
func (eu *EmailUpdate) ClearIdentities() *EmailUpdate {
	eu.mutation.ClearIdentities()
	return eu
}

// RemoveIdentityIDs removes the "identities" edge to Identity entities by IDs.
func (eu *EmailUpdate) RemoveIdentityIDs(ids ...uuid.UUID) *EmailUpdate {
	eu.mutation.RemoveIdentityIDs(ids...)
	return eu
}

// RemoveIdentities removes "identities" edges to Identity entities.
func (eu *EmailUpdate) RemoveIdentities(i ...*Identity) *EmailUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return eu.RemoveIdentityIDs(ids...)
}

// ClearPasscodes clears all "passcodes" edges to the Passcode entity.
func (eu *EmailUpdate) ClearPasscodes() *EmailUpdate {
	eu.mutation.ClearPasscodes()
	return eu
}

// RemovePasscodeIDs removes the "passcodes" edge to Passcode entities by IDs.
func (eu *EmailUpdate) RemovePasscodeIDs(ids ...uuid.UUID) *EmailUpdate {
	eu.mutation.RemovePasscodeIDs(ids...)
	return eu
}

// RemovePasscodes removes "passcodes" edges to Passcode entities.
func (eu *EmailUpdate) RemovePasscodes(p ...*Passcode) *EmailUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemovePasscodeIDs(ids...)
}

// ClearPrimaryEmail clears the "primary_email" edge to the PrimaryEmail entity.
func (eu *EmailUpdate) ClearPrimaryEmail() *EmailUpdate {
	eu.mutation.ClearPrimaryEmail()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmailUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks[int, EmailMutation](ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmailUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmailUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmailUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EmailUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := email.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

func (eu *EmailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(email.Table, email.Columns, sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Address(); ok {
		_spec.SetField(email.FieldAddress, field.TypeString, value)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(email.FieldUpdatedAt, field.TypeTime, value)
	}
	if eu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.IdentitiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedIdentitiesIDs(); len(nodes) > 0 && !eu.mutation.IdentitiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.IdentitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.PasscodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedPasscodesIDs(); len(nodes) > 0 && !eu.mutation.PasscodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.PasscodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.PrimaryEmailCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{email.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmailUpdateOne is the builder for updating a single Email entity.
type EmailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailMutation
}

// SetUserID sets the "user_id" field.
func (euo *EmailUpdateOne) SetUserID(u uuid.UUID) *EmailUpdateOne {
	euo.mutation.SetUserID(u)
	return euo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (euo *EmailUpdateOne) SetNillableUserID(u *uuid.UUID) *EmailUpdateOne {
	if u != nil {
		euo.SetUserID(*u)
	}
	return euo
}

// ClearUserID clears the value of the "user_id" field.
func (euo *EmailUpdateOne) ClearUserID() *EmailUpdateOne {
	euo.mutation.ClearUserID()
	return euo
}

// SetAddress sets the "address" field.
func (euo *EmailUpdateOne) SetAddress(s string) *EmailUpdateOne {
	euo.mutation.SetAddress(s)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EmailUpdateOne) SetUpdatedAt(t time.Time) *EmailUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetUser sets the "user" edge to the User entity.
func (euo *EmailUpdateOne) SetUser(u *User) *EmailUpdateOne {
	return euo.SetUserID(u.ID)
}

// AddIdentityIDs adds the "identities" edge to the Identity entity by IDs.
func (euo *EmailUpdateOne) AddIdentityIDs(ids ...uuid.UUID) *EmailUpdateOne {
	euo.mutation.AddIdentityIDs(ids...)
	return euo
}

// AddIdentities adds the "identities" edges to the Identity entity.
func (euo *EmailUpdateOne) AddIdentities(i ...*Identity) *EmailUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return euo.AddIdentityIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (euo *EmailUpdateOne) AddPasscodeIDs(ids ...uuid.UUID) *EmailUpdateOne {
	euo.mutation.AddPasscodeIDs(ids...)
	return euo
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (euo *EmailUpdateOne) AddPasscodes(p ...*Passcode) *EmailUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddPasscodeIDs(ids...)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (euo *EmailUpdateOne) SetPrimaryEmailID(id uuid.UUID) *EmailUpdateOne {
	euo.mutation.SetPrimaryEmailID(id)
	return euo
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (euo *EmailUpdateOne) SetNillablePrimaryEmailID(id *uuid.UUID) *EmailUpdateOne {
	if id != nil {
		euo = euo.SetPrimaryEmailID(*id)
	}
	return euo
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (euo *EmailUpdateOne) SetPrimaryEmail(p *PrimaryEmail) *EmailUpdateOne {
	return euo.SetPrimaryEmailID(p.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (euo *EmailUpdateOne) Mutation() *EmailMutation {
	return euo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (euo *EmailUpdateOne) ClearUser() *EmailUpdateOne {
	euo.mutation.ClearUser()
	return euo
}

// ClearIdentities clears all "identities" edges to the Identity entity.
func (euo *EmailUpdateOne) ClearIdentities() *EmailUpdateOne {
	euo.mutation.ClearIdentities()
	return euo
}

// RemoveIdentityIDs removes the "identities" edge to Identity entities by IDs.
func (euo *EmailUpdateOne) RemoveIdentityIDs(ids ...uuid.UUID) *EmailUpdateOne {
	euo.mutation.RemoveIdentityIDs(ids...)
	return euo
}

// RemoveIdentities removes "identities" edges to Identity entities.
func (euo *EmailUpdateOne) RemoveIdentities(i ...*Identity) *EmailUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return euo.RemoveIdentityIDs(ids...)
}

// ClearPasscodes clears all "passcodes" edges to the Passcode entity.
func (euo *EmailUpdateOne) ClearPasscodes() *EmailUpdateOne {
	euo.mutation.ClearPasscodes()
	return euo
}

// RemovePasscodeIDs removes the "passcodes" edge to Passcode entities by IDs.
func (euo *EmailUpdateOne) RemovePasscodeIDs(ids ...uuid.UUID) *EmailUpdateOne {
	euo.mutation.RemovePasscodeIDs(ids...)
	return euo
}

// RemovePasscodes removes "passcodes" edges to Passcode entities.
func (euo *EmailUpdateOne) RemovePasscodes(p ...*Passcode) *EmailUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemovePasscodeIDs(ids...)
}

// ClearPrimaryEmail clears the "primary_email" edge to the PrimaryEmail entity.
func (euo *EmailUpdateOne) ClearPrimaryEmail() *EmailUpdateOne {
	euo.mutation.ClearPrimaryEmail()
	return euo
}

// Where appends a list predicates to the EmailUpdate builder.
func (euo *EmailUpdateOne) Where(ps ...predicate.Email) *EmailUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmailUpdateOne) Select(field string, fields ...string) *EmailUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Email entity.
func (euo *EmailUpdateOne) Save(ctx context.Context) (*Email, error) {
	euo.defaults()
	return withHooks[*Email, EmailMutation](ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmailUpdateOne) SaveX(ctx context.Context) *Email {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmailUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmailUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EmailUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := email.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

func (euo *EmailUpdateOne) sqlSave(ctx context.Context) (_node *Email, err error) {
	_spec := sqlgraph.NewUpdateSpec(email.Table, email.Columns, sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Email.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, email.FieldID)
		for _, f := range fields {
			if !email.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != email.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Address(); ok {
		_spec.SetField(email.FieldAddress, field.TypeString, value)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(email.FieldUpdatedAt, field.TypeTime, value)
	}
	if euo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.IdentitiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedIdentitiesIDs(); len(nodes) > 0 && !euo.mutation.IdentitiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.IdentitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.PasscodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedPasscodesIDs(); len(nodes) > 0 && !euo.mutation.PasscodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.PasscodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.PrimaryEmailCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Email{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{email.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}