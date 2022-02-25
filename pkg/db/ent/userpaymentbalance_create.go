// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userpaymentbalance"
	"github.com/google/uuid"
)

// UserPaymentBalanceCreate is the builder for creating a UserPaymentBalance entity.
type UserPaymentBalanceCreate struct {
	config
	mutation *UserPaymentBalanceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (upbc *UserPaymentBalanceCreate) SetAppID(u uuid.UUID) *UserPaymentBalanceCreate {
	upbc.mutation.SetAppID(u)
	return upbc
}

// SetUserID sets the "user_id" field.
func (upbc *UserPaymentBalanceCreate) SetUserID(u uuid.UUID) *UserPaymentBalanceCreate {
	upbc.mutation.SetUserID(u)
	return upbc
}

// SetPaymentID sets the "payment_id" field.
func (upbc *UserPaymentBalanceCreate) SetPaymentID(u uuid.UUID) *UserPaymentBalanceCreate {
	upbc.mutation.SetPaymentID(u)
	return upbc
}

// SetAmount sets the "amount" field.
func (upbc *UserPaymentBalanceCreate) SetAmount(u uint64) *UserPaymentBalanceCreate {
	upbc.mutation.SetAmount(u)
	return upbc
}

// SetCreateAt sets the "create_at" field.
func (upbc *UserPaymentBalanceCreate) SetCreateAt(u uint32) *UserPaymentBalanceCreate {
	upbc.mutation.SetCreateAt(u)
	return upbc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (upbc *UserPaymentBalanceCreate) SetNillableCreateAt(u *uint32) *UserPaymentBalanceCreate {
	if u != nil {
		upbc.SetCreateAt(*u)
	}
	return upbc
}

// SetUpdateAt sets the "update_at" field.
func (upbc *UserPaymentBalanceCreate) SetUpdateAt(u uint32) *UserPaymentBalanceCreate {
	upbc.mutation.SetUpdateAt(u)
	return upbc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (upbc *UserPaymentBalanceCreate) SetNillableUpdateAt(u *uint32) *UserPaymentBalanceCreate {
	if u != nil {
		upbc.SetUpdateAt(*u)
	}
	return upbc
}

// SetDeleteAt sets the "delete_at" field.
func (upbc *UserPaymentBalanceCreate) SetDeleteAt(u uint32) *UserPaymentBalanceCreate {
	upbc.mutation.SetDeleteAt(u)
	return upbc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (upbc *UserPaymentBalanceCreate) SetNillableDeleteAt(u *uint32) *UserPaymentBalanceCreate {
	if u != nil {
		upbc.SetDeleteAt(*u)
	}
	return upbc
}

// SetID sets the "id" field.
func (upbc *UserPaymentBalanceCreate) SetID(u uuid.UUID) *UserPaymentBalanceCreate {
	upbc.mutation.SetID(u)
	return upbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (upbc *UserPaymentBalanceCreate) SetNillableID(u *uuid.UUID) *UserPaymentBalanceCreate {
	if u != nil {
		upbc.SetID(*u)
	}
	return upbc
}

// Mutation returns the UserPaymentBalanceMutation object of the builder.
func (upbc *UserPaymentBalanceCreate) Mutation() *UserPaymentBalanceMutation {
	return upbc.mutation
}

// Save creates the UserPaymentBalance in the database.
func (upbc *UserPaymentBalanceCreate) Save(ctx context.Context) (*UserPaymentBalance, error) {
	var (
		err  error
		node *UserPaymentBalance
	)
	upbc.defaults()
	if len(upbc.hooks) == 0 {
		if err = upbc.check(); err != nil {
			return nil, err
		}
		node, err = upbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserPaymentBalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upbc.check(); err != nil {
				return nil, err
			}
			upbc.mutation = mutation
			if node, err = upbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(upbc.hooks) - 1; i >= 0; i-- {
			if upbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (upbc *UserPaymentBalanceCreate) SaveX(ctx context.Context) *UserPaymentBalance {
	v, err := upbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upbc *UserPaymentBalanceCreate) Exec(ctx context.Context) error {
	_, err := upbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upbc *UserPaymentBalanceCreate) ExecX(ctx context.Context) {
	if err := upbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upbc *UserPaymentBalanceCreate) defaults() {
	if _, ok := upbc.mutation.CreateAt(); !ok {
		v := userpaymentbalance.DefaultCreateAt()
		upbc.mutation.SetCreateAt(v)
	}
	if _, ok := upbc.mutation.UpdateAt(); !ok {
		v := userpaymentbalance.DefaultUpdateAt()
		upbc.mutation.SetUpdateAt(v)
	}
	if _, ok := upbc.mutation.DeleteAt(); !ok {
		v := userpaymentbalance.DefaultDeleteAt()
		upbc.mutation.SetDeleteAt(v)
	}
	if _, ok := upbc.mutation.ID(); !ok {
		v := userpaymentbalance.DefaultID()
		upbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upbc *UserPaymentBalanceCreate) check() error {
	if _, ok := upbc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "UserPaymentBalance.app_id"`)}
	}
	if _, ok := upbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserPaymentBalance.user_id"`)}
	}
	if _, ok := upbc.mutation.PaymentID(); !ok {
		return &ValidationError{Name: "payment_id", err: errors.New(`ent: missing required field "UserPaymentBalance.payment_id"`)}
	}
	if _, ok := upbc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "UserPaymentBalance.amount"`)}
	}
	if _, ok := upbc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "UserPaymentBalance.create_at"`)}
	}
	if _, ok := upbc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "UserPaymentBalance.update_at"`)}
	}
	if _, ok := upbc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "UserPaymentBalance.delete_at"`)}
	}
	return nil
}

