// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/applang"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/predicate"
)

// AppLangDelete is the builder for deleting a AppLang entity.
type AppLangDelete struct {
	config
	hooks    []Hook
	mutation *AppLangMutation
}

// Where appends a list predicates to the AppLangDelete builder.
func (ald *AppLangDelete) Where(ps ...predicate.AppLang) *AppLangDelete {
	ald.mutation.Where(ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *AppLangDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ald.hooks) == 0 {
		affected, err = ald.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppLangMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ald.mutation = mutation
			affected, err = ald.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ald.hooks) - 1; i >= 0; i-- {
			if ald.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ald.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ald.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *AppLangDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *AppLangDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: applang.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applang.FieldID,
			},
		},
	}
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppLangDeleteOne is the builder for deleting a single AppLang entity.
type AppLangDeleteOne struct {
	ald *AppLangDelete
}

// Exec executes the deletion query.
func (aldo *AppLangDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{applang.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *AppLangDeleteOne) ExecX(ctx context.Context) {
	aldo.ald.ExecX(ctx)
}
