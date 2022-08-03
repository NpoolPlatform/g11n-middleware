// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/general"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/predicate"
)

// GeneralDelete is the builder for deleting a General entity.
type GeneralDelete struct {
	config
	hooks    []Hook
	mutation *GeneralMutation
}

// Where appends a list predicates to the GeneralDelete builder.
func (gd *GeneralDelete) Where(ps ...predicate.General) *GeneralDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GeneralDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gd.hooks) == 0 {
		affected, err = gd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gd.mutation = mutation
			affected, err = gd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gd.hooks) - 1; i >= 0; i-- {
			if gd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GeneralDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GeneralDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: general.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: general.FieldID,
			},
		},
	}
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GeneralDeleteOne is the builder for deleting a single General entity.
type GeneralDeleteOne struct {
	gd *GeneralDelete
}

// Exec executes the deletion query.
func (gdo *GeneralDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{general.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GeneralDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}