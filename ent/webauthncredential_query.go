// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/NoahJinnn/passkey_auth_svc/ent/predicate"
	"github.com/NoahJinnn/passkey_auth_svc/ent/user"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredential"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredentialtransport"
)

// WebauthnCredentialQuery is the builder for querying WebauthnCredential entities.
type WebauthnCredentialQuery struct {
	config
	ctx                              *QueryContext
	order                            []webauthncredential.OrderOption
	inters                           []Interceptor
	predicates                       []predicate.WebauthnCredential
	withWebauthnCredentialTransports *WebauthnCredentialTransportQuery
	withUser                         *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebauthnCredentialQuery builder.
func (wcq *WebauthnCredentialQuery) Where(ps ...predicate.WebauthnCredential) *WebauthnCredentialQuery {
	wcq.predicates = append(wcq.predicates, ps...)
	return wcq
}

// Limit the number of records to be returned by this query.
func (wcq *WebauthnCredentialQuery) Limit(limit int) *WebauthnCredentialQuery {
	wcq.ctx.Limit = &limit
	return wcq
}

// Offset to start from.
func (wcq *WebauthnCredentialQuery) Offset(offset int) *WebauthnCredentialQuery {
	wcq.ctx.Offset = &offset
	return wcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wcq *WebauthnCredentialQuery) Unique(unique bool) *WebauthnCredentialQuery {
	wcq.ctx.Unique = &unique
	return wcq
}

// Order specifies how the records should be ordered.
func (wcq *WebauthnCredentialQuery) Order(o ...webauthncredential.OrderOption) *WebauthnCredentialQuery {
	wcq.order = append(wcq.order, o...)
	return wcq
}

// QueryWebauthnCredentialTransports chains the current query on the "webauthn_credential_transports" edge.
func (wcq *WebauthnCredentialQuery) QueryWebauthnCredentialTransports() *WebauthnCredentialTransportQuery {
	query := (&WebauthnCredentialTransportClient{config: wcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(webauthncredential.Table, webauthncredential.FieldID, selector),
			sqlgraph.To(webauthncredentialtransport.Table, webauthncredentialtransport.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, webauthncredential.WebauthnCredentialTransportsTable, webauthncredential.WebauthnCredentialTransportsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (wcq *WebauthnCredentialQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: wcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(webauthncredential.Table, webauthncredential.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, webauthncredential.UserTable, webauthncredential.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WebauthnCredential entity from the query.
// Returns a *NotFoundError when no WebauthnCredential was found.
func (wcq *WebauthnCredentialQuery) First(ctx context.Context) (*WebauthnCredential, error) {
	nodes, err := wcq.Limit(1).All(setContextOp(ctx, wcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{webauthncredential.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) FirstX(ctx context.Context) *WebauthnCredential {
	node, err := wcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebauthnCredential ID from the query.
// Returns a *NotFoundError when no WebauthnCredential ID was found.
func (wcq *WebauthnCredentialQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wcq.Limit(1).IDs(setContextOp(ctx, wcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webauthncredential.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) FirstIDX(ctx context.Context) string {
	id, err := wcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebauthnCredential entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebauthnCredential entity is found.
// Returns a *NotFoundError when no WebauthnCredential entities are found.
func (wcq *WebauthnCredentialQuery) Only(ctx context.Context) (*WebauthnCredential, error) {
	nodes, err := wcq.Limit(2).All(setContextOp(ctx, wcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{webauthncredential.Label}
	default:
		return nil, &NotSingularError{webauthncredential.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) OnlyX(ctx context.Context) *WebauthnCredential {
	node, err := wcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebauthnCredential ID in the query.
// Returns a *NotSingularError when more than one WebauthnCredential ID is found.
// Returns a *NotFoundError when no entities are found.
func (wcq *WebauthnCredentialQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wcq.Limit(2).IDs(setContextOp(ctx, wcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webauthncredential.Label}
	default:
		err = &NotSingularError{webauthncredential.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) OnlyIDX(ctx context.Context) string {
	id, err := wcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebauthnCredentials.
func (wcq *WebauthnCredentialQuery) All(ctx context.Context) ([]*WebauthnCredential, error) {
	ctx = setContextOp(ctx, wcq.ctx, "All")
	if err := wcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebauthnCredential, *WebauthnCredentialQuery]()
	return withInterceptors[[]*WebauthnCredential](ctx, wcq, qr, wcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) AllX(ctx context.Context) []*WebauthnCredential {
	nodes, err := wcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebauthnCredential IDs.
func (wcq *WebauthnCredentialQuery) IDs(ctx context.Context) (ids []string, err error) {
	if wcq.ctx.Unique == nil && wcq.path != nil {
		wcq.Unique(true)
	}
	ctx = setContextOp(ctx, wcq.ctx, "IDs")
	if err = wcq.Select(webauthncredential.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) IDsX(ctx context.Context) []string {
	ids, err := wcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wcq *WebauthnCredentialQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Count")
	if err := wcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wcq, querierCount[*WebauthnCredentialQuery](), wcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) CountX(ctx context.Context) int {
	count, err := wcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wcq *WebauthnCredentialQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Exist")
	switch _, err := wcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wcq *WebauthnCredentialQuery) ExistX(ctx context.Context) bool {
	exist, err := wcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebauthnCredentialQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wcq *WebauthnCredentialQuery) Clone() *WebauthnCredentialQuery {
	if wcq == nil {
		return nil
	}
	return &WebauthnCredentialQuery{
		config:                           wcq.config,
		ctx:                              wcq.ctx.Clone(),
		order:                            append([]webauthncredential.OrderOption{}, wcq.order...),
		inters:                           append([]Interceptor{}, wcq.inters...),
		predicates:                       append([]predicate.WebauthnCredential{}, wcq.predicates...),
		withWebauthnCredentialTransports: wcq.withWebauthnCredentialTransports.Clone(),
		withUser:                         wcq.withUser.Clone(),
		// clone intermediate query.
		sql:  wcq.sql.Clone(),
		path: wcq.path,
	}
}

// WithWebauthnCredentialTransports tells the query-builder to eager-load the nodes that are connected to
// the "webauthn_credential_transports" edge. The optional arguments are used to configure the query builder of the edge.
func (wcq *WebauthnCredentialQuery) WithWebauthnCredentialTransports(opts ...func(*WebauthnCredentialTransportQuery)) *WebauthnCredentialQuery {
	query := (&WebauthnCredentialTransportClient{config: wcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcq.withWebauthnCredentialTransports = query
	return wcq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (wcq *WebauthnCredentialQuery) WithUser(opts ...func(*UserQuery)) *WebauthnCredentialQuery {
	query := (&UserClient{config: wcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcq.withUser = query
	return wcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebauthnCredential.Query().
//		GroupBy(webauthncredential.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wcq *WebauthnCredentialQuery) GroupBy(field string, fields ...string) *WebauthnCredentialGroupBy {
	wcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebauthnCredentialGroupBy{build: wcq}
	grbuild.flds = &wcq.ctx.Fields
	grbuild.label = webauthncredential.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.WebauthnCredential.Query().
//		Select(webauthncredential.FieldUserID).
//		Scan(ctx, &v)
func (wcq *WebauthnCredentialQuery) Select(fields ...string) *WebauthnCredentialSelect {
	wcq.ctx.Fields = append(wcq.ctx.Fields, fields...)
	sbuild := &WebauthnCredentialSelect{WebauthnCredentialQuery: wcq}
	sbuild.label = webauthncredential.Label
	sbuild.flds, sbuild.scan = &wcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebauthnCredentialSelect configured with the given aggregations.
func (wcq *WebauthnCredentialQuery) Aggregate(fns ...AggregateFunc) *WebauthnCredentialSelect {
	return wcq.Select().Aggregate(fns...)
}

func (wcq *WebauthnCredentialQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wcq); err != nil {
				return err
			}
		}
	}
	for _, f := range wcq.ctx.Fields {
		if !webauthncredential.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wcq.path != nil {
		prev, err := wcq.path(ctx)
		if err != nil {
			return err
		}
		wcq.sql = prev
	}
	return nil
}

func (wcq *WebauthnCredentialQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebauthnCredential, error) {
	var (
		nodes       = []*WebauthnCredential{}
		_spec       = wcq.querySpec()
		loadedTypes = [2]bool{
			wcq.withWebauthnCredentialTransports != nil,
			wcq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebauthnCredential).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebauthnCredential{config: wcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wcq.withWebauthnCredentialTransports; query != nil {
		if err := wcq.loadWebauthnCredentialTransports(ctx, query, nodes,
			func(n *WebauthnCredential) { n.Edges.WebauthnCredentialTransports = []*WebauthnCredentialTransport{} },
			func(n *WebauthnCredential, e *WebauthnCredentialTransport) {
				n.Edges.WebauthnCredentialTransports = append(n.Edges.WebauthnCredentialTransports, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := wcq.withUser; query != nil {
		if err := wcq.loadUser(ctx, query, nodes, nil,
			func(n *WebauthnCredential, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wcq *WebauthnCredentialQuery) loadWebauthnCredentialTransports(ctx context.Context, query *WebauthnCredentialTransportQuery, nodes []*WebauthnCredential, init func(*WebauthnCredential), assign func(*WebauthnCredential, *WebauthnCredentialTransport)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*WebauthnCredential)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(webauthncredentialtransport.FieldWebauthnCredentialID)
	}
	query.Where(predicate.WebauthnCredentialTransport(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(webauthncredential.WebauthnCredentialTransportsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WebauthnCredentialID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "webauthn_credential_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wcq *WebauthnCredentialQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*WebauthnCredential, init func(*WebauthnCredential), assign func(*WebauthnCredential, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WebauthnCredential)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wcq *WebauthnCredentialQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wcq.querySpec()
	_spec.Node.Columns = wcq.ctx.Fields
	if len(wcq.ctx.Fields) > 0 {
		_spec.Unique = wcq.ctx.Unique != nil && *wcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wcq.driver, _spec)
}

func (wcq *WebauthnCredentialQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeString))
	_spec.From = wcq.sql
	if unique := wcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wcq.path != nil {
		_spec.Unique = true
	}
	if fields := wcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webauthncredential.FieldID)
		for i := range fields {
			if fields[i] != webauthncredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wcq.withUser != nil {
			_spec.Node.AddColumnOnce(webauthncredential.FieldUserID)
		}
	}
	if ps := wcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wcq *WebauthnCredentialQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wcq.driver.Dialect())
	t1 := builder.Table(webauthncredential.Table)
	columns := wcq.ctx.Fields
	if len(columns) == 0 {
		columns = webauthncredential.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wcq.sql != nil {
		selector = wcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wcq.ctx.Unique != nil && *wcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wcq.predicates {
		p(selector)
	}
	for _, p := range wcq.order {
		p(selector)
	}
	if offset := wcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebauthnCredentialGroupBy is the group-by builder for WebauthnCredential entities.
type WebauthnCredentialGroupBy struct {
	selector
	build *WebauthnCredentialQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wcgb *WebauthnCredentialGroupBy) Aggregate(fns ...AggregateFunc) *WebauthnCredentialGroupBy {
	wcgb.fns = append(wcgb.fns, fns...)
	return wcgb
}

// Scan applies the selector query and scans the result into the given value.
func (wcgb *WebauthnCredentialGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcgb.build.ctx, "GroupBy")
	if err := wcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebauthnCredentialQuery, *WebauthnCredentialGroupBy](ctx, wcgb.build, wcgb, wcgb.build.inters, v)
}

func (wcgb *WebauthnCredentialGroupBy) sqlScan(ctx context.Context, root *WebauthnCredentialQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wcgb.fns))
	for _, fn := range wcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wcgb.flds)+len(wcgb.fns))
		for _, f := range *wcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebauthnCredentialSelect is the builder for selecting fields of WebauthnCredential entities.
type WebauthnCredentialSelect struct {
	*WebauthnCredentialQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wcs *WebauthnCredentialSelect) Aggregate(fns ...AggregateFunc) *WebauthnCredentialSelect {
	wcs.fns = append(wcs.fns, fns...)
	return wcs
}

// Scan applies the selector query and scans the result into the given value.
func (wcs *WebauthnCredentialSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcs.ctx, "Select")
	if err := wcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebauthnCredentialQuery, *WebauthnCredentialSelect](ctx, wcs.WebauthnCredentialQuery, wcs, wcs.inters, v)
}

func (wcs *WebauthnCredentialSelect) sqlScan(ctx context.Context, root *WebauthnCredentialQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wcs.fns))
	for _, fn := range wcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
