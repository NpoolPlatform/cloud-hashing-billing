// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodBenefitQuery is the builder for querying GoodBenefit entities.
type GoodBenefitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodBenefit
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodBenefitQuery builder.
func (gbq *GoodBenefitQuery) Where(ps ...predicate.GoodBenefit) *GoodBenefitQuery {
	gbq.predicates = append(gbq.predicates, ps...)
	return gbq
}

// Limit adds a limit step to the query.
func (gbq *GoodBenefitQuery) Limit(limit int) *GoodBenefitQuery {
	gbq.limit = &limit
	return gbq
}

// Offset adds an offset step to the query.
func (gbq *GoodBenefitQuery) Offset(offset int) *GoodBenefitQuery {
	gbq.offset = &offset
	return gbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gbq *GoodBenefitQuery) Unique(unique bool) *GoodBenefitQuery {
	gbq.unique = &unique
	return gbq
}

// Order adds an order step to the query.
func (gbq *GoodBenefitQuery) Order(o ...OrderFunc) *GoodBenefitQuery {
	gbq.order = append(gbq.order, o...)
	return gbq
}

// First returns the first GoodBenefit entity from the query.
// Returns a *NotFoundError when no GoodBenefit was found.
func (gbq *GoodBenefitQuery) First(ctx context.Context) (*GoodBenefit, error) {
	nodes, err := gbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodbenefit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gbq *GoodBenefitQuery) FirstX(ctx context.Context) *GoodBenefit {
	node, err := gbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodBenefit ID from the query.
// Returns a *NotFoundError when no GoodBenefit ID was found.
func (gbq *GoodBenefitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodbenefit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gbq *GoodBenefitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodBenefit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodBenefit entity is found.
// Returns a *NotFoundError when no GoodBenefit entities are found.
func (gbq *GoodBenefitQuery) Only(ctx context.Context) (*GoodBenefit, error) {
	nodes, err := gbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodbenefit.Label}
	default:
		return nil, &NotSingularError{goodbenefit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gbq *GoodBenefitQuery) OnlyX(ctx context.Context) *GoodBenefit {
	node, err := gbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodBenefit ID in the query.
// Returns a *NotSingularError when more than one GoodBenefit ID is found.
// Returns a *NotFoundError when no entities are found.
func (gbq *GoodBenefitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodbenefit.Label}
	default:
		err = &NotSingularError{goodbenefit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gbq *GoodBenefitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodBenefits.
func (gbq *GoodBenefitQuery) All(ctx context.Context) ([]*GoodBenefit, error) {
	if err := gbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gbq *GoodBenefitQuery) AllX(ctx context.Context) []*GoodBenefit {
	nodes, err := gbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodBenefit IDs.
func (gbq *GoodBenefitQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gbq.Select(goodbenefit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gbq *GoodBenefitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gbq *GoodBenefitQuery) Count(ctx context.Context) (int, error) {
	if err := gbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gbq *GoodBenefitQuery) CountX(ctx context.Context) int {
	count, err := gbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gbq *GoodBenefitQuery) Exist(ctx context.Context) (bool, error) {
	if err := gbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gbq *GoodBenefitQuery) ExistX(ctx context.Context) bool {
	exist, err := gbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodBenefitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gbq *GoodBenefitQuery) Clone() *GoodBenefitQuery {
	if gbq == nil {
		return nil
	}
	return &GoodBenefitQuery{
		config:     gbq.config,
		limit:      gbq.limit,
		offset:     gbq.offset,
		order:      append([]OrderFunc{}, gbq.order...),
		predicates: append([]predicate.GoodBenefit{}, gbq.predicates...),
		// clone intermediate query.
		sql:    gbq.sql.Clone(),
		path:   gbq.path,
		unique: gbq.unique,
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
//	client.GoodBenefit.Query().
//		GroupBy(goodbenefit.FieldGoodID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gbq *GoodBenefitQuery) GroupBy(field string, fields ...string) *GoodBenefitGroupBy {
	grbuild := &GoodBenefitGroupBy{config: gbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gbq.sqlQuery(ctx), nil
	}
	grbuild.label = goodbenefit.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
//	client.GoodBenefit.Query().
//		Select(goodbenefit.FieldGoodID).
//		Scan(ctx, &v)
//
func (gbq *GoodBenefitQuery) Select(fields ...string) *GoodBenefitSelect {
	gbq.fields = append(gbq.fields, fields...)
	selbuild := &GoodBenefitSelect{GoodBenefitQuery: gbq}
	selbuild.label = goodbenefit.Label
	selbuild.flds, selbuild.scan = &gbq.fields, selbuild.Scan
	return selbuild
}

func (gbq *GoodBenefitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gbq.fields {
		if !goodbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gbq.path != nil {
		prev, err := gbq.path(ctx)
		if err != nil {
			return err
		}
		gbq.sql = prev
	}
	return nil
}

func (gbq *GoodBenefitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodBenefit, error) {
	var (
		nodes = []*GoodBenefit{}
		_spec = gbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodBenefit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodBenefit{config: gbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gbq *GoodBenefitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gbq.querySpec()
	_spec.Node.Columns = gbq.fields
	if len(gbq.fields) > 0 {
		_spec.Unique = gbq.unique != nil && *gbq.unique
	}
	return sqlgraph.CountNodes(ctx, gbq.driver, _spec)
}

func (gbq *GoodBenefitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gbq *GoodBenefitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodbenefit.Table,
			Columns: goodbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodbenefit.FieldID,
			},
		},
		From:   gbq.sql,
		Unique: true,
	}
	if unique := gbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodbenefit.FieldID)
		for i := range fields {
			if fields[i] != goodbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gbq *GoodBenefitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gbq.driver.Dialect())
	t1 := builder.Table(goodbenefit.Table)
	columns := gbq.fields
	if len(columns) == 0 {
		columns = goodbenefit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gbq.sql != nil {
		selector = gbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gbq.unique != nil && *gbq.unique {
		selector.Distinct()
	}
	for _, p := range gbq.predicates {
		p(selector)
	}
	for _, p := range gbq.order {
		p(selector)
	}
	if offset := gbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodBenefitGroupBy is the group-by builder for GoodBenefit entities.
type GoodBenefitGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gbgb *GoodBenefitGroupBy) Aggregate(fns ...AggregateFunc) *GoodBenefitGroupBy {
	gbgb.fns = append(gbgb.fns, fns...)
	return gbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gbgb *GoodBenefitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gbgb.path(ctx)
	if err != nil {
		return err
	}
	gbgb.sql = query
	return gbgb.sqlScan(ctx, v)
}

func (gbgb *GoodBenefitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gbgb.fields {
		if !goodbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gbgb *GoodBenefitGroupBy) sqlQuery() *sql.Selector {
	selector := gbgb.sql.Select()
	aggregation := make([]string, 0, len(gbgb.fns))
	for _, fn := range gbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gbgb.fields)+len(gbgb.fns))
		for _, f := range gbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gbgb.fields...)...)
}

// GoodBenefitSelect is the builder for selecting fields of GoodBenefit entities.
type GoodBenefitSelect struct {
	*GoodBenefitQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gbs *GoodBenefitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gbs.prepareQuery(ctx); err != nil {
		return err
	}
	gbs.sql = gbs.GoodBenefitQuery.sqlQuery(ctx)
	return gbs.sqlScan(ctx, v)
}

func (gbs *GoodBenefitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gbs.sql.Query()
	if err := gbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
