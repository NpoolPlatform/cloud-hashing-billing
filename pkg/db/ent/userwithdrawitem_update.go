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
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdrawitem"
	"github.com/google/uuid"
)

// UserWithdrawItemUpdate is the builder for updating UserWithdrawItem entities.
type UserWithdrawItemUpdate struct {
	config
	hooks    []Hook
	mutation *UserWithdrawItemMutation
}

// Where appends a list predicates to the UserWithdrawItemUpdate builder.
func (uwiu *UserWithdrawItemUpdate) Where(ps ...predicate.UserWithdrawItem) *UserWithdrawItemUpdate {
	uwiu.mutation.Where(ps...)
	return uwiu
}

// SetAppID sets the "app_id" field.
func (uwiu *UserWithdrawItemUpdate) SetAppID(u uuid.UUID) *UserWithdrawItemUpdate {
	uwiu.mutation.SetAppID(u)
	return uwiu
}

// SetUserID sets the "user_id" field.
func (uwiu *UserWithdrawItemUpdate) SetUserID(u uuid.UUID) *UserWithdrawItemUpdate {
	uwiu.mutation.SetUserID(u)
	return uwiu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uwiu *UserWithdrawItemUpdate) SetCoinTypeID(u uuid.UUID) *UserWithdrawItemUpdate {
	uwiu.mutation.SetCoinTypeID(u)
	return uwiu
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (uwiu *UserWithdrawItemUpdate) SetWithdrawToAccountID(u uuid.UUID) *UserWithdrawItemUpdate {
	uwiu.mutation.SetWithdrawToAccountID(u)
	return uwiu
}

// SetAmount sets the "amount" field.
func (uwiu *UserWithdrawItemUpdate) SetAmount(u uint64) *UserWithdrawItemUpdate {
	uwiu.mutation.ResetAmount()
	uwiu.mutation.SetAmount(u)
	return uwiu
}

// AddAmount adds u to the "amount" field.
func (uwiu *UserWithdrawItemUpdate) AddAmount(u int64) *UserWithdrawItemUpdate {
	uwiu.mutation.AddAmount(u)
	return uwiu
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (uwiu *UserWithdrawItemUpdate) SetPlatformTransactionID(u uuid.UUID) *UserWithdrawItemUpdate {
	uwiu.mutation.SetPlatformTransactionID(u)
	return uwiu
}

// SetWithdrawType sets the "withdraw_type" field.
func (uwiu *UserWithdrawItemUpdate) SetWithdrawType(s string) *UserWithdrawItemUpdate {
	uwiu.mutation.SetWithdrawType(s)
	return uwiu
}

// SetCreateAt sets the "create_at" field.
func (uwiu *UserWithdrawItemUpdate) SetCreateAt(u uint32) *UserWithdrawItemUpdate {
	uwiu.mutation.ResetCreateAt()
	uwiu.mutation.SetCreateAt(u)
	return uwiu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uwiu *UserWithdrawItemUpdate) SetNillableCreateAt(u *uint32) *UserWithdrawItemUpdate {
	if u != nil {
		uwiu.SetCreateAt(*u)
	}
	return uwiu
}

// AddCreateAt adds u to the "create_at" field.
func (uwiu *UserWithdrawItemUpdate) AddCreateAt(u int32) *UserWithdrawItemUpdate {
	uwiu.mutation.AddCreateAt(u)
	return uwiu
}

// SetUpdateAt sets the "update_at" field.
func (uwiu *UserWithdrawItemUpdate) SetUpdateAt(u uint32) *UserWithdrawItemUpdate {
	uwiu.mutation.ResetUpdateAt()
	uwiu.mutation.SetUpdateAt(u)
	return uwiu
}

// AddUpdateAt adds u to the "update_at" field.
func (uwiu *UserWithdrawItemUpdate) AddUpdateAt(u int32) *UserWithdrawItemUpdate {
	uwiu.mutation.AddUpdateAt(u)
	return uwiu
}

// SetDeleteAt sets the "delete_at" field.
func (uwiu *UserWithdrawItemUpdate) SetDeleteAt(u uint32) *UserWithdrawItemUpdate {
	uwiu.mutation.ResetDeleteAt()
	uwiu.mutation.SetDeleteAt(u)
	return uwiu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uwiu *UserWithdrawItemUpdate) SetNillableDeleteAt(u *uint32) *UserWithdrawItemUpdate {
	if u != nil {
		uwiu.SetDeleteAt(*u)
	}
	return uwiu
}

// AddDeleteAt adds u to the "delete_at" field.
func (uwiu *UserWithdrawItemUpdate) AddDeleteAt(u int32) *UserWithdrawItemUpdate {
	uwiu.mutation.AddDeleteAt(u)
	return uwiu
}

// Mutation returns the UserWithdrawItemMutation object of the builder.
func (uwiu *UserWithdrawItemUpdate) Mutation() *UserWithdrawItemMutation {
	return uwiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uwiu *UserWithdrawItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uwiu.defaults()
	if len(uwiu.hooks) == 0 {
		affected, err = uwiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwiu.mutation = mutation
			affected, err = uwiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwiu.hooks) - 1; i >= 0; i-- {
			if uwiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwiu *UserWithdrawItemUpdate) SaveX(ctx context.Context) int {
	affected, err := uwiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uwiu *UserWithdrawItemUpdate) Exec(ctx context.Context) error {
	_, err := uwiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwiu *UserWithdrawItemUpdate) ExecX(ctx context.Context) {
	if err := uwiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwiu *UserWithdrawItemUpdate) defaults() {
	if _, ok := uwiu.mutation.UpdateAt(); !ok {
		v := userwithdrawitem.UpdateDefaultUpdateAt()
		uwiu.mutation.SetUpdateAt(v)
	}
}

func (uwiu *UserWithdrawItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithdrawitem.Table,
			Columns: userwithdrawitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdrawitem.FieldID,
			},
		},
	}
	if ps := uwiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwiu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldAppID,
		})
	}
	if value, ok := uwiu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldUserID,
		})
	}
	if value, ok := uwiu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldCoinTypeID,
		})
	}
	if value, ok := uwiu.mutation.WithdrawToAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldWithdrawToAccountID,
		})
	}
	if value, ok := uwiu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userwithdrawitem.FieldAmount,
		})
	}
	if value, ok := uwiu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userwithdrawitem.FieldAmount,
		})
	}
	if value, ok := uwiu.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldPlatformTransactionID,
		})
	}
	if value, ok := uwiu.mutation.WithdrawType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdrawitem.FieldWithdrawType,
		})
	}
	if value, ok := uwiu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldCreateAt,
		})
	}
	if value, ok := uwiu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldCreateAt,
		})
	}
	if value, ok := uwiu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldUpdateAt,
		})
	}
	if value, ok := uwiu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldUpdateAt,
		})
	}
	if value, ok := uwiu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldDeleteAt,
		})
	}
	if value, ok := uwiu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uwiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithdrawitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserWithdrawItemUpdateOne is the builder for updating a single UserWithdrawItem entity.
type UserWithdrawItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserWithdrawItemMutation
}

// SetAppID sets the "app_id" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetAppID(u uuid.UUID) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetAppID(u)
	return uwiuo
}

// SetUserID sets the "user_id" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetUserID(u uuid.UUID) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetUserID(u)
	return uwiuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetCoinTypeID(u uuid.UUID) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetCoinTypeID(u)
	return uwiuo
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetWithdrawToAccountID(u uuid.UUID) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetWithdrawToAccountID(u)
	return uwiuo
}

// SetAmount sets the "amount" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetAmount(u uint64) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.ResetAmount()
	uwiuo.mutation.SetAmount(u)
	return uwiuo
}

// AddAmount adds u to the "amount" field.
func (uwiuo *UserWithdrawItemUpdateOne) AddAmount(u int64) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.AddAmount(u)
	return uwiuo
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetPlatformTransactionID(u uuid.UUID) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetPlatformTransactionID(u)
	return uwiuo
}

// SetWithdrawType sets the "withdraw_type" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetWithdrawType(s string) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.SetWithdrawType(s)
	return uwiuo
}

// SetCreateAt sets the "create_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetCreateAt(u uint32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.ResetCreateAt()
	uwiuo.mutation.SetCreateAt(u)
	return uwiuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uwiuo *UserWithdrawItemUpdateOne) SetNillableCreateAt(u *uint32) *UserWithdrawItemUpdateOne {
	if u != nil {
		uwiuo.SetCreateAt(*u)
	}
	return uwiuo
}

