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
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodpayment"
	"github.com/google/uuid"
)

// GoodPaymentCreate is the builder for creating a GoodPayment entity.
type GoodPaymentCreate struct {
	config
	mutation *GoodPaymentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetGoodID sets the "good_id" field.
func (gpc *GoodPaymentCreate) SetGoodID(u uuid.UUID) *GoodPaymentCreate {
	gpc.mutation.SetGoodID(u)
	return gpc
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (gpc *GoodPaymentCreate) SetPaymentCoinTypeID(u uuid.UUID) *GoodPaymentCreate {
	gpc.mutation.SetPaymentCoinTypeID(u)
	return gpc
}

// SetAccountID sets the "account_id" field.
func (gpc *GoodPaymentCreate) SetAccountID(u uuid.UUID) *GoodPaymentCreate {
	gpc.mutation.SetAccountID(u)
	return gpc
}

// SetIdle sets the "idle" field.
func (gpc *GoodPaymentCreate) SetIdle(b bool) *GoodPaymentCreate {
	gpc.mutation.SetIdle(b)
	return gpc
}

// SetCreateAt sets the "create_at" field.
func (gpc *GoodPaymentCreate) SetCreateAt(u uint32) *GoodPaymentCreate {
	gpc.mutation.SetCreateAt(u)
	return gpc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (gpc *GoodPaymentCreate) SetNillableCreateAt(u *uint32) *GoodPaymentCreate {
	if u != nil {
		gpc.SetCreateAt(*u)
	}
	return gpc
}

// SetUpdateAt sets the "update_at" field.
func (gpc *GoodPaymentCreate) SetUpdateAt(u uint32) *GoodPaymentCreate {
	gpc.mutation.SetUpdateAt(u)
	return gpc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (gpc *GoodPaymentCreate) SetNillableUpdateAt(u *uint32) *GoodPaymentCreate {
	if u != nil {
		gpc.SetUpdateAt(*u)
	}
	return gpc
}

// SetDeleteAt sets the "delete_at" field.
func (gpc *GoodPaymentCreate) SetDeleteAt(u uint32) *GoodPaymentCreate {
	gpc.mutation.SetDeleteAt(u)
	return gpc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gpc *GoodPaymentCreate) SetNillableDeleteAt(u *uint32) *GoodPaymentCreate {
	if u != nil {
		gpc.SetDeleteAt(*u)
	}
	return gpc
}

// SetID sets the "id" field.
func (gpc *GoodPaymentCreate) SetID(u uuid.UUID) *GoodPaymentCreate {
	gpc.mutation.SetID(u)
	return gpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gpc *GoodPaymentCreate) SetNillableID(u *uuid.UUID) *GoodPaymentCreate {
	if u != nil {
		gpc.SetID(*u)
	}
	return gpc
}

// Mutation returns the GoodPaymentMutation object of the builder.
func (gpc *GoodPaymentCreate) Mutation() *GoodPaymentMutation {
	return gpc.mutation
}

// Save creates the GoodPayment in the database.
func (gpc *GoodPaymentCreate) Save(ctx context.Context) (*GoodPayment, error) {
	var (
		err  error
		node *GoodPayment
	)
	gpc.defaults()
	if len(gpc.hooks) == 0 {
		if err = gpc.check(); err != nil {
			return nil, err
		}
		node, err = gpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodPaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gpc.check(); err != nil {
				return nil, err
			}
			gpc.mutation = mutation
			if node, err = gpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gpc.hooks) - 1; i >= 0; i-- {
			if gpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gpc *GoodPaymentCreate) SaveX(ctx context.Context) *GoodPayment {
	v, err := gpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gpc *GoodPaymentCreate) Exec(ctx context.Context) error {
	_, err := gpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpc *GoodPaymentCreate) ExecX(ctx context.Context) {
	if err := gpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gpc *GoodPaymentCreate) defaults() {
	if _, ok := gpc.mutation.CreateAt(); !ok {
		v := goodpayment.DefaultCreateAt()
		gpc.mutation.SetCreateAt(v)
	}
	if _, ok := gpc.mutation.UpdateAt(); !ok {
		v := goodpayment.DefaultUpdateAt()
		gpc.mutation.SetUpdateAt(v)
	}
	if _, ok := gpc.mutation.DeleteAt(); !ok {
		v := goodpayment.DefaultDeleteAt()
		gpc.mutation.SetDeleteAt(v)
	}
	if _, ok := gpc.mutation.ID(); !ok {
		v := goodpayment.DefaultID()
		gpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gpc *GoodPaymentCreate) check() error {
	if _, ok := gpc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "GoodPayment.good_id"`)}
	}
	if _, ok := gpc.mutation.PaymentCoinTypeID(); !ok {
		return &ValidationError{Name: "payment_coin_type_id", err: errors.New(`ent: missing required field "GoodPayment.payment_coin_type_id"`)}
	}
	if _, ok := gpc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "GoodPayment.account_id"`)}
	}
	if _, ok := gpc.mutation.Idle(); !ok {
		return &ValidationError{Name: "idle", err: errors.New(`ent: missing required field "GoodPayment.idle"`)}
	}
	if _, ok := gpc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "GoodPayment.create_at"`)}
	}
	if _, ok := gpc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "GoodPayment.update_at"`)}
	}
	if _, ok := gpc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "GoodPayment.delete_at"`)}
	}
	return nil
}

func (gpc *GoodPaymentCreate) sqlSave(ctx context.Context) (*GoodPayment, error) {
	_node, _spec := gpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gpc.driver, _spec); err != nil {
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

func (gpc *GoodPaymentCreate) createSpec() (*GoodPayment, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodPayment{config: gpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodpayment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodpayment.FieldID,
			},
		}
	)
	_spec.OnConflict = gpc.conflict
	if id, ok := gpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gpc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodpayment.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := gpc.mutation.PaymentCoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodpayment.FieldPaymentCoinTypeID,
		})
		_node.PaymentCoinTypeID = value
	}
	if value, ok := gpc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodpayment.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := gpc.mutation.Idle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodpayment.FieldIdle,
		})
		_node.Idle = value
	}
	if value, ok := gpc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodpayment.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := gpc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodpayment.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := gpc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodpayment.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodPayment.Create().
//		SetGoodID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodPaymentUpsert) {
//			SetGoodID(v+v).
//		}).
//		Exec(ctx)
//
func (gpc *GoodPaymentCreate) OnConflict(opts ...sql.ConflictOption) *GoodPaymentUpsertOne {
	gpc.conflict = opts
	return &GoodPaymentUpsertOne{
		create: gpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodPayment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gpc *GoodPaymentCreate) OnConflictColumns(columns ...string) *GoodPaymentUpsertOne {
	gpc.conflict = append(gpc.conflict, sql.ConflictColumns(columns...))
	return &GoodPaymentUpsertOne{
		create: gpc,
	}
}

type (
	// GoodPaymentUpsertOne is the builder for "upsert"-ing
	//  one GoodPayment node.
	GoodPaymentUpsertOne struct {
		create *GoodPaymentCreate
	}

	// GoodPaymentUpsert is the "OnConflict" setter.
	GoodPaymentUpsert struct {
		*sql.UpdateSet
	}
)

// SetGoodID sets the "good_id" field.
func (u *GoodPaymentUpsert) SetGoodID(v uuid.UUID) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateGoodID() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldGoodID)
	return u
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *GoodPaymentUpsert) SetPaymentCoinTypeID(v uuid.UUID) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldPaymentCoinTypeID, v)
	return u
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdatePaymentCoinTypeID() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldPaymentCoinTypeID)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *GoodPaymentUpsert) SetAccountID(v uuid.UUID) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateAccountID() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldAccountID)
	return u
}

// SetIdle sets the "idle" field.
func (u *GoodPaymentUpsert) SetIdle(v bool) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldIdle, v)
	return u
}

// UpdateIdle sets the "idle" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateIdle() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldIdle)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *GoodPaymentUpsert) SetCreateAt(v uint32) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateCreateAt() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *GoodPaymentUpsert) AddCreateAt(v uint32) *GoodPaymentUpsert {
	u.Add(goodpayment.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodPaymentUpsert) SetUpdateAt(v uint32) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateUpdateAt() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *GoodPaymentUpsert) AddUpdateAt(v uint32) *GoodPaymentUpsert {
	u.Add(goodpayment.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodPaymentUpsert) SetDeleteAt(v uint32) *GoodPaymentUpsert {
	u.Set(goodpayment.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodPaymentUpsert) UpdateDeleteAt() *GoodPaymentUpsert {
	u.SetExcluded(goodpayment.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *GoodPaymentUpsert) AddDeleteAt(v uint32) *GoodPaymentUpsert {
	u.Add(goodpayment.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GoodPayment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodpayment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodPaymentUpsertOne) UpdateNewValues() *GoodPaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(goodpayment.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.GoodPayment.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *GoodPaymentUpsertOne) Ignore() *GoodPaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodPaymentUpsertOne) DoNothing() *GoodPaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodPaymentCreate.OnConflict
// documentation for more info.
func (u *GoodPaymentUpsertOne) Update(set func(*GoodPaymentUpsert)) *GoodPaymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodPaymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetGoodID sets the "good_id" field.
func (u *GoodPaymentUpsertOne) SetGoodID(v uuid.UUID) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateGoodID() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateGoodID()
	})
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *GoodPaymentUpsertOne) SetPaymentCoinTypeID(v uuid.UUID) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetPaymentCoinTypeID(v)
	})
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdatePaymentCoinTypeID() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdatePaymentCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *GoodPaymentUpsertOne) SetAccountID(v uuid.UUID) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateAccountID() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateAccountID()
	})
}

// SetIdle sets the "idle" field.
func (u *GoodPaymentUpsertOne) SetIdle(v bool) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetIdle(v)
	})
}

// UpdateIdle sets the "idle" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateIdle() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateIdle()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *GoodPaymentUpsertOne) SetCreateAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *GoodPaymentUpsertOne) AddCreateAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateCreateAt() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodPaymentUpsertOne) SetUpdateAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *GoodPaymentUpsertOne) AddUpdateAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateUpdateAt() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodPaymentUpsertOne) SetDeleteAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *GoodPaymentUpsertOne) AddDeleteAt(v uint32) *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertOne) UpdateDeleteAt() *GoodPaymentUpsertOne {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *GoodPaymentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodPaymentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodPaymentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoodPaymentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GoodPaymentUpsertOne.ID is not supported by MySQL driver. Use GoodPaymentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoodPaymentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoodPaymentCreateBulk is the builder for creating many GoodPayment entities in bulk.
type GoodPaymentCreateBulk struct {
	config
	builders []*GoodPaymentCreate
	conflict []sql.ConflictOption
}

// Save creates the GoodPayment entities in the database.
func (gpcb *GoodPaymentCreateBulk) Save(ctx context.Context) ([]*GoodPayment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gpcb.builders))
	nodes := make([]*GoodPayment, len(gpcb.builders))
	mutators := make([]Mutator, len(gpcb.builders))
	for i := range gpcb.builders {
		func(i int, root context.Context) {
			builder := gpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodPaymentMutation)
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
					_, err = mutators[i+1].Mutate(root, gpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gpcb *GoodPaymentCreateBulk) SaveX(ctx context.Context) []*GoodPayment {
	v, err := gpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gpcb *GoodPaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := gpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpcb *GoodPaymentCreateBulk) ExecX(ctx context.Context) {
	if err := gpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodPayment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodPaymentUpsert) {
//			SetGoodID(v+v).
//		}).
//		Exec(ctx)
//
func (gpcb *GoodPaymentCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoodPaymentUpsertBulk {
	gpcb.conflict = opts
	return &GoodPaymentUpsertBulk{
		create: gpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodPayment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gpcb *GoodPaymentCreateBulk) OnConflictColumns(columns ...string) *GoodPaymentUpsertBulk {
	gpcb.conflict = append(gpcb.conflict, sql.ConflictColumns(columns...))
	return &GoodPaymentUpsertBulk{
		create: gpcb,
	}
}

// GoodPaymentUpsertBulk is the builder for "upsert"-ing
// a bulk of GoodPayment nodes.
type GoodPaymentUpsertBulk struct {
	create *GoodPaymentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoodPayment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodpayment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodPaymentUpsertBulk) UpdateNewValues() *GoodPaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(goodpayment.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoodPayment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *GoodPaymentUpsertBulk) Ignore() *GoodPaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodPaymentUpsertBulk) DoNothing() *GoodPaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodPaymentCreateBulk.OnConflict
// documentation for more info.
func (u *GoodPaymentUpsertBulk) Update(set func(*GoodPaymentUpsert)) *GoodPaymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodPaymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetGoodID sets the "good_id" field.
func (u *GoodPaymentUpsertBulk) SetGoodID(v uuid.UUID) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateGoodID() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateGoodID()
	})
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *GoodPaymentUpsertBulk) SetPaymentCoinTypeID(v uuid.UUID) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetPaymentCoinTypeID(v)
	})
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdatePaymentCoinTypeID() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdatePaymentCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *GoodPaymentUpsertBulk) SetAccountID(v uuid.UUID) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateAccountID() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateAccountID()
	})
}

// SetIdle sets the "idle" field.
func (u *GoodPaymentUpsertBulk) SetIdle(v bool) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetIdle(v)
	})
}

// UpdateIdle sets the "idle" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateIdle() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateIdle()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *GoodPaymentUpsertBulk) SetCreateAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *GoodPaymentUpsertBulk) AddCreateAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateCreateAt() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodPaymentUpsertBulk) SetUpdateAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *GoodPaymentUpsertBulk) AddUpdateAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateUpdateAt() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodPaymentUpsertBulk) SetDeleteAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *GoodPaymentUpsertBulk) AddDeleteAt(v uint32) *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodPaymentUpsertBulk) UpdateDeleteAt() *GoodPaymentUpsertBulk {
	return u.Update(func(s *GoodPaymentUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *GoodPaymentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoodPaymentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodPaymentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodPaymentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
