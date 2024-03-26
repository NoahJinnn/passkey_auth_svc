// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NoahJinnn/passkey_auth_svc/ent/jwk"
	"github.com/NoahJinnn/passkey_auth_svc/ent/predicate"
)

// JwkQuery is the builder for querying Jwk entities.
type JwkQuery struct {
	config
	ctx        *QueryContext
	order      []jwk.OrderOption
	inters     []Interceptor
	predicates []predicate.Jwk
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JwkQuery builder.
func (jq *JwkQuery) Where(ps ...predicate.Jwk) *JwkQuery {
	jq.predicates = append(jq.predicates, ps...)
	return jq
}

// Limit the number of records to be returned by this query.
func (jq *JwkQuery) Limit(limit int) *JwkQuery {
	jq.ctx.Limit = &limit
	return jq
}

// Offset to start from.
func (jq *JwkQuery) Offset(offset int) *JwkQuery {
	jq.ctx.Offset = &offset
	return jq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jq *JwkQuery) Unique(unique bool) *JwkQuery {
	jq.ctx.Unique = &unique
	return jq
}

// Order specifies how the records should be ordered.
func (jq *JwkQuery) Order(o ...jwk.OrderOption) *JwkQuery {
	jq.order = append(jq.order, o...)
	return jq
}

// First returns the first Jwk entity from the query.
// Returns a *NotFoundError when no Jwk was found.
func (jq *JwkQuery) First(ctx context.Context) (*Jwk, error) {
	nodes, err := jq.Limit(1).All(setContextOp(ctx, jq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jwk.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jq *JwkQuery) FirstX(ctx context.Context) *Jwk {
	node, err := jq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Jwk ID from the query.
// Returns a *NotFoundError when no Jwk ID was found.
func (jq *JwkQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = jq.Limit(1).IDs(setContextOp(ctx, jq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jwk.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jq *JwkQuery) FirstIDX(ctx context.Context) uint {
	id, err := jq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Jwk entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Jwk entity is found.
// Returns a *NotFoundError when no Jwk entities are found.
func (jq *JwkQuery) Only(ctx context.Context) (*Jwk, error) {
	nodes, err := jq.Limit(2).All(setContextOp(ctx, jq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jwk.Label}
	default:
		return nil, &NotSingularError{jwk.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jq *JwkQuery) OnlyX(ctx context.Context) *Jwk {
	node, err := jq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Jwk ID in the query.
// Returns a *NotSingularError when more than one Jwk ID is found.
// Returns a *NotFoundError when no entities are found.
func (jq *JwkQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = jq.Limit(2).IDs(setContextOp(ctx, jq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jwk.Label}
	default:
		err = &NotSingularError{jwk.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jq *JwkQuery) OnlyIDX(ctx context.Context) uint {
	id, err := jq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Jwks.
func (jq *JwkQuery) All(ctx context.Context) ([]*Jwk, error) {
	ctx = setContextOp(ctx, jq.ctx, "All")
	if err := jq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Jwk, *JwkQuery]()
	return withInterceptors[[]*Jwk](ctx, jq, qr, jq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jq *JwkQuery) AllX(ctx context.Context) []*Jwk {
	nodes, err := jq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Jwk IDs.
func (jq *JwkQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if jq.ctx.Unique == nil && jq.path != nil {
		jq.Unique(true)
	}
	ctx = setContextOp(ctx, jq.ctx, "IDs")
	if err = jq.Select(jwk.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jq *JwkQuery) IDsX(ctx context.Context) []uint {
	ids, err := jq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jq *JwkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jq.ctx, "Count")
	if err := jq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jq, querierCount[*JwkQuery](), jq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jq *JwkQuery) CountX(ctx context.Context) int {
	count, err := jq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jq *JwkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jq.ctx, "Exist")
	switch _, err := jq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jq *JwkQuery) ExistX(ctx context.Context) bool {
	exist, err := jq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JwkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jq *JwkQuery) Clone() *JwkQuery {
	if jq == nil {
		return nil
	}
	return &JwkQuery{
		config:     jq.config,
		ctx:        jq.ctx.Clone(),
		order:      append([]jwk.OrderOption{}, jq.order...),
		inters:     append([]Interceptor{}, jq.inters...),
		predicates: append([]predicate.Jwk{}, jq.predicates...),
		// clone intermediate query.
		sql:  jq.sql.Clone(),
		path: jq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		KeyData string `json:"key_data,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Jwk.Query().
//		GroupBy(jwk.FieldKeyData).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jq *JwkQuery) GroupBy(field string, fields ...string) *JwkGroupBy {
	jq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JwkGroupBy{build: jq}
	grbuild.flds = &jq.ctx.Fields
	grbuild.label = jwk.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		KeyData string `json:"key_data,omitempty"`
//	}
//
//	client.Jwk.Query().
//		Select(jwk.FieldKeyData).
//		Scan(ctx, &v)
func (jq *JwkQuery) Select(fields ...string) *JwkSelect {
	jq.ctx.Fields = append(jq.ctx.Fields, fields...)
	sbuild := &JwkSelect{JwkQuery: jq}
	sbuild.label = jwk.Label
	sbuild.flds, sbuild.scan = &jq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JwkSelect configured with the given aggregations.
func (jq *JwkQuery) Aggregate(fns ...AggregateFunc) *JwkSelect {
	return jq.Select().Aggregate(fns...)
}

func (jq *JwkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jq); err != nil {
				return err
			}
		}
	}
	for _, f := range jq.ctx.Fields {
		if !jwk.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jq.path != nil {
		prev, err := jq.path(ctx)
		if err != nil {
			return err
		}
		jq.sql = prev
	}
	return nil
}

func (jq *JwkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Jwk, error) {
	var (
		nodes = []*Jwk{}
		_spec = jq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Jwk).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Jwk{config: jq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (jq *JwkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jq.querySpec()
	_spec.Node.Columns = jq.ctx.Fields
	if len(jq.ctx.Fields) > 0 {
		_spec.Unique = jq.ctx.Unique != nil && *jq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jq.driver, _spec)
}

func (jq *JwkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jwk.Table, jwk.Columns, sqlgraph.NewFieldSpec(jwk.FieldID, field.TypeUint))
	_spec.From = jq.sql
	if unique := jq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jq.path != nil {
		_spec.Unique = true
	}
	if fields := jq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jwk.FieldID)
		for i := range fields {
			if fields[i] != jwk.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jq *JwkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jq.driver.Dialect())
	t1 := builder.Table(jwk.Table)
	columns := jq.ctx.Fields
	if len(columns) == 0 {
		columns = jwk.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jq.sql != nil {
		selector = jq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jq.ctx.Unique != nil && *jq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jq.predicates {
		p(selector)
	}
	for _, p := range jq.order {
		p(selector)
	}
	if offset := jq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JwkGroupBy is the group-by builder for Jwk entities.
type JwkGroupBy struct {
	selector
	build *JwkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jgb *JwkGroupBy) Aggregate(fns ...AggregateFunc) *JwkGroupBy {
	jgb.fns = append(jgb.fns, fns...)
	return jgb
}

// Scan applies the selector query and scans the result into the given value.
func (jgb *JwkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jgb.build.ctx, "GroupBy")
	if err := jgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JwkQuery, *JwkGroupBy](ctx, jgb.build, jgb, jgb.build.inters, v)
}

func (jgb *JwkGroupBy) sqlScan(ctx context.Context, root *JwkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jgb.fns))
	for _, fn := range jgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jgb.flds)+len(jgb.fns))
		for _, f := range *jgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JwkSelect is the builder for selecting fields of Jwk entities.
type JwkSelect struct {
	*JwkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (js *JwkSelect) Aggregate(fns ...AggregateFunc) *JwkSelect {
	js.fns = append(js.fns, fns...)
	return js
}

// Scan applies the selector query and scans the result into the given value.
func (js *JwkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, js.ctx, "Select")
	if err := js.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JwkQuery, *JwkSelect](ctx, js.JwkQuery, js, js.inters, v)
}

func (js *JwkSelect) sqlScan(ctx context.Context, root *JwkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(js.fns))
	for _, fn := range js.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*js.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := js.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
