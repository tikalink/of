// Code generated by entc, DO NOT EDIT.

package kernel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tikafog/of/dbc/kernel/instruct"
)

// InstructCreate is the builder for creating a Instruct entity.
type InstructCreate struct {
	config
	mutation *InstructMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCurrentUnix sets the "current_unix" field.
func (ic *InstructCreate) SetCurrentUnix(i int64) *InstructCreate {
	ic.mutation.SetCurrentUnix(i)
	return ic
}

// SetNillableCurrentUnix sets the "current_unix" field if the given value is not nil.
func (ic *InstructCreate) SetNillableCurrentUnix(i *int64) *InstructCreate {
	if i != nil {
		ic.SetCurrentUnix(*i)
	}
	return ic
}

// SetUpdatedUnix sets the "updated_unix" field.
func (ic *InstructCreate) SetUpdatedUnix(i int64) *InstructCreate {
	ic.mutation.SetUpdatedUnix(i)
	return ic
}

// SetNillableUpdatedUnix sets the "updated_unix" field if the given value is not nil.
func (ic *InstructCreate) SetNillableUpdatedUnix(i *int64) *InstructCreate {
	if i != nil {
		ic.SetUpdatedUnix(*i)
	}
	return ic
}

// Mutation returns the InstructMutation object of the builder.
func (ic *InstructCreate) Mutation() *InstructMutation {
	return ic.mutation
}

// Save creates the Instruct in the database.
func (ic *InstructCreate) Save(ctx context.Context) (*Instruct, error) {
	var (
		err  error
		node *Instruct
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("kernel: uninitialized hook (forgotten import kernel/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstructCreate) SaveX(ctx context.Context) *Instruct {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InstructCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InstructCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InstructCreate) defaults() {
	if _, ok := ic.mutation.CurrentUnix(); !ok {
		v := instruct.DefaultCurrentUnix
		ic.mutation.SetCurrentUnix(v)
	}
	if _, ok := ic.mutation.UpdatedUnix(); !ok {
		v := instruct.DefaultUpdatedUnix
		ic.mutation.SetUpdatedUnix(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InstructCreate) check() error {
	if _, ok := ic.mutation.CurrentUnix(); !ok {
		return &ValidationError{Name: "current_unix", err: errors.New(`kernel: missing required field "Instruct.current_unix"`)}
	}
	if _, ok := ic.mutation.UpdatedUnix(); !ok {
		return &ValidationError{Name: "updated_unix", err: errors.New(`kernel: missing required field "Instruct.updated_unix"`)}
	}
	return nil
}

func (ic *InstructCreate) sqlSave(ctx context.Context) (*Instruct, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *InstructCreate) createSpec() (*Instruct, *sqlgraph.CreateSpec) {
	var (
		_node = &Instruct{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instruct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instruct.FieldID,
			},
		}
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.CurrentUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: instruct.FieldCurrentUnix,
		})
		_node.CurrentUnix = value
	}
	if value, ok := ic.mutation.UpdatedUnix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: instruct.FieldUpdatedUnix,
		})
		_node.UpdatedUnix = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Instruct.Create().
//		SetCurrentUnix(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InstructUpsert) {
//			SetCurrentUnix(v+v).
//		}).
//		Exec(ctx)
//
func (ic *InstructCreate) OnConflict(opts ...sql.ConflictOption) *InstructUpsertOne {
	ic.conflict = opts
	return &InstructUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Instruct.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ic *InstructCreate) OnConflictColumns(columns ...string) *InstructUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &InstructUpsertOne{
		create: ic,
	}
}

type (
	// InstructUpsertOne is the builder for "upsert"-ing
	//  one Instruct node.
	InstructUpsertOne struct {
		create *InstructCreate
	}

	// InstructUpsert is the "OnConflict" setter.
	InstructUpsert struct {
		*sql.UpdateSet
	}
)

// SetCurrentUnix sets the "current_unix" field.
func (u *InstructUpsert) SetCurrentUnix(v int64) *InstructUpsert {
	u.Set(instruct.FieldCurrentUnix, v)
	return u
}

// UpdateCurrentUnix sets the "current_unix" field to the value that was provided on create.
func (u *InstructUpsert) UpdateCurrentUnix() *InstructUpsert {
	u.SetExcluded(instruct.FieldCurrentUnix)
	return u
}

