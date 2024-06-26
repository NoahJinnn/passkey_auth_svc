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

// WebauthnSessionDataAllowedCredentialCreate is the builder for creating a WebauthnSessionDataAllowedCredential entity.
type WebauthnSessionDataAllowedCredentialCreate struct {
	config
	mutation *WebauthnSessionDataAllowedCredentialMutation
	hooks    []Hook
}

// SetCredentialID sets the "credential_id" field.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetCredentialID(s string) *WebauthnSessionDataAllowedCredentialCreate {
	wsdacc.mutation.SetCredentialID(s)
	return wsdacc
}

// SetWebauthnSessionDataID sets the "webauthn_session_data_id" field.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetWebauthnSessionDataID(u uuid.UUID) *WebauthnSessionDataAllowedCredentialCreate {
	wsdacc.mutation.SetWebauthnSessionDataID(u)
	return wsdacc
}

// SetNillableWebauthnSessionDataID sets the "webauthn_session_data_id" field if the given value is not nil.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetNillableWebauthnSessionDataID(u *uuid.UUID) *WebauthnSessionDataAllowedCredentialCreate {
	if u != nil {
		wsdacc.SetWebauthnSessionDataID(*u)
	}
	return wsdacc
}

// SetCreatedAt sets the "created_at" field.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetCreatedAt(t time.Time) *WebauthnSessionDataAllowedCredentialCreate {
	wsdacc.mutation.SetCreatedAt(t)
	return wsdacc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetNillableCreatedAt(t *time.Time) *WebauthnSessionDataAllowedCredentialCreate {
	if t != nil {
		wsdacc.SetCreatedAt(*t)
	}
	return wsdacc
}

// SetUpdatedAt sets the "updated_at" field.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetUpdatedAt(t time.Time) *WebauthnSessionDataAllowedCredentialCreate {
	wsdacc.mutation.SetUpdatedAt(t)
	return wsdacc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetNillableUpdatedAt(t *time.Time) *WebauthnSessionDataAllowedCredentialCreate {
	if t != nil {
		wsdacc.SetUpdatedAt(*t)
	}
	return wsdacc
}

// SetID sets the "id" field.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetID(u uuid.UUID) *WebauthnSessionDataAllowedCredentialCreate {
	wsdacc.mutation.SetID(u)
	return wsdacc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetNillableID(u *uuid.UUID) *WebauthnSessionDataAllowedCredentialCreate {
	if u != nil {
		wsdacc.SetID(*u)
	}
	return wsdacc
}

// SetWebauthnSessionData sets the "webauthn_session_data" edge to the WebauthnSessionData entity.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SetWebauthnSessionData(w *WebauthnSessionData) *WebauthnSessionDataAllowedCredentialCreate {
	return wsdacc.SetWebauthnSessionDataID(w.ID)
}

// Mutation returns the WebauthnSessionDataAllowedCredentialMutation object of the builder.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) Mutation() *WebauthnSessionDataAllowedCredentialMutation {
	return wsdacc.mutation
}

// Save creates the WebauthnSessionDataAllowedCredential in the database.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) Save(ctx context.Context) (*WebauthnSessionDataAllowedCredential, error) {
	wsdacc.defaults()
	return withHooks(ctx, wsdacc.sqlSave, wsdacc.mutation, wsdacc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) SaveX(ctx context.Context) *WebauthnSessionDataAllowedCredential {
	v, err := wsdacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) Exec(ctx context.Context) error {
	_, err := wsdacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) ExecX(ctx context.Context) {
	if err := wsdacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) defaults() {
	if _, ok := wsdacc.mutation.CreatedAt(); !ok {
		v := webauthnsessiondataallowedcredential.DefaultCreatedAt()
		wsdacc.mutation.SetCreatedAt(v)
	}
	if _, ok := wsdacc.mutation.UpdatedAt(); !ok {
		v := webauthnsessiondataallowedcredential.DefaultUpdatedAt()
		wsdacc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wsdacc.mutation.ID(); !ok {
		v := webauthnsessiondataallowedcredential.DefaultID()
		wsdacc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) check() error {
	if _, ok := wsdacc.mutation.CredentialID(); !ok {
		return &ValidationError{Name: "credential_id", err: errors.New(`ent: missing required field "WebauthnSessionDataAllowedCredential.credential_id"`)}
	}
	if _, ok := wsdacc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WebauthnSessionDataAllowedCredential.created_at"`)}
	}
	if _, ok := wsdacc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WebauthnSessionDataAllowedCredential.updated_at"`)}
	}
	return nil
}

func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) sqlSave(ctx context.Context) (*WebauthnSessionDataAllowedCredential, error) {
	if err := wsdacc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wsdacc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wsdacc.driver, _spec); err != nil {
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
	wsdacc.mutation.id = &_node.ID
	wsdacc.mutation.done = true
	return _node, nil
}

func (wsdacc *WebauthnSessionDataAllowedCredentialCreate) createSpec() (*WebauthnSessionDataAllowedCredential, *sqlgraph.CreateSpec) {
	var (
		_node = &WebauthnSessionDataAllowedCredential{config: wsdacc.config}
		_spec = sqlgraph.NewCreateSpec(webauthnsessiondataallowedcredential.Table, sqlgraph.NewFieldSpec(webauthnsessiondataallowedcredential.FieldID, field.TypeUUID))
	)
	if id, ok := wsdacc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wsdacc.mutation.CredentialID(); ok {
		_spec.SetField(webauthnsessiondataallowedcredential.FieldCredentialID, field.TypeString, value)
		_node.CredentialID = value
	}
	if value, ok := wsdacc.mutation.CreatedAt(); ok {
		_spec.SetField(webauthnsessiondataallowedcredential.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wsdacc.mutation.UpdatedAt(); ok {
		_spec.SetField(webauthnsessiondataallowedcredential.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := wsdacc.mutation.WebauthnSessionDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthnsessiondataallowedcredential.WebauthnSessionDataTable,
			Columns: []string{webauthnsessiondataallowedcredential.WebauthnSessionDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthnsessiondata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WebauthnSessionDataID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebauthnSessionDataAllowedCredentialCreateBulk is the builder for creating many WebauthnSessionDataAllowedCredential entities in bulk.
type WebauthnSessionDataAllowedCredentialCreateBulk struct {
	config
	builders []*WebauthnSessionDataAllowedCredentialCreate
}

// Save creates the WebauthnSessionDataAllowedCredential entities in the database.
func (wsdaccb *WebauthnSessionDataAllowedCredentialCreateBulk) Save(ctx context.Context) ([]*WebauthnSessionDataAllowedCredential, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wsdaccb.builders))
	nodes := make([]*WebauthnSessionDataAllowedCredential, len(wsdaccb.builders))
	mutators := make([]Mutator, len(wsdaccb.builders))
	for i := range wsdaccb.builders {
		func(i int, root context.Context) {
			builder := wsdaccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebauthnSessionDataAllowedCredentialMutation)
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
					_, err = mutators[i+1].Mutate(root, wsdaccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wsdaccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wsdaccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wsdaccb *WebauthnSessionDataAllowedCredentialCreateBulk) SaveX(ctx context.Context) []*WebauthnSessionDataAllowedCredential {
	v, err := wsdaccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wsdaccb *WebauthnSessionDataAllowedCredentialCreateBulk) Exec(ctx context.Context) error {
	_, err := wsdaccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdaccb *WebauthnSessionDataAllowedCredentialCreateBulk) ExecX(ctx context.Context) {
	if err := wsdaccb.Exec(ctx); err != nil {
		panic(err)
	}
}