func (upbc *UserPaymentBalanceCreate) sqlSave(ctx context.Context) (*UserPaymentBalance, error) {
	_node, _spec := upbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (upbc *UserPaymentBalanceCreate) createSpec() (*UserPaymentBalance, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPaymentBalance{config: upbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userpaymentbalance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userpaymentbalance.FieldID,
			},
		}
	)
	_spec.OnConflict = upbc.conflict
	if id, ok := upbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := upbc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userpaymentbalance.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := upbc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userpaymentbalance.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := upbc.mutation.PaymentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userpaymentbalance.FieldPaymentID,
		})
		_node.PaymentID = value
	}
	if value, ok := upbc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userpaymentbalance.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := upbc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userpaymentbalance.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := upbc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userpaymentbalance.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := upbc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userpaymentbalance.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserPaymentBalance.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserPaymentBalanceUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (upbc *UserPaymentBalanceCreate) OnConflict(opts ...sql.ConflictOption) *UserPaymentBalanceUpsertOne {
	upbc.conflict = opts
	return &UserPaymentBalanceUpsertOne{
		create: upbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserPaymentBalance.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (upbc *UserPaymentBalanceCreate) OnConflictColumns(columns ...string) *UserPaymentBalanceUpsertOne {
	upbc.conflict = append(upbc.conflict, sql.ConflictColumns(columns...))
	return &UserPaymentBalanceUpsertOne{
		create: upbc,
	}
}

type (
	// UserPaymentBalanceUpsertOne is the builder for "upsert"-ing
	//  one UserPaymentBalance node.
	UserPaymentBalanceUpsertOne struct {
		create *UserPaymentBalanceCreate
	}

	// UserPaymentBalanceUpsert is the "OnConflict" setter.
	UserPaymentBalanceUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *UserPaymentBalanceUpsert) SetAppID(v uuid.UUID) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateAppID() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserPaymentBalanceUpsert) SetUserID(v uuid.UUID) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateUserID() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldUserID)
	return u
}

// SetPaymentID sets the "payment_id" field.
func (u *UserPaymentBalanceUpsert) SetPaymentID(v uuid.UUID) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldPaymentID, v)
	return u
}

// UpdatePaymentID sets the "payment_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdatePaymentID() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldPaymentID)
	return u
}

// SetAmount sets the "amount" field.
func (u *UserPaymentBalanceUpsert) SetAmount(v uint64) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateAmount() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *UserPaymentBalanceUpsert) AddAmount(v uint64) *UserPaymentBalanceUpsert {
	u.Add(userpaymentbalance.FieldAmount, v)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *UserPaymentBalanceUpsert) SetCreateAt(v uint32) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateCreateAt() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserPaymentBalanceUpsert) AddCreateAt(v uint32) *UserPaymentBalanceUpsert {
	u.Add(userpaymentbalance.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *UserPaymentBalanceUpsert) SetUpdateAt(v uint32) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateUpdateAt() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserPaymentBalanceUpsert) AddUpdateAt(v uint32) *UserPaymentBalanceUpsert {
	u.Add(userpaymentbalance.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserPaymentBalanceUpsert) SetDeleteAt(v uint32) *UserPaymentBalanceUpsert {
	u.Set(userpaymentbalance.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsert) UpdateDeleteAt() *UserPaymentBalanceUpsert {
	u.SetExcluded(userpaymentbalance.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserPaymentBalanceUpsert) AddDeleteAt(v uint32) *UserPaymentBalanceUpsert {
	u.Add(userpaymentbalance.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserPaymentBalance.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userpaymentbalance.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserPaymentBalanceUpsertOne) UpdateNewValues() *UserPaymentBalanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userpaymentbalance.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserPaymentBalance.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserPaymentBalanceUpsertOne) Ignore() *UserPaymentBalanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserPaymentBalanceUpsertOne) DoNothing() *UserPaymentBalanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserPaymentBalanceCreate.OnConflict
// documentation for more info.
func (u *UserPaymentBalanceUpsertOne) Update(set func(*UserPaymentBalanceUpsert)) *UserPaymentBalanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserPaymentBalanceUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserPaymentBalanceUpsertOne) SetAppID(v uuid.UUID) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateAppID() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserPaymentBalanceUpsertOne) SetUserID(v uuid.UUID) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateUserID() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateUserID()
	})
}