// AddCurrentUnix adds v to the "current_unix" field.
func (u *InstructUpsert) AddCurrentUnix(v int64) *InstructUpsert {
	u.Add(instruct.FieldCurrentUnix, v)
	return u
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *InstructUpsert) SetUpdatedUnix(v int64) *InstructUpsert {
	u.Set(instruct.FieldUpdatedUnix, v)
	return u
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *InstructUpsert) UpdateUpdatedUnix() *InstructUpsert {
	u.SetExcluded(instruct.FieldUpdatedUnix)
	return u
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *InstructUpsert) AddUpdatedUnix(v int64) *InstructUpsert {
	u.Add(instruct.FieldUpdatedUnix, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Instruct.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *InstructUpsertOne) UpdateNewValues() *InstructUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Instruct.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *InstructUpsertOne) Ignore() *InstructUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InstructUpsertOne) DoNothing() *InstructUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InstructCreate.OnConflict
// documentation for more info.
func (u *InstructUpsertOne) Update(set func(*InstructUpsert)) *InstructUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InstructUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrentUnix sets the "current_unix" field.
func (u *InstructUpsertOne) SetCurrentUnix(v int64) *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.SetCurrentUnix(v)
	})
}

// AddCurrentUnix adds v to the "current_unix" field.
func (u *InstructUpsertOne) AddCurrentUnix(v int64) *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.AddCurrentUnix(v)
	})
}

// UpdateCurrentUnix sets the "current_unix" field to the value that was provided on create.
func (u *InstructUpsertOne) UpdateCurrentUnix() *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.UpdateCurrentUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *InstructUpsertOne) SetUpdatedUnix(v int64) *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *InstructUpsertOne) AddUpdatedUnix(v int64) *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *InstructUpsertOne) UpdateUpdatedUnix() *InstructUpsertOne {
	return u.Update(func(s *InstructUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// Exec executes the query.
func (u *InstructUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("kernel: missing options for InstructCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InstructUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InstructUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InstructUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InstructCreateBulk is the builder for creating many Instruct entities in bulk.
type InstructCreateBulk struct {
	config
	builders []*InstructCreate
	conflict []sql.ConflictOption
}

// Save creates the Instruct entities in the database.
func (icb *InstructCreateBulk) Save(ctx context.Context) ([]*Instruct, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Instruct, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstructMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InstructCreateBulk) SaveX(ctx context.Context) []*Instruct {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InstructCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InstructCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Instruct.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InstructUpsert) {
//			SetCurrentUnix(v+v).
//		}).
//		Exec(ctx)
//
func (icb *InstructCreateBulk) OnConflict(opts ...sql.ConflictOption) *InstructUpsertBulk {
	icb.conflict = opts
	return &InstructUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Instruct.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (icb *InstructCreateBulk) OnConflictColumns(columns ...string) *InstructUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &InstructUpsertBulk{
		create: icb,
	}
}

// InstructUpsertBulk is the builder for "upsert"-ing
// a bulk of Instruct nodes.
type InstructUpsertBulk struct {
	create *InstructCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Instruct.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *InstructUpsertBulk) UpdateNewValues() *InstructUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Instruct.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *InstructUpsertBulk) Ignore() *InstructUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InstructUpsertBulk) DoNothing() *InstructUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InstructCreateBulk.OnConflict
// documentation for more info.
func (u *InstructUpsertBulk) Update(set func(*InstructUpsert)) *InstructUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InstructUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrentUnix sets the "current_unix" field.
func (u *InstructUpsertBulk) SetCurrentUnix(v int64) *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.SetCurrentUnix(v)
	})
}

// AddCurrentUnix adds v to the "current_unix" field.
func (u *InstructUpsertBulk) AddCurrentUnix(v int64) *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.AddCurrentUnix(v)
	})
}

// UpdateCurrentUnix sets the "current_unix" field to the value that was provided on create.
func (u *InstructUpsertBulk) UpdateCurrentUnix() *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.UpdateCurrentUnix()
	})
}

// SetUpdatedUnix sets the "updated_unix" field.
func (u *InstructUpsertBulk) SetUpdatedUnix(v int64) *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.SetUpdatedUnix(v)
	})
}

// AddUpdatedUnix adds v to the "updated_unix" field.
func (u *InstructUpsertBulk) AddUpdatedUnix(v int64) *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.AddUpdatedUnix(v)
	})
}

// UpdateUpdatedUnix sets the "updated_unix" field to the value that was provided on create.
func (u *InstructUpsertBulk) UpdateUpdatedUnix() *InstructUpsertBulk {
	return u.Update(func(s *InstructUpsert) {
		s.UpdateUpdatedUnix()
	})
}

// Exec executes the query.
func (u *InstructUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("kernel: OnConflict was set for builder %d. Set it on the InstructCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("kernel: missing options for InstructCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InstructUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
