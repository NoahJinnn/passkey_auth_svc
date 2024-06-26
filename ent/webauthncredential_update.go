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
	"github.com/NoahJinnn/passkey_auth_svc/ent/predicate"
	"github.com/NoahJinnn/passkey_auth_svc/ent/user"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredential"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredentialtransport"
)

// WebauthnCredentialUpdate is the builder for updating WebauthnCredential entities.
type WebauthnCredentialUpdate struct {
	config
	hooks    []Hook
	mutation *WebauthnCredentialMutation
}

// Where appends a list predicates to the WebauthnCredentialUpdate builder.
func (wcu *WebauthnCredentialUpdate) Where(ps ...predicate.WebauthnCredential) *WebauthnCredentialUpdate {
	wcu.mutation.Where(ps...)
	return wcu
}

// SetUserID sets the "user_id" field.
func (wcu *WebauthnCredentialUpdate) SetUserID(u uuid.UUID) *WebauthnCredentialUpdate {
	wcu.mutation.SetUserID(u)
	return wcu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcu *WebauthnCredentialUpdate) SetNillableUserID(u *uuid.UUID) *WebauthnCredentialUpdate {
	if u != nil {
		wcu.SetUserID(*u)
	}
	return wcu
}

// ClearUserID clears the value of the "user_id" field.
func (wcu *WebauthnCredentialUpdate) ClearUserID() *WebauthnCredentialUpdate {
	wcu.mutation.ClearUserID()
	return wcu
}

// SetPublicKey sets the "public_key" field.
func (wcu *WebauthnCredentialUpdate) SetPublicKey(s string) *WebauthnCredentialUpdate {
	wcu.mutation.SetPublicKey(s)
	return wcu
}

// SetAttestationType sets the "attestation_type" field.
func (wcu *WebauthnCredentialUpdate) SetAttestationType(s string) *WebauthnCredentialUpdate {
	wcu.mutation.SetAttestationType(s)
	return wcu
}

// SetAaguid sets the "aaguid" field.
func (wcu *WebauthnCredentialUpdate) SetAaguid(u uuid.UUID) *WebauthnCredentialUpdate {
	wcu.mutation.SetAaguid(u)
	return wcu
}

// SetSignCount sets the "sign_count" field.
func (wcu *WebauthnCredentialUpdate) SetSignCount(i int32) *WebauthnCredentialUpdate {
	wcu.mutation.ResetSignCount()
	wcu.mutation.SetSignCount(i)
	return wcu
}

// AddSignCount adds i to the "sign_count" field.
func (wcu *WebauthnCredentialUpdate) AddSignCount(i int32) *WebauthnCredentialUpdate {
	wcu.mutation.AddSignCount(i)
	return wcu
}

// SetUpdatedAt sets the "updated_at" field.
func (wcu *WebauthnCredentialUpdate) SetUpdatedAt(t time.Time) *WebauthnCredentialUpdate {
	wcu.mutation.SetUpdatedAt(t)
	return wcu
}

// SetName sets the "name" field.
func (wcu *WebauthnCredentialUpdate) SetName(s string) *WebauthnCredentialUpdate {
	wcu.mutation.SetName(s)
	return wcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wcu *WebauthnCredentialUpdate) SetNillableName(s *string) *WebauthnCredentialUpdate {
	if s != nil {
		wcu.SetName(*s)
	}
	return wcu
}

// ClearName clears the value of the "name" field.
func (wcu *WebauthnCredentialUpdate) ClearName() *WebauthnCredentialUpdate {
	wcu.mutation.ClearName()
	return wcu
}

// SetBackupEligible sets the "backup_eligible" field.
func (wcu *WebauthnCredentialUpdate) SetBackupEligible(b bool) *WebauthnCredentialUpdate {
	wcu.mutation.SetBackupEligible(b)
	return wcu
}

// SetBackupState sets the "backup_state" field.
func (wcu *WebauthnCredentialUpdate) SetBackupState(b bool) *WebauthnCredentialUpdate {
	wcu.mutation.SetBackupState(b)
	return wcu
}

