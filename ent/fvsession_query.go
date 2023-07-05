// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/fvsession"
	"github.com/hellohq/hqservice/ent/predicate"
	"github.com/hellohq/hqservice/ent/user"
)

// FvSessionQuery is the builder for querying FvSession entities.
type FvSessionQuery struct {
	config
	ctx        *QueryContext
	order      []fvsession.OrderOption
	inters     []Interceptor
	predicates []predicate.FvSession
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FvSessionQuery builder.
func (fsq *FvSessionQuery) Where(ps ...predicate.FvSession) *FvSessionQuery {
	fsq.predicates = append(fsq.predicates, ps...)
	return fsq
}

// Limit the number of records to be returned by this query.
func (fsq *FvSessionQuery) Limit(limit int) *FvSessionQuery {
	fsq.ctx.Limit = &limit
	return fsq
}

// Offset to start from.
func (fsq *FvSessionQuery) Offset(offset int) *FvSessionQuery {
	fsq.ctx.Offset = &offset
	return fsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fsq *FvSessionQuery) Unique(unique bool) *FvSessionQuery {
	fsq.ctx.Unique = &unique
	return fsq
}

// Order specifies how the records should be ordered.
func (fsq *FvSessionQuery) Order(o ...fvsession.OrderOption) *FvSessionQuery {
	fsq.order = append(fsq.order, o...)
	return fsq
}

// QueryUser chains the current query on the "user" edge.
func (fsq *FvSessionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: fsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fvsession.Table, fvsession.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, fvsession.UserTable, fvsession.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FvSession entity from the query.
// Returns a *NotFoundError when no FvSession was found.
func (fsq *FvSessionQuery) First(ctx context.Context) (*FvSession, error) {
	nodes, err := fsq.Limit(1).All(setContextOp(ctx, fsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fvsession.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fsq *FvSessionQuery) FirstX(ctx context.Context) *FvSession {
	node, err := fsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FvSession ID from the query.
// Returns a *NotFoundError when no FvSession ID was found.
func (fsq *FvSessionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fsq.Limit(1).IDs(setContextOp(ctx, fsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fvsession.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fsq *FvSessionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FvSession entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FvSession entity is found.
// Returns a *NotFoundError when no FvSession entities are found.
func (fsq *FvSessionQuery) Only(ctx context.Context) (*FvSession, error) {
	nodes, err := fsq.Limit(2).All(setContextOp(ctx, fsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fvsession.Label}
	default:
		return nil, &NotSingularError{fvsession.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fsq *FvSessionQuery) OnlyX(ctx context.Context) *FvSession {
	node, err := fsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FvSession ID in the query.
// Returns a *NotSingularError when more than one FvSession ID is found.
// Returns a *NotFoundError when no entities are found.
func (fsq *FvSessionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fsq.Limit(2).IDs(setContextOp(ctx, fsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fvsession.Label}
	default:
		err = &NotSingularError{fvsession.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fsq *FvSessionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FvSessions.
func (fsq *FvSessionQuery) All(ctx context.Context) ([]*FvSession, error) {
	ctx = setContextOp(ctx, fsq.ctx, "All")
	if err := fsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FvSession, *FvSessionQuery]()
	return withInterceptors[[]*FvSession](ctx, fsq, qr, fsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fsq *FvSessionQuery) AllX(ctx context.Context) []*FvSession {
	nodes, err := fsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FvSession IDs.
func (fsq *FvSessionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if fsq.ctx.Unique == nil && fsq.path != nil {
		fsq.Unique(true)
	}
	ctx = setContextOp(ctx, fsq.ctx, "IDs")
	if err = fsq.Select(fvsession.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fsq *FvSessionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fsq *FvSessionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fsq.ctx, "Count")
	if err := fsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fsq, querierCount[*FvSessionQuery](), fsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fsq *FvSessionQuery) CountX(ctx context.Context) int {
	count, err := fsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fsq *FvSessionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fsq.ctx, "Exist")
	switch _, err := fsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fsq *FvSessionQuery) ExistX(ctx context.Context) bool {
	exist, err := fsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FvSessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fsq *FvSessionQuery) Clone() *FvSessionQuery {
	if fsq == nil {
		return nil
	}
	return &FvSessionQuery{
		config:     fsq.config,
		ctx:        fsq.ctx.Clone(),
		order:      append([]fvsession.OrderOption{}, fsq.order...),
		inters:     append([]Interceptor{}, fsq.inters...),
		predicates: append([]predicate.FvSession{}, fsq.predicates...),
		withUser:   fsq.withUser.Clone(),
		// clone intermediate query.
		sql:  fsq.sql.Clone(),
		path: fsq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FvSessionQuery) WithUser(opts ...func(*UserQuery)) *FvSessionQuery {
	query := (&UserClient{config: fsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fsq.withUser = query
	return fsq
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
//	client.FvSession.Query().
//		GroupBy(fvsession.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fsq *FvSessionQuery) GroupBy(field string, fields ...string) *FvSessionGroupBy {
	fsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FvSessionGroupBy{build: fsq}
	grbuild.flds = &fsq.ctx.Fields
	grbuild.label = fvsession.Label
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
//	client.FvSession.Query().
//		Select(fvsession.FieldUserID).
//		Scan(ctx, &v)
func (fsq *FvSessionQuery) Select(fields ...string) *FvSessionSelect {
	fsq.ctx.Fields = append(fsq.ctx.Fields, fields...)
	sbuild := &FvSessionSelect{FvSessionQuery: fsq}
	sbuild.label = fvsession.Label
	sbuild.flds, sbuild.scan = &fsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FvSessionSelect configured with the given aggregations.
func (fsq *FvSessionQuery) Aggregate(fns ...AggregateFunc) *FvSessionSelect {
	return fsq.Select().Aggregate(fns...)
}

func (fsq *FvSessionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fsq); err != nil {
				return err
			}
		}
	}
	for _, f := range fsq.ctx.Fields {
		if !fvsession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fsq.path != nil {
		prev, err := fsq.path(ctx)
		if err != nil {
			return err
		}
		fsq.sql = prev
	}
	return nil
}

func (fsq *FvSessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FvSession, error) {
	var (
		nodes       = []*FvSession{}
		_spec       = fsq.querySpec()
		loadedTypes = [1]bool{
			fsq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FvSession).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FvSession{config: fsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fsq.withUser; query != nil {
		if err := fsq.loadUser(ctx, query, nodes, nil,
			func(n *FvSession, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fsq *FvSessionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*FvSession, init func(*FvSession), assign func(*FvSession, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FvSession)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
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

func (fsq *FvSessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fsq.querySpec()
	_spec.Node.Columns = fsq.ctx.Fields
	if len(fsq.ctx.Fields) > 0 {
		_spec.Unique = fsq.ctx.Unique != nil && *fsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fsq.driver, _spec)
}

func (fsq *FvSessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fvsession.Table, fvsession.Columns, sqlgraph.NewFieldSpec(fvsession.FieldID, field.TypeUUID))
	_spec.From = fsq.sql
	if unique := fsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fsq.path != nil {
		_spec.Unique = true
	}
	if fields := fsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fvsession.FieldID)
		for i := range fields {
			if fields[i] != fvsession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fsq.withUser != nil {
			_spec.Node.AddColumnOnce(fvsession.FieldUserID)
		}
	}
	if ps := fsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fsq *FvSessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fsq.driver.Dialect())
	t1 := builder.Table(fvsession.Table)
	columns := fsq.ctx.Fields
	if len(columns) == 0 {
		columns = fvsession.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fsq.sql != nil {
		selector = fsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fsq.ctx.Unique != nil && *fsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fsq.predicates {
		p(selector)
	}
	for _, p := range fsq.order {
		p(selector)
	}
	if offset := fsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FvSessionGroupBy is the group-by builder for FvSession entities.
type FvSessionGroupBy struct {
	selector
	build *FvSessionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fsgb *FvSessionGroupBy) Aggregate(fns ...AggregateFunc) *FvSessionGroupBy {
	fsgb.fns = append(fsgb.fns, fns...)
	return fsgb
}

// Scan applies the selector query and scans the result into the given value.
func (fsgb *FvSessionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fsgb.build.ctx, "GroupBy")
	if err := fsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FvSessionQuery, *FvSessionGroupBy](ctx, fsgb.build, fsgb, fsgb.build.inters, v)
}

func (fsgb *FvSessionGroupBy) sqlScan(ctx context.Context, root *FvSessionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fsgb.fns))
	for _, fn := range fsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fsgb.flds)+len(fsgb.fns))
		for _, f := range *fsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FvSessionSelect is the builder for selecting fields of FvSession entities.
type FvSessionSelect struct {
	*FvSessionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fss *FvSessionSelect) Aggregate(fns ...AggregateFunc) *FvSessionSelect {
	fss.fns = append(fss.fns, fns...)
	return fss
}

// Scan applies the selector query and scans the result into the given value.
func (fss *FvSessionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fss.ctx, "Select")
	if err := fss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FvSessionQuery, *FvSessionSelect](ctx, fss.FvSessionQuery, fss, fss.inters, v)
}

func (fss *FvSessionSelect) sqlScan(ctx context.Context, root *FvSessionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fss.fns))
	for _, fn := range fss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}