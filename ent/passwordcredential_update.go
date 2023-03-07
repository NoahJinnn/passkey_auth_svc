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
	"github.com/hellohq/hqservice/ent/passwordcredential"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/user"
)

// PasswordCredentialUpdate is the builder for updating PasswordCredential entities.
type PasswordCredentialUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordCredentialMutation
}

// Where appends a list predicates to the PasswordCredentialUpdate builder.
func (pcu *PasswordCredentialUpdate) Where(ps ...predicate.PasswordCredential) *PasswordCredentialUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetUserID sets the "user_id" field.
func (pcu *PasswordCredentialUpdate) SetUserID(u uuid.UUID) *PasswordCredentialUpdate {
	pcu.mutation.SetUserID(u)
	return pcu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pcu *PasswordCredentialUpdate) SetNillableUserID(u *uuid.UUID) *PasswordCredentialUpdate {
	if u != nil {
		pcu.SetUserID(*u)
	}
	return pcu
}

// ClearUserID clears the value of the "user_id" field.
func (pcu *PasswordCredentialUpdate) ClearUserID() *PasswordCredentialUpdate {
	pcu.mutation.ClearUserID()
	return pcu
}

// SetPassword sets the "password" field.
func (pcu *PasswordCredentialUpdate) SetPassword(s string) *PasswordCredentialUpdate {
	pcu.mutation.SetPassword(s)
	return pcu
}

// SetCreatedAt sets the "created_at" field.
func (pcu *PasswordCredentialUpdate) SetCreatedAt(t time.Time) *PasswordCredentialUpdate {
	pcu.mutation.SetCreatedAt(t)
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *PasswordCredentialUpdate) SetUpdatedAt(t time.Time) *PasswordCredentialUpdate {
	pcu.mutation.SetUpdatedAt(t)
	return pcu
}

// SetUser sets the "user" edge to the User entity.
func (pcu *PasswordCredentialUpdate) SetUser(u *User) *PasswordCredentialUpdate {
	return pcu.SetUserID(u.ID)
}

// Mutation returns the PasswordCredentialMutation object of the builder.
func (pcu *PasswordCredentialUpdate) Mutation() *PasswordCredentialMutation {
	return pcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pcu *PasswordCredentialUpdate) ClearUser() *PasswordCredentialUpdate {
	pcu.mutation.ClearUser()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PasswordCredentialUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PasswordCredentialMutation](ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *PasswordCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *PasswordCredentialUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *PasswordCredentialUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *PasswordCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(passwordcredential.Table, passwordcredential.Columns, sqlgraph.NewFieldSpec(passwordcredential.FieldID, field.TypeUUID))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Password(); ok {
		_spec.SetField(passwordcredential.FieldPassword, field.TypeString, value)
	}
	if value, ok := pcu.mutation.CreatedAt(); ok {
		_spec.SetField(passwordcredential.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(passwordcredential.FieldUpdatedAt, field.TypeTime, value)
	}
	if pcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   passwordcredential.UserTable,
			Columns: []string{passwordcredential.UserColumn},
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
	if nodes := pcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   passwordcredential.UserTable,
			Columns: []string{passwordcredential.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// PasswordCredentialUpdateOne is the builder for updating a single PasswordCredential entity.
type PasswordCredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordCredentialMutation
}

// SetUserID sets the "user_id" field.
func (pcuo *PasswordCredentialUpdateOne) SetUserID(u uuid.UUID) *PasswordCredentialUpdateOne {
	pcuo.mutation.SetUserID(u)
	return pcuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pcuo *PasswordCredentialUpdateOne) SetNillableUserID(u *uuid.UUID) *PasswordCredentialUpdateOne {
	if u != nil {
		pcuo.SetUserID(*u)
	}
	return pcuo
}

// ClearUserID clears the value of the "user_id" field.
func (pcuo *PasswordCredentialUpdateOne) ClearUserID() *PasswordCredentialUpdateOne {
	pcuo.mutation.ClearUserID()
	return pcuo
}

// SetPassword sets the "password" field.
func (pcuo *PasswordCredentialUpdateOne) SetPassword(s string) *PasswordCredentialUpdateOne {
	pcuo.mutation.SetPassword(s)
	return pcuo
}

// SetCreatedAt sets the "created_at" field.
func (pcuo *PasswordCredentialUpdateOne) SetCreatedAt(t time.Time) *PasswordCredentialUpdateOne {
	pcuo.mutation.SetCreatedAt(t)
	return pcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *PasswordCredentialUpdateOne) SetUpdatedAt(t time.Time) *PasswordCredentialUpdateOne {
	pcuo.mutation.SetUpdatedAt(t)
	return pcuo
}

// SetUser sets the "user" edge to the User entity.
func (pcuo *PasswordCredentialUpdateOne) SetUser(u *User) *PasswordCredentialUpdateOne {
	return pcuo.SetUserID(u.ID)
}

// Mutation returns the PasswordCredentialMutation object of the builder.
func (pcuo *PasswordCredentialUpdateOne) Mutation() *PasswordCredentialMutation {
	return pcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pcuo *PasswordCredentialUpdateOne) ClearUser() *PasswordCredentialUpdateOne {
	pcuo.mutation.ClearUser()
	return pcuo
}

// Where appends a list predicates to the PasswordCredentialUpdate builder.
func (pcuo *PasswordCredentialUpdateOne) Where(ps ...predicate.PasswordCredential) *PasswordCredentialUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *PasswordCredentialUpdateOne) Select(field string, fields ...string) *PasswordCredentialUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated PasswordCredential entity.
func (pcuo *PasswordCredentialUpdateOne) Save(ctx context.Context) (*PasswordCredential, error) {
	return withHooks[*PasswordCredential, PasswordCredentialMutation](ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *PasswordCredentialUpdateOne) SaveX(ctx context.Context) *PasswordCredential {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *PasswordCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *PasswordCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *PasswordCredentialUpdateOne) sqlSave(ctx context.Context) (_node *PasswordCredential, err error) {
	_spec := sqlgraph.NewUpdateSpec(passwordcredential.Table, passwordcredential.Columns, sqlgraph.NewFieldSpec(passwordcredential.FieldID, field.TypeUUID))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PasswordCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwordcredential.FieldID)
		for _, f := range fields {
			if !passwordcredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passwordcredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Password(); ok {
		_spec.SetField(passwordcredential.FieldPassword, field.TypeString, value)
	}
	if value, ok := pcuo.mutation.CreatedAt(); ok {
		_spec.SetField(passwordcredential.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(passwordcredential.FieldUpdatedAt, field.TypeTime, value)
	}
	if pcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   passwordcredential.UserTable,
			Columns: []string{passwordcredential.UserColumn},
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
	if nodes := pcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   passwordcredential.UserTable,
			Columns: []string{passwordcredential.UserColumn},
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
	_node = &PasswordCredential{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
