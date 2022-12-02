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
	"github.com/ogen-go/ent2ogen/internal/test/ent/predicate"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemaa"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemab"
)

// SchemaAQuery is the builder for querying SchemaA entities.
type SchemaAQuery struct {
	config
	limit                                 *int
	offset                                *int
	unique                                *bool
	order                                 []OrderFunc
	fields                                []string
	predicates                            []predicate.SchemaA
	withEdgeSchemabUniqueRequired         *SchemaBQuery
	withEdgeSchemabUniqueRequiredBindtoBs *SchemaBQuery
	withEdgeSchemabUniqueOptional         *SchemaBQuery
	withEdgeSchemab                       *SchemaBQuery
	withFKs                               bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SchemaAQuery builder.
func (sa *SchemaAQuery) Where(ps ...predicate.SchemaA) *SchemaAQuery {
	sa.predicates = append(sa.predicates, ps...)
	return sa
}

// Limit adds a limit step to the query.
func (sa *SchemaAQuery) Limit(limit int) *SchemaAQuery {
	sa.limit = &limit
	return sa
}

// Offset adds an offset step to the query.
func (sa *SchemaAQuery) Offset(offset int) *SchemaAQuery {
	sa.offset = &offset
	return sa
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sa *SchemaAQuery) Unique(unique bool) *SchemaAQuery {
	sa.unique = &unique
	return sa
}

// Order adds an order step to the query.
func (sa *SchemaAQuery) Order(o ...OrderFunc) *SchemaAQuery {
	sa.order = append(sa.order, o...)
	return sa
}

// QueryEdgeSchemabUniqueRequired chains the current query on the "edge_schemab_unique_required" edge.
func (sa *SchemaAQuery) QueryEdgeSchemabUniqueRequired() *SchemaBQuery {
	query := &SchemaBQuery{config: sa.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sa.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sa.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, selector),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueRequiredTable, schemaa.EdgeSchemabUniqueRequiredColumn),
		)
		fromU = sqlgraph.SetNeighbors(sa.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgeSchemabUniqueRequiredBindtoBs chains the current query on the "edge_schemab_unique_required_bindto_bs" edge.
func (sa *SchemaAQuery) QueryEdgeSchemabUniqueRequiredBindtoBs() *SchemaBQuery {
	query := &SchemaBQuery{config: sa.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sa.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sa.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, selector),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueRequiredBindtoBsTable, schemaa.EdgeSchemabUniqueRequiredBindtoBsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sa.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgeSchemabUniqueOptional chains the current query on the "edge_schemab_unique_optional" edge.
func (sa *SchemaAQuery) QueryEdgeSchemabUniqueOptional() *SchemaBQuery {
	query := &SchemaBQuery{config: sa.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sa.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sa.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, selector),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueOptionalTable, schemaa.EdgeSchemabUniqueOptionalColumn),
		)
		fromU = sqlgraph.SetNeighbors(sa.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgeSchemab chains the current query on the "edge_schemab" edge.
func (sa *SchemaAQuery) QueryEdgeSchemab() *SchemaBQuery {
	query := &SchemaBQuery{config: sa.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sa.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sa.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, selector),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, schemaa.EdgeSchemabTable, schemaa.EdgeSchemabColumn),
		)
		fromU = sqlgraph.SetNeighbors(sa.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SchemaA entity from the query.
// Returns a *NotFoundError when no SchemaA was found.
func (sa *SchemaAQuery) First(ctx context.Context) (*SchemaA, error) {
	nodes, err := sa.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{schemaa.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sa *SchemaAQuery) FirstX(ctx context.Context) *SchemaA {
	node, err := sa.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SchemaA ID from the query.
// Returns a *NotFoundError when no SchemaA ID was found.
func (sa *SchemaAQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sa.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{schemaa.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sa *SchemaAQuery) FirstIDX(ctx context.Context) int {
	id, err := sa.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SchemaA entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SchemaA entity is found.
// Returns a *NotFoundError when no SchemaA entities are found.
func (sa *SchemaAQuery) Only(ctx context.Context) (*SchemaA, error) {
	nodes, err := sa.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{schemaa.Label}
	default:
		return nil, &NotSingularError{schemaa.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sa *SchemaAQuery) OnlyX(ctx context.Context) *SchemaA {
	node, err := sa.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SchemaA ID in the query.
// Returns a *NotSingularError when more than one SchemaA ID is found.
// Returns a *NotFoundError when no entities are found.
func (sa *SchemaAQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sa.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{schemaa.Label}
	default:
		err = &NotSingularError{schemaa.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sa *SchemaAQuery) OnlyIDX(ctx context.Context) int {
	id, err := sa.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SchemaAs.
func (sa *SchemaAQuery) All(ctx context.Context) ([]*SchemaA, error) {
	if err := sa.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sa.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sa *SchemaAQuery) AllX(ctx context.Context) []*SchemaA {
	nodes, err := sa.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SchemaA IDs.
func (sa *SchemaAQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sa.Select(schemaa.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sa *SchemaAQuery) IDsX(ctx context.Context) []int {
	ids, err := sa.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sa *SchemaAQuery) Count(ctx context.Context) (int, error) {
	if err := sa.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sa.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sa *SchemaAQuery) CountX(ctx context.Context) int {
	count, err := sa.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sa *SchemaAQuery) Exist(ctx context.Context) (bool, error) {
	if err := sa.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sa.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sa *SchemaAQuery) ExistX(ctx context.Context) bool {
	exist, err := sa.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SchemaAQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sa *SchemaAQuery) Clone() *SchemaAQuery {
	if sa == nil {
		return nil
	}
	return &SchemaAQuery{
		config:                                sa.config,
		limit:                                 sa.limit,
		offset:                                sa.offset,
		order:                                 append([]OrderFunc{}, sa.order...),
		predicates:                            append([]predicate.SchemaA{}, sa.predicates...),
		withEdgeSchemabUniqueRequired:         sa.withEdgeSchemabUniqueRequired.Clone(),
		withEdgeSchemabUniqueRequiredBindtoBs: sa.withEdgeSchemabUniqueRequiredBindtoBs.Clone(),
		withEdgeSchemabUniqueOptional:         sa.withEdgeSchemabUniqueOptional.Clone(),
		withEdgeSchemab:                       sa.withEdgeSchemab.Clone(),
		// clone intermediate query.
		sql:    sa.sql.Clone(),
		path:   sa.path,
		unique: sa.unique,
	}
}

// WithEdgeSchemabUniqueRequired tells the query-builder to eager-load the nodes that are connected to
// the "edge_schemab_unique_required" edge. The optional arguments are used to configure the query builder of the edge.
func (sa *SchemaAQuery) WithEdgeSchemabUniqueRequired(opts ...func(*SchemaBQuery)) *SchemaAQuery {
	query := &SchemaBQuery{config: sa.config}
	for _, opt := range opts {
		opt(query)
	}
	sa.withEdgeSchemabUniqueRequired = query
	return sa
}

// WithEdgeSchemabUniqueRequiredBindtoBs tells the query-builder to eager-load the nodes that are connected to
// the "edge_schemab_unique_required_bindto_bs" edge. The optional arguments are used to configure the query builder of the edge.
func (sa *SchemaAQuery) WithEdgeSchemabUniqueRequiredBindtoBs(opts ...func(*SchemaBQuery)) *SchemaAQuery {
	query := &SchemaBQuery{config: sa.config}
	for _, opt := range opts {
		opt(query)
	}
	sa.withEdgeSchemabUniqueRequiredBindtoBs = query
	return sa
}

// WithEdgeSchemabUniqueOptional tells the query-builder to eager-load the nodes that are connected to
// the "edge_schemab_unique_optional" edge. The optional arguments are used to configure the query builder of the edge.
func (sa *SchemaAQuery) WithEdgeSchemabUniqueOptional(opts ...func(*SchemaBQuery)) *SchemaAQuery {
	query := &SchemaBQuery{config: sa.config}
	for _, opt := range opts {
		opt(query)
	}
	sa.withEdgeSchemabUniqueOptional = query
	return sa
}

// WithEdgeSchemab tells the query-builder to eager-load the nodes that are connected to
// the "edge_schemab" edge. The optional arguments are used to configure the query builder of the edge.
func (sa *SchemaAQuery) WithEdgeSchemab(opts ...func(*SchemaBQuery)) *SchemaAQuery {
	query := &SchemaBQuery{config: sa.config}
	for _, opt := range opts {
		opt(query)
	}
	sa.withEdgeSchemab = query
	return sa
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int64 int64 `json:"int64,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SchemaA.Query().
//		GroupBy(schemaa.FieldInt64).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sa *SchemaAQuery) GroupBy(field string, fields ...string) *SchemaAGroupBy {
	grbuild := &SchemaAGroupBy{config: sa.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sa.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sa.sqlQuery(ctx), nil
	}
	grbuild.label = schemaa.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Int64 int64 `json:"int64,omitempty"`
//	}
//
//	client.SchemaA.Query().
//		Select(schemaa.FieldInt64).
//		Scan(ctx, &v)
func (sa *SchemaAQuery) Select(fields ...string) *SchemaASelect {
	sa.fields = append(sa.fields, fields...)
	selbuild := &SchemaASelect{SchemaAQuery: sa}
	selbuild.label = schemaa.Label
	selbuild.flds, selbuild.scan = &sa.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SchemaASelect configured with the given aggregations.
func (sa *SchemaAQuery) Aggregate(fns ...AggregateFunc) *SchemaASelect {
	return sa.Select().Aggregate(fns...)
}

func (sa *SchemaAQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sa.fields {
		if !schemaa.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sa.path != nil {
		prev, err := sa.path(ctx)
		if err != nil {
			return err
		}
		sa.sql = prev
	}
	return nil
}

func (sa *SchemaAQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SchemaA, error) {
	var (
		nodes       = []*SchemaA{}
		withFKs     = sa.withFKs
		_spec       = sa.querySpec()
		loadedTypes = [4]bool{
			sa.withEdgeSchemabUniqueRequired != nil,
			sa.withEdgeSchemabUniqueRequiredBindtoBs != nil,
			sa.withEdgeSchemabUniqueOptional != nil,
			sa.withEdgeSchemab != nil,
		}
	)
	if sa.withEdgeSchemabUniqueRequired != nil || sa.withEdgeSchemabUniqueRequiredBindtoBs != nil || sa.withEdgeSchemabUniqueOptional != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, schemaa.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SchemaA).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SchemaA{config: sa.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sa.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sa.withEdgeSchemabUniqueRequired; query != nil {
		if err := sa.loadEdgeSchemabUniqueRequired(ctx, query, nodes, nil,
			func(n *SchemaA, e *SchemaB) { n.Edges.EdgeSchemabUniqueRequired = e }); err != nil {
			return nil, err
		}
	}
	if query := sa.withEdgeSchemabUniqueRequiredBindtoBs; query != nil {
		if err := sa.loadEdgeSchemabUniqueRequiredBindtoBs(ctx, query, nodes, nil,
			func(n *SchemaA, e *SchemaB) { n.Edges.EdgeSchemabUniqueRequiredBindtoBs = e }); err != nil {
			return nil, err
		}
	}
	if query := sa.withEdgeSchemabUniqueOptional; query != nil {
		if err := sa.loadEdgeSchemabUniqueOptional(ctx, query, nodes, nil,
			func(n *SchemaA, e *SchemaB) { n.Edges.EdgeSchemabUniqueOptional = e }); err != nil {
			return nil, err
		}
	}
	if query := sa.withEdgeSchemab; query != nil {
		if err := sa.loadEdgeSchemab(ctx, query, nodes,
			func(n *SchemaA) { n.Edges.EdgeSchemab = []*SchemaB{} },
			func(n *SchemaA, e *SchemaB) { n.Edges.EdgeSchemab = append(n.Edges.EdgeSchemab, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sa *SchemaAQuery) loadEdgeSchemabUniqueRequired(ctx context.Context, query *SchemaBQuery, nodes []*SchemaA, init func(*SchemaA), assign func(*SchemaA, *SchemaB)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*SchemaA)
	for i := range nodes {
		if nodes[i].schemaa_edge_schemab_unique_required == nil {
			continue
		}
		fk := *nodes[i].schemaa_edge_schemab_unique_required
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(schemab.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "schemaa_edge_schemab_unique_required" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sa *SchemaAQuery) loadEdgeSchemabUniqueRequiredBindtoBs(ctx context.Context, query *SchemaBQuery, nodes []*SchemaA, init func(*SchemaA), assign func(*SchemaA, *SchemaB)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*SchemaA)
	for i := range nodes {
		if nodes[i].schemaa_edge_schemab_unique_required_bindto_bs == nil {
			continue
		}
		fk := *nodes[i].schemaa_edge_schemab_unique_required_bindto_bs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(schemab.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "schemaa_edge_schemab_unique_required_bindto_bs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sa *SchemaAQuery) loadEdgeSchemabUniqueOptional(ctx context.Context, query *SchemaBQuery, nodes []*SchemaA, init func(*SchemaA), assign func(*SchemaA, *SchemaB)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*SchemaA)
	for i := range nodes {
		if nodes[i].schemaa_edge_schemab_unique_optional == nil {
			continue
		}
		fk := *nodes[i].schemaa_edge_schemab_unique_optional
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(schemab.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "schemaa_edge_schemab_unique_optional" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sa *SchemaAQuery) loadEdgeSchemab(ctx context.Context, query *SchemaBQuery, nodes []*SchemaA, init func(*SchemaA), assign func(*SchemaA, *SchemaB)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*SchemaA)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SchemaB(func(s *sql.Selector) {
		s.Where(sql.InValues(schemaa.EdgeSchemabColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.schemaa_edge_schemab
		if fk == nil {
			return fmt.Errorf(`foreign-key "schemaa_edge_schemab" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "schemaa_edge_schemab" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sa *SchemaAQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sa.querySpec()
	_spec.Node.Columns = sa.fields
	if len(sa.fields) > 0 {
		_spec.Unique = sa.unique != nil && *sa.unique
	}
	return sqlgraph.CountNodes(ctx, sa.driver, _spec)
}

func (sa *SchemaAQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sa.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sa *SchemaAQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schemaa.Table,
			Columns: schemaa.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schemaa.FieldID,
			},
		},
		From:   sa.sql,
		Unique: true,
	}
	if unique := sa.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sa.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schemaa.FieldID)
		for i := range fields {
			if fields[i] != schemaa.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sa.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sa.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sa.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sa.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sa *SchemaAQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sa.driver.Dialect())
	t1 := builder.Table(schemaa.Table)
	columns := sa.fields
	if len(columns) == 0 {
		columns = schemaa.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sa.sql != nil {
		selector = sa.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sa.unique != nil && *sa.unique {
		selector.Distinct()
	}
	for _, p := range sa.predicates {
		p(selector)
	}
	for _, p := range sa.order {
		p(selector)
	}
	if offset := sa.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sa.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SchemaAGroupBy is the group-by builder for SchemaA entities.
type SchemaAGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sab *SchemaAGroupBy) Aggregate(fns ...AggregateFunc) *SchemaAGroupBy {
	sab.fns = append(sab.fns, fns...)
	return sab
}

// Scan applies the group-by query and scans the result into the given value.
func (sab *SchemaAGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sab.path(ctx)
	if err != nil {
		return err
	}
	sab.sql = query
	return sab.sqlScan(ctx, v)
}

func (sab *SchemaAGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sab.fields {
		if !schemaa.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sab.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sab.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sab *SchemaAGroupBy) sqlQuery() *sql.Selector {
	selector := sab.sql.Select()
	aggregation := make([]string, 0, len(sab.fns))
	for _, fn := range sab.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sab.fields)+len(sab.fns))
		for _, f := range sab.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sab.fields...)...)
}

// SchemaASelect is the builder for selecting fields of SchemaA entities.
type SchemaASelect struct {
	*SchemaAQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sa *SchemaASelect) Aggregate(fns ...AggregateFunc) *SchemaASelect {
	sa.fns = append(sa.fns, fns...)
	return sa
}

// Scan applies the selector query and scans the result into the given value.
func (sa *SchemaASelect) Scan(ctx context.Context, v any) error {
	if err := sa.prepareQuery(ctx); err != nil {
		return err
	}
	sa.sql = sa.SchemaAQuery.sqlQuery(ctx)
	return sa.sqlScan(ctx, v)
}

func (sa *SchemaASelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(sa.fns))
	for _, fn := range sa.fns {
		aggregation = append(aggregation, fn(sa.sql))
	}
	switch n := len(*sa.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		sa.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		sa.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := sa.sql.Query()
	if err := sa.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
