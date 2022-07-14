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
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
	"github.com/google/uuid"
)

// UserWithdrawQuery is the builder for querying UserWithdraw entities.
type UserWithdrawQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserWithdraw
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserWithdrawQuery builder.
func (uwq *UserWithdrawQuery) Where(ps ...predicate.UserWithdraw) *UserWithdrawQuery {
	uwq.predicates = append(uwq.predicates, ps...)
	return uwq
}

// Limit adds a limit step to the query.
func (uwq *UserWithdrawQuery) Limit(limit int) *UserWithdrawQuery {
	uwq.limit = &limit
	return uwq
}

// Offset adds an offset step to the query.
func (uwq *UserWithdrawQuery) Offset(offset int) *UserWithdrawQuery {
	uwq.offset = &offset
	return uwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uwq *UserWithdrawQuery) Unique(unique bool) *UserWithdrawQuery {
	uwq.unique = &unique
	return uwq
}

// Order adds an order step to the query.
func (uwq *UserWithdrawQuery) Order(o ...OrderFunc) *UserWithdrawQuery {
	uwq.order = append(uwq.order, o...)
	return uwq
}

// First returns the first UserWithdraw entity from the query.
// Returns a *NotFoundError when no UserWithdraw was found.
func (uwq *UserWithdrawQuery) First(ctx context.Context) (*UserWithdraw, error) {
	nodes, err := uwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userwithdraw.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uwq *UserWithdrawQuery) FirstX(ctx context.Context) *UserWithdraw {
	node, err := uwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserWithdraw ID from the query.
// Returns a *NotFoundError when no UserWithdraw ID was found.
func (uwq *UserWithdrawQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userwithdraw.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uwq *UserWithdrawQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserWithdraw entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserWithdraw entity is found.
// Returns a *NotFoundError when no UserWithdraw entities are found.
func (uwq *UserWithdrawQuery) Only(ctx context.Context) (*UserWithdraw, error) {
	nodes, err := uwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userwithdraw.Label}
	default:
		return nil, &NotSingularError{userwithdraw.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uwq *UserWithdrawQuery) OnlyX(ctx context.Context) *UserWithdraw {
	node, err := uwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserWithdraw ID in the query.
// Returns a *NotSingularError when more than one UserWithdraw ID is found.
// Returns a *NotFoundError when no entities are found.
func (uwq *UserWithdrawQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userwithdraw.Label}
	default:
		err = &NotSingularError{userwithdraw.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uwq *UserWithdrawQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserWithdraws.
func (uwq *UserWithdrawQuery) All(ctx context.Context) ([]*UserWithdraw, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uwq *UserWithdrawQuery) AllX(ctx context.Context) []*UserWithdraw {
	nodes, err := uwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserWithdraw IDs.
func (uwq *UserWithdrawQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := uwq.Select(userwithdraw.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uwq *UserWithdrawQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uwq *UserWithdrawQuery) Count(ctx context.Context) (int, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uwq *UserWithdrawQuery) CountX(ctx context.Context) int {
	count, err := uwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uwq *UserWithdrawQuery) Exist(ctx context.Context) (bool, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uwq *UserWithdrawQuery) ExistX(ctx context.Context) bool {
	exist, err := uwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserWithdrawQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uwq *UserWithdrawQuery) Clone() *UserWithdrawQuery {
	if uwq == nil {
		return nil
	}
	return &UserWithdrawQuery{
		config:     uwq.config,
		limit:      uwq.limit,
		offset:     uwq.offset,
		order:      append([]OrderFunc{}, uwq.order...),
		predicates: append([]predicate.UserWithdraw{}, uwq.predicates...),
		// clone intermediate query.
		sql:    uwq.sql.Clone(),
		path:   uwq.path,
		unique: uwq.unique,
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
//	client.UserWithdraw.Query().
//		GroupBy(userwithdraw.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uwq *UserWithdrawQuery) GroupBy(field string, fields ...string) *UserWithdrawGroupBy {
	grbuild := &UserWithdrawGroupBy{config: uwq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uwq.sqlQuery(ctx), nil
	}
	grbuild.label = userwithdraw.Label
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
//	client.UserWithdraw.Query().
//		Select(userwithdraw.FieldAppID).
//		Scan(ctx, &v)
//
func (uwq *UserWithdrawQuery) Select(fields ...string) *UserWithdrawSelect {
	uwq.fields = append(uwq.fields, fields...)
	selbuild := &UserWithdrawSelect{UserWithdrawQuery: uwq}
	selbuild.label = userwithdraw.Label
	selbuild.flds, selbuild.scan = &uwq.fields, selbuild.Scan
	return selbuild
}

func (uwq *UserWithdrawQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uwq.fields {
		if !userwithdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uwq.path != nil {
		prev, err := uwq.path(ctx)
		if err != nil {
			return err
		}
		uwq.sql = prev
	}
	return nil
}

func (uwq *UserWithdrawQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserWithdraw, error) {
	var (
		nodes = []*UserWithdraw{}
		_spec = uwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserWithdraw).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserWithdraw{config: uwq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uwq *UserWithdrawQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uwq.querySpec()
	_spec.Node.Columns = uwq.fields
	if len(uwq.fields) > 0 {
		_spec.Unique = uwq.unique != nil && *uwq.unique
	}
	return sqlgraph.CountNodes(ctx, uwq.driver, _spec)
}

func (uwq *UserWithdrawQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uwq *UserWithdrawQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithdraw.Table,
			Columns: userwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdraw.FieldID,
			},
		},
		From:   uwq.sql,
		Unique: true,
	}
	if unique := uwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwithdraw.FieldID)
		for i := range fields {
			if fields[i] != userwithdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uwq *UserWithdrawQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uwq.driver.Dialect())
	t1 := builder.Table(userwithdraw.Table)
	columns := uwq.fields
	if len(columns) == 0 {
		columns = userwithdraw.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uwq.sql != nil {
		selector = uwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uwq.unique != nil && *uwq.unique {
		selector.Distinct()
	}
	for _, p := range uwq.predicates {
		p(selector)
	}
	for _, p := range uwq.order {
		p(selector)
	}
	if offset := uwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserWithdrawGroupBy is the group-by builder for UserWithdraw entities.
type UserWithdrawGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uwgb *UserWithdrawGroupBy) Aggregate(fns ...AggregateFunc) *UserWithdrawGroupBy {
	uwgb.fns = append(uwgb.fns, fns...)
	return uwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uwgb *UserWithdrawGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uwgb.path(ctx)
	if err != nil {
		return err
	}
	uwgb.sql = query
	return uwgb.sqlScan(ctx, v)
}

func (uwgb *UserWithdrawGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uwgb.fields {
		if !userwithdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uwgb *UserWithdrawGroupBy) sqlQuery() *sql.Selector {
	selector := uwgb.sql.Select()
	aggregation := make([]string, 0, len(uwgb.fns))
	for _, fn := range uwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uwgb.fields)+len(uwgb.fns))
		for _, f := range uwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uwgb.fields...)...)
}

// UserWithdrawSelect is the builder for selecting fields of UserWithdraw entities.
type UserWithdrawSelect struct {
	*UserWithdrawQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uws *UserWithdrawSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uws.prepareQuery(ctx); err != nil {
		return err
	}
	uws.sql = uws.UserWithdrawQuery.sqlQuery(ctx)
	return uws.sqlScan(ctx, v)
}

func (uws *UserWithdrawSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uws.sql.Query()
	if err := uws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
