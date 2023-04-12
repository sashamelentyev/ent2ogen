// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// SwitchModelQuery is the builder for querying SwitchModel entities.
type SwitchModelQuery struct {
	config
	ctx        *QueryContext
	order      []switchmodel.Order
	inters     []Interceptor
	predicates []predicate.SwitchModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SwitchModelQuery builder.
func (smq *SwitchModelQuery) Where(ps ...predicate.SwitchModel) *SwitchModelQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit the number of records to be returned by this query.
func (smq *SwitchModelQuery) Limit(limit int) *SwitchModelQuery {
	smq.ctx.Limit = &limit
	return smq
}

// Offset to start from.
func (smq *SwitchModelQuery) Offset(offset int) *SwitchModelQuery {
	smq.ctx.Offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *SwitchModelQuery) Unique(unique bool) *SwitchModelQuery {
	smq.ctx.Unique = &unique
	return smq
}

// Order specifies how the records should be ordered.
func (smq *SwitchModelQuery) Order(o ...switchmodel.Order) *SwitchModelQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// First returns the first SwitchModel entity from the query.
// Returns a *NotFoundError when no SwitchModel was found.
func (smq *SwitchModelQuery) First(ctx context.Context) (*SwitchModel, error) {
	nodes, err := smq.Limit(1).All(setContextOp(ctx, smq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{switchmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *SwitchModelQuery) FirstX(ctx context.Context) *SwitchModel {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SwitchModel ID from the query.
// Returns a *NotFoundError when no SwitchModel ID was found.
func (smq *SwitchModelQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(1).IDs(setContextOp(ctx, smq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{switchmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *SwitchModelQuery) FirstIDX(ctx context.Context) int64 {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SwitchModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SwitchModel entity is found.
// Returns a *NotFoundError when no SwitchModel entities are found.
func (smq *SwitchModelQuery) Only(ctx context.Context) (*SwitchModel, error) {
	nodes, err := smq.Limit(2).All(setContextOp(ctx, smq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{switchmodel.Label}
	default:
		return nil, &NotSingularError{switchmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *SwitchModelQuery) OnlyX(ctx context.Context) *SwitchModel {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SwitchModel ID in the query.
// Returns a *NotSingularError when more than one SwitchModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (smq *SwitchModelQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(2).IDs(setContextOp(ctx, smq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{switchmodel.Label}
	default:
		err = &NotSingularError{switchmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *SwitchModelQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SwitchModels.
func (smq *SwitchModelQuery) All(ctx context.Context) ([]*SwitchModel, error) {
	ctx = setContextOp(ctx, smq.ctx, "All")
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SwitchModel, *SwitchModelQuery]()
	return withInterceptors[[]*SwitchModel](ctx, smq, qr, smq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smq *SwitchModelQuery) AllX(ctx context.Context) []*SwitchModel {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SwitchModel IDs.
func (smq *SwitchModelQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if smq.ctx.Unique == nil && smq.path != nil {
		smq.Unique(true)
	}
	ctx = setContextOp(ctx, smq.ctx, "IDs")
	if err = smq.Select(switchmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *SwitchModelQuery) IDsX(ctx context.Context) []int64 {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *SwitchModelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smq.ctx, "Count")
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smq, querierCount[*SwitchModelQuery](), smq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smq *SwitchModelQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *SwitchModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smq.ctx, "Exist")
	switch _, err := smq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *SwitchModelQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SwitchModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *SwitchModelQuery) Clone() *SwitchModelQuery {
	if smq == nil {
		return nil
	}
	return &SwitchModelQuery{
		config:     smq.config,
		ctx:        smq.ctx.Clone(),
		order:      append([]switchmodel.Order{}, smq.order...),
		inters:     append([]Interceptor{}, smq.inters...),
		predicates: append([]predicate.SwitchModel{}, smq.predicates...),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
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
//	client.SwitchModel.Query().
//		GroupBy(switchmodel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (smq *SwitchModelQuery) GroupBy(field string, fields ...string) *SwitchModelGroupBy {
	smq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SwitchModelGroupBy{build: smq}
	grbuild.flds = &smq.ctx.Fields
	grbuild.label = switchmodel.Label
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
//	client.SwitchModel.Query().
//		Select(switchmodel.FieldName).
//		Scan(ctx, &v)
func (smq *SwitchModelQuery) Select(fields ...string) *SwitchModelSelect {
	smq.ctx.Fields = append(smq.ctx.Fields, fields...)
	sbuild := &SwitchModelSelect{SwitchModelQuery: smq}
	sbuild.label = switchmodel.Label
	sbuild.flds, sbuild.scan = &smq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SwitchModelSelect configured with the given aggregations.
func (smq *SwitchModelQuery) Aggregate(fns ...AggregateFunc) *SwitchModelSelect {
	return smq.Select().Aggregate(fns...)
}

func (smq *SwitchModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smq); err != nil {
				return err
			}
		}
	}
	for _, f := range smq.ctx.Fields {
		if !switchmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *SwitchModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SwitchModel, error) {
	var (
		nodes = []*SwitchModel{}
		_spec = smq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SwitchModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SwitchModel{config: smq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (smq *SwitchModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	_spec.Node.Columns = smq.ctx.Fields
	if len(smq.ctx.Fields) > 0 {
		_spec.Unique = smq.ctx.Unique != nil && *smq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *SwitchModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(switchmodel.Table, switchmodel.Columns, sqlgraph.NewFieldSpec(switchmodel.FieldID, field.TypeInt64))
	_spec.From = smq.sql
	if unique := smq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smq.path != nil {
		_spec.Unique = true
	}
	if fields := smq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, switchmodel.FieldID)
		for i := range fields {
			if fields[i] != switchmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *SwitchModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(switchmodel.Table)
	columns := smq.ctx.Fields
	if len(columns) == 0 {
		columns = switchmodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.ctx.Unique != nil && *smq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SwitchModelGroupBy is the group-by builder for SwitchModel entities.
type SwitchModelGroupBy struct {
	selector
	build *SwitchModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *SwitchModelGroupBy) Aggregate(fns ...AggregateFunc) *SwitchModelGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the selector query and scans the result into the given value.
func (smgb *SwitchModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smgb.build.ctx, "GroupBy")
	if err := smgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SwitchModelQuery, *SwitchModelGroupBy](ctx, smgb.build, smgb, smgb.build.inters, v)
}

func (smgb *SwitchModelGroupBy) sqlScan(ctx context.Context, root *SwitchModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smgb.flds)+len(smgb.fns))
		for _, f := range *smgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SwitchModelSelect is the builder for selecting fields of SwitchModel entities.
type SwitchModelSelect struct {
	*SwitchModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sms *SwitchModelSelect) Aggregate(fns ...AggregateFunc) *SwitchModelSelect {
	sms.fns = append(sms.fns, fns...)
	return sms
}

// Scan applies the selector query and scans the result into the given value.
func (sms *SwitchModelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sms.ctx, "Select")
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SwitchModelQuery, *SwitchModelSelect](ctx, sms.SwitchModelQuery, sms, sms.inters, v)
}

func (sms *SwitchModelSelect) sqlScan(ctx context.Context, root *SwitchModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sms.fns))
	for _, fn := range sms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
