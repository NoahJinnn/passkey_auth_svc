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
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/passwordcredential"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddEmailIDs adds the "emails" edge to the Email entity by IDs.
func (uu *UserUpdate) AddEmailIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddEmailIDs(ids...)
	return uu
}

// AddEmails adds the "emails" edges to the Email entity.
func (uu *UserUpdate) AddEmails(e ...*Email) *UserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddEmailIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (uu *UserUpdate) AddPasscodeIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddPasscodeIDs(ids...)
	return uu
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (uu *UserUpdate) AddPasscodes(p ...*Passcode) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPasscodeIDs(ids...)
}

// SetPasswordCredentialID sets the "password_credential" edge to the PasswordCredential entity by ID.
func (uu *UserUpdate) SetPasswordCredentialID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetPasswordCredentialID(id)
	return uu
}

// SetNillablePasswordCredentialID sets the "password_credential" edge to the PasswordCredential entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswordCredentialID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetPasswordCredentialID(*id)
	}
	return uu
}

// SetPasswordCredential sets the "password_credential" edge to the PasswordCredential entity.
func (uu *UserUpdate) SetPasswordCredential(p *PasswordCredential) *UserUpdate {
	return uu.SetPasswordCredentialID(p.ID)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (uu *UserUpdate) SetPrimaryEmailID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetPrimaryEmailID(id)
	return uu
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePrimaryEmailID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetPrimaryEmailID(*id)
	}
	return uu
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (uu *UserUpdate) SetPrimaryEmail(p *PrimaryEmail) *UserUpdate {
	return uu.SetPrimaryEmailID(p.ID)
}

// AddWebauthnCredentialIDs adds the "webauthn_credentials" edge to the WebauthnCredential entity by IDs.
func (uu *UserUpdate) AddWebauthnCredentialIDs(ids ...string) *UserUpdate {
	uu.mutation.AddWebauthnCredentialIDs(ids...)
	return uu
}

// AddWebauthnCredentials adds the "webauthn_credentials" edges to the WebauthnCredential entity.
func (uu *UserUpdate) AddWebauthnCredentials(w ...*WebauthnCredential) *UserUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.AddWebauthnCredentialIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearEmails clears all "emails" edges to the Email entity.
func (uu *UserUpdate) ClearEmails() *UserUpdate {
	uu.mutation.ClearEmails()
	return uu
}

// RemoveEmailIDs removes the "emails" edge to Email entities by IDs.
func (uu *UserUpdate) RemoveEmailIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveEmailIDs(ids...)
	return uu
}

// RemoveEmails removes "emails" edges to Email entities.
func (uu *UserUpdate) RemoveEmails(e ...*Email) *UserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveEmailIDs(ids...)
}

// ClearPasscodes clears all "passcodes" edges to the Passcode entity.
func (uu *UserUpdate) ClearPasscodes() *UserUpdate {
	uu.mutation.ClearPasscodes()
	return uu
}

// RemovePasscodeIDs removes the "passcodes" edge to Passcode entities by IDs.
func (uu *UserUpdate) RemovePasscodeIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemovePasscodeIDs(ids...)
	return uu
}

// RemovePasscodes removes "passcodes" edges to Passcode entities.
func (uu *UserUpdate) RemovePasscodes(p ...*Passcode) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePasscodeIDs(ids...)
}

// ClearPasswordCredential clears the "password_credential" edge to the PasswordCredential entity.
func (uu *UserUpdate) ClearPasswordCredential() *UserUpdate {
	uu.mutation.ClearPasswordCredential()
	return uu
}

// ClearPrimaryEmail clears the "primary_email" edge to the PrimaryEmail entity.
func (uu *UserUpdate) ClearPrimaryEmail() *UserUpdate {
	uu.mutation.ClearPrimaryEmail()
	return uu
}

// ClearWebauthnCredentials clears all "webauthn_credentials" edges to the WebauthnCredential entity.
func (uu *UserUpdate) ClearWebauthnCredentials() *UserUpdate {
	uu.mutation.ClearWebauthnCredentials()
	return uu
}

