// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userdirectbenefit"
	"github.com/google/uuid"
)

// UserDirectBenefitQuery is the builder for querying UserDirectBenefit entities.
type UserDirectBenefitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserDirectBenefit
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserDirectBenefitQuery builder.
func (udbq *UserDirectBenefitQuery) Where(ps ...predicate.UserDirectBenefit) *UserDirectBenefitQuery {
	udbq.predicates = append(udbq.predicates, ps...)
	return udbq
}

// Limit adds a limit step to the query.
func (udbq *UserDirectBenefitQuery) Limit(limit int) *UserDirectBenefitQuery {
	udbq.limit = &limit
	return udbq
}

// Offset adds an offset step to the query.
func (udbq *UserDirectBenefitQuery) Offset(offset int) *UserDirectBenefitQuery {
	udbq.offset = &offset
	return udbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (udbq *UserDirectBenefitQuery) Unique(unique bool) *UserDirectBenefitQuery {
	udbq.unique = &unique
	return udbq
}

// Order adds an order step to the query.
func (udbq *UserDirectBenefitQuery) Order(o ...OrderFunc) *UserDirectBenefitQuery {
	udbq.order = append(udbq.order, o...)
	return udbq
}

// First returns the first UserDirectBenefit entity from the query.
// Returns a *NotFoundError when no UserDirectBenefit was found.
func (udbq *UserDirectBenefitQuery) First(ctx context.Context) (*UserDirectBenefit, error) {
	nodes, err := udbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userdirectbenefit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) FirstX(ctx context.Context) *UserDirectBenefit {
	node, err := udbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserDirectBenefit ID from the query.
// Returns a *NotFoundError when no UserDirectBenefit ID was found.
func (udbq *UserDirectBenefitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = udbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userdirectbenefit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := udbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserDirectBenefit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserDirectBenefit entity is found.
// Returns a *NotFoundError when no UserDirectBenefit entities are found.
func (udbq *UserDirectBenefitQuery) Only(ctx context.Context) (*UserDirectBenefit, error) {
	nodes, err := udbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userdirectbenefit.Label}
	default:
		return nil, &NotSingularError{userdirectbenefit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) OnlyX(ctx context.Context) *UserDirectBenefit {
	node, err := udbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserDirectBenefit ID in the query.
// Returns a *NotSingularError when more than one UserDirectBenefit ID is found.
// Returns a *NotFoundError when no entities are found.
func (udbq *UserDirectBenefitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = udbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userdirectbenefit.Label}
	default:
		err = &NotSingularError{userdirectbenefit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := udbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserDirectBenefits.
func (udbq *UserDirectBenefitQuery) All(ctx context.Context) ([]*UserDirectBenefit, error) {
	if err := udbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return udbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) AllX(ctx context.Context) []*UserDirectBenefit {
	nodes, err := udbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserDirectBenefit IDs.
func (udbq *UserDirectBenefitQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := udbq.Select(userdirectbenefit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := udbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (udbq *UserDirectBenefitQuery) Count(ctx context.Context) (int, error) {
	if err := udbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return udbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) CountX(ctx context.Context) int {
	count, err := udbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (udbq *UserDirectBenefitQuery) Exist(ctx context.Context) (bool, error) {
	if err := udbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return udbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (udbq *UserDirectBenefitQuery) ExistX(ctx context.Context) bool {
	exist, err := udbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserDirectBenefitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (udbq *UserDirectBenefitQuery) Clone() *UserDirectBenefitQuery {
	if udbq == nil {
		return nil
	}
	return &UserDirectBenefitQuery{
		config:     udbq.config,
		limit:      udbq.limit,
		offset:     udbq.offset,
		order:      append([]OrderFunc{}, udbq.order...),
		predicates: append([]predicate.UserDirectBenefit{}, udbq.predicates...),
		// clone intermediate query.
		sql:    udbq.sql.Clone(),
		path:   udbq.path,
		unique: udbq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserDirectBenefit.Query().
//		GroupBy(userdirectbenefit.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (udbq *UserDirectBenefitQuery) GroupBy(field string, fields ...string) *UserDirectBenefitGroupBy {
	grbuild := &UserDirectBenefitGroupBy{config: udbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := udbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return udbq.sqlQuery(ctx), nil
	}
	grbuild.label = userdirectbenefit.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.UserDirectBenefit.Query().
//		Select(userdirectbenefit.FieldAppID).
//		Scan(ctx, &v)
//
func (udbq *UserDirectBenefitQuery) Select(fields ...string) *UserDirectBenefitSelect {
	udbq.fields = append(udbq.fields, fields...)
	selbuild := &UserDirectBenefitSelect{UserDirectBenefitQuery: udbq}
	selbuild.label = userdirectbenefit.Label
	selbuild.flds, selbuild.scan = &udbq.fields, selbuild.Scan
	return selbuild
}

func (udbq *UserDirectBenefitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range udbq.fields {
		if !userdirectbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if udbq.path != nil {
		prev, err := udbq.path(ctx)
		if err != nil {
			return err
		}
		udbq.sql = prev
	}
	return nil
}

func (udbq *UserDirectBenefitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserDirectBenefit, error) {
	var (
		nodes = []*UserDirectBenefit{}
		_spec = udbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserDirectBenefit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserDirectBenefit{config: udbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, udbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (udbq *UserDirectBenefitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := udbq.querySpec()
	_spec.Node.Columns = udbq.fields
	if len(udbq.fields) > 0 {
		_spec.Unique = udbq.unique != nil && *udbq.unique
	}
	return sqlgraph.CountNodes(ctx, udbq.driver, _spec)
}

func (udbq *UserDirectBenefitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := udbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (udbq *UserDirectBenefitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userdirectbenefit.Table,
			Columns: userdirectbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userdirectbenefit.FieldID,
			},
		},
		From:   udbq.sql,
		Unique: true,
	}
	if unique := udbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := udbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdirectbenefit.FieldID)
		for i := range fields {
			if fields[i] != userdirectbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := udbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := udbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := udbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := udbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (udbq *UserDirectBenefitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(udbq.driver.Dialect())
	t1 := builder.Table(userdirectbenefit.Table)
	columns := udbq.fields
	if len(columns) == 0 {
		columns = userdirectbenefit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if udbq.sql != nil {
		selector = udbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if udbq.unique != nil && *udbq.unique {
		selector.Distinct()
	}
	for _, p := range udbq.predicates {
		p(selector)
	}
	for _, p := range udbq.order {
		p(selector)
	}
	if offset := udbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := udbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserDirectBenefitGroupBy is the group-by builder for UserDirectBenefit entities.
type UserDirectBenefitGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (udbgb *UserDirectBenefitGroupBy) Aggregate(fns ...AggregateFunc) *UserDirectBenefitGroupBy {
	udbgb.fns = append(udbgb.fns, fns...)
	return udbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (udbgb *UserDirectBenefitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := udbgb.path(ctx)
	if err != nil {
		return err
	}
	udbgb.sql = query
	return udbgb.sqlScan(ctx, v)
}

func (udbgb *UserDirectBenefitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range udbgb.fields {
		if !userdirectbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := udbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := udbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (udbgb *UserDirectBenefitGroupBy) sqlQuery() *sql.Selector {
	selector := udbgb.sql.Select()
	aggregation := make([]string, 0, len(udbgb.fns))
	for _, fn := range udbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(udbgb.fields)+len(udbgb.fns))
		for _, f := range udbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(udbgb.fields...)...)
}

// UserDirectBenefitSelect is the builder for selecting fields of UserDirectBenefit entities.
type UserDirectBenefitSelect struct {
	*UserDirectBenefitQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (udbs *UserDirectBenefitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := udbs.prepareQuery(ctx); err != nil {
		return err
	}
	udbs.sql = udbs.UserDirectBenefitQuery.sqlQuery(ctx)
	return udbs.sqlScan(ctx, v)
}

func (udbs *UserDirectBenefitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := udbs.sql.Query()
	if err := udbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
