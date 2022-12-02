// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/keyboard"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// KeyboardQuery is the builder for querying Keyboard entities.
type KeyboardQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Keyboard
	withSwitches *SwitchModelQuery
	withKeycaps  *KeycapModelQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeyboardQuery builder.
func (kq *KeyboardQuery) Where(ps ...predicate.Keyboard) *KeyboardQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit adds a limit step to the query.
func (kq *KeyboardQuery) Limit(limit int) *KeyboardQuery {
	kq.limit = &limit
	return kq
}

// Offset adds an offset step to the query.
func (kq *KeyboardQuery) Offset(offset int) *KeyboardQuery {
	kq.offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KeyboardQuery) Unique(unique bool) *KeyboardQuery {
	kq.unique = &unique
	return kq
}

// Order adds an order step to the query.
func (kq *KeyboardQuery) Order(o ...OrderFunc) *KeyboardQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// QuerySwitches chains the current query on the "switches" edge.
func (kq *KeyboardQuery) QuerySwitches() *SwitchModelQuery {
	query := &SwitchModelQuery{config: kq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(keyboard.Table, keyboard.FieldID, selector),
			sqlgraph.To(switchmodel.Table, switchmodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, keyboard.SwitchesTable, keyboard.SwitchesColumn),
		)
		fromU = sqlgraph.SetNeighbors(kq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKeycaps chains the current query on the "keycaps" edge.
func (kq *KeyboardQuery) QueryKeycaps() *KeycapModelQuery {
	query := &KeycapModelQuery{config: kq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(keyboard.Table, keyboard.FieldID, selector),
			sqlgraph.To(keycapmodel.Table, keycapmodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, keyboard.KeycapsTable, keyboard.KeycapsColumn),
		)
		fromU = sqlgraph.SetNeighbors(kq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Keyboard entity from the query.
// Returns a *NotFoundError when no Keyboard was found.
func (kq *KeyboardQuery) First(ctx context.Context) (*Keyboard, error) {
	nodes, err := kq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keyboard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KeyboardQuery) FirstX(ctx context.Context) *Keyboard {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Keyboard ID from the query.
// Returns a *NotFoundError when no Keyboard ID was found.
func (kq *KeyboardQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = kq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keyboard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KeyboardQuery) FirstIDX(ctx context.Context) int64 {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Keyboard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Keyboard entity is found.
// Returns a *NotFoundError when no Keyboard entities are found.
func (kq *KeyboardQuery) Only(ctx context.Context) (*Keyboard, error) {
	nodes, err := kq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keyboard.Label}
	default:
		return nil, &NotSingularError{keyboard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KeyboardQuery) OnlyX(ctx context.Context) *Keyboard {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Keyboard ID in the query.
// Returns a *NotSingularError when more than one Keyboard ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KeyboardQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = kq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keyboard.Label}
	default:
		err = &NotSingularError{keyboard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KeyboardQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Keyboards.
func (kq *KeyboardQuery) All(ctx context.Context) ([]*Keyboard, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kq *KeyboardQuery) AllX(ctx context.Context) []*Keyboard {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Keyboard IDs.
func (kq *KeyboardQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := kq.Select(keyboard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KeyboardQuery) IDsX(ctx context.Context) []int64 {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KeyboardQuery) Count(ctx context.Context) (int, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KeyboardQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KeyboardQuery) Exist(ctx context.Context) (bool, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KeyboardQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeyboardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KeyboardQuery) Clone() *KeyboardQuery {
	if kq == nil {
		return nil
	}
	return &KeyboardQuery{
		config:       kq.config,
		limit:        kq.limit,
		offset:       kq.offset,
		order:        append([]OrderFunc{}, kq.order...),
		predicates:   append([]predicate.Keyboard{}, kq.predicates...),
		withSwitches: kq.withSwitches.Clone(),
		withKeycaps:  kq.withKeycaps.Clone(),
		// clone intermediate query.
		sql:    kq.sql.Clone(),
		path:   kq.path,
		unique: kq.unique,
	}
}

// WithSwitches tells the query-builder to eager-load the nodes that are connected to
// the "switches" edge. The optional arguments are used to configure the query builder of the edge.
func (kq *KeyboardQuery) WithSwitches(opts ...func(*SwitchModelQuery)) *KeyboardQuery {
	query := &SwitchModelQuery{config: kq.config}
	for _, opt := range opts {
		opt(query)
	}
	kq.withSwitches = query
	return kq
}

// WithKeycaps tells the query-builder to eager-load the nodes that are connected to
// the "keycaps" edge. The optional arguments are used to configure the query builder of the edge.
func (kq *KeyboardQuery) WithKeycaps(opts ...func(*KeycapModelQuery)) *KeyboardQuery {
	query := &KeycapModelQuery{config: kq.config}
	for _, opt := range opts {
		opt(query)
	}
	kq.withKeycaps = query
	return kq
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
//	client.Keyboard.Query().
//		GroupBy(keyboard.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kq *KeyboardQuery) GroupBy(field string, fields ...string) *KeyboardGroupBy {
	grbuild := &KeyboardGroupBy{config: kq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kq.sqlQuery(ctx), nil
	}
	grbuild.label = keyboard.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Keyboard.Query().
//		Select(keyboard.FieldName).
//		Scan(ctx, &v)
func (kq *KeyboardQuery) Select(fields ...string) *KeyboardSelect {
	kq.fields = append(kq.fields, fields...)
	selbuild := &KeyboardSelect{KeyboardQuery: kq}
	selbuild.label = keyboard.Label
	selbuild.flds, selbuild.scan = &kq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a KeyboardSelect configured with the given aggregations.
func (kq *KeyboardQuery) Aggregate(fns ...AggregateFunc) *KeyboardSelect {
	return kq.Select().Aggregate(fns...)
}

func (kq *KeyboardQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kq.fields {
		if !keyboard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	return nil
}

func (kq *KeyboardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Keyboard, error) {
	var (
		nodes       = []*Keyboard{}
		withFKs     = kq.withFKs
		_spec       = kq.querySpec()
		loadedTypes = [2]bool{
			kq.withSwitches != nil,
			kq.withKeycaps != nil,
		}
	)
	if kq.withSwitches != nil || kq.withKeycaps != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, keyboard.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Keyboard).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Keyboard{config: kq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := kq.withSwitches; query != nil {
		if err := kq.loadSwitches(ctx, query, nodes, nil,
			func(n *Keyboard, e *SwitchModel) { n.Edges.Switches = e }); err != nil {
			return nil, err
		}
	}
	if query := kq.withKeycaps; query != nil {
		if err := kq.loadKeycaps(ctx, query, nodes, nil,
			func(n *Keyboard, e *KeycapModel) { n.Edges.Keycaps = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (kq *KeyboardQuery) loadSwitches(ctx context.Context, query *SwitchModelQuery, nodes []*Keyboard, init func(*Keyboard), assign func(*Keyboard, *SwitchModel)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Keyboard)
	for i := range nodes {
		if nodes[i].keyboard_switches == nil {
			continue
		}
		fk := *nodes[i].keyboard_switches
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(switchmodel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "keyboard_switches" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (kq *KeyboardQuery) loadKeycaps(ctx context.Context, query *KeycapModelQuery, nodes []*Keyboard, init func(*Keyboard), assign func(*Keyboard, *KeycapModel)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Keyboard)
	for i := range nodes {
		if nodes[i].keyboard_keycaps == nil {
			continue
		}
		fk := *nodes[i].keyboard_keycaps
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(keycapmodel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "keyboard_keycaps" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (kq *KeyboardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	_spec.Node.Columns = kq.fields
	if len(kq.fields) > 0 {
		_spec.Unique = kq.unique != nil && *kq.unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KeyboardQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := kq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (kq *KeyboardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keyboard.Table,
			Columns: keyboard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: keyboard.FieldID,
			},
		},
		From:   kq.sql,
		Unique: true,
	}
	if unique := kq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keyboard.FieldID)
		for i := range fields {
			if fields[i] != keyboard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KeyboardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(keyboard.Table)
	columns := kq.fields
	if len(columns) == 0 {
		columns = keyboard.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.unique != nil && *kq.unique {
		selector.Distinct()
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeyboardGroupBy is the group-by builder for Keyboard entities.
type KeyboardGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KeyboardGroupBy) Aggregate(fns ...AggregateFunc) *KeyboardGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kgb *KeyboardGroupBy) Scan(ctx context.Context, v any) error {
	query, err := kgb.path(ctx)
	if err != nil {
		return err
	}
	kgb.sql = query
	return kgb.sqlScan(ctx, v)
}

func (kgb *KeyboardGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range kgb.fields {
		if !keyboard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kgb *KeyboardGroupBy) sqlQuery() *sql.Selector {
	selector := kgb.sql.Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kgb.fields)+len(kgb.fns))
		for _, f := range kgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kgb.fields...)...)
}

// KeyboardSelect is the builder for selecting fields of Keyboard entities.
type KeyboardSelect struct {
	*KeyboardQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ks *KeyboardSelect) Aggregate(fns ...AggregateFunc) *KeyboardSelect {
	ks.fns = append(ks.fns, fns...)
	return ks
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KeyboardSelect) Scan(ctx context.Context, v any) error {
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	ks.sql = ks.KeyboardQuery.sqlQuery(ctx)
	return ks.sqlScan(ctx, v)
}

func (ks *KeyboardSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ks.fns))
	for _, fn := range ks.fns {
		aggregation = append(aggregation, fn(ks.sql))
	}
	switch n := len(*ks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ks.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ks.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ks.sql.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
