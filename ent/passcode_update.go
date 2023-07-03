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
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/user"
)

// PasscodeUpdate is the builder for updating Passcode entities.
type PasscodeUpdate struct {
	config
	hooks    []Hook
	mutation *PasscodeMutation
}

// Where appends a list predicates to the PasscodeUpdate builder.
func (pu *PasscodeUpdate) Where(ps ...predicate.Passcode) *PasscodeUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PasscodeUpdate) SetUserID(u uuid.UUID) *PasscodeUpdate {
	pu.mutation.SetUserID(u)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PasscodeUpdate) SetNillableUserID(u *uuid.UUID) *PasscodeUpdate {
	if u != nil {
		pu.SetUserID(*u)
	}
	return pu
}

// ClearUserID clears the value of the "user_id" field.
func (pu *PasscodeUpdate) ClearUserID() *PasscodeUpdate {
	pu.mutation.ClearUserID()
	return pu
}

// SetTTL sets the "ttl" field.
func (pu *PasscodeUpdate) SetTTL(i int32) *PasscodeUpdate {
	pu.mutation.ResetTTL()
	pu.mutation.SetTTL(i)
	return pu
}

// AddTTL adds i to the "ttl" field.
func (pu *PasscodeUpdate) AddTTL(i int32) *PasscodeUpdate {
	pu.mutation.AddTTL(i)
	return pu
}

// SetCode sets the "code" field.
func (pu *PasscodeUpdate) SetCode(s string) *PasscodeUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetTryCount sets the "try_count" field.
func (pu *PasscodeUpdate) SetTryCount(i int32) *PasscodeUpdate {
	pu.mutation.ResetTryCount()
	pu.mutation.SetTryCount(i)
	return pu
}

// AddTryCount adds i to the "try_count" field.
func (pu *PasscodeUpdate) AddTryCount(i int32) *PasscodeUpdate {
	pu.mutation.AddTryCount(i)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PasscodeUpdate) SetUpdatedAt(t time.Time) *PasscodeUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetEmailID sets the "email_id" field.
func (pu *PasscodeUpdate) SetEmailID(u uuid.UUID) *PasscodeUpdate {
	pu.mutation.SetEmailID(u)
	return pu
}

// SetNillableEmailID sets the "email_id" field if the given value is not nil.
func (pu *PasscodeUpdate) SetNillableEmailID(u *uuid.UUID) *PasscodeUpdate {
	if u != nil {
		pu.SetEmailID(*u)
	}
	return pu
}

// ClearEmailID clears the value of the "email_id" field.
func (pu *PasscodeUpdate) ClearEmailID() *PasscodeUpdate {
	pu.mutation.ClearEmailID()
	return pu
}

// SetEmail sets the "email" edge to the Email entity.
func (pu *PasscodeUpdate) SetEmail(e *Email) *PasscodeUpdate {
	return pu.SetEmailID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pu *PasscodeUpdate) SetUser(u *User) *PasscodeUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PasscodeMutation object of the builder.
func (pu *PasscodeUpdate) Mutation() *PasscodeMutation {
	return pu.mutation
}

