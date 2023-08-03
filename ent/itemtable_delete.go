// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hellohq/hqservice/ent/itemtable"
	"github.com/hellohq/hqservice/ent/predicate"
)

// ItemTableDelete is the builder for deleting a ItemTable entity.
type ItemTableDelete struct {
	config
	hooks    []Hook
	mutation *ItemTableMutation
}

// Where appends a list predicates to the ItemTableDelete builder.
func (itd *ItemTableDelete) Where(ps ...predicate.ItemTable) *ItemTableDelete {
	itd.mutation.Where(ps...)
	return itd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (itd *ItemTableDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, itd.sqlExec, itd.mutation, itd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (itd *ItemTableDelete) ExecX(ctx context.Context) int {
	n, err := itd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (itd *ItemTableDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(itemtable.Table, sqlgraph.NewFieldSpec(itemtable.FieldID, field.TypeUUID))
	if ps := itd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, itd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	itd.mutation.done = true
	return affected, err
}

// ItemTableDeleteOne is the builder for deleting a single ItemTable entity.
type ItemTableDeleteOne struct {
	itd *ItemTableDelete
}

// Where appends a list predicates to the ItemTableDelete builder.
func (itdo *ItemTableDeleteOne) Where(ps ...predicate.ItemTable) *ItemTableDeleteOne {
	itdo.itd.mutation.Where(ps...)
	return itdo
}

// Exec executes the deletion query.
func (itdo *ItemTableDeleteOne) Exec(ctx context.Context) error {
	n, err := itdo.itd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itemtable.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (itdo *ItemTableDeleteOne) ExecX(ctx context.Context) {
	if err := itdo.Exec(ctx); err != nil {
		panic(err)
	}
}