// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinSettingUpdate is the builder for updating CoinSetting entities.
type CoinSettingUpdate struct {
	config
	hooks    []Hook
	mutation *CoinSettingMutation
}

// Where appends a list predicates to the CoinSettingUpdate builder.
func (csu *CoinSettingUpdate) Where(ps ...predicate.CoinSetting) *CoinSettingUpdate {
	csu.mutation.Where(ps...)
	return csu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (csu *CoinSettingUpdate) SetCoinTypeID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetCoinTypeID(u)
	return csu
}

// SetWarmAccountCoinAmount sets the "warm_account_coin_amount" field.
func (csu *CoinSettingUpdate) SetWarmAccountCoinAmount(u uint64) *CoinSettingUpdate {
	csu.mutation.ResetWarmAccountCoinAmount()
	csu.mutation.SetWarmAccountCoinAmount(u)
	return csu
}

// AddWarmAccountCoinAmount adds u to the "warm_account_coin_amount" field.
func (csu *CoinSettingUpdate) AddWarmAccountCoinAmount(u int64) *CoinSettingUpdate {
	csu.mutation.AddWarmAccountCoinAmount(u)
	return csu
}

// SetPaymentAccountCoinAmount sets the "payment_account_coin_amount" field.
func (csu *CoinSettingUpdate) SetPaymentAccountCoinAmount(u uint64) *CoinSettingUpdate {
	csu.mutation.ResetPaymentAccountCoinAmount()
	csu.mutation.SetPaymentAccountCoinAmount(u)
	return csu
}

// AddPaymentAccountCoinAmount adds u to the "payment_account_coin_amount" field.
func (csu *CoinSettingUpdate) AddPaymentAccountCoinAmount(u int64) *CoinSettingUpdate {
	csu.mutation.AddPaymentAccountCoinAmount(u)
	return csu
}

// SetPlatformOfflineAccountID sets the "platform_offline_account_id" field.
func (csu *CoinSettingUpdate) SetPlatformOfflineAccountID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetPlatformOfflineAccountID(u)
	return csu
}

// SetUserOnlineAccountID sets the "user_online_account_id" field.
func (csu *CoinSettingUpdate) SetUserOnlineAccountID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetUserOnlineAccountID(u)
	return csu
}

// SetUserOfflineAccountID sets the "user_offline_account_id" field.
func (csu *CoinSettingUpdate) SetUserOfflineAccountID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetUserOfflineAccountID(u)
	return csu
}

// SetGoodIncomingAccountID sets the "good_incoming_account_id" field.
func (csu *CoinSettingUpdate) SetGoodIncomingAccountID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetGoodIncomingAccountID(u)
	return csu
}

// SetGasProviderAccountID sets the "gas_provider_account_id" field.
func (csu *CoinSettingUpdate) SetGasProviderAccountID(u uuid.UUID) *CoinSettingUpdate {
	csu.mutation.SetGasProviderAccountID(u)
	return csu
}

