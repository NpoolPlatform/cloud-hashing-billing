// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodincoming"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodIncomingQuery is the builder for querying GoodIncoming entities.
type GoodIncomingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodIncoming
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodIncomingQuery builder.
func (giq *GoodIncomingQuery) Where(ps ...predicate.GoodIncoming) *GoodIncomingQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GoodIncomingQuery) Limit(limit int) *GoodIncomingQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GoodIncomingQuery) Offset(offset int) *GoodIncomingQuery {
	giq.offset = &offset
	return giq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (giq *GoodIncomingQuery) Unique(unique bool) *GoodIncomingQuery {
	giq.unique = &unique
	return giq
}

// Order adds an order step to the query.
func (giq *GoodIncomingQuery) Order(o ...OrderFunc) *GoodIncomingQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// First returns the first GoodIncoming entity from the query.
// Returns a *NotFoundError when no GoodIncoming was found.
func (giq *GoodIncomingQuery) First(ctx context.Context) (*GoodIncoming, error) {
	nodes, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodincoming.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GoodIncomingQuery) FirstX(ctx context.Context) *GoodIncoming {
	node, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodIncoming ID from the query.
// Returns a *NotFoundError when no GoodIncoming ID was found.
func (giq *GoodIncomingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodincoming.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (giq *GoodIncomingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodIncoming entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodIncoming entity is not found.
// Returns a *NotFoundError when no GoodIncoming entities are found.
func (giq *GoodIncomingQuery) Only(ctx context.Context) (*GoodIncoming, error) {
	nodes, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodincoming.Label}
	default:
		return nil, &NotSingularError{goodincoming.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GoodIncomingQuery) OnlyX(ctx context.Context) *GoodIncoming {
	node, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodIncoming ID in the query.
// Returns a *NotSingularError when exactly one GoodIncoming ID is not found.
// Returns a *NotFoundError when no entities are found.
func (giq *GoodIncomingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = &NotSingularError{goodincoming.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (giq *GoodIncomingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodIncomings.
func (giq *GoodIncomingQuery) All(ctx context.Context) ([]*GoodIncoming, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return giq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (giq *GoodIncomingQuery) AllX(ctx context.Context) []*GoodIncoming {
	nodes, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodIncoming IDs.
func (giq *GoodIncomingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := giq.Select(goodincoming.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GoodIncomingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GoodIncomingQuery) Count(ctx context.Context) (int, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return giq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GoodIncomingQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GoodIncomingQuery) Exist(ctx context.Context) (bool, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return giq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GoodIncomingQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodIncomingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GoodIncomingQuery) Clone() *GoodIncomingQuery {
	if giq == nil {
		return nil
	}
	return &GoodIncomingQuery{
		config:     giq.config,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]OrderFunc{}, giq.order...),
		predicates: append([]predicate.GoodIncoming{}, giq.predicates...),
		// clone intermediate query.
		sql:  giq.sql.Clone(),
		path: giq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodIncoming.Query().
//		GroupBy(goodincoming.FieldGoodID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GoodIncomingQuery) GroupBy(field string, fields ...string) *GoodIncomingGroupBy {
	group := &GoodIncomingGroupBy{config: giq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := giq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return giq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//	}
//
//	client.GoodIncoming.Query().
//		Select(goodincoming.FieldGoodID).
//		Scan(ctx, &v)
//
func (giq *GoodIncomingQuery) Select(fields ...string) *GoodIncomingSelect {
	giq.fields = append(giq.fields, fields...)
	return &GoodIncomingSelect{GoodIncomingQuery: giq}
}

func (giq *GoodIncomingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range giq.fields {
		if !goodincoming.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if giq.path != nil {
		prev, err := giq.path(ctx)
		if err != nil {
			return err
		}
		giq.sql = prev
	}
	return nil
}

func (giq *GoodIncomingQuery) sqlAll(ctx context.Context) ([]*GoodIncoming, error) {
	var (
		nodes = []*GoodIncoming{}
		_spec = giq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodIncoming{config: giq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, giq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (giq *GoodIncomingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := giq.querySpec()
	_spec.Node.Columns = giq.fields
	if len(giq.fields) > 0 {
		_spec.Unique = giq.unique != nil && *giq.unique
	}
	return sqlgraph.CountNodes(ctx, giq.driver, _spec)
}

func (giq *GoodIncomingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := giq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (giq *GoodIncomingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodincoming.Table,
			Columns: goodincoming.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodincoming.FieldID,
			},
		},
		From:   giq.sql,
		Unique: true,
	}
	if unique := giq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := giq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodincoming.FieldID)
		for i := range fields {
			if fields[i] != goodincoming.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := giq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := giq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := giq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := giq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (giq *GoodIncomingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(giq.driver.Dialect())
	t1 := builder.Table(goodincoming.Table)
	columns := giq.fields
	if len(columns) == 0 {
		columns = goodincoming.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if giq.sql != nil {
		selector = giq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if giq.unique != nil && *giq.unique {
		selector.Distinct()
	}
	for _, p := range giq.predicates {
		p(selector)
	}
	for _, p := range giq.order {
		p(selector)
	}
	if offset := giq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := giq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodIncomingGroupBy is the group-by builder for GoodIncoming entities.
type GoodIncomingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GoodIncomingGroupBy) Aggregate(fns ...AggregateFunc) *GoodIncomingGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scans the result into the given value.
func (gigb *GoodIncomingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gigb.path(ctx)
	if err != nil {
		return err
	}
	gigb.sql = query
	return gigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) StringX(ctx context.Context) string {
	v, err := gigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) IntX(ctx context.Context) int {
	v, err := gigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodIncomingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gigb *GoodIncomingGroupBy) BoolX(ctx context.Context) bool {
	v, err := gigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GoodIncomingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gigb.fields {
		if !goodincoming.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gigb *GoodIncomingGroupBy) sqlQuery() *sql.Selector {
	selector := gigb.sql.Select()
	aggregation := make([]string, 0, len(gigb.fns))
	for _, fn := range gigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gigb.fields)+len(gigb.fns))
		for _, f := range gigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gigb.fields...)...)
}

// GoodIncomingSelect is the builder for selecting fields of GoodIncoming entities.
type GoodIncomingSelect struct {
	*GoodIncomingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gis *GoodIncomingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gis.prepareQuery(ctx); err != nil {
		return err
	}
	gis.sql = gis.GoodIncomingQuery.sqlQuery(ctx)
	return gis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gis *GoodIncomingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gis *GoodIncomingSelect) StringsX(ctx context.Context) []string {
	v, err := gis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gis *GoodIncomingSelect) StringX(ctx context.Context) string {
	v, err := gis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gis *GoodIncomingSelect) IntsX(ctx context.Context) []int {
	v, err := gis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gis *GoodIncomingSelect) IntX(ctx context.Context) int {
	v, err := gis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gis *GoodIncomingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gis *GoodIncomingSelect) Float64X(ctx context.Context) float64 {
	v, err := gis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodIncomingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gis *GoodIncomingSelect) BoolsX(ctx context.Context) []bool {
	v, err := gis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gis *GoodIncomingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodincoming.Label}
	default:
		err = fmt.Errorf("ent: GoodIncomingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gis *GoodIncomingSelect) BoolX(ctx context.Context) bool {
	v, err := gis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gis *GoodIncomingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gis.sql.Query()
	if err := gis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
