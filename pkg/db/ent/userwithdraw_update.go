// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
	"github.com/google/uuid"
)

// UserWithdrawUpdate is the builder for updating UserWithdraw entities.
type UserWithdrawUpdate struct {
	config
	hooks    []Hook
	mutation *UserWithdrawMutation
}

// Where appends a list predicates to the UserWithdrawUpdate builder.
func (uwu *UserWithdrawUpdate) Where(ps ...predicate.UserWithdraw) *UserWithdrawUpdate {
	uwu.mutation.Where(ps...)
	return uwu
}

// SetAppID sets the "app_id" field.
func (uwu *UserWithdrawUpdate) SetAppID(u uuid.UUID) *UserWithdrawUpdate {
	uwu.mutation.SetAppID(u)
	return uwu
}

// SetUserID sets the "user_id" field.
func (uwu *UserWithdrawUpdate) SetUserID(u uuid.UUID) *UserWithdrawUpdate {
	uwu.mutation.SetUserID(u)
	return uwu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uwu *UserWithdrawUpdate) SetCoinTypeID(u uuid.UUID) *UserWithdrawUpdate {
	uwu.mutation.SetCoinTypeID(u)
	return uwu
}

// SetAccountID sets the "account_id" field.
func (uwu *UserWithdrawUpdate) SetAccountID(u uuid.UUID) *UserWithdrawUpdate {
	uwu.mutation.SetAccountID(u)
	return uwu
}

// SetName sets the "name" field.
func (uwu *UserWithdrawUpdate) SetName(s string) *UserWithdrawUpdate {
	uwu.mutation.SetName(s)
	return uwu
}

// SetMessage sets the "message" field.
func (uwu *UserWithdrawUpdate) SetMessage(s string) *UserWithdrawUpdate {
	uwu.mutation.SetMessage(s)
	return uwu
}

// SetCreateAt sets the "create_at" field.
func (uwu *UserWithdrawUpdate) SetCreateAt(u uint32) *UserWithdrawUpdate {
	uwu.mutation.ResetCreateAt()
	uwu.mutation.SetCreateAt(u)
	return uwu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uwu *UserWithdrawUpdate) SetNillableCreateAt(u *uint32) *UserWithdrawUpdate {
	if u != nil {
		uwu.SetCreateAt(*u)
	}
	return uwu
}

// AddCreateAt adds u to the "create_at" field.
func (uwu *UserWithdrawUpdate) AddCreateAt(u int32) *UserWithdrawUpdate {
	uwu.mutation.AddCreateAt(u)
	return uwu
}

// SetUpdateAt sets the "update_at" field.
func (uwu *UserWithdrawUpdate) SetUpdateAt(u uint32) *UserWithdrawUpdate {
	uwu.mutation.ResetUpdateAt()
	uwu.mutation.SetUpdateAt(u)
	return uwu
}

// AddUpdateAt adds u to the "update_at" field.
func (uwu *UserWithdrawUpdate) AddUpdateAt(u int32) *UserWithdrawUpdate {
	uwu.mutation.AddUpdateAt(u)
	return uwu
}

// SetDeleteAt sets the "delete_at" field.
func (uwu *UserWithdrawUpdate) SetDeleteAt(u uint32) *UserWithdrawUpdate {
	uwu.mutation.ResetDeleteAt()
	uwu.mutation.SetDeleteAt(u)
	return uwu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uwu *UserWithdrawUpdate) SetNillableDeleteAt(u *uint32) *UserWithdrawUpdate {
	if u != nil {
		uwu.SetDeleteAt(*u)
	}
	return uwu
}

// AddDeleteAt adds u to the "delete_at" field.
func (uwu *UserWithdrawUpdate) AddDeleteAt(u int32) *UserWithdrawUpdate {
	uwu.mutation.AddDeleteAt(u)
	return uwu
}

