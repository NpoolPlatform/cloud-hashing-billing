// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userdirectbenefit"
)

// UserDirectBenefitDelete is the builder for deleting a UserDirectBenefit entity.
type UserDirectBenefitDelete struct {
	config
	hooks    []Hook
	mutation *UserDirectBenefitMutation
}

// Where appends a list predicates to the UserDirectBenefitDelete builder.
func (udbd *UserDirectBenefitDelete) Where(ps ...predicate.UserDirectBenefit) *UserDirectBenefitDelete {
	udbd.mutation.Where(ps...)
	return udbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (udbd *UserDirectBenefitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(udbd.hooks) == 0 {
		affected, err = udbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserDirectBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			udbd.mutation = mutation
			affected, err = udbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(udbd.hooks) - 1; i >= 0; i-- {
			if udbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = udbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, udbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (udbd *UserDirectBenefitDelete) ExecX(ctx context.Context) int {
	n, err := udbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (udbd *UserDirectBenefitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userdirectbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userdirectbenefit.FieldID,
			},
		},
	}
	if ps := udbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, udbd.driver, _spec)
}

// UserDirectBenefitDeleteOne is the builder for deleting a single UserDirectBenefit entity.
type UserDirectBenefitDeleteOne struct {
	udbd *UserDirectBenefitDelete
}

// Exec executes the deletion query.
func (udbdo *UserDirectBenefitDeleteOne) Exec(ctx context.Context) error {
	n, err := udbdo.udbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userdirectbenefit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udbdo *UserDirectBenefitDeleteOne) ExecX(ctx context.Context) {
	udbdo.udbd.ExecX(ctx)
}
