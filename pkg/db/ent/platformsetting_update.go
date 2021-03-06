// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
)

// PlatformSettingUpdate is the builder for updating PlatformSetting entities.
type PlatformSettingUpdate struct {
	config
	hooks    []Hook
	mutation *PlatformSettingMutation
}

// Where appends a list predicates to the PlatformSettingUpdate builder.
func (psu *PlatformSettingUpdate) Where(ps ...predicate.PlatformSetting) *PlatformSettingUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (psu *PlatformSettingUpdate) SetWarmAccountUsdAmount(u uint64) *PlatformSettingUpdate {
	psu.mutation.ResetWarmAccountUsdAmount()
	psu.mutation.SetWarmAccountUsdAmount(u)
	return psu
}

// AddWarmAccountUsdAmount adds u to the "warm_account_usd_amount" field.
func (psu *PlatformSettingUpdate) AddWarmAccountUsdAmount(u int64) *PlatformSettingUpdate {
	psu.mutation.AddWarmAccountUsdAmount(u)
	return psu
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (psu *PlatformSettingUpdate) SetPaymentAccountUsdAmount(u uint64) *PlatformSettingUpdate {
	psu.mutation.ResetPaymentAccountUsdAmount()
	psu.mutation.SetPaymentAccountUsdAmount(u)
	return psu
}

// AddPaymentAccountUsdAmount adds u to the "payment_account_usd_amount" field.
func (psu *PlatformSettingUpdate) AddPaymentAccountUsdAmount(u int64) *PlatformSettingUpdate {
	psu.mutation.AddPaymentAccountUsdAmount(u)
	return psu
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (psu *PlatformSettingUpdate) SetWithdrawAutoReviewUsdAmount(u uint64) *PlatformSettingUpdate {
	psu.mutation.ResetWithdrawAutoReviewUsdAmount()
	psu.mutation.SetWithdrawAutoReviewUsdAmount(u)
	return psu
}

// AddWithdrawAutoReviewUsdAmount adds u to the "withdraw_auto_review_usd_amount" field.
func (psu *PlatformSettingUpdate) AddWithdrawAutoReviewUsdAmount(u int64) *PlatformSettingUpdate {
	psu.mutation.AddWithdrawAutoReviewUsdAmount(u)
	return psu
}

// SetCreateAt sets the "create_at" field.
func (psu *PlatformSettingUpdate) SetCreateAt(u uint32) *PlatformSettingUpdate {
	psu.mutation.ResetCreateAt()
	psu.mutation.SetCreateAt(u)
	return psu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (psu *PlatformSettingUpdate) SetNillableCreateAt(u *uint32) *PlatformSettingUpdate {
	if u != nil {
		psu.SetCreateAt(*u)
	}
	return psu
}

// AddCreateAt adds u to the "create_at" field.
func (psu *PlatformSettingUpdate) AddCreateAt(u int32) *PlatformSettingUpdate {
	psu.mutation.AddCreateAt(u)
	return psu
}

// SetUpdateAt sets the "update_at" field.
func (psu *PlatformSettingUpdate) SetUpdateAt(u uint32) *PlatformSettingUpdate {
	psu.mutation.ResetUpdateAt()
	psu.mutation.SetUpdateAt(u)
	return psu
}

// AddUpdateAt adds u to the "update_at" field.
func (psu *PlatformSettingUpdate) AddUpdateAt(u int32) *PlatformSettingUpdate {
	psu.mutation.AddUpdateAt(u)
	return psu
}

// SetDeleteAt sets the "delete_at" field.
func (psu *PlatformSettingUpdate) SetDeleteAt(u uint32) *PlatformSettingUpdate {
	psu.mutation.ResetDeleteAt()
	psu.mutation.SetDeleteAt(u)
	return psu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (psu *PlatformSettingUpdate) SetNillableDeleteAt(u *uint32) *PlatformSettingUpdate {
	if u != nil {
		psu.SetDeleteAt(*u)
	}
	return psu
}

// AddDeleteAt adds u to the "delete_at" field.
func (psu *PlatformSettingUpdate) AddDeleteAt(u int32) *PlatformSettingUpdate {
	psu.mutation.AddDeleteAt(u)
	return psu
}

// Mutation returns the PlatformSettingMutation object of the builder.
func (psu *PlatformSettingUpdate) Mutation() *PlatformSettingMutation {
	return psu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PlatformSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	psu.defaults()
	if len(psu.hooks) == 0 {
		affected, err = psu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psu.mutation = mutation
			affected, err = psu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psu.hooks) - 1; i >= 0; i-- {
			if psu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PlatformSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PlatformSettingUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PlatformSettingUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *PlatformSettingUpdate) defaults() {
	if _, ok := psu.mutation.UpdateAt(); !ok {
		v := platformsetting.UpdateDefaultUpdateAt()
		psu.mutation.SetUpdateAt(v)
	}
}

func (psu *PlatformSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platformsetting.Table,
			Columns: platformsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformsetting.FieldID,
			},
		},
	}
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.WarmAccountUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWarmAccountUsdAmount,
		})
	}
	if value, ok := psu.mutation.AddedWarmAccountUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWarmAccountUsdAmount,
		})
	}
	if value, ok := psu.mutation.PaymentAccountUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldPaymentAccountUsdAmount,
		})
	}
	if value, ok := psu.mutation.AddedPaymentAccountUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldPaymentAccountUsdAmount,
		})
	}
	if value, ok := psu.mutation.WithdrawAutoReviewUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWithdrawAutoReviewUsdAmount,
		})
	}
	if value, ok := psu.mutation.AddedWithdrawAutoReviewUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWithdrawAutoReviewUsdAmount,
		})
	}
	if value, ok := psu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldCreateAt,
		})
	}
	if value, ok := psu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldCreateAt,
		})
	}
	if value, ok := psu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldUpdateAt,
		})
	}
	if value, ok := psu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldUpdateAt,
		})
	}
	if value, ok := psu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldDeleteAt,
		})
	}
	if value, ok := psu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platformsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlatformSettingUpdateOne is the builder for updating a single PlatformSetting entity.
type PlatformSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlatformSettingMutation
}

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) SetWarmAccountUsdAmount(u uint64) *PlatformSettingUpdateOne {
	psuo.mutation.ResetWarmAccountUsdAmount()
	psuo.mutation.SetWarmAccountUsdAmount(u)
	return psuo
}

// AddWarmAccountUsdAmount adds u to the "warm_account_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) AddWarmAccountUsdAmount(u int64) *PlatformSettingUpdateOne {
	psuo.mutation.AddWarmAccountUsdAmount(u)
	return psuo
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) SetPaymentAccountUsdAmount(u uint64) *PlatformSettingUpdateOne {
	psuo.mutation.ResetPaymentAccountUsdAmount()
	psuo.mutation.SetPaymentAccountUsdAmount(u)
	return psuo
}

// AddPaymentAccountUsdAmount adds u to the "payment_account_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) AddPaymentAccountUsdAmount(u int64) *PlatformSettingUpdateOne {
	psuo.mutation.AddPaymentAccountUsdAmount(u)
	return psuo
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) SetWithdrawAutoReviewUsdAmount(u uint64) *PlatformSettingUpdateOne {
	psuo.mutation.ResetWithdrawAutoReviewUsdAmount()
	psuo.mutation.SetWithdrawAutoReviewUsdAmount(u)
	return psuo
}

// AddWithdrawAutoReviewUsdAmount adds u to the "withdraw_auto_review_usd_amount" field.
func (psuo *PlatformSettingUpdateOne) AddWithdrawAutoReviewUsdAmount(u int64) *PlatformSettingUpdateOne {
	psuo.mutation.AddWithdrawAutoReviewUsdAmount(u)
	return psuo
}