// Mutation returns the UserWithdrawMutation object of the builder.
func (uwu *UserWithdrawUpdate) Mutation() *UserWithdrawMutation {
	return uwu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uwu *UserWithdrawUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uwu.defaults()
	if len(uwu.hooks) == 0 {
		affected, err = uwu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwu.mutation = mutation
			affected, err = uwu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwu.hooks) - 1; i >= 0; i-- {
			if uwu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwu *UserWithdrawUpdate) SaveX(ctx context.Context) int {
	affected, err := uwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uwu *UserWithdrawUpdate) Exec(ctx context.Context) error {
	_, err := uwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwu *UserWithdrawUpdate) ExecX(ctx context.Context) {
	if err := uwu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwu *UserWithdrawUpdate) defaults() {
	if _, ok := uwu.mutation.UpdateAt(); !ok {
		v := userwithdraw.UpdateDefaultUpdateAt()
		uwu.mutation.SetUpdateAt(v)
	}
}

func (uwu *UserWithdrawUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithdraw.Table,
			Columns: userwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdraw.FieldID,
			},
		},
	}
	if ps := uwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldAppID,
		})
	}
	if value, ok := uwu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldUserID,
		})
	}
	if value, ok := uwu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldCoinTypeID,
		})
	}
	if value, ok := uwu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldAccountID,
		})
	}
	if value, ok := uwu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdraw.FieldName,
		})
	}
	if value, ok := uwu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdraw.FieldMessage,
		})
	}
	if value, ok := uwu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldCreateAt,
		})
	}
	if value, ok := uwu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldCreateAt,
		})
	}
	if value, ok := uwu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldUpdateAt,
		})
	}
	if value, ok := uwu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldUpdateAt,
		})
	}
	if value, ok := uwu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldDeleteAt,
		})
	}
	if value, ok := uwu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserWithdrawUpdateOne is the builder for updating a single UserWithdraw entity.
type UserWithdrawUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserWithdrawMutation
}

// SetAppID sets the "app_id" field.
func (uwuo *UserWithdrawUpdateOne) SetAppID(u uuid.UUID) *UserWithdrawUpdateOne {
	uwuo.mutation.SetAppID(u)
	return uwuo
}

// SetUserID sets the "user_id" field.
func (uwuo *UserWithdrawUpdateOne) SetUserID(u uuid.UUID) *UserWithdrawUpdateOne {
	uwuo.mutation.SetUserID(u)
	return uwuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uwuo *UserWithdrawUpdateOne) SetCoinTypeID(u uuid.UUID) *UserWithdrawUpdateOne {
	uwuo.mutation.SetCoinTypeID(u)
	return uwuo
}

// SetAccountID sets the "account_id" field.
func (uwuo *UserWithdrawUpdateOne) SetAccountID(u uuid.UUID) *UserWithdrawUpdateOne {
	uwuo.mutation.SetAccountID(u)
	return uwuo
}

// SetName sets the "name" field.
func (uwuo *UserWithdrawUpdateOne) SetName(s string) *UserWithdrawUpdateOne {
	uwuo.mutation.SetName(s)
	return uwuo
}

// SetMessage sets the "message" field.
func (uwuo *UserWithdrawUpdateOne) SetMessage(s string) *UserWithdrawUpdateOne {
	uwuo.mutation.SetMessage(s)
	return uwuo
}

// SetCreateAt sets the "create_at" field.
func (uwuo *UserWithdrawUpdateOne) SetCreateAt(u uint32) *UserWithdrawUpdateOne {
	uwuo.mutation.ResetCreateAt()
	uwuo.mutation.SetCreateAt(u)
	return uwuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uwuo *UserWithdrawUpdateOne) SetNillableCreateAt(u *uint32) *UserWithdrawUpdateOne {
	if u != nil {
		uwuo.SetCreateAt(*u)
	}
	return uwuo
}

// AddCreateAt adds u to the "create_at" field.
func (uwuo *UserWithdrawUpdateOne) AddCreateAt(u int32) *UserWithdrawUpdateOne {
	uwuo.mutation.AddCreateAt(u)
	return uwuo
}