// SetLastUsedAt sets the "last_used_at" field.
func (wcu *WebauthnCredentialUpdate) SetLastUsedAt(t time.Time) *WebauthnCredentialUpdate {
	wcu.mutation.SetLastUsedAt(t)
	return wcu
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (wcu *WebauthnCredentialUpdate) SetNillableLastUsedAt(t *time.Time) *WebauthnCredentialUpdate {
	if t != nil {
		wcu.SetLastUsedAt(*t)
	}
	return wcu
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (wcu *WebauthnCredentialUpdate) ClearLastUsedAt() *WebauthnCredentialUpdate {
	wcu.mutation.ClearLastUsedAt()
	return wcu
}

// AddWebauthnCredentialTransportIDs adds the "webauthn_credential_transports" edge to the WebauthnCredentialTransport entity by IDs.
func (wcu *WebauthnCredentialUpdate) AddWebauthnCredentialTransportIDs(ids ...uuid.UUID) *WebauthnCredentialUpdate {
	wcu.mutation.AddWebauthnCredentialTransportIDs(ids...)
	return wcu
}

// AddWebauthnCredentialTransports adds the "webauthn_credential_transports" edges to the WebauthnCredentialTransport entity.
func (wcu *WebauthnCredentialUpdate) AddWebauthnCredentialTransports(w ...*WebauthnCredentialTransport) *WebauthnCredentialUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcu.AddWebauthnCredentialTransportIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (wcu *WebauthnCredentialUpdate) SetUser(u *User) *WebauthnCredentialUpdate {
	return wcu.SetUserID(u.ID)
}

// Mutation returns the WebauthnCredentialMutation object of the builder.
func (wcu *WebauthnCredentialUpdate) Mutation() *WebauthnCredentialMutation {
	return wcu.mutation
}

// ClearWebauthnCredentialTransports clears all "webauthn_credential_transports" edges to the WebauthnCredentialTransport entity.
func (wcu *WebauthnCredentialUpdate) ClearWebauthnCredentialTransports() *WebauthnCredentialUpdate {
	wcu.mutation.ClearWebauthnCredentialTransports()
	return wcu
}

// RemoveWebauthnCredentialTransportIDs removes the "webauthn_credential_transports" edge to WebauthnCredentialTransport entities by IDs.
func (wcu *WebauthnCredentialUpdate) RemoveWebauthnCredentialTransportIDs(ids ...uuid.UUID) *WebauthnCredentialUpdate {
	wcu.mutation.RemoveWebauthnCredentialTransportIDs(ids...)
	return wcu
}

// RemoveWebauthnCredentialTransports removes "webauthn_credential_transports" edges to WebauthnCredentialTransport entities.
func (wcu *WebauthnCredentialUpdate) RemoveWebauthnCredentialTransports(w ...*WebauthnCredentialTransport) *WebauthnCredentialUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcu.RemoveWebauthnCredentialTransportIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (wcu *WebauthnCredentialUpdate) ClearUser() *WebauthnCredentialUpdate {
	wcu.mutation.ClearUser()
	return wcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wcu *WebauthnCredentialUpdate) Save(ctx context.Context) (int, error) {
	wcu.defaults()
	return withHooks(ctx, wcu.sqlSave, wcu.mutation, wcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcu *WebauthnCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := wcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wcu *WebauthnCredentialUpdate) Exec(ctx context.Context) error {
	_, err := wcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcu *WebauthnCredentialUpdate) ExecX(ctx context.Context) {
	if err := wcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcu *WebauthnCredentialUpdate) defaults() {
	if _, ok := wcu.mutation.UpdatedAt(); !ok {
		v := webauthncredential.UpdateDefaultUpdatedAt()
		wcu.mutation.SetUpdatedAt(v)
	}
}

func (wcu *WebauthnCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString))
	if ps := wcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcu.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeString, value)
	}
	if value, ok := wcu.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
	}
	if value, ok := wcu.mutation.Aaguid(); ok {
		_spec.SetField(webauthncredential.FieldAaguid, field.TypeUUID, value)
	}
	if value, ok := wcu.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeInt32, value)
	}
	if value, ok := wcu.mutation.AddedSignCount(); ok {
		_spec.AddField(webauthncredential.FieldSignCount, field.TypeInt32, value)
	}
	if value, ok := wcu.mutation.UpdatedAt(); ok {
		_spec.SetField(webauthncredential.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wcu.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
	}
	if wcu.mutation.NameCleared() {
		_spec.ClearField(webauthncredential.FieldName, field.TypeString)
	}
	if value, ok := wcu.mutation.BackupEligible(); ok {
		_spec.SetField(webauthncredential.FieldBackupEligible, field.TypeBool, value)
	}
	if value, ok := wcu.mutation.BackupState(); ok {
		_spec.SetField(webauthncredential.FieldBackupState, field.TypeBool, value)
	}
	if value, ok := wcu.mutation.LastUsedAt(); ok {
		_spec.SetField(webauthncredential.FieldLastUsedAt, field.TypeTime, value)
	}
	if wcu.mutation.LastUsedAtCleared() {
		_spec.ClearField(webauthncredential.FieldLastUsedAt, field.TypeTime)
	}
	if wcu.mutation.WebauthnCredentialTransportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.RemovedWebauthnCredentialTransportsIDs(); len(nodes) > 0 && !wcu.mutation.WebauthnCredentialTransportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.WebauthnCredentialTransportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, wcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webauthncredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wcu.mutation.done = true
	return n, nil
}

// WebauthnCredentialUpdateOne is the builder for updating a single WebauthnCredential entity.
type WebauthnCredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebauthnCredentialMutation
}

