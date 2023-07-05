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
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/predicate"
)

// ConnectionUpdate is the builder for updating Connection entities.
type ConnectionUpdate struct {
	config
	hooks    []Hook
	mutation *ConnectionMutation
}

// Where appends a list predicates to the ConnectionUpdate builder.
func (cu *ConnectionUpdate) Where(ps ...predicate.Connection) *ConnectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetInstitutionID sets the "institution_id" field.
func (cu *ConnectionUpdate) SetInstitutionID(u uuid.UUID) *ConnectionUpdate {
	cu.mutation.SetInstitutionID(u)
	return cu
}

// SetNillableInstitutionID sets the "institution_id" field if the given value is not nil.
func (cu *ConnectionUpdate) SetNillableInstitutionID(u *uuid.UUID) *ConnectionUpdate {
	if u != nil {
		cu.SetInstitutionID(*u)
	}
	return cu
}

// ClearInstitutionID clears the value of the "institution_id" field.
func (cu *ConnectionUpdate) ClearInstitutionID() *ConnectionUpdate {
	cu.mutation.ClearInstitutionID()
	return cu
}

// SetData sets the "data" field.
func (cu *ConnectionUpdate) SetData(s string) *ConnectionUpdate {
	cu.mutation.SetData(s)
	return cu
}

// SetEnv sets the "env" field.
func (cu *ConnectionUpdate) SetEnv(s string) *ConnectionUpdate {
	cu.mutation.SetEnv(s)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ConnectionUpdate) SetUpdatedAt(t time.Time) *ConnectionUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (cu *ConnectionUpdate) SetInstitution(i *Institution) *ConnectionUpdate {
	return cu.SetInstitutionID(i.ID)
}

// Mutation returns the ConnectionMutation object of the builder.
func (cu *ConnectionUpdate) Mutation() *ConnectionMutation {
	return cu.mutation
}

// ClearInstitution clears the "institution" edge to the Institution entity.
func (cu *ConnectionUpdate) ClearInstitution() *ConnectionUpdate {
	cu.mutation.ClearInstitution()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectionUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConnectionUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := connection.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ConnectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(connection.Table, connection.Columns, sqlgraph.NewFieldSpec(connection.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Data(); ok {
		_spec.SetField(connection.FieldData, field.TypeString, value)
	}
	if value, ok := cu.mutation.Env(); ok {
		_spec.SetField(connection.FieldEnv, field.TypeString, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(connection.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   connection.InstitutionTable,
			Columns: []string{connection.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   connection.InstitutionTable,
			Columns: []string{connection.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConnectionUpdateOne is the builder for updating a single Connection entity.
type ConnectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConnectionMutation
}

// SetInstitutionID sets the "institution_id" field.
func (cuo *ConnectionUpdateOne) SetInstitutionID(u uuid.UUID) *ConnectionUpdateOne {
	cuo.mutation.SetInstitutionID(u)
	return cuo
}

// SetNillableInstitutionID sets the "institution_id" field if the given value is not nil.
func (cuo *ConnectionUpdateOne) SetNillableInstitutionID(u *uuid.UUID) *ConnectionUpdateOne {
	if u != nil {
		cuo.SetInstitutionID(*u)
	}
	return cuo
}

// ClearInstitutionID clears the value of the "institution_id" field.
func (cuo *ConnectionUpdateOne) ClearInstitutionID() *ConnectionUpdateOne {
	cuo.mutation.ClearInstitutionID()
	return cuo
}

// SetData sets the "data" field.
func (cuo *ConnectionUpdateOne) SetData(s string) *ConnectionUpdateOne {
	cuo.mutation.SetData(s)
	return cuo
}

// SetEnv sets the "env" field.
func (cuo *ConnectionUpdateOne) SetEnv(s string) *ConnectionUpdateOne {
	cuo.mutation.SetEnv(s)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ConnectionUpdateOne) SetUpdatedAt(t time.Time) *ConnectionUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (cuo *ConnectionUpdateOne) SetInstitution(i *Institution) *ConnectionUpdateOne {
	return cuo.SetInstitutionID(i.ID)
}

// Mutation returns the ConnectionMutation object of the builder.
func (cuo *ConnectionUpdateOne) Mutation() *ConnectionMutation {
	return cuo.mutation
}

// ClearInstitution clears the "institution" edge to the Institution entity.
func (cuo *ConnectionUpdateOne) ClearInstitution() *ConnectionUpdateOne {
	cuo.mutation.ClearInstitution()
	return cuo
}

// Where appends a list predicates to the ConnectionUpdate builder.
func (cuo *ConnectionUpdateOne) Where(ps ...predicate.Connection) *ConnectionUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectionUpdateOne) Select(field string, fields ...string) *ConnectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connection entity.
func (cuo *ConnectionUpdateOne) Save(ctx context.Context) (*Connection, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectionUpdateOne) SaveX(ctx context.Context) *Connection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConnectionUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := connection.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ConnectionUpdateOne) sqlSave(ctx context.Context) (_node *Connection, err error) {
	_spec := sqlgraph.NewUpdateSpec(connection.Table, connection.Columns, sqlgraph.NewFieldSpec(connection.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Connection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connection.FieldID)
		for _, f := range fields {
			if !connection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != connection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Data(); ok {
		_spec.SetField(connection.FieldData, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Env(); ok {
		_spec.SetField(connection.FieldEnv, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(connection.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   connection.InstitutionTable,
			Columns: []string{connection.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   connection.InstitutionTable,
			Columns: []string{connection.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Connection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}