// SetUpdateAt sets the "update_at" field.
func (uwuo *UserWithdrawUpdateOne) SetUpdateAt(u uint32) *UserWithdrawUpdateOne {
	uwuo.mutation.ResetUpdateAt()
	uwuo.mutation.SetUpdateAt(u)
	return uwuo
}

// AddUpdateAt adds u to the "update_at" field.
func (uwuo *UserWithdrawUpdateOne) AddUpdateAt(u int32) *UserWithdrawUpdateOne {
	uwuo.mutation.AddUpdateAt(u)
	return uwuo
}

// SetDeleteAt sets the "delete_at" field.
func (uwuo *UserWithdrawUpdateOne) SetDeleteAt(u uint32) *UserWithdrawUpdateOne {
	uwuo.mutation.ResetDeleteAt()
	uwuo.mutation.SetDeleteAt(u)
	return uwuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uwuo *UserWithdrawUpdateOne) SetNillableDeleteAt(u *uint32) *UserWithdrawUpdateOne {
	if u != nil {
		uwuo.SetDeleteAt(*u)
	}
	return uwuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (uwuo *UserWithdrawUpdateOne) AddDeleteAt(u int32) *UserWithdrawUpdateOne {
	uwuo.mutation.AddDeleteAt(u)
	return uwuo
}

// Mutation returns the UserWithdrawMutation object of the builder.
func (uwuo *UserWithdrawUpdateOne) Mutation() *UserWithdrawMutation {
	return uwuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uwuo *UserWithdrawUpdateOne) Select(field string, fields ...string) *UserWithdrawUpdateOne {
	uwuo.fields = append([]string{field}, fields...)
	return uwuo
}

// Save executes the query and returns the updated UserWithdraw entity.
func (uwuo *UserWithdrawUpdateOne) Save(ctx context.Context) (*UserWithdraw, error) {
	var (
		err  error
		node *UserWithdraw
	)
	uwuo.defaults()
	if len(uwuo.hooks) == 0 {
		node, err = uwuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwuo.mutation = mutation
			node, err = uwuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uwuo.hooks) - 1; i >= 0; i-- {
			if uwuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwuo *UserWithdrawUpdateOne) SaveX(ctx context.Context) *UserWithdraw {
	node, err := uwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uwuo *UserWithdrawUpdateOne) Exec(ctx context.Context) error {
	_, err := uwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwuo *UserWithdrawUpdateOne) ExecX(ctx context.Context) {
	if err := uwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwuo *UserWithdrawUpdateOne) defaults() {
	if _, ok := uwuo.mutation.UpdateAt(); !ok {
		v := userwithdraw.UpdateDefaultUpdateAt()
		uwuo.mutation.SetUpdateAt(v)
	}
}

func (uwuo *UserWithdrawUpdateOne) sqlSave(ctx context.Context) (_node *UserWithdraw, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithdraw.Table,
			Columns: userwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdraw.FieldID,
			},
		},
	}
	id, ok := uwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserWithdraw.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwithdraw.FieldID)
		for _, f := range fields {
			if !userwithdraw.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userwithdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldAppID,
		})
	}
	if value, ok := uwuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldUserID,
		})
	}
	if value, ok := uwuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldCoinTypeID,
		})
	}
	if value, ok := uwuo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdraw.FieldAccountID,
		})
	}
	if value, ok := uwuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdraw.FieldName,
		})
	}
	if value, ok := uwuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdraw.FieldMessage,
		})
	}
	if value, ok := uwuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldCreateAt,
		})
	}
	if value, ok := uwuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldCreateAt,
		})
	}
	if value, ok := uwuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldUpdateAt,
		})
	}
	if value, ok := uwuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldUpdateAt,
		})
	}
	if value, ok := uwuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldDeleteAt,
		})
	}
	if value, ok := uwuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdraw.FieldDeleteAt,
		})
	}
	_node = &UserWithdraw{config: uwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}