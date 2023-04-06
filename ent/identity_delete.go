// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hellohq/hqservice/ent/identity"
	"github.com/hellohq/hqservice/ent/predicate"
)

// IdentityDelete is the builder for deleting a Identity entity.
type IdentityDelete struct {
	config
	hooks    []Hook
	mutation *IdentityMutation
}

// Where appends a list predicates to the IdentityDelete builder.
func (id *IdentityDelete) Where(ps ...predicate.Identity) *IdentityDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IdentityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, IdentityMutation](ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IdentityDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IdentityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(identity.Table, sqlgraph.NewFieldSpec(identity.FieldID, field.TypeUUID))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// IdentityDeleteOne is the builder for deleting a single Identity entity.
type IdentityDeleteOne struct {
	id *IdentityDelete
}

// Where appends a list predicates to the IdentityDelete builder.
func (ido *IdentityDeleteOne) Where(ps ...predicate.Identity) *IdentityDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *IdentityDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{identity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IdentityDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
