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
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthnsessiondata"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthnsessiondataallowedcredential"
)

// WebauthnSessionDataCreate is the builder for creating a WebauthnSessionData entity.
type WebauthnSessionDataCreate struct {
	config
	mutation *WebauthnSessionDataMutation
	hooks    []Hook
}

// SetChallenge sets the "challenge" field.
func (wsdc *WebauthnSessionDataCreate) SetChallenge(s string) *WebauthnSessionDataCreate {
	wsdc.mutation.SetChallenge(s)
	return wsdc
}

// SetUserID sets the "user_id" field.
func (wsdc *WebauthnSessionDataCreate) SetUserID(u uuid.UUID) *WebauthnSessionDataCreate {
	wsdc.mutation.SetUserID(u)
	return wsdc
}

// SetUserVerification sets the "user_verification" field.
func (wsdc *WebauthnSessionDataCreate) SetUserVerification(s string) *WebauthnSessionDataCreate {
	wsdc.mutation.SetUserVerification(s)
	return wsdc
}

// SetOperation sets the "operation" field.
func (wsdc *WebauthnSessionDataCreate) SetOperation(s string) *WebauthnSessionDataCreate {
	wsdc.mutation.SetOperation(s)
	return wsdc
}

// SetCreatedAt sets the "created_at" field.
func (wsdc *WebauthnSessionDataCreate) SetCreatedAt(t time.Time) *WebauthnSessionDataCreate {
	wsdc.mutation.SetCreatedAt(t)
	return wsdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wsdc *WebauthnSessionDataCreate) SetNillableCreatedAt(t *time.Time) *WebauthnSessionDataCreate {
	if t != nil {
		wsdc.SetCreatedAt(*t)
	}
	return wsdc
}

// SetUpdatedAt sets the "updated_at" field.
func (wsdc *WebauthnSessionDataCreate) SetUpdatedAt(t time.Time) *WebauthnSessionDataCreate {
	wsdc.mutation.SetUpdatedAt(t)
	return wsdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wsdc *WebauthnSessionDataCreate) SetNillableUpdatedAt(t *time.Time) *WebauthnSessionDataCreate {
	if t != nil {
		wsdc.SetUpdatedAt(*t)
	}
	return wsdc
}

// SetID sets the "id" field.
func (wsdc *WebauthnSessionDataCreate) SetID(u uuid.UUID) *WebauthnSessionDataCreate {
	wsdc.mutation.SetID(u)
	return wsdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wsdc *WebauthnSessionDataCreate) SetNillableID(u *uuid.UUID) *WebauthnSessionDataCreate {
	if u != nil {
		wsdc.SetID(*u)
	}
	return wsdc
}

// AddWebauthnSessionDataAllowedCredentialIDs adds the "webauthn_session_data_allowed_credentials" edge to the WebauthnSessionDataAllowedCredential entity by IDs.
func (wsdc *WebauthnSessionDataCreate) AddWebauthnSessionDataAllowedCredentialIDs(ids ...uuid.UUID) *WebauthnSessionDataCreate {
	wsdc.mutation.AddWebauthnSessionDataAllowedCredentialIDs(ids...)
	return wsdc
}

// AddWebauthnSessionDataAllowedCredentials adds the "webauthn_session_data_allowed_credentials" edges to the WebauthnSessionDataAllowedCredential entity.
func (wsdc *WebauthnSessionDataCreate) AddWebauthnSessionDataAllowedCredentials(w ...*WebauthnSessionDataAllowedCredential) *WebauthnSessionDataCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsdc.AddWebauthnSessionDataAllowedCredentialIDs(ids...)
}

// Mutation returns the WebauthnSessionDataMutation object of the builder.
func (wsdc *WebauthnSessionDataCreate) Mutation() *WebauthnSessionDataMutation {
	return wsdc.mutation
}

