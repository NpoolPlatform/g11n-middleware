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
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/lang"
	"github.com/google/uuid"
)

// LangCreate is the builder for creating a Lang entity.
type LangCreate struct {
	config
	mutation *LangMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (lc *LangCreate) SetCreatedAt(u uint32) *LangCreate {
	lc.mutation.SetCreatedAt(u)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableCreatedAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetCreatedAt(*u)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LangCreate) SetUpdatedAt(u uint32) *LangCreate {
	lc.mutation.SetUpdatedAt(u)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableUpdatedAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetUpdatedAt(*u)
	}
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LangCreate) SetDeletedAt(u uint32) *LangCreate {
	lc.mutation.SetDeletedAt(u)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LangCreate) SetNillableDeletedAt(u *uint32) *LangCreate {
	if u != nil {
		lc.SetDeletedAt(*u)
	}
	return lc
}

// SetLang sets the "lang" field.
func (lc *LangCreate) SetLang(s string) *LangCreate {
	lc.mutation.SetLang(s)
	return lc
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (lc *LangCreate) SetNillableLang(s *string) *LangCreate {
	if s != nil {
		lc.SetLang(*s)
	}
	return lc
}

// SetLogo sets the "logo" field.
func (lc *LangCreate) SetLogo(s string) *LangCreate {
	lc.mutation.SetLogo(s)
	return lc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (lc *LangCreate) SetNillableLogo(s *string) *LangCreate {
	if s != nil {
		lc.SetLogo(*s)
	}
	return lc
}

// SetName sets the "name" field.
func (lc *LangCreate) SetName(s string) *LangCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lc *LangCreate) SetNillableName(s *string) *LangCreate {
	if s != nil {
		lc.SetName(*s)
	}
	return lc
}

// SetShort sets the "short" field.
func (lc *LangCreate) SetShort(s string) *LangCreate {
	lc.mutation.SetShort(s)
	return lc
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (lc *LangCreate) SetNillableShort(s *string) *LangCreate {
	if s != nil {
		lc.SetShort(*s)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LangCreate) SetID(u uuid.UUID) *LangCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LangCreate) SetNillableID(u *uuid.UUID) *LangCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// Mutation returns the LangMutation object of the builder.
func (lc *LangCreate) Mutation() *LangMutation {
	return lc.mutation
}

// Save creates the Lang in the database.
func (lc *LangCreate) Save(ctx context.Context) (*Lang, error) {
	var (
		err  error
		node *Lang
	)
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LangMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Lang)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LangMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LangCreate) SaveX(ctx context.Context) *Lang {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LangCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LangCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LangCreate) defaults() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		if lang.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized lang.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := lang.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		if lang.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized lang.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := lang.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.DeletedAt(); !ok {
		if lang.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized lang.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := lang.DefaultDeletedAt()
		lc.mutation.SetDeletedAt(v)
	}
	if _, ok := lc.mutation.Lang(); !ok {
		v := lang.DefaultLang
		lc.mutation.SetLang(v)
	}
	if _, ok := lc.mutation.Logo(); !ok {
		v := lang.DefaultLogo
		lc.mutation.SetLogo(v)
	}
	if _, ok := lc.mutation.Name(); !ok {
		v := lang.DefaultName
		lc.mutation.SetName(v)
	}
	if _, ok := lc.mutation.Short(); !ok {
		v := lang.DefaultShort
		lc.mutation.SetShort(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		if lang.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized lang.DefaultID (forgotten import ent/runtime?)")
		}
		v := lang.DefaultID()
		lc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LangCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Lang.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Lang.updated_at"`)}
	}
	if _, ok := lc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Lang.deleted_at"`)}
	}
	return nil
}

func (lc *LangCreate) sqlSave(ctx context.Context) (*Lang, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
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

func (lc *LangCreate) createSpec() (*Lang, *sqlgraph.CreateSpec) {
	var (
		_node = &Lang{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lang.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lang.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lang.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := lc.mutation.Lang(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLang,
		})
		_node.Lang = value
	}
	if value, ok := lc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.Short(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lang.FieldShort,
		})
		_node.Short = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lang.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LangUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (lc *LangCreate) OnConflict(opts ...sql.ConflictOption) *LangUpsertOne {
	lc.conflict = opts
	return &LangUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *LangCreate) OnConflictColumns(columns ...string) *LangUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LangUpsertOne{
		create: lc,
	}
}

