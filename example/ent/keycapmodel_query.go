// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
)

// KeycapModelQuery is the builder for querying KeycapModel entities.
type KeycapModelQuery struct {
	config
	ctx        *QueryContext
	order      []keycapmodel.Order
	inters     []Interceptor
	predicates []predicate.KeycapModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeycapModelQuery builder.
func (kmq *KeycapModelQuery) Where(ps ...predicate.KeycapModel) *KeycapModelQuery {
	kmq.predicates = append(kmq.predicates, ps...)
	return kmq
}

// Limit the number of records to be returned by this query.
func (kmq *KeycapModelQuery) Limit(limit int) *KeycapModelQuery {
	kmq.ctx.Limit = &limit
	return kmq
}

// Offset to start from.
func (kmq *KeycapModelQuery) Offset(offset int) *KeycapModelQuery {
	kmq.ctx.Offset = &offset
	return kmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kmq *KeycapModelQuery) Unique(unique bool) *KeycapModelQuery {
	kmq.ctx.Unique = &unique
	return kmq
}

// Order specifies how the records should be ordered.
func (kmq *KeycapModelQuery) Order(o ...keycapmodel.Order) *KeycapModelQuery {
	kmq.order = append(kmq.order, o...)
	return kmq
}

// First returns the first KeycapModel entity from the query.
// Returns a *NotFoundError when no KeycapModel was found.
func (kmq *KeycapModelQuery) First(ctx context.Context) (*KeycapModel, error) {
	nodes, err := kmq.Limit(1).All(setContextOp(ctx, kmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keycapmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kmq *KeycapModelQuery) FirstX(ctx context.Context) *KeycapModel {
	node, err := kmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KeycapModel ID from the query.
// Returns a *NotFoundError when no KeycapModel ID was found.
func (kmq *KeycapModelQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = kmq.Limit(1).IDs(setContextOp(ctx, kmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keycapmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kmq *KeycapModelQuery) FirstIDX(ctx context.Context) int64 {
	id, err := kmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KeycapModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one KeycapModel entity is found.
// Returns a *NotFoundError when no KeycapModel entities are found.
func (kmq *KeycapModelQuery) Only(ctx context.Context) (*KeycapModel, error) {
	nodes, err := kmq.Limit(2).All(setContextOp(ctx, kmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keycapmodel.Label}
	default:
		return nil, &NotSingularError{keycapmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kmq *KeycapModelQuery) OnlyX(ctx context.Context) *KeycapModel {
	node, err := kmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KeycapModel ID in the query.
// Returns a *NotSingularError when more than one KeycapModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (kmq *KeycapModelQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = kmq.Limit(2).IDs(setContextOp(ctx, kmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keycapmodel.Label}
	default:
		err = &NotSingularError{keycapmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kmq *KeycapModelQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := kmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeycapModels.
func (kmq *KeycapModelQuery) All(ctx context.Context) ([]*KeycapModel, error) {
	ctx = setContextOp(ctx, kmq.ctx, "All")
	if err := kmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*KeycapModel, *KeycapModelQuery]()
	return withInterceptors[[]*KeycapModel](ctx, kmq, qr, kmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (kmq *KeycapModelQuery) AllX(ctx context.Context) []*KeycapModel {
	nodes, err := kmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KeycapModel IDs.
func (kmq *KeycapModelQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if kmq.ctx.Unique == nil && kmq.path != nil {
		kmq.Unique(true)
	}
	ctx = setContextOp(ctx, kmq.ctx, "IDs")
	if err = kmq.Select(keycapmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kmq *KeycapModelQuery) IDsX(ctx context.Context) []int64 {
	ids, err := kmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kmq *KeycapModelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, kmq.ctx, "Count")
	if err := kmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, kmq, querierCount[*KeycapModelQuery](), kmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (kmq *KeycapModelQuery) CountX(ctx context.Context) int {
	count, err := kmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kmq *KeycapModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, kmq.ctx, "Exist")
	switch _, err := kmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (kmq *KeycapModelQuery) ExistX(ctx context.Context) bool {
	exist, err := kmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeycapModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kmq *KeycapModelQuery) Clone() *KeycapModelQuery {
	if kmq == nil {
		return nil
	}
	return &KeycapModelQuery{
		config:     kmq.config,
		ctx:        kmq.ctx.Clone(),
		order:      append([]keycapmodel.Order{}, kmq.order...),
		inters:     append([]Interceptor{}, kmq.inters...),
		predicates: append([]predicate.KeycapModel{}, kmq.predicates...),
		// clone intermediate query.
		sql:  kmq.sql.Clone(),
		path: kmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KeycapModel.Query().
//		GroupBy(keycapmodel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kmq *KeycapModelQuery) GroupBy(field string, fields ...string) *KeycapModelGroupBy {
	kmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &KeycapModelGroupBy{build: kmq}
	grbuild.flds = &kmq.ctx.Fields
	grbuild.label = keycapmodel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.KeycapModel.Query().
//		Select(keycapmodel.FieldName).
//		Scan(ctx, &v)
func (kmq *KeycapModelQuery) Select(fields ...string) *KeycapModelSelect {
	kmq.ctx.Fields = append(kmq.ctx.Fields, fields...)
	sbuild := &KeycapModelSelect{KeycapModelQuery: kmq}
	sbuild.label = keycapmodel.Label
	sbuild.flds, sbuild.scan = &kmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a KeycapModelSelect configured with the given aggregations.
func (kmq *KeycapModelQuery) Aggregate(fns ...AggregateFunc) *KeycapModelSelect {
	return kmq.Select().Aggregate(fns...)
}

func (kmq *KeycapModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range kmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, kmq); err != nil {
				return err
			}
		}
	}
	for _, f := range kmq.ctx.Fields {
		if !keycapmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kmq.path != nil {
		prev, err := kmq.path(ctx)
		if err != nil {
			return err
		}
		kmq.sql = prev
	}
	return nil
}

func (kmq *KeycapModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*KeycapModel, error) {
	var (
		nodes = []*KeycapModel{}
		_spec = kmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*KeycapModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &KeycapModel{config: kmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kmq *KeycapModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kmq.querySpec()
	_spec.Node.Columns = kmq.ctx.Fields
	if len(kmq.ctx.Fields) > 0 {
		_spec.Unique = kmq.ctx.Unique != nil && *kmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, kmq.driver, _spec)
}

func (kmq *KeycapModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(keycapmodel.Table, keycapmodel.Columns, sqlgraph.NewFieldSpec(keycapmodel.FieldID, field.TypeInt64))
	_spec.From = kmq.sql
	if unique := kmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if kmq.path != nil {
		_spec.Unique = true
	}
	if fields := kmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keycapmodel.FieldID)
		for i := range fields {
			if fields[i] != keycapmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kmq *KeycapModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kmq.driver.Dialect())
	t1 := builder.Table(keycapmodel.Table)
	columns := kmq.ctx.Fields
	if len(columns) == 0 {
		columns = keycapmodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kmq.sql != nil {
		selector = kmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kmq.ctx.Unique != nil && *kmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range kmq.predicates {
		p(selector)
	}
	for _, p := range kmq.order {
		p(selector)
	}
	if offset := kmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeycapModelGroupBy is the group-by builder for KeycapModel entities.
type KeycapModelGroupBy struct {
	selector
	build *KeycapModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kmgb *KeycapModelGroupBy) Aggregate(fns ...AggregateFunc) *KeycapModelGroupBy {
	kmgb.fns = append(kmgb.fns, fns...)
	return kmgb
}

// Scan applies the selector query and scans the result into the given value.
func (kmgb *KeycapModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kmgb.build.ctx, "GroupBy")
	if err := kmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeycapModelQuery, *KeycapModelGroupBy](ctx, kmgb.build, kmgb, kmgb.build.inters, v)
}

func (kmgb *KeycapModelGroupBy) sqlScan(ctx context.Context, root *KeycapModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(kmgb.fns))
	for _, fn := range kmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*kmgb.flds)+len(kmgb.fns))
		for _, f := range *kmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*kmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// KeycapModelSelect is the builder for selecting fields of KeycapModel entities.
type KeycapModelSelect struct {
	*KeycapModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (kms *KeycapModelSelect) Aggregate(fns ...AggregateFunc) *KeycapModelSelect {
	kms.fns = append(kms.fns, fns...)
	return kms
}

// Scan applies the selector query and scans the result into the given value.
func (kms *KeycapModelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kms.ctx, "Select")
	if err := kms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeycapModelQuery, *KeycapModelSelect](ctx, kms.KeycapModelQuery, kms, kms.inters, v)
}

func (kms *KeycapModelSelect) sqlScan(ctx context.Context, root *KeycapModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(kms.fns))
	for _, fn := range kms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*kms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
