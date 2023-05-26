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
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/country"
	"github.com/google/uuid"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CountryCreate) SetCreatedAt(u uint32) *CountryCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCreatedAt(u *uint32) *CountryCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CountryCreate) SetUpdatedAt(u uint32) *CountryCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedAt(u *uint32) *CountryCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CountryCreate) SetDeletedAt(u uint32) *CountryCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableDeletedAt(u *uint32) *CountryCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetCountry sets the "country" field.
func (cc *CountryCreate) SetCountry(s string) *CountryCreate {
	cc.mutation.SetCountry(s)
	return cc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCountry(s *string) *CountryCreate {
	if s != nil {
		cc.SetCountry(*s)
	}
	return cc
}

// SetFlag sets the "flag" field.
func (cc *CountryCreate) SetFlag(s string) *CountryCreate {
	cc.mutation.SetFlag(s)
	return cc
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (cc *CountryCreate) SetNillableFlag(s *string) *CountryCreate {
	if s != nil {
		cc.SetFlag(*s)
	}
	return cc
}

// SetCode sets the "code" field.
func (cc *CountryCreate) SetCode(s string) *CountryCreate {
	cc.mutation.SetCode(s)
	return cc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCode(s *string) *CountryCreate {
	if s != nil {
		cc.SetCode(*s)
	}
	return cc
}

// SetShort sets the "short" field.
func (cc *CountryCreate) SetShort(s string) *CountryCreate {
	cc.mutation.SetShort(s)
	return cc
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (cc *CountryCreate) SetNillableShort(s *string) *CountryCreate {
	if s != nil {
		cc.SetShort(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CountryCreate) SetID(u uuid.UUID) *CountryCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CountryCreate) SetNillableID(u *uuid.UUID) *CountryCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	var (
		err  error
		node *Country
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Country)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CountryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CountryCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if country.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if country.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if country.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.Country(); !ok {
		v := country.DefaultCountry
		cc.mutation.SetCountry(v)
	}
	if _, ok := cc.mutation.Flag(); !ok {
		v := country.DefaultFlag
		cc.mutation.SetFlag(v)
	}
	if _, ok := cc.mutation.Code(); !ok {
		v := country.DefaultCode
		cc.mutation.SetCode(v)
	}
	if _, ok := cc.mutation.Short(); !ok {
		v := country.DefaultShort
		cc.mutation.SetShort(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if country.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultID (forgotten import ent/runtime?)")
		}
		v := country.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Country.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Country.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Country.deleted_at"`)}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: country.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: country.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: country.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: country.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: country.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := cc.mutation.Flag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldFlag,
		})
		_node.Flag = value
	}
	if value, ok := cc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := cc.mutation.Short(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldShort,
		})
		_node.Short = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Country.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CountryCreate) OnConflict(opts ...sql.ConflictOption) *CountryUpsertOne {
	cc.conflict = opts
	return &CountryUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CountryCreate) OnConflictColumns(columns ...string) *CountryUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CountryUpsertOne{
		create: cc,
	}
}