// RemoveWebauthnCredentialIDs removes the "webauthn_credentials" edge to WebauthnCredential entities by IDs.
func (uu *UserUpdate) RemoveWebauthnCredentialIDs(ids ...string) *UserUpdate {
	uu.mutation.RemoveWebauthnCredentialIDs(ids...)
	return uu
}

// RemoveWebauthnCredentials removes "webauthn_credentials" edges to WebauthnCredential entities.
func (uu *UserUpdate) RemoveWebauthnCredentials(w ...*WebauthnCredential) *UserUpdate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.RemoveWebauthnCredentialIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.EmailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: email.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedEmailsIDs(); len(nodes) > 0 && !uu.mutation.EmailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EmailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PasscodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if nodes := uu.mutation.RemovedPasscodesIDs(); len(nodes) > 0 && !uu.mutation.PasscodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if nodes := uu.mutation.PasscodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if uu.mutation.PasswordCredentialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PasswordCredentialTable,
			Columns: []string{user.PasswordCredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passwordcredential.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PasswordCredentialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PasswordCredentialTable,
			Columns: []string{user.PasswordCredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passwordcredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PrimaryEmailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PrimaryEmailTable,
			Columns: []string{user.PrimaryEmailColumn},
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
	if nodes := uu.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PrimaryEmailTable,
			Columns: []string{user.PrimaryEmailColumn},
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
	if uu.mutation.WebauthnCredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedWebauthnCredentialsIDs(); len(nodes) > 0 && !uu.mutation.WebauthnCredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WebauthnCredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddEmailIDs adds the "emails" edge to the Email entity by IDs.
func (uuo *UserUpdateOne) AddEmailIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddEmailIDs(ids...)
	return uuo
}

// AddEmails adds the "emails" edges to the Email entity.
func (uuo *UserUpdateOne) AddEmails(e ...*Email) *UserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddEmailIDs(ids...)
}

// AddPasscodeIDs adds the "passcodes" edge to the Passcode entity by IDs.
func (uuo *UserUpdateOne) AddPasscodeIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddPasscodeIDs(ids...)
	return uuo
}

// AddPasscodes adds the "passcodes" edges to the Passcode entity.
func (uuo *UserUpdateOne) AddPasscodes(p ...*Passcode) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPasscodeIDs(ids...)
}

// SetPasswordCredentialID sets the "password_credential" edge to the PasswordCredential entity by ID.
func (uuo *UserUpdateOne) SetPasswordCredentialID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetPasswordCredentialID(id)
	return uuo
}

// SetNillablePasswordCredentialID sets the "password_credential" edge to the PasswordCredential entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswordCredentialID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPasswordCredentialID(*id)
	}
	return uuo
}

// SetPasswordCredential sets the "password_credential" edge to the PasswordCredential entity.
func (uuo *UserUpdateOne) SetPasswordCredential(p *PasswordCredential) *UserUpdateOne {
	return uuo.SetPasswordCredentialID(p.ID)
}

// SetPrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID.
func (uuo *UserUpdateOne) SetPrimaryEmailID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetPrimaryEmailID(id)
	return uuo
}

// SetNillablePrimaryEmailID sets the "primary_email" edge to the PrimaryEmail entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePrimaryEmailID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPrimaryEmailID(*id)
	}
	return uuo
}

// SetPrimaryEmail sets the "primary_email" edge to the PrimaryEmail entity.
func (uuo *UserUpdateOne) SetPrimaryEmail(p *PrimaryEmail) *UserUpdateOne {
	return uuo.SetPrimaryEmailID(p.ID)
}

// AddWebauthnCredentialIDs adds the "webauthn_credentials" edge to the WebauthnCredential entity by IDs.
func (uuo *UserUpdateOne) AddWebauthnCredentialIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddWebauthnCredentialIDs(ids...)
	return uuo
}