// SetCreateAt sets the "create_at" field.
func (csu *CoinSettingUpdate) SetCreateAt(u uint32) *CoinSettingUpdate {
	csu.mutation.ResetCreateAt()
	csu.mutation.SetCreateAt(u)
	return csu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (csu *CoinSettingUpdate) SetNillableCreateAt(u *uint32) *CoinSettingUpdate {
	if u != nil {
		csu.SetCreateAt(*u)
	}
	return csu
}

// AddCreateAt adds u to the "create_at" field.
func (csu *CoinSettingUpdate) AddCreateAt(u int32) *CoinSettingUpdate {
	csu.mutation.AddCreateAt(u)
	return csu
}

// SetUpdateAt sets the "update_at" field.
func (csu *CoinSettingUpdate) SetUpdateAt(u uint32) *CoinSettingUpdate {
	csu.mutation.ResetUpdateAt()
	csu.mutation.SetUpdateAt(u)
	return csu
}

// AddUpdateAt adds u to the "update_at" field.
func (csu *CoinSettingUpdate) AddUpdateAt(u int32) *CoinSettingUpdate {
	csu.mutation.AddUpdateAt(u)
	return csu
}

// SetDeleteAt sets the "delete_at" field.
func (csu *CoinSettingUpdate) SetDeleteAt(u uint32) *CoinSettingUpdate {
	csu.mutation.ResetDeleteAt()
	csu.mutation.SetDeleteAt(u)
	return csu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (csu *CoinSettingUpdate) SetNillableDeleteAt(u *uint32) *CoinSettingUpdate {
	if u != nil {
		csu.SetDeleteAt(*u)
	}
	return csu
}

// AddDeleteAt adds u to the "delete_at" field.
func (csu *CoinSettingUpdate) AddDeleteAt(u int32) *CoinSettingUpdate {
	csu.mutation.AddDeleteAt(u)
	return csu
}

// Mutation returns the CoinSettingMutation object of the builder.
func (csu *CoinSettingUpdate) Mutation() *CoinSettingMutation {
	return csu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *CoinSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	csu.defaults()
	if len(csu.hooks) == 0 {
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			if csu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *CoinSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *CoinSettingUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *CoinSettingUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csu *CoinSettingUpdate) defaults() {
	if _, ok := csu.mutation.UpdateAt(); !ok {
		v := coinsetting.UpdateDefaultUpdateAt()
		csu.mutation.SetUpdateAt(v)
	}
}

func (csu *CoinSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinsetting.Table,
			Columns: coinsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsetting.FieldID,
			},
		},
	}
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldCoinTypeID,
		})
	}
	if value, ok := csu.mutation.WarmAccountCoinAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldWarmAccountCoinAmount,
		})
	}
	if value, ok := csu.mutation.AddedWarmAccountCoinAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldWarmAccountCoinAmount,
		})
	}
	if value, ok := csu.mutation.PaymentAccountCoinAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldPaymentAccountCoinAmount,
		})
	}
	if value, ok := csu.mutation.AddedPaymentAccountCoinAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldPaymentAccountCoinAmount,
		})
	}
	if value, ok := csu.mutation.PlatformOfflineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldPlatformOfflineAccountID,
		})
	}
	if value, ok := csu.mutation.UserOnlineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldUserOnlineAccountID,
		})
	}
	if value, ok := csu.mutation.UserOfflineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldUserOfflineAccountID,
		})
	}
	if value, ok := csu.mutation.GoodIncomingAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldGoodIncomingAccountID,
		})
	}
	if value, ok := csu.mutation.GasProviderAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldGasProviderAccountID,
		})
	}
	if value, ok := csu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldCreateAt,
		})
	}
	if value, ok := csu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldCreateAt,
		})
	}
	if value, ok := csu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldUpdateAt,
		})
	}
	if value, ok := csu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldUpdateAt,
		})
	}
	if value, ok := csu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldDeleteAt,
		})
	}
	if value, ok := csu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinSettingUpdateOne is the builder for updating a single CoinSetting entity.
type CoinSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoinSettingMutation
}

// SetCoinTypeID sets the "coin_type_id" field.
func (csuo *CoinSettingUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetCoinTypeID(u)
	return csuo
}

// SetWarmAccountCoinAmount sets the "warm_account_coin_amount" field.
func (csuo *CoinSettingUpdateOne) SetWarmAccountCoinAmount(u uint64) *CoinSettingUpdateOne {
	csuo.mutation.ResetWarmAccountCoinAmount()
	csuo.mutation.SetWarmAccountCoinAmount(u)
	return csuo
}

// AddWarmAccountCoinAmount adds u to the "warm_account_coin_amount" field.
func (csuo *CoinSettingUpdateOne) AddWarmAccountCoinAmount(u int64) *CoinSettingUpdateOne {
	csuo.mutation.AddWarmAccountCoinAmount(u)
	return csuo
}

// SetPaymentAccountCoinAmount sets the "payment_account_coin_amount" field.
func (csuo *CoinSettingUpdateOne) SetPaymentAccountCoinAmount(u uint64) *CoinSettingUpdateOne {
	csuo.mutation.ResetPaymentAccountCoinAmount()
	csuo.mutation.SetPaymentAccountCoinAmount(u)
	return csuo
}

// AddPaymentAccountCoinAmount adds u to the "payment_account_coin_amount" field.
func (csuo *CoinSettingUpdateOne) AddPaymentAccountCoinAmount(u int64) *CoinSettingUpdateOne {
	csuo.mutation.AddPaymentAccountCoinAmount(u)
	return csuo
}

// SetPlatformOfflineAccountID sets the "platform_offline_account_id" field.
func (csuo *CoinSettingUpdateOne) SetPlatformOfflineAccountID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetPlatformOfflineAccountID(u)
	return csuo
}

// SetUserOnlineAccountID sets the "user_online_account_id" field.
func (csuo *CoinSettingUpdateOne) SetUserOnlineAccountID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetUserOnlineAccountID(u)
	return csuo
}

// SetUserOfflineAccountID sets the "user_offline_account_id" field.
func (csuo *CoinSettingUpdateOne) SetUserOfflineAccountID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetUserOfflineAccountID(u)
	return csuo
}

// SetGoodIncomingAccountID sets the "good_incoming_account_id" field.
func (csuo *CoinSettingUpdateOne) SetGoodIncomingAccountID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetGoodIncomingAccountID(u)
	return csuo
}

// SetGasProviderAccountID sets the "gas_provider_account_id" field.
func (csuo *CoinSettingUpdateOne) SetGasProviderAccountID(u uuid.UUID) *CoinSettingUpdateOne {
	csuo.mutation.SetGasProviderAccountID(u)
	return csuo
}