type (
	// CountryUpsertOne is the builder for "upsert"-ing
	//  one Country node.
	CountryUpsertOne struct {
		create *CountryCreate
	}

	// CountryUpsert is the "OnConflict" setter.
	CountryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CountryUpsert) SetCreatedAt(v uint32) *CountryUpsert {
	u.Set(country.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CountryUpsert) UpdateCreatedAt() *CountryUpsert {
	u.SetExcluded(country.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CountryUpsert) AddCreatedAt(v uint32) *CountryUpsert {
	u.Add(country.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsert) SetUpdatedAt(v uint32) *CountryUpsert {
	u.Set(country.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsert) UpdateUpdatedAt() *CountryUpsert {
	u.SetExcluded(country.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CountryUpsert) AddUpdatedAt(v uint32) *CountryUpsert {
	u.Add(country.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CountryUpsert) SetDeletedAt(v uint32) *CountryUpsert {
	u.Set(country.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CountryUpsert) UpdateDeletedAt() *CountryUpsert {
	u.SetExcluded(country.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CountryUpsert) AddDeletedAt(v uint32) *CountryUpsert {
	u.Add(country.FieldDeletedAt, v)
	return u
}

// SetCountry sets the "country" field.
func (u *CountryUpsert) SetCountry(v string) *CountryUpsert {
	u.Set(country.FieldCountry, v)
	return u
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *CountryUpsert) UpdateCountry() *CountryUpsert {
	u.SetExcluded(country.FieldCountry)
	return u
}

// ClearCountry clears the value of the "country" field.
func (u *CountryUpsert) ClearCountry() *CountryUpsert {
	u.SetNull(country.FieldCountry)
	return u
}

// SetFlag sets the "flag" field.
func (u *CountryUpsert) SetFlag(v string) *CountryUpsert {
	u.Set(country.FieldFlag, v)
	return u
}

// UpdateFlag sets the "flag" field to the value that was provided on create.
func (u *CountryUpsert) UpdateFlag() *CountryUpsert {
	u.SetExcluded(country.FieldFlag)
	return u
}

// ClearFlag clears the value of the "flag" field.
func (u *CountryUpsert) ClearFlag() *CountryUpsert {
	u.SetNull(country.FieldFlag)
	return u
}

// SetCode sets the "code" field.
func (u *CountryUpsert) SetCode(v string) *CountryUpsert {
	u.Set(country.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsert) UpdateCode() *CountryUpsert {
	u.SetExcluded(country.FieldCode)
	return u
}

// ClearCode clears the value of the "code" field.
func (u *CountryUpsert) ClearCode() *CountryUpsert {
	u.SetNull(country.FieldCode)
	return u
}

// SetShort sets the "short" field.
func (u *CountryUpsert) SetShort(v string) *CountryUpsert {
	u.Set(country.FieldShort, v)
	return u
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *CountryUpsert) UpdateShort() *CountryUpsert {
	u.SetExcluded(country.FieldShort)
	return u
}

// ClearShort clears the value of the "short" field.
func (u *CountryUpsert) ClearShort() *CountryUpsert {
	u.SetNull(country.FieldShort)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(country.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CountryUpsertOne) UpdateNewValues() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(country.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Country.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CountryUpsertOne) Ignore() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryUpsertOne) DoNothing() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryCreate.OnConflict
// documentation for more info.
func (u *CountryUpsertOne) Update(set func(*CountryUpsert)) *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CountryUpsertOne) SetCreatedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CountryUpsertOne) AddCreatedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateCreatedAt() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsertOne) SetUpdatedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CountryUpsertOne) AddUpdatedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateUpdatedAt() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CountryUpsertOne) SetDeletedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CountryUpsertOne) AddDeletedAt(v uint32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateDeletedAt() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCountry sets the "country" field.
func (u *CountryUpsertOne) SetCountry(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateCountry() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCountry()
	})
}

// ClearCountry clears the value of the "country" field.
func (u *CountryUpsertOne) ClearCountry() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearCountry()
	})
}

// SetFlag sets the "flag" field.
func (u *CountryUpsertOne) SetFlag(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetFlag(v)
	})
}

// UpdateFlag sets the "flag" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateFlag() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateFlag()
	})
}

// ClearFlag clears the value of the "flag" field.
func (u *CountryUpsertOne) ClearFlag() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearFlag()
	})
}

// SetCode sets the "code" field.
func (u *CountryUpsertOne) SetCode(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateCode() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *CountryUpsertOne) ClearCode() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearCode()
	})
}

// SetShort sets the "short" field.
func (u *CountryUpsertOne) SetShort(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateShort() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *CountryUpsertOne) ClearShort() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *CountryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CountryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CountryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CountryUpsertOne.ID is not supported by MySQL driver. Use CountryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CountryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	builders []*CountryCreate
	conflict []sql.ConflictOption
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Country.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CountryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CountryUpsertBulk {
	ccb.conflict = opts
	return &CountryUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CountryCreateBulk) OnConflictColumns(columns ...string) *CountryUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CountryUpsertBulk{
		create: ccb,
	}
}

// CountryUpsertBulk is the builder for "upsert"-ing
// a bulk of Country nodes.
type CountryUpsertBulk struct {
	create *CountryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(country.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CountryUpsertBulk) UpdateNewValues() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(country.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CountryUpsertBulk) Ignore() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryUpsertBulk) DoNothing() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryCreateBulk.OnConflict
// documentation for more info.
func (u *CountryUpsertBulk) Update(set func(*CountryUpsert)) *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CountryUpsertBulk) SetCreatedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CountryUpsertBulk) AddCreatedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateCreatedAt() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsertBulk) SetUpdatedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CountryUpsertBulk) AddUpdatedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateUpdatedAt() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CountryUpsertBulk) SetDeletedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CountryUpsertBulk) AddDeletedAt(v uint32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateDeletedAt() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCountry sets the "country" field.
func (u *CountryUpsertBulk) SetCountry(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateCountry() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCountry()
	})
}

// ClearCountry clears the value of the "country" field.
func (u *CountryUpsertBulk) ClearCountry() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearCountry()
	})
}

// SetFlag sets the "flag" field.
func (u *CountryUpsertBulk) SetFlag(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetFlag(v)
	})
}

// UpdateFlag sets the "flag" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateFlag() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateFlag()
	})
}

// ClearFlag clears the value of the "flag" field.
func (u *CountryUpsertBulk) ClearFlag() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearFlag()
	})
}

// SetCode sets the "code" field.
func (u *CountryUpsertBulk) SetCode(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateCode() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *CountryUpsertBulk) ClearCode() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearCode()
	})
}

// SetShort sets the "short" field.
func (u *CountryUpsertBulk) SetShort(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateShort() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *CountryUpsertBulk) ClearShort() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *CountryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CountryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CountryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
