// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/appwithdrawsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
)

// AppWithdrawSettingDelete is the builder for deleting a AppWithdrawSetting entity.
type AppWithdrawSettingDelete struct {
	config
	hooks    []Hook
	mutation *AppWithdrawSettingMutation
}

// Where appends a list predicates to the AppWithdrawSettingDelete builder.
func (awsd *AppWithdrawSettingDelete) Where(ps ...predicate.AppWithdrawSetting) *AppWithdrawSettingDelete {
	awsd.mutation.Where(ps...)
	return awsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (awsd *AppWithdrawSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(awsd.hooks) == 0 {
		affected, err = awsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppWithdrawSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			awsd.mutation = mutation
			affected, err = awsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(awsd.hooks) - 1; i >= 0; i-- {
			if awsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = awsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (awsd *AppWithdrawSettingDelete) ExecX(ctx context.Context) int {
	n, err := awsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (awsd *AppWithdrawSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appwithdrawsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appwithdrawsetting.FieldID,
			},
		},
	}
	if ps := awsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, awsd.driver, _spec)
}

// AppWithdrawSettingDeleteOne is the builder for deleting a single AppWithdrawSetting entity.
type AppWithdrawSettingDeleteOne struct {
	awsd *AppWithdrawSettingDelete
}

// Exec executes the deletion query.
func (awsdo *AppWithdrawSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := awsdo.awsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appwithdrawsetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (awsdo *AppWithdrawSettingDeleteOne) ExecX(ctx context.Context) {
	awsdo.awsd.ExecX(ctx)
}