// SetCreateAt sets the "create_at" field.
func (psuo *PlatformSettingUpdateOne) SetCreateAt(u uint32) *PlatformSettingUpdateOne {
	psuo.mutation.ResetCreateAt()
	psuo.mutation.SetCreateAt(u)
	return psuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (psuo *PlatformSettingUpdateOne) SetNillableCreateAt(u *uint32) *PlatformSettingUpdateOne {
	if u != nil {
		psuo.SetCreateAt(*u)
	}
	return psuo
}

// AddCreateAt adds u to the "create_at" field.
func (psuo *PlatformSettingUpdateOne) AddCreateAt(u int32) *PlatformSettingUpdateOne {
	psuo.mutation.AddCreateAt(u)
	return psuo
}

// SetUpdateAt sets the "update_at" field.
func (psuo *PlatformSettingUpdateOne) SetUpdateAt(u uint32) *PlatformSettingUpdateOne {
	psuo.mutation.ResetUpdateAt()
	psuo.mutation.SetUpdateAt(u)
	return psuo
}

// AddUpdateAt adds u to the "update_at" field.
func (psuo *PlatformSettingUpdateOne) AddUpdateAt(u int32) *PlatformSettingUpdateOne {
	psuo.mutation.AddUpdateAt(u)
	return psuo
}

// SetDeleteAt sets the "delete_at" field.
func (psuo *PlatformSettingUpdateOne) SetDeleteAt(u uint32) *PlatformSettingUpdateOne {
	psuo.mutation.ResetDeleteAt()
	psuo.mutation.SetDeleteAt(u)
	return psuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (psuo *PlatformSettingUpdateOne) SetNillableDeleteAt(u *uint32) *PlatformSettingUpdateOne {
	if u != nil {
		psuo.SetDeleteAt(*u)
	}
	return psuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (psuo *PlatformSettingUpdateOne) AddDeleteAt(u int32) *PlatformSettingUpdateOne {
	psuo.mutation.AddDeleteAt(u)
	return psuo
}

// Mutation returns the PlatformSettingMutation object of the builder.
func (psuo *PlatformSettingUpdateOne) Mutation() *PlatformSettingMutation {
	return psuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PlatformSettingUpdateOne) Select(field string, fields ...string) *PlatformSettingUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PlatformSetting entity.
func (psuo *PlatformSettingUpdateOne) Save(ctx context.Context) (*PlatformSetting, error) {
	var (
		err  error
		node *PlatformSetting
	)
	psuo.defaults()
	if len(psuo.hooks) == 0 {
		node, err = psuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psuo.mutation = mutation
			node, err = psuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psuo.hooks) - 1; i >= 0; i-- {
			if psuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, psuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PlatformSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlatformSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PlatformSettingUpdateOne) SaveX(ctx context.Context) *PlatformSetting {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PlatformSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PlatformSettingUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *PlatformSettingUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdateAt(); !ok {
		v := platformsetting.UpdateDefaultUpdateAt()
		psuo.mutation.SetUpdateAt(v)
	}
}

func (psuo *PlatformSettingUpdateOne) sqlSave(ctx context.Context) (_node *PlatformSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platformsetting.Table,
			Columns: platformsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformsetting.FieldID,
			},
		},
	}
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlatformSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platformsetting.FieldID)
		for _, f := range fields {
			if !platformsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != platformsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.WarmAccountUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWarmAccountUsdAmount,
		})
	}
	if value, ok := psuo.mutation.AddedWarmAccountUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWarmAccountUsdAmount,
		})
	}
	if value, ok := psuo.mutation.PaymentAccountUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldPaymentAccountUsdAmount,
		})
	}
	if value, ok := psuo.mutation.AddedPaymentAccountUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldPaymentAccountUsdAmount,
		})
	}
	if value, ok := psuo.mutation.WithdrawAutoReviewUsdAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWithdrawAutoReviewUsdAmount,
		})
	}
	if value, ok := psuo.mutation.AddedWithdrawAutoReviewUsdAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWithdrawAutoReviewUsdAmount,
		})
	}
	if value, ok := psuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldCreateAt,
		})
	}
	if value, ok := psuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldCreateAt,
		})
	}
	if value, ok := psuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldUpdateAt,
		})
	}
	if value, ok := psuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldUpdateAt,
		})
	}
	if value, ok := psuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldDeleteAt,
		})
	}
	if value, ok := psuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldDeleteAt,
		})
	}
	_node = &PlatformSetting{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platformsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
