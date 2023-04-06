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
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
	"github.com/hellohq/hqservice/ent/webauthncredentialtransport"
)

// WebauthnCredentialCreate is the builder for creating a WebauthnCredential entity.
type WebauthnCredentialCreate struct {
	config
	mutation *WebauthnCredentialMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (wcc *WebauthnCredentialCreate) SetUserID(u uuid.UUID) *WebauthnCredentialCreate {
	wcc.mutation.SetUserID(u)
	return wcc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcc *WebauthnCredentialCreate) SetNillableUserID(u *uuid.UUID) *WebauthnCredentialCreate {
	if u != nil {
		wcc.SetUserID(*u)
	}
	return wcc
}

// SetPublicKey sets the "public_key" field.
func (wcc *WebauthnCredentialCreate) SetPublicKey(s string) *WebauthnCredentialCreate {
	wcc.mutation.SetPublicKey(s)
	return wcc
}

// SetAttestationType sets the "attestation_type" field.
func (wcc *WebauthnCredentialCreate) SetAttestationType(s string) *WebauthnCredentialCreate {
	wcc.mutation.SetAttestationType(s)
	return wcc
}

// SetAaguid sets the "aaguid" field.
func (wcc *WebauthnCredentialCreate) SetAaguid(u uuid.UUID) *WebauthnCredentialCreate {
	wcc.mutation.SetAaguid(u)
	return wcc
}

// SetSignCount sets the "sign_count" field.
func (wcc *WebauthnCredentialCreate) SetSignCount(i int32) *WebauthnCredentialCreate {
	wcc.mutation.SetSignCount(i)
	return wcc
}

// SetCreatedAt sets the "created_at" field.
func (wcc *WebauthnCredentialCreate) SetCreatedAt(t time.Time) *WebauthnCredentialCreate {
	wcc.mutation.SetCreatedAt(t)
	return wcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wcc *WebauthnCredentialCreate) SetNillableCreatedAt(t *time.Time) *WebauthnCredentialCreate {
	if t != nil {
		wcc.SetCreatedAt(*t)
	}
	return wcc
}

// SetUpdatedAt sets the "updated_at" field.
func (wcc *WebauthnCredentialCreate) SetUpdatedAt(t time.Time) *WebauthnCredentialCreate {
	wcc.mutation.SetUpdatedAt(t)
	return wcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wcc *WebauthnCredentialCreate) SetNillableUpdatedAt(t *time.Time) *WebauthnCredentialCreate {
	if t != nil {
		wcc.SetUpdatedAt(*t)
	}
	return wcc
}

// SetName sets the "name" field.
func (wcc *WebauthnCredentialCreate) SetName(s string) *WebauthnCredentialCreate {
	wcc.mutation.SetName(s)
	return wcc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wcc *WebauthnCredentialCreate) SetNillableName(s *string) *WebauthnCredentialCreate {
	if s != nil {
		wcc.SetName(*s)
	}
	return wcc
}

// SetBackupEligible sets the "backup_eligible" field.
func (wcc *WebauthnCredentialCreate) SetBackupEligible(b bool) *WebauthnCredentialCreate {
	wcc.mutation.SetBackupEligible(b)
	return wcc
}

// SetBackupState sets the "backup_state" field.
func (wcc *WebauthnCredentialCreate) SetBackupState(b bool) *WebauthnCredentialCreate {
	wcc.mutation.SetBackupState(b)
	return wcc
}

// SetLastUsedAt sets the "last_used_at" field.
func (wcc *WebauthnCredentialCreate) SetLastUsedAt(t time.Time) *WebauthnCredentialCreate {
	wcc.mutation.SetLastUsedAt(t)
	return wcc
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (wcc *WebauthnCredentialCreate) SetNillableLastUsedAt(t *time.Time) *WebauthnCredentialCreate {
	if t != nil {
		wcc.SetLastUsedAt(*t)
	}
	return wcc
}

// SetID sets the "id" field.
func (wcc *WebauthnCredentialCreate) SetID(s string) *WebauthnCredentialCreate {
	wcc.mutation.SetID(s)
	return wcc
}

// AddWebauthnCredentialTransportIDs adds the "webauthn_credential_transports" edge to the WebauthnCredentialTransport entity by IDs.
func (wcc *WebauthnCredentialCreate) AddWebauthnCredentialTransportIDs(ids ...uuid.UUID) *WebauthnCredentialCreate {
	wcc.mutation.AddWebauthnCredentialTransportIDs(ids...)
	return wcc
}

// AddWebauthnCredentialTransports adds the "webauthn_credential_transports" edges to the WebauthnCredentialTransport entity.
func (wcc *WebauthnCredentialCreate) AddWebauthnCredentialTransports(w ...*WebauthnCredentialTransport) *WebauthnCredentialCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcc.AddWebauthnCredentialTransportIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (wcc *WebauthnCredentialCreate) SetUser(u *User) *WebauthnCredentialCreate {
	return wcc.SetUserID(u.ID)
}

// Mutation returns the WebauthnCredentialMutation object of the builder.
func (wcc *WebauthnCredentialCreate) Mutation() *WebauthnCredentialMutation {
	return wcc.mutation
}

// Save creates the WebauthnCredential in the database.
func (wcc *WebauthnCredentialCreate) Save(ctx context.Context) (*WebauthnCredential, error) {
	wcc.defaults()
	return withHooks[*WebauthnCredential, WebauthnCredentialMutation](ctx, wcc.sqlSave, wcc.mutation, wcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcc *WebauthnCredentialCreate) SaveX(ctx context.Context) *WebauthnCredential {
	v, err := wcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcc *WebauthnCredentialCreate) Exec(ctx context.Context) error {
	_, err := wcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcc *WebauthnCredentialCreate) ExecX(ctx context.Context) {
	if err := wcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcc *WebauthnCredentialCreate) defaults() {
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		v := webauthncredential.DefaultCreatedAt()
		wcc.mutation.SetCreatedAt(v)
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		v := webauthncredential.DefaultUpdatedAt()
		wcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcc *WebauthnCredentialCreate) check() error {
	if _, ok := wcc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "WebauthnCredential.public_key"`)}
	}
	if _, ok := wcc.mutation.AttestationType(); !ok {
		return &ValidationError{Name: "attestation_type", err: errors.New(`ent: missing required field "WebauthnCredential.attestation_type"`)}
	}
	if _, ok := wcc.mutation.Aaguid(); !ok {
		return &ValidationError{Name: "aaguid", err: errors.New(`ent: missing required field "WebauthnCredential.aaguid"`)}
	}
	if _, ok := wcc.mutation.SignCount(); !ok {
		return &ValidationError{Name: "sign_count", err: errors.New(`ent: missing required field "WebauthnCredential.sign_count"`)}
	}
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WebauthnCredential.created_at"`)}
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WebauthnCredential.updated_at"`)}
	}
	if _, ok := wcc.mutation.BackupEligible(); !ok {
		return &ValidationError{Name: "backup_eligible", err: errors.New(`ent: missing required field "WebauthnCredential.backup_eligible"`)}
	}
	if _, ok := wcc.mutation.BackupState(); !ok {
		return &ValidationError{Name: "backup_state", err: errors.New(`ent: missing required field "WebauthnCredential.backup_state"`)}
	}
	return nil
}