type (
	// LangUpsertOne is the builder for "upsert"-ing
	//  one Lang node.
	LangUpsertOne struct {
		create *LangCreate
	}

	// LangUpsert is the "OnConflict" setter.
	LangUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *LangUpsert) SetCreatedAt(v uint32) *LangUpsert {
	u.Set(lang.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateCreatedAt() *LangUpsert {
	u.SetExcluded(lang.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LangUpsert) AddCreatedAt(v uint32) *LangUpsert {
	u.Add(lang.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LangUpsert) SetUpdatedAt(v uint32) *LangUpsert {
	u.Set(lang.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateUpdatedAt() *LangUpsert {
	u.SetExcluded(lang.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LangUpsert) AddUpdatedAt(v uint32) *LangUpsert {
	u.Add(lang.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LangUpsert) SetDeletedAt(v uint32) *LangUpsert {
	u.Set(lang.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LangUpsert) UpdateDeletedAt() *LangUpsert {
	u.SetExcluded(lang.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LangUpsert) AddDeletedAt(v uint32) *LangUpsert {
	u.Add(lang.FieldDeletedAt, v)
	return u
}

// SetLang sets the "lang" field.
func (u *LangUpsert) SetLang(v string) *LangUpsert {
	u.Set(lang.FieldLang, v)
	return u
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsert) UpdateLang() *LangUpsert {
	u.SetExcluded(lang.FieldLang)
	return u
}

// ClearLang clears the value of the "lang" field.
func (u *LangUpsert) ClearLang() *LangUpsert {
	u.SetNull(lang.FieldLang)
	return u
}

// SetLogo sets the "logo" field.
func (u *LangUpsert) SetLogo(v string) *LangUpsert {
	u.Set(lang.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *LangUpsert) UpdateLogo() *LangUpsert {
	u.SetExcluded(lang.FieldLogo)
	return u
}

// ClearLogo clears the value of the "logo" field.
func (u *LangUpsert) ClearLogo() *LangUpsert {
	u.SetNull(lang.FieldLogo)
	return u
}

// SetName sets the "name" field.
func (u *LangUpsert) SetName(v string) *LangUpsert {
	u.Set(lang.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LangUpsert) UpdateName() *LangUpsert {
	u.SetExcluded(lang.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *LangUpsert) ClearName() *LangUpsert {
	u.SetNull(lang.FieldName)
	return u
}

// SetShort sets the "short" field.
func (u *LangUpsert) SetShort(v string) *LangUpsert {
	u.Set(lang.FieldShort, v)
	return u
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *LangUpsert) UpdateShort() *LangUpsert {
	u.SetExcluded(lang.FieldShort)
	return u
}

// ClearShort clears the value of the "short" field.
func (u *LangUpsert) ClearShort() *LangUpsert {
	u.SetNull(lang.FieldShort)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lang.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LangUpsertOne) UpdateNewValues() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(lang.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Lang.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LangUpsertOne) Ignore() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LangUpsertOne) DoNothing() *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LangCreate.OnConflict
// documentation for more info.
func (u *LangUpsertOne) Update(set func(*LangUpsert)) *LangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LangUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LangUpsertOne) SetCreatedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LangUpsertOne) AddCreatedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateCreatedAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LangUpsertOne) SetUpdatedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LangUpsertOne) AddUpdatedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateUpdatedAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LangUpsertOne) SetDeletedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LangUpsertOne) AddDeletedAt(v uint32) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateDeletedAt() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetLang sets the "lang" field.
func (u *LangUpsertOne) SetLang(v string) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateLang() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLang()
	})
}

// ClearLang clears the value of the "lang" field.
func (u *LangUpsertOne) ClearLang() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.ClearLang()
	})
}

// SetLogo sets the "logo" field.
func (u *LangUpsertOne) SetLogo(v string) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateLogo() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *LangUpsertOne) ClearLogo() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.ClearLogo()
	})
}

// SetName sets the "name" field.
func (u *LangUpsertOne) SetName(v string) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateName() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *LangUpsertOne) ClearName() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.ClearName()
	})
}

// SetShort sets the "short" field.
func (u *LangUpsertOne) SetShort(v string) *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *LangUpsertOne) UpdateShort() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *LangUpsertOne) ClearShort() *LangUpsertOne {
	return u.Update(func(s *LangUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *LangUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LangCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LangUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LangUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LangUpsertOne.ID is not supported by MySQL driver. Use LangUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LangUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LangCreateBulk is the builder for creating many Lang entities in bulk.
type LangCreateBulk struct {
	config
	builders []*LangCreate
	conflict []sql.ConflictOption
}

// Save creates the Lang entities in the database.
func (lcb *LangCreateBulk) Save(ctx context.Context) ([]*Lang, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lang, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LangMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LangCreateBulk) SaveX(ctx context.Context) []*Lang {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LangCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LangCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lang.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LangUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *LangCreateBulk) OnConflict(opts ...sql.ConflictOption) *LangUpsertBulk {
	lcb.conflict = opts
	return &LangUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *LangCreateBulk) OnConflictColumns(columns ...string) *LangUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LangUpsertBulk{
		create: lcb,
	}
}

// LangUpsertBulk is the builder for "upsert"-ing
// a bulk of Lang nodes.
type LangUpsertBulk struct {
	create *LangCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lang.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LangUpsertBulk) UpdateNewValues() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(lang.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Lang.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LangUpsertBulk) Ignore() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LangUpsertBulk) DoNothing() *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LangCreateBulk.OnConflict
// documentation for more info.
func (u *LangUpsertBulk) Update(set func(*LangUpsert)) *LangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LangUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LangUpsertBulk) SetCreatedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LangUpsertBulk) AddCreatedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateCreatedAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LangUpsertBulk) SetUpdatedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LangUpsertBulk) AddUpdatedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateUpdatedAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LangUpsertBulk) SetDeletedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LangUpsertBulk) AddDeletedAt(v uint32) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateDeletedAt() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetLang sets the "lang" field.
func (u *LangUpsertBulk) SetLang(v string) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateLang() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLang()
	})
}

// ClearLang clears the value of the "lang" field.
func (u *LangUpsertBulk) ClearLang() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.ClearLang()
	})
}

// SetLogo sets the "logo" field.
func (u *LangUpsertBulk) SetLogo(v string) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateLogo() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *LangUpsertBulk) ClearLogo() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.ClearLogo()
	})
}

// SetName sets the "name" field.
func (u *LangUpsertBulk) SetName(v string) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateName() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *LangUpsertBulk) ClearName() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.ClearName()
	})
}

// SetShort sets the "short" field.
func (u *LangUpsertBulk) SetShort(v string) *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *LangUpsertBulk) UpdateShort() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *LangUpsertBulk) ClearShort() *LangUpsertBulk {
	return u.Update(func(s *LangUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *LangUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LangCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LangCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LangUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
