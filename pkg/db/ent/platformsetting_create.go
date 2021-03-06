// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/google/uuid"
)

// PlatformSettingCreate is the builder for creating a PlatformSetting entity.
type PlatformSettingCreate struct {
	config
	mutation *PlatformSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (psc *PlatformSettingCreate) SetWarmAccountUsdAmount(u uint64) *PlatformSettingCreate {
	psc.mutation.SetWarmAccountUsdAmount(u)
	return psc
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (psc *PlatformSettingCreate) SetPaymentAccountUsdAmount(u uint64) *PlatformSettingCreate {
	psc.mutation.SetPaymentAccountUsdAmount(u)
	return psc
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (psc *PlatformSettingCreate) SetWithdrawAutoReviewUsdAmount(u uint64) *PlatformSettingCreate {
	psc.mutation.SetWithdrawAutoReviewUsdAmount(u)
	return psc
}

// SetCreateAt sets the "create_at" field.
func (psc *PlatformSettingCreate) SetCreateAt(u uint32) *PlatformSettingCreate {
	psc.mutation.SetCreateAt(u)
	return psc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (psc *PlatformSettingCreate) SetNillableCreateAt(u *uint32) *PlatformSettingCreate {
	if u != nil {
		psc.SetCreateAt(*u)
	}
	return psc
}

// SetUpdateAt sets the "update_at" field.
func (psc *PlatformSettingCreate) SetUpdateAt(u uint32) *PlatformSettingCreate {
	psc.mutation.SetUpdateAt(u)
	return psc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (psc *PlatformSettingCreate) SetNillableUpdateAt(u *uint32) *PlatformSettingCreate {
	if u != nil {
		psc.SetUpdateAt(*u)
	}
	return psc
}

// SetDeleteAt sets the "delete_at" field.
func (psc *PlatformSettingCreate) SetDeleteAt(u uint32) *PlatformSettingCreate {
	psc.mutation.SetDeleteAt(u)
	return psc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (psc *PlatformSettingCreate) SetNillableDeleteAt(u *uint32) *PlatformSettingCreate {
	if u != nil {
		psc.SetDeleteAt(*u)
	}
	return psc
}

// SetID sets the "id" field.
func (psc *PlatformSettingCreate) SetID(u uuid.UUID) *PlatformSettingCreate {
	psc.mutation.SetID(u)
	return psc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psc *PlatformSettingCreate) SetNillableID(u *uuid.UUID) *PlatformSettingCreate {
	if u != nil {
		psc.SetID(*u)
	}
	return psc
}

// Mutation returns the PlatformSettingMutation object of the builder.
func (psc *PlatformSettingCreate) Mutation() *PlatformSettingMutation {
	return psc.mutation
}

// Save creates the PlatformSetting in the database.
func (psc *PlatformSettingCreate) Save(ctx context.Context) (*PlatformSetting, error) {
	var (
		err  error
		node *PlatformSetting
	)
	psc.defaults()
	if len(psc.hooks) == 0 {
		if err = psc.check(); err != nil {
			return nil, err
		}
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psc.check(); err != nil {
				return nil, err
			}
			psc.mutation = mutation
			if node, err = psc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			if psc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, psc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (psc *PlatformSettingCreate) SaveX(ctx context.Context) *PlatformSetting {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PlatformSettingCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PlatformSettingCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *PlatformSettingCreate) defaults() {
	if _, ok := psc.mutation.CreateAt(); !ok {
		v := platformsetting.DefaultCreateAt()
		psc.mutation.SetCreateAt(v)
	}
	if _, ok := psc.mutation.UpdateAt(); !ok {
		v := platformsetting.DefaultUpdateAt()
		psc.mutation.SetUpdateAt(v)
	}
	if _, ok := psc.mutation.DeleteAt(); !ok {
		v := platformsetting.DefaultDeleteAt()
		psc.mutation.SetDeleteAt(v)
	}
	if _, ok := psc.mutation.ID(); !ok {
		v := platformsetting.DefaultID()
		psc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PlatformSettingCreate) check() error {
	if _, ok := psc.mutation.WarmAccountUsdAmount(); !ok {
		return &ValidationError{Name: "warm_account_usd_amount", err: errors.New(`ent: missing required field "PlatformSetting.warm_account_usd_amount"`)}
	}
	if _, ok := psc.mutation.PaymentAccountUsdAmount(); !ok {
		return &ValidationError{Name: "payment_account_usd_amount", err: errors.New(`ent: missing required field "PlatformSetting.payment_account_usd_amount"`)}
	}
	if _, ok := psc.mutation.WithdrawAutoReviewUsdAmount(); !ok {
		return &ValidationError{Name: "withdraw_auto_review_usd_amount", err: errors.New(`ent: missing required field "PlatformSetting.withdraw_auto_review_usd_amount"`)}
	}
	if _, ok := psc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "PlatformSetting.create_at"`)}
	}
	if _, ok := psc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "PlatformSetting.update_at"`)}
	}
	if _, ok := psc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "PlatformSetting.delete_at"`)}
	}
	return nil
}

func (psc *PlatformSettingCreate) sqlSave(ctx context.Context) (*PlatformSetting, error) {
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (psc *PlatformSettingCreate) createSpec() (*PlatformSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &PlatformSetting{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: platformsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = psc.conflict
	if id, ok := psc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psc.mutation.WarmAccountUsdAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWarmAccountUsdAmount,
		})
		_node.WarmAccountUsdAmount = value
	}
	if value, ok := psc.mutation.PaymentAccountUsdAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldPaymentAccountUsdAmount,
		})
		_node.PaymentAccountUsdAmount = value
	}
	if value, ok := psc.mutation.WithdrawAutoReviewUsdAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformsetting.FieldWithdrawAutoReviewUsdAmount,
		})
		_node.WithdrawAutoReviewUsdAmount = value
	}
	if value, ok := psc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := psc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := psc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformsetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlatformSetting.Create().
//		SetWarmAccountUsdAmount(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformSettingUpsert) {
//			SetWarmAccountUsdAmount(v+v).
//		}).
//		Exec(ctx)
//
func (psc *PlatformSettingCreate) OnConflict(opts ...sql.ConflictOption) *PlatformSettingUpsertOne {
	psc.conflict = opts
	return &PlatformSettingUpsertOne{
		create: psc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlatformSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (psc *PlatformSettingCreate) OnConflictColumns(columns ...string) *PlatformSettingUpsertOne {
	psc.conflict = append(psc.conflict, sql.ConflictColumns(columns...))
	return &PlatformSettingUpsertOne{
		create: psc,
	}
}

type (
	// PlatformSettingUpsertOne is the builder for "upsert"-ing
	//  one PlatformSetting node.
	PlatformSettingUpsertOne struct {
		create *PlatformSettingCreate
	}

	// PlatformSettingUpsert is the "OnConflict" setter.
	PlatformSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsert) SetWarmAccountUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldWarmAccountUsdAmount, v)
	return u
}

// UpdateWarmAccountUsdAmount sets the "warm_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdateWarmAccountUsdAmount() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldWarmAccountUsdAmount)
	return u
}

// AddWarmAccountUsdAmount adds v to the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsert) AddWarmAccountUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldWarmAccountUsdAmount, v)
	return u
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsert) SetPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldPaymentAccountUsdAmount, v)
	return u
}

// UpdatePaymentAccountUsdAmount sets the "payment_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdatePaymentAccountUsdAmount() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldPaymentAccountUsdAmount)
	return u
}

// AddPaymentAccountUsdAmount adds v to the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsert) AddPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldPaymentAccountUsdAmount, v)
	return u
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsert) SetWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldWithdrawAutoReviewUsdAmount, v)
	return u
}

// UpdateWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdateWithdrawAutoReviewUsdAmount() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldWithdrawAutoReviewUsdAmount)
	return u
}

// AddWithdrawAutoReviewUsdAmount adds v to the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsert) AddWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldWithdrawAutoReviewUsdAmount, v)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformSettingUpsert) SetCreateAt(v uint32) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdateCreateAt() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformSettingUpsert) AddCreateAt(v uint32) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformSettingUpsert) SetUpdateAt(v uint32) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdateUpdateAt() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformSettingUpsert) AddUpdateAt(v uint32) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformSettingUpsert) SetDeleteAt(v uint32) *PlatformSettingUpsert {
	u.Set(platformsetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformSettingUpsert) UpdateDeleteAt() *PlatformSettingUpsert {
	u.SetExcluded(platformsetting.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformSettingUpsert) AddDeleteAt(v uint32) *PlatformSettingUpsert {
	u.Add(platformsetting.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PlatformSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(platformsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PlatformSettingUpsertOne) UpdateNewValues() *PlatformSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(platformsetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.PlatformSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PlatformSettingUpsertOne) Ignore() *PlatformSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformSettingUpsertOne) DoNothing() *PlatformSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformSettingCreate.OnConflict
// documentation for more info.
func (u *PlatformSettingUpsertOne) Update(set func(*PlatformSettingUpsert)) *PlatformSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsertOne) SetWarmAccountUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetWarmAccountUsdAmount(v)
	})
}

// AddWarmAccountUsdAmount adds v to the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsertOne) AddWarmAccountUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddWarmAccountUsdAmount(v)
	})
}

// UpdateWarmAccountUsdAmount sets the "warm_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdateWarmAccountUsdAmount() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateWarmAccountUsdAmount()
	})
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsertOne) SetPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetPaymentAccountUsdAmount(v)
	})
}

// AddPaymentAccountUsdAmount adds v to the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsertOne) AddPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddPaymentAccountUsdAmount(v)
	})
}

// UpdatePaymentAccountUsdAmount sets the "payment_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdatePaymentAccountUsdAmount() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdatePaymentAccountUsdAmount()
	})
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsertOne) SetWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetWithdrawAutoReviewUsdAmount(v)
	})
}

// AddWithdrawAutoReviewUsdAmount adds v to the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsertOne) AddWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddWithdrawAutoReviewUsdAmount(v)
	})
}

// UpdateWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdateWithdrawAutoReviewUsdAmount() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateWithdrawAutoReviewUsdAmount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformSettingUpsertOne) SetCreateAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformSettingUpsertOne) AddCreateAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdateCreateAt() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformSettingUpsertOne) SetUpdateAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformSettingUpsertOne) AddUpdateAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdateUpdateAt() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformSettingUpsertOne) SetDeleteAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformSettingUpsertOne) AddDeleteAt(v uint32) *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertOne) UpdateDeleteAt() *PlatformSettingUpsertOne {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PlatformSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlatformSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlatformSettingUpsertOne.ID is not supported by MySQL driver. Use PlatformSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlatformSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlatformSettingCreateBulk is the builder for creating many PlatformSetting entities in bulk.
type PlatformSettingCreateBulk struct {
	config
	builders []*PlatformSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the PlatformSetting entities in the database.
func (pscb *PlatformSettingCreateBulk) Save(ctx context.Context) ([]*PlatformSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PlatformSetting, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlatformSettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PlatformSettingCreateBulk) SaveX(ctx context.Context) []*PlatformSetting {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PlatformSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PlatformSettingCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlatformSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformSettingUpsert) {
//			SetWarmAccountUsdAmount(v+v).
//		}).
//		Exec(ctx)
//
func (pscb *PlatformSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlatformSettingUpsertBulk {
	pscb.conflict = opts
	return &PlatformSettingUpsertBulk{
		create: pscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlatformSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pscb *PlatformSettingCreateBulk) OnConflictColumns(columns ...string) *PlatformSettingUpsertBulk {
	pscb.conflict = append(pscb.conflict, sql.ConflictColumns(columns...))
	return &PlatformSettingUpsertBulk{
		create: pscb,
	}
}

// PlatformSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of PlatformSetting nodes.
type PlatformSettingUpsertBulk struct {
	create *PlatformSettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PlatformSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(platformsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PlatformSettingUpsertBulk) UpdateNewValues() *PlatformSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(platformsetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlatformSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PlatformSettingUpsertBulk) Ignore() *PlatformSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformSettingUpsertBulk) DoNothing() *PlatformSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformSettingCreateBulk.OnConflict
// documentation for more info.
func (u *PlatformSettingUpsertBulk) Update(set func(*PlatformSettingUpsert)) *PlatformSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetWarmAccountUsdAmount sets the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsertBulk) SetWarmAccountUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetWarmAccountUsdAmount(v)
	})
}

// AddWarmAccountUsdAmount adds v to the "warm_account_usd_amount" field.
func (u *PlatformSettingUpsertBulk) AddWarmAccountUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddWarmAccountUsdAmount(v)
	})
}

// UpdateWarmAccountUsdAmount sets the "warm_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdateWarmAccountUsdAmount() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateWarmAccountUsdAmount()
	})
}

// SetPaymentAccountUsdAmount sets the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsertBulk) SetPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetPaymentAccountUsdAmount(v)
	})
}

// AddPaymentAccountUsdAmount adds v to the "payment_account_usd_amount" field.
func (u *PlatformSettingUpsertBulk) AddPaymentAccountUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddPaymentAccountUsdAmount(v)
	})
}

// UpdatePaymentAccountUsdAmount sets the "payment_account_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdatePaymentAccountUsdAmount() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdatePaymentAccountUsdAmount()
	})
}

// SetWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsertBulk) SetWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetWithdrawAutoReviewUsdAmount(v)
	})
}

// AddWithdrawAutoReviewUsdAmount adds v to the "withdraw_auto_review_usd_amount" field.
func (u *PlatformSettingUpsertBulk) AddWithdrawAutoReviewUsdAmount(v uint64) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddWithdrawAutoReviewUsdAmount(v)
	})
}

// UpdateWithdrawAutoReviewUsdAmount sets the "withdraw_auto_review_usd_amount" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdateWithdrawAutoReviewUsdAmount() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateWithdrawAutoReviewUsdAmount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformSettingUpsertBulk) SetCreateAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformSettingUpsertBulk) AddCreateAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdateCreateAt() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformSettingUpsertBulk) SetUpdateAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformSettingUpsertBulk) AddUpdateAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdateUpdateAt() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformSettingUpsertBulk) SetDeleteAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformSettingUpsertBulk) AddDeleteAt(v uint32) *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformSettingUpsertBulk) UpdateDeleteAt() *PlatformSettingUpsertBulk {
	return u.Update(func(s *PlatformSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PlatformSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlatformSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