// SetUserID sets the "user_id" field.
func (wcuo *WebauthnCredentialUpdateOne) SetUserID(u uuid.UUID) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetUserID(u)
	return wcuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcuo *WebauthnCredentialUpdateOne) SetNillableUserID(u *uuid.UUID) *WebauthnCredentialUpdateOne {
	if u != nil {
		wcuo.SetUserID(*u)
	}
	return wcuo
}

// ClearUserID clears the value of the "user_id" field.
func (wcuo *WebauthnCredentialUpdateOne) ClearUserID() *WebauthnCredentialUpdateOne {
	wcuo.mutation.ClearUserID()
	return wcuo
}

// SetPublicKey sets the "public_key" field.
func (wcuo *WebauthnCredentialUpdateOne) SetPublicKey(s string) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetPublicKey(s)
	return wcuo
}

// SetAttestationType sets the "attestation_type" field.
func (wcuo *WebauthnCredentialUpdateOne) SetAttestationType(s string) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetAttestationType(s)
	return wcuo
}

// SetAaguid sets the "aaguid" field.
func (wcuo *WebauthnCredentialUpdateOne) SetAaguid(u uuid.UUID) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetAaguid(u)
	return wcuo
}

// SetSignCount sets the "sign_count" field.
func (wcuo *WebauthnCredentialUpdateOne) SetSignCount(i int32) *WebauthnCredentialUpdateOne {
	wcuo.mutation.ResetSignCount()
	wcuo.mutation.SetSignCount(i)
	return wcuo
}

// AddSignCount adds i to the "sign_count" field.
func (wcuo *WebauthnCredentialUpdateOne) AddSignCount(i int32) *WebauthnCredentialUpdateOne {
	wcuo.mutation.AddSignCount(i)
	return wcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wcuo *WebauthnCredentialUpdateOne) SetUpdatedAt(t time.Time) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetUpdatedAt(t)
	return wcuo
}

// SetName sets the "name" field.
func (wcuo *WebauthnCredentialUpdateOne) SetName(s string) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetName(s)
	return wcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wcuo *WebauthnCredentialUpdateOne) SetNillableName(s *string) *WebauthnCredentialUpdateOne {
	if s != nil {
		wcuo.SetName(*s)
	}
	return wcuo
}

// ClearName clears the value of the "name" field.
func (wcuo *WebauthnCredentialUpdateOne) ClearName() *WebauthnCredentialUpdateOne {
	wcuo.mutation.ClearName()
	return wcuo
}

// SetBackupEligible sets the "backup_eligible" field.
func (wcuo *WebauthnCredentialUpdateOne) SetBackupEligible(b bool) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetBackupEligible(b)
	return wcuo
}

// SetBackupState sets the "backup_state" field.
func (wcuo *WebauthnCredentialUpdateOne) SetBackupState(b bool) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetBackupState(b)
	return wcuo
}