// SetCreateAt sets the "create_at" field.
func (csuo *CoinSettingUpdateOne) SetCreateAt(u uint32) *CoinSettingUpdateOne {
	csuo.mutation.ResetCreateAt()
	csuo.mutation.SetCreateAt(u)
	return csuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (csuo *CoinSettingUpdateOne) SetNillableCreateAt(u *uint32) *CoinSettingUpdateOne {
	if u != nil {
		csuo.SetCreateAt(*u)
	}
	return csuo
}

// AddCreateAt adds u to the "create_at" field.
func (csuo *CoinSettingUpdateOne) AddCreateAt(u int32) *CoinSettingUpdateOne {
	csuo.mutation.AddCreateAt(u)
	return csuo
}

// SetUpdateAt sets the "update_at" field.
func (csuo *CoinSettingUpdateOne) SetUpdateAt(u uint32) *CoinSettingUpdateOne {
	csuo.mutation.ResetUpdateAt()
	csuo.mutation.SetUpdateAt(u)
	return csuo
}

// AddUpdateAt adds u to the "update_at" field.
func (csuo *CoinSettingUpdateOne) AddUpdateAt(u int32) *CoinSettingUpdateOne {
	csuo.mutation.AddUpdateAt(u)
	return csuo
}

// SetDeleteAt sets the "delete_at" field.
func (csuo *CoinSettingUpdateOne) SetDeleteAt(u uint32) *CoinSettingUpdateOne {
	csuo.mutation.ResetDeleteAt()
	csuo.mutation.SetDeleteAt(u)
	return csuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (csuo *CoinSettingUpdateOne) SetNillableDeleteAt(u *uint32) *CoinSettingUpdateOne {
	if u != nil {
		csuo.SetDeleteAt(*u)
	}
	return csuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (csuo *CoinSettingUpdateOne) AddDeleteAt(u int32) *CoinSettingUpdateOne {
	csuo.mutation.AddDeleteAt(u)
	return csuo
}

// Mutation returns the CoinSettingMutation object of the builder.
func (csuo *CoinSettingUpdateOne) Mutation() *CoinSettingMutation {
	return csuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csuo *CoinSettingUpdateOne) Select(field string, fields ...string) *CoinSettingUpdateOne {
	csuo.fields = append([]string{field}, fields...)
	return csuo
}

// Save executes the query and returns the updated CoinSetting entity.
func (csuo *CoinSettingUpdateOne) Save(ctx context.Context) (*CoinSetting, error) {
	var (
		err  error
		node *CoinSetting
	)
	csuo.defaults()
	if len(csuo.hooks) == 0 {
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			if csuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, csuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *CoinSettingUpdateOne) SaveX(ctx context.Context) *CoinSetting {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *CoinSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *CoinSettingUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csuo *CoinSettingUpdateOne) defaults() {
	if _, ok := csuo.mutation.UpdateAt(); !ok {
		v := coinsetting.UpdateDefaultUpdateAt()
		csuo.mutation.SetUpdateAt(v)
	}
}

func (csuo *CoinSettingUpdateOne) sqlSave(ctx context.Context) (_node *CoinSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinsetting.Table,
			Columns: coinsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsetting.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinsetting.FieldID)
		for _, f := range fields {
			if !coinsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldCoinTypeID,
		})
	}
	if value, ok := csuo.mutation.WarmAccountCoinAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldWarmAccountCoinAmount,
		})
	}
	if value, ok := csuo.mutation.AddedWarmAccountCoinAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldWarmAccountCoinAmount,
		})
	}
	if value, ok := csuo.mutation.PaymentAccountCoinAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldPaymentAccountCoinAmount,
		})
	}
	if value, ok := csuo.mutation.AddedPaymentAccountCoinAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coinsetting.FieldPaymentAccountCoinAmount,
		})
	}
	if value, ok := csuo.mutation.PlatformOfflineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldPlatformOfflineAccountID,
		})
	}
	if value, ok := csuo.mutation.UserOnlineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldUserOnlineAccountID,
		})
	}
	if value, ok := csuo.mutation.UserOfflineAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldUserOfflineAccountID,
		})
	}
	if value, ok := csuo.mutation.GoodIncomingAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldGoodIncomingAccountID,
		})
	}
	if value, ok := csuo.mutation.GasProviderAccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinsetting.FieldGasProviderAccountID,
		})
	}
	if value, ok := csuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldCreateAt,
		})
	}
	if value, ok := csuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldCreateAt,
		})
	}
	if value, ok := csuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldUpdateAt,
		})
	}
	if value, ok := csuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldUpdateAt,
		})
	}
	if value, ok := csuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldDeleteAt,
		})
	}
	if value, ok := csuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsetting.FieldDeleteAt,
		})
	}
	_node = &CoinSetting{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
