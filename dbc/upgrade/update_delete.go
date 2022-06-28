// Code generated by entc, DO NOT EDIT.

package upgrade

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/upgrade/predicate"
	"github.com/tikafog/of/dbc/upgrade/update"
)

// UpdateDelete is the builder for deleting a Update entity.
type UpdateDelete struct {
	config
	hooks    []Hook
	mutation *UpdateMutation
}

// Where appends a list predicates to the UpdateDelete builder.
func (ud *UpdateDelete) Where(ps ...predicate.Update) *UpdateDelete {
	ud.mutation.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UpdateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ud.hooks) == 0 {
		affected, err = ud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UpdateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ud.mutation = mutation
			affected, err = ud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ud.hooks) - 1; i >= 0; i-- {
			if ud.hooks[i] == nil {
				return 0, fmt.Errorf("upgrade: uninitialized hook (forgotten import upgrade/runtime?)")
			}
			mut = ud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UpdateDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UpdateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: update.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: update.FieldID,
			},
		},
	}
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
}

// UpdateDeleteOne is the builder for deleting a single Update entity.
type UpdateDeleteOne struct {
	ud *UpdateDelete
}

// Exec executes the deletion query.
func (udo *UpdateDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{update.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UpdateDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
