// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/appcountry"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppCountryUpdate is the builder for updating AppCountry entities.
type AppCountryUpdate struct {
	config
	hooks     []Hook
	mutation  *AppCountryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppCountryUpdate builder.
func (acu *AppCountryUpdate) Where(ps ...predicate.AppCountry) *AppCountryUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetCreatedAt sets the "created_at" field.
func (acu *AppCountryUpdate) SetCreatedAt(u uint32) *AppCountryUpdate {
	acu.mutation.ResetCreatedAt()
	acu.mutation.SetCreatedAt(u)
	return acu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acu *AppCountryUpdate) SetNillableCreatedAt(u *uint32) *AppCountryUpdate {
	if u != nil {
		acu.SetCreatedAt(*u)
	}
	return acu
}

// AddCreatedAt adds u to the "created_at" field.
func (acu *AppCountryUpdate) AddCreatedAt(u int32) *AppCountryUpdate {
	acu.mutation.AddCreatedAt(u)
	return acu
}

// SetUpdatedAt sets the "updated_at" field.
func (acu *AppCountryUpdate) SetUpdatedAt(u uint32) *AppCountryUpdate {
	acu.mutation.ResetUpdatedAt()
	acu.mutation.SetUpdatedAt(u)
	return acu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (acu *AppCountryUpdate) AddUpdatedAt(u int32) *AppCountryUpdate {
	acu.mutation.AddUpdatedAt(u)
	return acu
}

// SetDeletedAt sets the "deleted_at" field.
func (acu *AppCountryUpdate) SetDeletedAt(u uint32) *AppCountryUpdate {
	acu.mutation.ResetDeletedAt()
	acu.mutation.SetDeletedAt(u)
	return acu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acu *AppCountryUpdate) SetNillableDeletedAt(u *uint32) *AppCountryUpdate {
	if u != nil {
		acu.SetDeletedAt(*u)
	}
	return acu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (acu *AppCountryUpdate) AddDeletedAt(u int32) *AppCountryUpdate {
	acu.mutation.AddDeletedAt(u)
	return acu
}

// SetEntID sets the "ent_id" field.
func (acu *AppCountryUpdate) SetEntID(u uuid.UUID) *AppCountryUpdate {
	acu.mutation.SetEntID(u)
	return acu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (acu *AppCountryUpdate) SetNillableEntID(u *uuid.UUID) *AppCountryUpdate {
	if u != nil {
		acu.SetEntID(*u)
	}
	return acu
}

// SetAppID sets the "app_id" field.
func (acu *AppCountryUpdate) SetAppID(u uuid.UUID) *AppCountryUpdate {
	acu.mutation.SetAppID(u)
	return acu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (acu *AppCountryUpdate) SetNillableAppID(u *uuid.UUID) *AppCountryUpdate {
	if u != nil {
		acu.SetAppID(*u)
	}
	return acu
}

// ClearAppID clears the value of the "app_id" field.
func (acu *AppCountryUpdate) ClearAppID() *AppCountryUpdate {
	acu.mutation.ClearAppID()
	return acu
}

// SetCountryID sets the "country_id" field.
func (acu *AppCountryUpdate) SetCountryID(u uuid.UUID) *AppCountryUpdate {
	acu.mutation.SetCountryID(u)
	return acu
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (acu *AppCountryUpdate) SetNillableCountryID(u *uuid.UUID) *AppCountryUpdate {
	if u != nil {
		acu.SetCountryID(*u)
	}
	return acu
}

// ClearCountryID clears the value of the "country_id" field.
func (acu *AppCountryUpdate) ClearCountryID() *AppCountryUpdate {
	acu.mutation.ClearCountryID()
	return acu
}

// Mutation returns the AppCountryMutation object of the builder.
func (acu *AppCountryUpdate) Mutation() *AppCountryMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AppCountryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := acu.defaults(); err != nil {
		return 0, err
	}
	if len(acu.hooks) == 0 {
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			if acu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AppCountryUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AppCountryUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AppCountryUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acu *AppCountryUpdate) defaults() error {
	if _, ok := acu.mutation.UpdatedAt(); !ok {
		if appcountry.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcountry.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appcountry.UpdateDefaultUpdatedAt()
		acu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acu *AppCountryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppCountryUpdate {
	acu.modifiers = append(acu.modifiers, modifiers...)
	return acu
}

func (acu *AppCountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcountry.Table,
			Columns: appcountry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appcountry.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldCreatedAt,
		})
	}
	if value, ok := acu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldCreatedAt,
		})
	}
	if value, ok := acu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldUpdatedAt,
		})
	}
	if value, ok := acu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldUpdatedAt,
		})
	}
	if value, ok := acu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldDeletedAt,
		})
	}
	if value, ok := acu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldDeletedAt,
		})
	}
	if value, ok := acu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldEntID,
		})
	}
	if value, ok := acu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldAppID,
		})
	}
	if acu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appcountry.FieldAppID,
		})
	}
	if value, ok := acu.mutation.CountryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldCountryID,
		})
	}
	if acu.mutation.CountryIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appcountry.FieldCountryID,
		})
	}
	_spec.Modifiers = acu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcountry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppCountryUpdateOne is the builder for updating a single AppCountry entity.
type AppCountryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppCountryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (acuo *AppCountryUpdateOne) SetCreatedAt(u uint32) *AppCountryUpdateOne {
	acuo.mutation.ResetCreatedAt()
	acuo.mutation.SetCreatedAt(u)
	return acuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acuo *AppCountryUpdateOne) SetNillableCreatedAt(u *uint32) *AppCountryUpdateOne {
	if u != nil {
		acuo.SetCreatedAt(*u)
	}
	return acuo
}