// ClearEmail clears the "email" edge to the Email entity.
func (pu *PasscodeUpdate) ClearEmail() *PasscodeUpdate {
	pu.mutation.ClearEmail()
	return pu
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PasscodeUpdate) ClearUser() *PasscodeUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PasscodeUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PasscodeUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PasscodeUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PasscodeUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PasscodeUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := passcode.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PasscodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(passcode.Table, passcode.Columns, sqlgraph.NewFieldSpec(passcode.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.TTL(); ok {
		_spec.SetField(passcode.FieldTTL, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedTTL(); ok {
		_spec.AddField(passcode.FieldTTL, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(passcode.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.TryCount(); ok {
		_spec.SetField(passcode.FieldTryCount, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedTryCount(); ok {
		_spec.AddField(passcode.FieldTryCount, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(passcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.EmailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.EmailTable,
			Columns: []string{passcode.EmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.EmailTable,
			Columns: []string{passcode.EmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.UserTable,
			Columns: []string{passcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.UserTable,
			Columns: []string{passcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PasscodeUpdateOne is the builder for updating a single Passcode entity.
type PasscodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasscodeMutation
}

// SetUserID sets the "user_id" field.
func (puo *PasscodeUpdateOne) SetUserID(u uuid.UUID) *PasscodeUpdateOne {
	puo.mutation.SetUserID(u)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PasscodeUpdateOne) SetNillableUserID(u *uuid.UUID) *PasscodeUpdateOne {
	if u != nil {
		puo.SetUserID(*u)
	}
	return puo
}

// ClearUserID clears the value of the "user_id" field.
func (puo *PasscodeUpdateOne) ClearUserID() *PasscodeUpdateOne {
	puo.mutation.ClearUserID()
	return puo
}

// SetTTL sets the "ttl" field.
func (puo *PasscodeUpdateOne) SetTTL(i int32) *PasscodeUpdateOne {
	puo.mutation.ResetTTL()
	puo.mutation.SetTTL(i)
	return puo
}

// AddTTL adds i to the "ttl" field.
func (puo *PasscodeUpdateOne) AddTTL(i int32) *PasscodeUpdateOne {
	puo.mutation.AddTTL(i)
	return puo
}

// SetCode sets the "code" field.
func (puo *PasscodeUpdateOne) SetCode(s string) *PasscodeUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetTryCount sets the "try_count" field.
func (puo *PasscodeUpdateOne) SetTryCount(i int32) *PasscodeUpdateOne {
	puo.mutation.ResetTryCount()
	puo.mutation.SetTryCount(i)
	return puo
}

// AddTryCount adds i to the "try_count" field.
func (puo *PasscodeUpdateOne) AddTryCount(i int32) *PasscodeUpdateOne {
	puo.mutation.AddTryCount(i)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PasscodeUpdateOne) SetUpdatedAt(t time.Time) *PasscodeUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetEmailID sets the "email_id" field.
func (puo *PasscodeUpdateOne) SetEmailID(u uuid.UUID) *PasscodeUpdateOne {
	puo.mutation.SetEmailID(u)
	return puo
}

// SetNillableEmailID sets the "email_id" field if the given value is not nil.
func (puo *PasscodeUpdateOne) SetNillableEmailID(u *uuid.UUID) *PasscodeUpdateOne {
	if u != nil {
		puo.SetEmailID(*u)
	}
	return puo
}

// ClearEmailID clears the value of the "email_id" field.
func (puo *PasscodeUpdateOne) ClearEmailID() *PasscodeUpdateOne {
	puo.mutation.ClearEmailID()
	return puo
}

// SetEmail sets the "email" edge to the Email entity.
func (puo *PasscodeUpdateOne) SetEmail(e *Email) *PasscodeUpdateOne {
	return puo.SetEmailID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (puo *PasscodeUpdateOne) SetUser(u *User) *PasscodeUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PasscodeMutation object of the builder.
func (puo *PasscodeUpdateOne) Mutation() *PasscodeMutation {
	return puo.mutation
}

// ClearEmail clears the "email" edge to the Email entity.
func (puo *PasscodeUpdateOne) ClearEmail() *PasscodeUpdateOne {
	puo.mutation.ClearEmail()
	return puo
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PasscodeUpdateOne) ClearUser() *PasscodeUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PasscodeUpdate builder.
func (puo *PasscodeUpdateOne) Where(ps ...predicate.Passcode) *PasscodeUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PasscodeUpdateOne) Select(field string, fields ...string) *PasscodeUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Passcode entity.
func (puo *PasscodeUpdateOne) Save(ctx context.Context) (*Passcode, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PasscodeUpdateOne) SaveX(ctx context.Context) *Passcode {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PasscodeUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PasscodeUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PasscodeUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := passcode.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PasscodeUpdateOne) sqlSave(ctx context.Context) (_node *Passcode, err error) {
	_spec := sqlgraph.NewUpdateSpec(passcode.Table, passcode.Columns, sqlgraph.NewFieldSpec(passcode.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Passcode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passcode.FieldID)
		for _, f := range fields {
			if !passcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.TTL(); ok {
		_spec.SetField(passcode.FieldTTL, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedTTL(); ok {
		_spec.AddField(passcode.FieldTTL, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(passcode.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.TryCount(); ok {
		_spec.SetField(passcode.FieldTryCount, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedTryCount(); ok {
		_spec.AddField(passcode.FieldTryCount, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(passcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.EmailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.EmailTable,
			Columns: []string{passcode.EmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EmailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.EmailTable,
			Columns: []string{passcode.EmailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(email.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.UserTable,
			Columns: []string{passcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passcode.UserTable,
			Columns: []string{passcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Passcode{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