// SetPaymentID sets the "payment_id" field.
func (u *UserPaymentBalanceUpsertOne) SetPaymentID(v uuid.UUID) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetPaymentID(v)
	})
}

// UpdatePaymentID sets the "payment_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdatePaymentID() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdatePaymentID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserPaymentBalanceUpsertOne) SetAmount(v uint64) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UserPaymentBalanceUpsertOne) AddAmount(v uint64) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateAmount() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateAmount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserPaymentBalanceUpsertOne) SetCreateAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserPaymentBalanceUpsertOne) AddCreateAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateCreateAt() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserPaymentBalanceUpsertOne) SetUpdateAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserPaymentBalanceUpsertOne) AddUpdateAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateUpdateAt() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserPaymentBalanceUpsertOne) SetDeleteAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserPaymentBalanceUpsertOne) AddDeleteAt(v uint32) *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertOne) UpdateDeleteAt() *UserPaymentBalanceUpsertOne {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserPaymentBalanceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserPaymentBalanceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserPaymentBalanceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserPaymentBalanceUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserPaymentBalanceUpsertOne.ID is not supported by MySQL driver. Use UserPaymentBalanceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserPaymentBalanceUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserPaymentBalanceCreateBulk is the builder for creating many UserPaymentBalance entities in bulk.
type UserPaymentBalanceCreateBulk struct {
	config
	builders []*UserPaymentBalanceCreate
	conflict []sql.ConflictOption
}

// Save creates the UserPaymentBalance entities in the database.
func (upbcb *UserPaymentBalanceCreateBulk) Save(ctx context.Context) ([]*UserPaymentBalance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upbcb.builders))
	nodes := make([]*UserPaymentBalance, len(upbcb.builders))
	mutators := make([]Mutator, len(upbcb.builders))
	for i := range upbcb.builders {
		func(i int, root context.Context) {
			builder := upbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPaymentBalanceMutation)
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
					_, err = mutators[i+1].Mutate(root, upbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = upbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, upbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upbcb *UserPaymentBalanceCreateBulk) SaveX(ctx context.Context) []*UserPaymentBalance {
	v, err := upbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upbcb *UserPaymentBalanceCreateBulk) Exec(ctx context.Context) error {
	_, err := upbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upbcb *UserPaymentBalanceCreateBulk) ExecX(ctx context.Context) {
	if err := upbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserPaymentBalance.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserPaymentBalanceUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (upbcb *UserPaymentBalanceCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserPaymentBalanceUpsertBulk {
	upbcb.conflict = opts
	return &UserPaymentBalanceUpsertBulk{
		create: upbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserPaymentBalance.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (upbcb *UserPaymentBalanceCreateBulk) OnConflictColumns(columns ...string) *UserPaymentBalanceUpsertBulk {
	upbcb.conflict = append(upbcb.conflict, sql.ConflictColumns(columns...))
	return &UserPaymentBalanceUpsertBulk{
		create: upbcb,
	}
}

// UserPaymentBalanceUpsertBulk is the builder for "upsert"-ing
// a bulk of UserPaymentBalance nodes.
type UserPaymentBalanceUpsertBulk struct {
	create *UserPaymentBalanceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserPaymentBalance.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userpaymentbalance.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserPaymentBalanceUpsertBulk) UpdateNewValues() *UserPaymentBalanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userpaymentbalance.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserPaymentBalance.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserPaymentBalanceUpsertBulk) Ignore() *UserPaymentBalanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserPaymentBalanceUpsertBulk) DoNothing() *UserPaymentBalanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserPaymentBalanceCreateBulk.OnConflict
// documentation for more info.
func (u *UserPaymentBalanceUpsertBulk) Update(set func(*UserPaymentBalanceUpsert)) *UserPaymentBalanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserPaymentBalanceUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserPaymentBalanceUpsertBulk) SetAppID(v uuid.UUID) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateAppID() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserPaymentBalanceUpsertBulk) SetUserID(v uuid.UUID) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateUserID() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateUserID()
	})
}

// SetPaymentID sets the "payment_id" field.
func (u *UserPaymentBalanceUpsertBulk) SetPaymentID(v uuid.UUID) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetPaymentID(v)
	})
}

// UpdatePaymentID sets the "payment_id" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdatePaymentID() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdatePaymentID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserPaymentBalanceUpsertBulk) SetAmount(v uint64) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UserPaymentBalanceUpsertBulk) AddAmount(v uint64) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateAmount() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateAmount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserPaymentBalanceUpsertBulk) SetCreateAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserPaymentBalanceUpsertBulk) AddCreateAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateCreateAt() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserPaymentBalanceUpsertBulk) SetUpdateAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserPaymentBalanceUpsertBulk) AddUpdateAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateUpdateAt() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserPaymentBalanceUpsertBulk) SetDeleteAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserPaymentBalanceUpsertBulk) AddDeleteAt(v uint32) *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserPaymentBalanceUpsertBulk) UpdateDeleteAt() *UserPaymentBalanceUpsertBulk {
	return u.Update(func(s *UserPaymentBalanceUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserPaymentBalanceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserPaymentBalanceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserPaymentBalanceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserPaymentBalanceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}