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
	"github.com/hellohq/hqservice/ent/assettable"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/user"
)

// AssetTableUpdate is the builder for updating AssetTable entities.
type AssetTableUpdate struct {
	config
	hooks    []Hook
	mutation *AssetTableMutation
}

// Where appends a list predicates to the AssetTableUpdate builder.
func (atu *AssetTableUpdate) Where(ps ...predicate.AssetTable) *AssetTableUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUserID sets the "user_id" field.
func (atu *AssetTableUpdate) SetUserID(u uuid.UUID) *AssetTableUpdate {
	atu.mutation.SetUserID(u)
	return atu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atu *AssetTableUpdate) SetNillableUserID(u *uuid.UUID) *AssetTableUpdate {
	if u != nil {
		atu.SetUserID(*u)
	}
	return atu
}

// ClearUserID clears the value of the "user_id" field.
func (atu *AssetTableUpdate) ClearUserID() *AssetTableUpdate {
	atu.mutation.ClearUserID()
	return atu
}

// SetSheet sets the "sheet" field.
func (atu *AssetTableUpdate) SetSheet(i int32) *AssetTableUpdate {
	atu.mutation.ResetSheet()
	atu.mutation.SetSheet(i)
	return atu
}

// SetNillableSheet sets the "sheet" field if the given value is not nil.
func (atu *AssetTableUpdate) SetNillableSheet(i *int32) *AssetTableUpdate {
	if i != nil {
		atu.SetSheet(*i)
	}
	return atu
}

// AddSheet adds i to the "sheet" field.
func (atu *AssetTableUpdate) AddSheet(i int32) *AssetTableUpdate {
	atu.mutation.AddSheet(i)
	return atu
}

// ClearSheet clears the value of the "sheet" field.
func (atu *AssetTableUpdate) ClearSheet() *AssetTableUpdate {
	atu.mutation.ClearSheet()
	return atu
}

// SetSection sets the "section" field.
func (atu *AssetTableUpdate) SetSection(i int32) *AssetTableUpdate {
	atu.mutation.ResetSection()
	atu.mutation.SetSection(i)
	return atu
}

// SetNillableSection sets the "section" field if the given value is not nil.
func (atu *AssetTableUpdate) SetNillableSection(i *int32) *AssetTableUpdate {
	if i != nil {
		atu.SetSection(*i)
	}
	return atu
}

// AddSection adds i to the "section" field.
func (atu *AssetTableUpdate) AddSection(i int32) *AssetTableUpdate {
	atu.mutation.AddSection(i)
	return atu
}

// ClearSection clears the value of the "section" field.
func (atu *AssetTableUpdate) ClearSection() *AssetTableUpdate {
	atu.mutation.ClearSection()
	return atu
}

// SetDescription sets the "description" field.
func (atu *AssetTableUpdate) SetDescription(s string) *AssetTableUpdate {
	atu.mutation.SetDescription(s)
	return atu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atu *AssetTableUpdate) SetNillableDescription(s *string) *AssetTableUpdate {
	if s != nil {
		atu.SetDescription(*s)
	}
	return atu
}

// ClearDescription clears the value of the "description" field.
func (atu *AssetTableUpdate) ClearDescription() *AssetTableUpdate {
	atu.mutation.ClearDescription()
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *AssetTableUpdate) SetUpdatedAt(t time.Time) *AssetTableUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *AssetTableUpdate) SetUser(u *User) *AssetTableUpdate {
	return atu.SetUserID(u.ID)
}