// AddCreatedAt adds u to the "created_at" field.
func (acuo *AppCountryUpdateOne) AddCreatedAt(u int32) *AppCountryUpdateOne {
	acuo.mutation.AddCreatedAt(u)
	return acuo
}

// SetUpdatedAt sets the "updated_at" field.
func (acuo *AppCountryUpdateOne) SetUpdatedAt(u uint32) *AppCountryUpdateOne {
	acuo.mutation.ResetUpdatedAt()
	acuo.mutation.SetUpdatedAt(u)
	return acuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (acuo *AppCountryUpdateOne) AddUpdatedAt(u int32) *AppCountryUpdateOne {
	acuo.mutation.AddUpdatedAt(u)
	return acuo
}

// SetDeletedAt sets the "deleted_at" field.
func (acuo *AppCountryUpdateOne) SetDeletedAt(u uint32) *AppCountryUpdateOne {
	acuo.mutation.ResetDeletedAt()
	acuo.mutation.SetDeletedAt(u)
	return acuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acuo *AppCountryUpdateOne) SetNillableDeletedAt(u *uint32) *AppCountryUpdateOne {
	if u != nil {
		acuo.SetDeletedAt(*u)
	}
	return acuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (acuo *AppCountryUpdateOne) AddDeletedAt(u int32) *AppCountryUpdateOne {
	acuo.mutation.AddDeletedAt(u)
	return acuo
}

// SetEntID sets the "ent_id" field.
func (acuo *AppCountryUpdateOne) SetEntID(u uuid.UUID) *AppCountryUpdateOne {
	acuo.mutation.SetEntID(u)
	return acuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (acuo *AppCountryUpdateOne) SetNillableEntID(u *uuid.UUID) *AppCountryUpdateOne {
	if u != nil {
		acuo.SetEntID(*u)
	}
	return acuo
}

// SetAppID sets the "app_id" field.
func (acuo *AppCountryUpdateOne) SetAppID(u uuid.UUID) *AppCountryUpdateOne {
	acuo.mutation.SetAppID(u)
	return acuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (acuo *AppCountryUpdateOne) SetNillableAppID(u *uuid.UUID) *AppCountryUpdateOne {
	if u != nil {
		acuo.SetAppID(*u)
	}
	return acuo
}

// ClearAppID clears the value of the "app_id" field.
func (acuo *AppCountryUpdateOne) ClearAppID() *AppCountryUpdateOne {
	acuo.mutation.ClearAppID()
	return acuo
}

// SetCountryID sets the "country_id" field.
func (acuo *AppCountryUpdateOne) SetCountryID(u uuid.UUID) *AppCountryUpdateOne {
	acuo.mutation.SetCountryID(u)
	return acuo
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (acuo *AppCountryUpdateOne) SetNillableCountryID(u *uuid.UUID) *AppCountryUpdateOne {
	if u != nil {
		acuo.SetCountryID(*u)
	}
	return acuo
}

// ClearCountryID clears the value of the "country_id" field.
func (acuo *AppCountryUpdateOne) ClearCountryID() *AppCountryUpdateOne {
	acuo.mutation.ClearCountryID()
	return acuo
}

// Mutation returns the AppCountryMutation object of the builder.
func (acuo *AppCountryUpdateOne) Mutation() *AppCountryMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AppCountryUpdateOne) Select(field string, fields ...string) *AppCountryUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AppCountry entity.
func (acuo *AppCountryUpdateOne) Save(ctx context.Context) (*AppCountry, error) {
	var (
		err  error
		node *AppCountry
	)
	if err := acuo.defaults(); err != nil {
		return nil, err
	}
	if len(acuo.hooks) == 0 {
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			if acuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppCountry)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppCountryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AppCountryUpdateOne) SaveX(ctx context.Context) *AppCountry {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AppCountryUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AppCountryUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acuo *AppCountryUpdateOne) defaults() error {
	if _, ok := acuo.mutation.UpdatedAt(); !ok {
		if appcountry.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcountry.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appcountry.UpdateDefaultUpdatedAt()
		acuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acuo *AppCountryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppCountryUpdateOne {
	acuo.modifiers = append(acuo.modifiers, modifiers...)
	return acuo
}

func (acuo *AppCountryUpdateOne) sqlSave(ctx context.Context) (_node *AppCountry, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcountry.Table,
			Columns: appcountry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appcountry.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppCountry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcountry.FieldID)
		for _, f := range fields {
			if !appcountry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appcountry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldCreatedAt,
		})
	}
	if value, ok := acuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldCreatedAt,
		})
	}
	if value, ok := acuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldUpdatedAt,
		})
	}
	if value, ok := acuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldUpdatedAt,
		})
	}
	if value, ok := acuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldDeletedAt,
		})
	}
	if value, ok := acuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldDeletedAt,
		})
	}
	if value, ok := acuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldEntID,
		})
	}
	if value, ok := acuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldAppID,
		})
	}
	if acuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appcountry.FieldAppID,
		})
	}
	if value, ok := acuo.mutation.CountryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldCountryID,
		})
	}
	if acuo.mutation.CountryIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appcountry.FieldCountryID,
		})
	}
	_spec.Modifiers = acuo.modifiers
	_node = &AppCountry{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcountry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