// AddWebauthnCredentials adds the "webauthn_credentials" edges to the WebauthnCredential entity.
func (uuo *UserUpdateOne) AddWebauthnCredentials(w ...*WebauthnCredential) *UserUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.AddWebauthnCredentialIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearEmails clears all "emails" edges to the Email entity.
func (uuo *UserUpdateOne) ClearEmails() *UserUpdateOne {
	uuo.mutation.ClearEmails()
	return uuo
}

// RemoveEmailIDs removes the "emails" edge to Email entities by IDs.
func (uuo *UserUpdateOne) RemoveEmailIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveEmailIDs(ids...)
	return uuo
}

// RemoveEmails removes "emails" edges to Email entities.
func (uuo *UserUpdateOne) RemoveEmails(e ...*Email) *UserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveEmailIDs(ids...)
}

// ClearPasscodes clears all "passcodes" edges to the Passcode entity.
func (uuo *UserUpdateOne) ClearPasscodes() *UserUpdateOne {
	uuo.mutation.ClearPasscodes()
	return uuo
}

// RemovePasscodeIDs removes the "passcodes" edge to Passcode entities by IDs.
func (uuo *UserUpdateOne) RemovePasscodeIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemovePasscodeIDs(ids...)
	return uuo
}

// RemovePasscodes removes "passcodes" edges to Passcode entities.
func (uuo *UserUpdateOne) RemovePasscodes(p ...*Passcode) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePasscodeIDs(ids...)
}

// ClearPasswordCredential clears the "password_credential" edge to the PasswordCredential entity.
func (uuo *UserUpdateOne) ClearPasswordCredential() *UserUpdateOne {
	uuo.mutation.ClearPasswordCredential()
	return uuo
}

// ClearPrimaryEmail clears the "primary_email" edge to the PrimaryEmail entity.
func (uuo *UserUpdateOne) ClearPrimaryEmail() *UserUpdateOne {
	uuo.mutation.ClearPrimaryEmail()
	return uuo
}

// ClearWebauthnCredentials clears all "webauthn_credentials" edges to the WebauthnCredential entity.
func (uuo *UserUpdateOne) ClearWebauthnCredentials() *UserUpdateOne {
	uuo.mutation.ClearWebauthnCredentials()
	return uuo
}

// RemoveWebauthnCredentialIDs removes the "webauthn_credentials" edge to WebauthnCredential entities by IDs.
func (uuo *UserUpdateOne) RemoveWebauthnCredentialIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemoveWebauthnCredentialIDs(ids...)
	return uuo
}

// RemoveWebauthnCredentials removes "webauthn_credentials" edges to WebauthnCredential entities.
func (uuo *UserUpdateOne) RemoveWebauthnCredentials(w ...*WebauthnCredential) *UserUpdateOne {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.RemoveWebauthnCredentialIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.EmailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: email.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedEmailsIDs(); len(nodes) > 0 && !uuo.mutation.EmailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EmailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PasscodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if nodes := uuo.mutation.RemovedPasscodesIDs(); len(nodes) > 0 && !uuo.mutation.PasscodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if nodes := uuo.mutation.PasscodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasscodesTable,
			Columns: []string{user.PasscodesColumn},
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
	if uuo.mutation.PasswordCredentialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PasswordCredentialTable,
			Columns: []string{user.PasswordCredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passwordcredential.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PasswordCredentialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PasswordCredentialTable,
			Columns: []string{user.PasswordCredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passwordcredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PrimaryEmailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PrimaryEmailTable,
			Columns: []string{user.PrimaryEmailColumn},
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
	if nodes := uuo.mutation.PrimaryEmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PrimaryEmailTable,
			Columns: []string{user.PrimaryEmailColumn},
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
	if uuo.mutation.WebauthnCredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedWebauthnCredentialsIDs(); len(nodes) > 0 && !uuo.mutation.WebauthnCredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WebauthnCredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WebauthnCredentialsTable,
			Columns: []string{user.WebauthnCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: webauthncredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