// Mutation returns the AssetTableMutation object of the builder.
func (atu *AssetTableUpdate) Mutation() *AssetTableMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *AssetTableUpdate) ClearUser() *AssetTableUpdate {
	atu.mutation.ClearUser()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AssetTableUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AssetTableUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AssetTableUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AssetTableUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *AssetTableUpdate) defaults() {
	if _, ok := atu.mutation.UpdatedAt(); !ok {
		v := assettable.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
}

func (atu *AssetTableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(assettable.Table, assettable.Columns, sqlgraph.NewFieldSpec(assettable.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Sheet(); ok {
		_spec.SetField(assettable.FieldSheet, field.TypeInt32, value)
	}
	if value, ok := atu.mutation.AddedSheet(); ok {
		_spec.AddField(assettable.FieldSheet, field.TypeInt32, value)
	}
	if atu.mutation.SheetCleared() {
		_spec.ClearField(assettable.FieldSheet, field.TypeInt32)
	}
	if value, ok := atu.mutation.Section(); ok {
		_spec.SetField(assettable.FieldSection, field.TypeInt32, value)
	}
	if value, ok := atu.mutation.AddedSection(); ok {
		_spec.AddField(assettable.FieldSection, field.TypeInt32, value)
	}
	if atu.mutation.SectionCleared() {
		_spec.ClearField(assettable.FieldSection, field.TypeInt32)
	}
	if value, ok := atu.mutation.Description(); ok {
		_spec.SetField(assettable.FieldDescription, field.TypeString, value)
	}
	if atu.mutation.DescriptionCleared() {
		_spec.ClearField(assettable.FieldDescription, field.TypeString)
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(assettable.FieldUpdatedAt, field.TypeTime, value)
	}
	if atu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assettable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AssetTableUpdateOne is the builder for updating a single AssetTable entity.
type AssetTableUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetTableMutation
}

// SetUserID sets the "user_id" field.
func (atuo *AssetTableUpdateOne) SetUserID(u uuid.UUID) *AssetTableUpdateOne {
	atuo.mutation.SetUserID(u)
	return atuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atuo *AssetTableUpdateOne) SetNillableUserID(u *uuid.UUID) *AssetTableUpdateOne {
	if u != nil {
		atuo.SetUserID(*u)
	}
	return atuo
}

// ClearUserID clears the value of the "user_id" field.
func (atuo *AssetTableUpdateOne) ClearUserID() *AssetTableUpdateOne {
	atuo.mutation.ClearUserID()
	return atuo
}

// SetSheet sets the "sheet" field.
func (atuo *AssetTableUpdateOne) SetSheet(i int32) *AssetTableUpdateOne {
	atuo.mutation.ResetSheet()
	atuo.mutation.SetSheet(i)
	return atuo
}

// SetNillableSheet sets the "sheet" field if the given value is not nil.
func (atuo *AssetTableUpdateOne) SetNillableSheet(i *int32) *AssetTableUpdateOne {
	if i != nil {
		atuo.SetSheet(*i)
	}
	return atuo
}

// AddSheet adds i to the "sheet" field.
func (atuo *AssetTableUpdateOne) AddSheet(i int32) *AssetTableUpdateOne {
	atuo.mutation.AddSheet(i)
	return atuo
}

// ClearSheet clears the value of the "sheet" field.
func (atuo *AssetTableUpdateOne) ClearSheet() *AssetTableUpdateOne {
	atuo.mutation.ClearSheet()
	return atuo
}

// SetSection sets the "section" field.
func (atuo *AssetTableUpdateOne) SetSection(i int32) *AssetTableUpdateOne {
	atuo.mutation.ResetSection()
	atuo.mutation.SetSection(i)
	return atuo
}

// SetNillableSection sets the "section" field if the given value is not nil.
func (atuo *AssetTableUpdateOne) SetNillableSection(i *int32) *AssetTableUpdateOne {
	if i != nil {
		atuo.SetSection(*i)
	}
	return atuo
}

// AddSection adds i to the "section" field.
func (atuo *AssetTableUpdateOne) AddSection(i int32) *AssetTableUpdateOne {
	atuo.mutation.AddSection(i)
	return atuo
}

// ClearSection clears the value of the "section" field.
func (atuo *AssetTableUpdateOne) ClearSection() *AssetTableUpdateOne {
	atuo.mutation.ClearSection()
	return atuo
}

// SetDescription sets the "description" field.
func (atuo *AssetTableUpdateOne) SetDescription(s string) *AssetTableUpdateOne {
	atuo.mutation.SetDescription(s)
	return atuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atuo *AssetTableUpdateOne) SetNillableDescription(s *string) *AssetTableUpdateOne {
	if s != nil {
		atuo.SetDescription(*s)
	}
	return atuo
}

// ClearDescription clears the value of the "description" field.
func (atuo *AssetTableUpdateOne) ClearDescription() *AssetTableUpdateOne {
	atuo.mutation.ClearDescription()
	return atuo
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *AssetTableUpdateOne) SetUpdatedAt(t time.Time) *AssetTableUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *AssetTableUpdateOne) SetUser(u *User) *AssetTableUpdateOne {
	return atuo.SetUserID(u.ID)
}

// Mutation returns the AssetTableMutation object of the builder.
func (atuo *AssetTableUpdateOne) Mutation() *AssetTableMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *AssetTableUpdateOne) ClearUser() *AssetTableUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// Where appends a list predicates to the AssetTableUpdate builder.
func (atuo *AssetTableUpdateOne) Where(ps ...predicate.AssetTable) *AssetTableUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AssetTableUpdateOne) Select(field string, fields ...string) *AssetTableUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AssetTable entity.
func (atuo *AssetTableUpdateOne) Save(ctx context.Context) (*AssetTable, error) {
	atuo.defaults()
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AssetTableUpdateOne) SaveX(ctx context.Context) *AssetTable {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AssetTableUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AssetTableUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *AssetTableUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdatedAt(); !ok {
		v := assettable.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
}

func (atuo *AssetTableUpdateOne) sqlSave(ctx context.Context) (_node *AssetTable, err error) {
	_spec := sqlgraph.NewUpdateSpec(assettable.Table, assettable.Columns, sqlgraph.NewFieldSpec(assettable.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssetTable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assettable.FieldID)
		for _, f := range fields {
			if !assettable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != assettable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.Sheet(); ok {
		_spec.SetField(assettable.FieldSheet, field.TypeInt32, value)
	}
	if value, ok := atuo.mutation.AddedSheet(); ok {
		_spec.AddField(assettable.FieldSheet, field.TypeInt32, value)
	}
	if atuo.mutation.SheetCleared() {
		_spec.ClearField(assettable.FieldSheet, field.TypeInt32)
	}
	if value, ok := atuo.mutation.Section(); ok {
		_spec.SetField(assettable.FieldSection, field.TypeInt32, value)
	}
	if value, ok := atuo.mutation.AddedSection(); ok {
		_spec.AddField(assettable.FieldSection, field.TypeInt32, value)
	}
	if atuo.mutation.SectionCleared() {
		_spec.ClearField(assettable.FieldSection, field.TypeInt32)
	}
	if value, ok := atuo.mutation.Description(); ok {
		_spec.SetField(assettable.FieldDescription, field.TypeString, value)
	}
	if atuo.mutation.DescriptionCleared() {
		_spec.ClearField(assettable.FieldDescription, field.TypeString)
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(assettable.FieldUpdatedAt, field.TypeTime, value)
	}
	if atuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AssetTable{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assettable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