// Save creates the WebauthnSessionData in the database.
func (wsdc *WebauthnSessionDataCreate) Save(ctx context.Context) (*WebauthnSessionData, error) {
	wsdc.defaults()
	return withHooks(ctx, wsdc.sqlSave, wsdc.mutation, wsdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wsdc *WebauthnSessionDataCreate) SaveX(ctx context.Context) *WebauthnSessionData {
	v, err := wsdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wsdc *WebauthnSessionDataCreate) Exec(ctx context.Context) error {
	_, err := wsdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdc *WebauthnSessionDataCreate) ExecX(ctx context.Context) {
	if err := wsdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsdc *WebauthnSessionDataCreate) defaults() {
	if _, ok := wsdc.mutation.CreatedAt(); !ok {
		v := webauthnsessiondata.DefaultCreatedAt()
		wsdc.mutation.SetCreatedAt(v)
	}
	if _, ok := wsdc.mutation.UpdatedAt(); !ok {
		v := webauthnsessiondata.DefaultUpdatedAt()
		wsdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wsdc.mutation.ID(); !ok {
		v := webauthnsessiondata.DefaultID()
		wsdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wsdc *WebauthnSessionDataCreate) check() error {
	if _, ok := wsdc.mutation.Challenge(); !ok {
		return &ValidationError{Name: "challenge", err: errors.New(`ent: missing required field "WebauthnSessionData.challenge"`)}
	}
	if _, ok := wsdc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "WebauthnSessionData.user_id"`)}
	}
	if _, ok := wsdc.mutation.UserVerification(); !ok {
		return &ValidationError{Name: "user_verification", err: errors.New(`ent: missing required field "WebauthnSessionData.user_verification"`)}
	}
	if _, ok := wsdc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "WebauthnSessionData.operation"`)}
	}
	if _, ok := wsdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WebauthnSessionData.created_at"`)}
	}
	if _, ok := wsdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WebauthnSessionData.updated_at"`)}
	}
	return nil
}

func (wsdc *WebauthnSessionDataCreate) sqlSave(ctx context.Context) (*WebauthnSessionData, error) {
	if err := wsdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wsdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wsdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wsdc.mutation.id = &_node.ID
	wsdc.mutation.done = true
	return _node, nil
}

func (wsdc *WebauthnSessionDataCreate) createSpec() (*WebauthnSessionData, *sqlgraph.CreateSpec) {
	var (
		_node = &WebauthnSessionData{config: wsdc.config}
		_spec = sqlgraph.NewCreateSpec(webauthnsessiondata.Table, sqlgraph.NewFieldSpec(webauthnsessiondata.FieldID, field.TypeUUID))
	)
	if id, ok := wsdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wsdc.mutation.Challenge(); ok {
		_spec.SetField(webauthnsessiondata.FieldChallenge, field.TypeString, value)
		_node.Challenge = value
	}
	if value, ok := wsdc.mutation.UserID(); ok {
		_spec.SetField(webauthnsessiondata.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := wsdc.mutation.UserVerification(); ok {
		_spec.SetField(webauthnsessiondata.FieldUserVerification, field.TypeString, value)
		_node.UserVerification = value
	}
	if value, ok := wsdc.mutation.Operation(); ok {
		_spec.SetField(webauthnsessiondata.FieldOperation, field.TypeString, value)
		_node.Operation = value
	}
	if value, ok := wsdc.mutation.CreatedAt(); ok {
		_spec.SetField(webauthnsessiondata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wsdc.mutation.UpdatedAt(); ok {
		_spec.SetField(webauthnsessiondata.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := wsdc.mutation.WebauthnSessionDataAllowedCredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthnsessiondata.WebauthnSessionDataAllowedCredentialsTable,
			Columns: []string{webauthnsessiondata.WebauthnSessionDataAllowedCredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthnsessiondataallowedcredential.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebauthnSessionDataCreateBulk is the builder for creating many WebauthnSessionData entities in bulk.
type WebauthnSessionDataCreateBulk struct {
	config
	builders []*WebauthnSessionDataCreate
}

// Save creates the WebauthnSessionData entities in the database.
func (wsdcb *WebauthnSessionDataCreateBulk) Save(ctx context.Context) ([]*WebauthnSessionData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wsdcb.builders))
	nodes := make([]*WebauthnSessionData, len(wsdcb.builders))
	mutators := make([]Mutator, len(wsdcb.builders))
	for i := range wsdcb.builders {
		func(i int, root context.Context) {
			builder := wsdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebauthnSessionDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wsdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wsdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wsdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wsdcb *WebauthnSessionDataCreateBulk) SaveX(ctx context.Context) []*WebauthnSessionData {
	v, err := wsdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wsdcb *WebauthnSessionDataCreateBulk) Exec(ctx context.Context) error {
	_, err := wsdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdcb *WebauthnSessionDataCreateBulk) ExecX(ctx context.Context) {
	if err := wsdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
