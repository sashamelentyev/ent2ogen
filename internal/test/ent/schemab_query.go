// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/internal/test/ent/predicate"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemab"
)

// SchemaBQuery is the builder for querying SchemaB entities.
type SchemaBQuery struct {
	config
	ctx        *QueryContext
	order      []schemab.Order
	inters     []Interceptor
	predicates []predicate.SchemaB
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SchemaBQuery builder.
func (sb *SchemaBQuery) Where(ps ...predicate.SchemaB) *SchemaBQuery {
	sb.predicates = append(sb.predicates, ps...)
	return sb
}

// Limit the number of records to be returned by this query.
func (sb *SchemaBQuery) Limit(limit int) *SchemaBQuery {
	sb.ctx.Limit = &limit
	return sb
}

// Offset to start from.
func (sb *SchemaBQuery) Offset(offset int) *SchemaBQuery {
	sb.ctx.Offset = &offset
	return sb
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sb *SchemaBQuery) Unique(unique bool) *SchemaBQuery {
	sb.ctx.Unique = &unique
	return sb
}

// Order specifies how the records should be ordered.
func (sb *SchemaBQuery) Order(o ...schemab.Order) *SchemaBQuery {
	sb.order = append(sb.order, o...)
	return sb
}

// First returns the first SchemaB entity from the query.
// Returns a *NotFoundError when no SchemaB was found.
func (sb *SchemaBQuery) First(ctx context.Context) (*SchemaB, error) {
	nodes, err := sb.Limit(1).All(setContextOp(ctx, sb.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{schemab.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sb *SchemaBQuery) FirstX(ctx context.Context) *SchemaB {
	node, err := sb.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SchemaB ID from the query.
// Returns a *NotFoundError when no SchemaB ID was found.
func (sb *SchemaBQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sb.Limit(1).IDs(setContextOp(ctx, sb.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{schemab.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sb *SchemaBQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sb.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SchemaB entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SchemaB entity is found.
// Returns a *NotFoundError when no SchemaB entities are found.
func (sb *SchemaBQuery) Only(ctx context.Context) (*SchemaB, error) {
	nodes, err := sb.Limit(2).All(setContextOp(ctx, sb.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{schemab.Label}
	default:
		return nil, &NotSingularError{schemab.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sb *SchemaBQuery) OnlyX(ctx context.Context) *SchemaB {
	node, err := sb.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SchemaB ID in the query.
// Returns a *NotSingularError when more than one SchemaB ID is found.
// Returns a *NotFoundError when no entities are found.
func (sb *SchemaBQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sb.Limit(2).IDs(setContextOp(ctx, sb.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{schemab.Label}
	default:
		err = &NotSingularError{schemab.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sb *SchemaBQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sb.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SchemaBs.
func (sb *SchemaBQuery) All(ctx context.Context) ([]*SchemaB, error) {
	ctx = setContextOp(ctx, sb.ctx, "All")
	if err := sb.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SchemaB, *SchemaBQuery]()
	return withInterceptors[[]*SchemaB](ctx, sb, qr, sb.inters)
}

// AllX is like All, but panics if an error occurs.
func (sb *SchemaBQuery) AllX(ctx context.Context) []*SchemaB {
	nodes, err := sb.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SchemaB IDs.
func (sb *SchemaBQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sb.ctx.Unique == nil && sb.path != nil {
		sb.Unique(true)
	}
	ctx = setContextOp(ctx, sb.ctx, "IDs")
	if err = sb.Select(schemab.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sb *SchemaBQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sb.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sb *SchemaBQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sb.ctx, "Count")
	if err := sb.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sb, querierCount[*SchemaBQuery](), sb.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sb *SchemaBQuery) CountX(ctx context.Context) int {
	count, err := sb.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sb *SchemaBQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sb.ctx, "Exist")
	switch _, err := sb.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sb *SchemaBQuery) ExistX(ctx context.Context) bool {
	exist, err := sb.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SchemaBQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sb *SchemaBQuery) Clone() *SchemaBQuery {
	if sb == nil {
		return nil
	}
	return &SchemaBQuery{
		config:     sb.config,
		ctx:        sb.ctx.Clone(),
		order:      append([]schemab.Order{}, sb.order...),
		inters:     append([]Interceptor{}, sb.inters...),
		predicates: append([]predicate.SchemaB{}, sb.predicates...),
		// clone intermediate query.
		sql:  sb.sql.Clone(),
		path: sb.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sb *SchemaBQuery) GroupBy(field string, fields ...string) *SchemaBGroupBy {
	sb.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SchemaBGroupBy{build: sb}
	grbuild.flds = &sb.ctx.Fields
	grbuild.label = schemab.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sb *SchemaBQuery) Select(fields ...string) *SchemaBSelect {
	sb.ctx.Fields = append(sb.ctx.Fields, fields...)
	sbuild := &SchemaBSelect{SchemaBQuery: sb}
	sbuild.label = schemab.Label
	sbuild.flds, sbuild.scan = &sb.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SchemaBSelect configured with the given aggregations.
func (sb *SchemaBQuery) Aggregate(fns ...AggregateFunc) *SchemaBSelect {
	return sb.Select().Aggregate(fns...)
}

func (sb *SchemaBQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sb.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sb); err != nil {
				return err
			}
		}
	}
	for _, f := range sb.ctx.Fields {
		if !schemab.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sb.path != nil {
		prev, err := sb.path(ctx)
		if err != nil {
			return err
		}
		sb.sql = prev
	}
	return nil
}

func (sb *SchemaBQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SchemaB, error) {
	var (
		nodes   = []*SchemaB{}
		withFKs = sb.withFKs
		_spec   = sb.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, schemab.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SchemaB).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SchemaB{config: sb.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sb.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sb *SchemaBQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sb.querySpec()
	_spec.Node.Columns = sb.ctx.Fields
	if len(sb.ctx.Fields) > 0 {
		_spec.Unique = sb.ctx.Unique != nil && *sb.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sb.driver, _spec)
}

func (sb *SchemaBQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(schemab.Table, schemab.Columns, sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64))
	_spec.From = sb.sql
	if unique := sb.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sb.path != nil {
		_spec.Unique = true
	}
	if fields := sb.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schemab.FieldID)
		for i := range fields {
			if fields[i] != schemab.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sb.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sb.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sb.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sb.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sb *SchemaBQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sb.driver.Dialect())
	t1 := builder.Table(schemab.Table)
	columns := sb.ctx.Fields
	if len(columns) == 0 {
		columns = schemab.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sb.sql != nil {
		selector = sb.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sb.ctx.Unique != nil && *sb.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sb.predicates {
		p(selector)
	}
	for _, p := range sb.order {
		p(selector)
	}
	if offset := sb.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sb.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SchemaBGroupBy is the group-by builder for SchemaB entities.
type SchemaBGroupBy struct {
	selector
	build *SchemaBQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sbb *SchemaBGroupBy) Aggregate(fns ...AggregateFunc) *SchemaBGroupBy {
	sbb.fns = append(sbb.fns, fns...)
	return sbb
}

// Scan applies the selector query and scans the result into the given value.
func (sbb *SchemaBGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbb.build.ctx, "GroupBy")
	if err := sbb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SchemaBQuery, *SchemaBGroupBy](ctx, sbb.build, sbb, sbb.build.inters, v)
}

func (sbb *SchemaBGroupBy) sqlScan(ctx context.Context, root *SchemaBQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sbb.fns))
	for _, fn := range sbb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sbb.flds)+len(sbb.fns))
		for _, f := range *sbb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sbb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SchemaBSelect is the builder for selecting fields of SchemaB entities.
type SchemaBSelect struct {
	*SchemaBQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sb *SchemaBSelect) Aggregate(fns ...AggregateFunc) *SchemaBSelect {
	sb.fns = append(sb.fns, fns...)
	return sb
}

// Scan applies the selector query and scans the result into the given value.
func (sb *SchemaBSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sb.ctx, "Select")
	if err := sb.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SchemaBQuery, *SchemaBSelect](ctx, sb.SchemaBQuery, sb, sb.inters, v)
}

func (sb *SchemaBSelect) sqlScan(ctx context.Context, root *SchemaBQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sb.fns))
	for _, fn := range sb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sb.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
