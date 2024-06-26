// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NoahJinnn/passkey_auth_svc/ent/predicate"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredential"
)

// WebauthnCredentialDelete is the builder for deleting a WebauthnCredential entity.
type WebauthnCredentialDelete struct {
	config
	hooks    []Hook
	mutation *WebauthnCredentialMutation
}

// Where appends a list predicates to the WebauthnCredentialDelete builder.
func (wcd *WebauthnCredentialDelete) Where(ps ...predicate.WebauthnCredential) *WebauthnCredentialDelete {
	wcd.mutation.Where(ps...)
	return wcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wcd *WebauthnCredentialDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wcd.sqlExec, wcd.mutation, wcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wcd *WebauthnCredentialDelete) ExecX(ctx context.Context) int {
	n, err := wcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wcd *WebauthnCredentialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(webauthncredential.Table, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString))
	if ps := wcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wcd.mutation.done = true
	return affected, err
}

// WebauthnCredentialDeleteOne is the builder for deleting a single WebauthnCredential entity.
type WebauthnCredentialDeleteOne struct {
	wcd *WebauthnCredentialDelete
}

// Where appends a list predicates to the WebauthnCredentialDelete builder.
func (wcdo *WebauthnCredentialDeleteOne) Where(ps ...predicate.WebauthnCredential) *WebauthnCredentialDeleteOne {
	wcdo.wcd.mutation.Where(ps...)
	return wcdo
}

// Exec executes the deletion query.
func (wcdo *WebauthnCredentialDeleteOne) Exec(ctx context.Context) error {
	n, err := wcdo.wcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{webauthncredential.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wcdo *WebauthnCredentialDeleteOne) ExecX(ctx context.Context) {
	if err := wcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