// SetLastUsedAt sets the "last_used_at" field.
func (wcuo *WebauthnCredentialUpdateOne) SetLastUsedAt(t time.Time) *WebauthnCredentialUpdateOne {
	wcuo.mutation.SetLastUsedAt(t)
	return wcuo
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (wcuo *WebauthnCredentialUpdateOne) SetNillableLastUsedAt(t *time.Time) *WebauthnCredentialUpdateOne {
	if t != nil {
		wcuo.SetLastUsedAt(*t)
	}
	return wcuo
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (wcuo *WebauthnCredentialUpdateOne) ClearLastUsedAt() *WebauthnCredentialUpdateOne {
	wcuo.mutation.ClearLastUsedAt()
	return wcuo
}

// AddWebauthnCredentialTransportIDs adds the "webauthn_credential_transports" edge to the WebauthnCredentialTransport entity by IDs.
func (wcuo *WebauthnCredentialUpdateOne) AddWebauthnCredentialTransportIDs(ids ...uuid.UUID) *WebauthnCredentialUpdateOne {
	wcuo.mutation.AddWebauthnCredentialTransportIDs(ids...)
	return wcuo
}

// AddWebauthnCredentialTransports adds the "webauthn_credential_transports" edges to the WebauthnCredentialTransport entity.
func (wcuo *WebauthnCredentialUpdateOne) AddWebauthnCredentialTransports(w ...*WebauthnCredentialTransport) *WebauthnCredentialUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcuo.AddWebauthnCredentialTransportIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (wcuo *WebauthnCredentialUpdateOne) SetUser(u *User) *WebauthnCredentialUpdateOne {
	return wcuo.SetUserID(u.ID)
}

// Mutation returns the WebauthnCredentialMutation object of the builder.
func (wcuo *WebauthnCredentialUpdateOne) Mutation() *WebauthnCredentialMutation {
	return wcuo.mutation
}

// ClearWebauthnCredentialTransports clears all "webauthn_credential_transports" edges to the WebauthnCredentialTransport entity.
func (wcuo *WebauthnCredentialUpdateOne) ClearWebauthnCredentialTransports() *WebauthnCredentialUpdateOne {
	wcuo.mutation.ClearWebauthnCredentialTransports()
	return wcuo
}

// RemoveWebauthnCredentialTransportIDs removes the "webauthn_credential_transports" edge to WebauthnCredentialTransport entities by IDs.
func (wcuo *WebauthnCredentialUpdateOne) RemoveWebauthnCredentialTransportIDs(ids ...uuid.UUID) *WebauthnCredentialUpdateOne {
	wcuo.mutation.RemoveWebauthnCredentialTransportIDs(ids...)
	return wcuo
}

// RemoveWebauthnCredentialTransports removes "webauthn_credential_transports" edges to WebauthnCredentialTransport entities.
func (wcuo *WebauthnCredentialUpdateOne) RemoveWebauthnCredentialTransports(w ...*WebauthnCredentialTransport) *WebauthnCredentialUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcuo.RemoveWebauthnCredentialTransportIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (wcuo *WebauthnCredentialUpdateOne) ClearUser() *WebauthnCredentialUpdateOne {
	wcuo.mutation.ClearUser()
	return wcuo
}

// Where appends a list predicates to the WebauthnCredentialUpdate builder.
func (wcuo *WebauthnCredentialUpdateOne) Where(ps ...predicate.WebauthnCredential) *WebauthnCredentialUpdateOne {
	wcuo.mutation.Where(ps...)
	return wcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wcuo *WebauthnCredentialUpdateOne) Select(field string, fields ...string) *WebauthnCredentialUpdateOne {
	wcuo.fields = append([]string{field}, fields...)
	return wcuo
}

// Save executes the query and returns the updated WebauthnCredential entity.
func (wcuo *WebauthnCredentialUpdateOne) Save(ctx context.Context) (*WebauthnCredential, error) {
	wcuo.defaults()
	return withHooks(ctx, wcuo.sqlSave, wcuo.mutation, wcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcuo *WebauthnCredentialUpdateOne) SaveX(ctx context.Context) *WebauthnCredential {
	node, err := wcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wcuo *WebauthnCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := wcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcuo *WebauthnCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := wcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcuo *WebauthnCredentialUpdateOne) defaults() {
	if _, ok := wcuo.mutation.UpdatedAt(); !ok {
		v := webauthncredential.UpdateDefaultUpdatedAt()
		wcuo.mutation.SetUpdatedAt(v)
	}
}

func (wcuo *WebauthnCredentialUpdateOne) sqlSave(ctx context.Context) (_node *WebauthnCredential, err error) {
	_spec := sqlgraph.NewUpdateSpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString))
	id, ok := wcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WebauthnCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webauthncredential.FieldID)
		for _, f := range fields {
			if !webauthncredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != webauthncredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcuo.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeString, value)
	}
	if value, ok := wcuo.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
	}
	if value, ok := wcuo.mutation.Aaguid(); ok {
		_spec.SetField(webauthncredential.FieldAaguid, field.TypeUUID, value)
	}
	if value, ok := wcuo.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeInt32, value)
	}
	if value, ok := wcuo.mutation.AddedSignCount(); ok {
		_spec.AddField(webauthncredential.FieldSignCount, field.TypeInt32, value)
	}
	if value, ok := wcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(webauthncredential.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wcuo.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
	}
	if wcuo.mutation.NameCleared() {
		_spec.ClearField(webauthncredential.FieldName, field.TypeString)
	}
	if value, ok := wcuo.mutation.BackupEligible(); ok {
		_spec.SetField(webauthncredential.FieldBackupEligible, field.TypeBool, value)
	}
	if value, ok := wcuo.mutation.BackupState(); ok {
		_spec.SetField(webauthncredential.FieldBackupState, field.TypeBool, value)
	}
	if value, ok := wcuo.mutation.LastUsedAt(); ok {
		_spec.SetField(webauthncredential.FieldLastUsedAt, field.TypeTime, value)
	}
	if wcuo.mutation.LastUsedAtCleared() {
		_spec.ClearField(webauthncredential.FieldLastUsedAt, field.TypeTime)
	}
	if wcuo.mutation.WebauthnCredentialTransportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.RemovedWebauthnCredentialTransportsIDs(); len(nodes) > 0 && !wcuo.mutation.WebauthnCredentialTransportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.WebauthnCredentialTransportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   webauthncredential.WebauthnCredentialTransportsTable,
			Columns: []string{webauthncredential.WebauthnCredentialTransportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(webauthncredentialtransport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
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
	_node = &WebauthnCredential{config: wcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webauthncredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wcuo.mutation.done = true
	return _node, nil
}