func (wcc *WebauthnCredentialCreate) sqlSave(ctx context.Context) (*WebauthnCredential, error) {
	if err := wcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected WebauthnCredential.ID type: %T", _spec.ID.Value)
		}
	}
	wcc.mutation.id = &_node.ID
	wcc.mutation.done = true
	return _node, nil
}

func (wcc *WebauthnCredentialCreate) createSpec() (*WebauthnCredential, *sqlgraph.CreateSpec) {
	var (
		_node = &WebauthnCredential{config: wcc.config}
		_spec = sqlgraph.NewCreateSpec(webauthncredential.Table, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString))
	)
	if id, ok := wcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wcc.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeString, value)
		_node.PublicKey = value
	}
	if value, ok := wcc.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
		_node.AttestationType = value
	}
	if value, ok := wcc.mutation.Aaguid(); ok {
		_spec.SetField(webauthncredential.FieldAaguid, field.TypeUUID, value)
		_node.Aaguid = value
	}
	if value, ok := wcc.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeInt32, value)
		_node.SignCount = value
	}
	if value, ok := wcc.mutation.CreatedAt(); ok {
		_spec.SetField(webauthncredential.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wcc.mutation.UpdatedAt(); ok {
		_spec.SetField(webauthncredential.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wcc.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wcc.mutation.BackupEligible(); ok {
		_spec.SetField(webauthncredential.FieldBackupEligible, field.TypeBool, value)
		_node.BackupEligible = value
	}
	if value, ok := wcc.mutation.BackupState(); ok {
		_spec.SetField(webauthncredential.FieldBackupState, field.TypeBool, value)
		_node.BackupState = value
	}
	if value, ok := wcc.mutation.LastUsedAt(); ok {
		_spec.SetField(webauthncredential.FieldLastUsedAt, field.TypeTime, value)
		_node.LastUsedAt = value
	}
	if nodes := wcc.mutation.WebauthnCredentialTransportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: webauthncredentialtransport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
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
	return _node, _spec
}

// WebauthnCredentialCreateBulk is the builder for creating many WebauthnCredential entities in bulk.
type WebauthnCredentialCreateBulk struct {
	config
	builders []*WebauthnCredentialCreate
}

// Save creates the WebauthnCredential entities in the database.
func (wccb *WebauthnCredentialCreateBulk) Save(ctx context.Context) ([]*WebauthnCredential, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wccb.builders))
	nodes := make([]*WebauthnCredential, len(wccb.builders))
	mutators := make([]Mutator, len(wccb.builders))
	for i := range wccb.builders {
		func(i int, root context.Context) {
			builder := wccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebauthnCredentialMutation)
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
					_, err = mutators[i+1].Mutate(root, wccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wccb *WebauthnCredentialCreateBulk) SaveX(ctx context.Context) []*WebauthnCredential {
	v, err := wccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wccb *WebauthnCredentialCreateBulk) Exec(ctx context.Context) error {
	_, err := wccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wccb *WebauthnCredentialCreateBulk) ExecX(ctx context.Context) {
	if err := wccb.Exec(ctx); err != nil {
		panic(err)
	}
}