// AddCreateAt adds u to the "create_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) AddCreateAt(u int32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.AddCreateAt(u)
	return uwiuo
}

// SetUpdateAt sets the "update_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetUpdateAt(u uint32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.ResetUpdateAt()
	uwiuo.mutation.SetUpdateAt(u)
	return uwiuo
}

// AddUpdateAt adds u to the "update_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) AddUpdateAt(u int32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.AddUpdateAt(u)
	return uwiuo
}

// SetDeleteAt sets the "delete_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) SetDeleteAt(u uint32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.ResetDeleteAt()
	uwiuo.mutation.SetDeleteAt(u)
	return uwiuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uwiuo *UserWithdrawItemUpdateOne) SetNillableDeleteAt(u *uint32) *UserWithdrawItemUpdateOne {
	if u != nil {
		uwiuo.SetDeleteAt(*u)
	}
	return uwiuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (uwiuo *UserWithdrawItemUpdateOne) AddDeleteAt(u int32) *UserWithdrawItemUpdateOne {
	uwiuo.mutation.AddDeleteAt(u)
	return uwiuo
}

// Mutation returns the UserWithdrawItemMutation object of the builder.
func (uwiuo *UserWithdrawItemUpdateOne) Mutation() *UserWithdrawItemMutation {
	return uwiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uwiuo *UserWithdrawItemUpdateOne) Select(field string, fields ...string) *UserWithdrawItemUpdateOne {
	uwiuo.fields = append([]string{field}, fields...)
	return uwiuo
}

// Save executes the query and returns the updated UserWithdrawItem entity.
func (uwiuo *UserWithdrawItemUpdateOne) Save(ctx context.Context) (*UserWithdrawItem, error) {
	var (
		err  error
		node *UserWithdrawItem
	)
	uwiuo.defaults()
	if len(uwiuo.hooks) == 0 {
		node, err = uwiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwiuo.mutation = mutation
			node, err = uwiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uwiuo.hooks) - 1; i >= 0; i-- {
			if uwiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwiuo *UserWithdrawItemUpdateOne) SaveX(ctx context.Context) *UserWithdrawItem {
	node, err := uwiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uwiuo *UserWithdrawItemUpdateOne) Exec(ctx context.Context) error {
	_, err := uwiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwiuo *UserWithdrawItemUpdateOne) ExecX(ctx context.Context) {
	if err := uwiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwiuo *UserWithdrawItemUpdateOne) defaults() {
	if _, ok := uwiuo.mutation.UpdateAt(); !ok {
		v := userwithdrawitem.UpdateDefaultUpdateAt()
		uwiuo.mutation.SetUpdateAt(v)
	}
}

func (uwiuo *UserWithdrawItemUpdateOne) sqlSave(ctx context.Context) (_node *UserWithdrawItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithdrawitem.Table,
			Columns: userwithdrawitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdrawitem.FieldID,
			},
		},
	}
	id, ok := uwiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserWithdrawItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uwiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwithdrawitem.FieldID)
		for _, f := range fields {
			if !userwithdrawitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userwithdrawitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uwiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwiuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldAppID,
		})
	}
	if value, ok := uwiuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldUserID,
		})
	}
	if value, ok := uwiuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldCoinTypeID,
		})
	}
	if value, ok := uwiuo.mutation.WithdrawToAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldWithdrawToAccountID,
		})
	}
	if value, ok := uwiuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userwithdrawitem.FieldAmount,
		})
	}
	if value, ok := uwiuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userwithdrawitem.FieldAmount,
		})
	}
	if value, ok := uwiuo.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldPlatformTransactionID,
		})
	}
	if value, ok := uwiuo.mutation.WithdrawType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwithdrawitem.FieldWithdrawType,
		})
	}
	if value, ok := uwiuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldCreateAt,
		})
	}
	if value, ok := uwiuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldCreateAt,
		})
	}
	if value, ok := uwiuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldUpdateAt,
		})
	}
	if value, ok := uwiuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldUpdateAt,
		})
	}
	if value, ok := uwiuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldDeleteAt,
		})
	}
	if value, ok := uwiuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldDeleteAt,
		})
	}
	_node = &UserWithdrawItem{config: uwiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uwiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithdrawitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
