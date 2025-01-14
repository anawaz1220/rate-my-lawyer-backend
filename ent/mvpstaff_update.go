// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/mvpstaff"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// MvpStaffUpdate is the builder for updating MvpStaff entities.
type MvpStaffUpdate struct {
	config
	hooks    []Hook
	mutation *MvpStaffMutation
}

// Where appends a list predicates to the MvpStaffUpdate builder.
func (msu *MvpStaffUpdate) Where(ps ...predicate.MvpStaff) *MvpStaffUpdate {
	msu.mutation.Where(ps...)
	return msu
}

// SetName sets the "name" field.
func (msu *MvpStaffUpdate) SetName(s string) *MvpStaffUpdate {
	msu.mutation.SetName(s)
	return msu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillableName(s *string) *MvpStaffUpdate {
	if s != nil {
		msu.SetName(*s)
	}
	return msu
}

// SetRole sets the "role" field.
func (msu *MvpStaffUpdate) SetRole(s string) *MvpStaffUpdate {
	msu.mutation.SetRole(s)
	return msu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillableRole(s *string) *MvpStaffUpdate {
	if s != nil {
		msu.SetRole(*s)
	}
	return msu
}

// SetEmail sets the "email" field.
func (msu *MvpStaffUpdate) SetEmail(s string) *MvpStaffUpdate {
	msu.mutation.SetEmail(s)
	return msu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillableEmail(s *string) *MvpStaffUpdate {
	if s != nil {
		msu.SetEmail(*s)
	}
	return msu
}

// ClearEmail clears the value of the "email" field.
func (msu *MvpStaffUpdate) ClearEmail() *MvpStaffUpdate {
	msu.mutation.ClearEmail()
	return msu
}

// SetPhone sets the "phone" field.
func (msu *MvpStaffUpdate) SetPhone(s string) *MvpStaffUpdate {
	msu.mutation.SetPhone(s)
	return msu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillablePhone(s *string) *MvpStaffUpdate {
	if s != nil {
		msu.SetPhone(*s)
	}
	return msu
}

// ClearPhone clears the value of the "phone" field.
func (msu *MvpStaffUpdate) ClearPhone() *MvpStaffUpdate {
	msu.mutation.ClearPhone()
	return msu
}

// SetLastName sets the "last_name" field.
func (msu *MvpStaffUpdate) SetLastName(s string) *MvpStaffUpdate {
	msu.mutation.SetLastName(s)
	return msu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillableLastName(s *string) *MvpStaffUpdate {
	if s != nil {
		msu.SetLastName(*s)
	}
	return msu
}

// SetBirthday sets the "birthday" field.
func (msu *MvpStaffUpdate) SetBirthday(t time.Time) *MvpStaffUpdate {
	msu.mutation.SetBirthday(t)
	return msu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (msu *MvpStaffUpdate) SetNillableBirthday(t *time.Time) *MvpStaffUpdate {
	if t != nil {
		msu.SetBirthday(*t)
	}
	return msu
}

// ClearBirthday clears the value of the "birthday" field.
func (msu *MvpStaffUpdate) ClearBirthday() *MvpStaffUpdate {
	msu.mutation.ClearBirthday()
	return msu
}

// Mutation returns the MvpStaffMutation object of the builder.
func (msu *MvpStaffUpdate) Mutation() *MvpStaffMutation {
	return msu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (msu *MvpStaffUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, msu.sqlSave, msu.mutation, msu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msu *MvpStaffUpdate) SaveX(ctx context.Context) int {
	affected, err := msu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (msu *MvpStaffUpdate) Exec(ctx context.Context) error {
	_, err := msu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msu *MvpStaffUpdate) ExecX(ctx context.Context) {
	if err := msu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msu *MvpStaffUpdate) check() error {
	if v, ok := msu.mutation.Role(); ok {
		if err := mvpstaff.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "MvpStaff.role": %w`, err)}
		}
	}
	return nil
}

func (msu *MvpStaffUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := msu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mvpstaff.Table, mvpstaff.Columns, sqlgraph.NewFieldSpec(mvpstaff.FieldID, field.TypeInt))
	if ps := msu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msu.mutation.Name(); ok {
		_spec.SetField(mvpstaff.FieldName, field.TypeString, value)
	}
	if value, ok := msu.mutation.Role(); ok {
		_spec.SetField(mvpstaff.FieldRole, field.TypeString, value)
	}
	if value, ok := msu.mutation.Email(); ok {
		_spec.SetField(mvpstaff.FieldEmail, field.TypeString, value)
	}
	if msu.mutation.EmailCleared() {
		_spec.ClearField(mvpstaff.FieldEmail, field.TypeString)
	}
	if value, ok := msu.mutation.Phone(); ok {
		_spec.SetField(mvpstaff.FieldPhone, field.TypeString, value)
	}
	if msu.mutation.PhoneCleared() {
		_spec.ClearField(mvpstaff.FieldPhone, field.TypeString)
	}
	if value, ok := msu.mutation.LastName(); ok {
		_spec.SetField(mvpstaff.FieldLastName, field.TypeString, value)
	}
	if value, ok := msu.mutation.Birthday(); ok {
		_spec.SetField(mvpstaff.FieldBirthday, field.TypeTime, value)
	}
	if msu.mutation.BirthdayCleared() {
		_spec.ClearField(mvpstaff.FieldBirthday, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, msu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mvpstaff.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	msu.mutation.done = true
	return n, nil
}

// MvpStaffUpdateOne is the builder for updating a single MvpStaff entity.
type MvpStaffUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MvpStaffMutation
}

// SetName sets the "name" field.
func (msuo *MvpStaffUpdateOne) SetName(s string) *MvpStaffUpdateOne {
	msuo.mutation.SetName(s)
	return msuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillableName(s *string) *MvpStaffUpdateOne {
	if s != nil {
		msuo.SetName(*s)
	}
	return msuo
}

// SetRole sets the "role" field.
func (msuo *MvpStaffUpdateOne) SetRole(s string) *MvpStaffUpdateOne {
	msuo.mutation.SetRole(s)
	return msuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillableRole(s *string) *MvpStaffUpdateOne {
	if s != nil {
		msuo.SetRole(*s)
	}
	return msuo
}

// SetEmail sets the "email" field.
func (msuo *MvpStaffUpdateOne) SetEmail(s string) *MvpStaffUpdateOne {
	msuo.mutation.SetEmail(s)
	return msuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillableEmail(s *string) *MvpStaffUpdateOne {
	if s != nil {
		msuo.SetEmail(*s)
	}
	return msuo
}

// ClearEmail clears the value of the "email" field.
func (msuo *MvpStaffUpdateOne) ClearEmail() *MvpStaffUpdateOne {
	msuo.mutation.ClearEmail()
	return msuo
}

// SetPhone sets the "phone" field.
func (msuo *MvpStaffUpdateOne) SetPhone(s string) *MvpStaffUpdateOne {
	msuo.mutation.SetPhone(s)
	return msuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillablePhone(s *string) *MvpStaffUpdateOne {
	if s != nil {
		msuo.SetPhone(*s)
	}
	return msuo
}

// ClearPhone clears the value of the "phone" field.
func (msuo *MvpStaffUpdateOne) ClearPhone() *MvpStaffUpdateOne {
	msuo.mutation.ClearPhone()
	return msuo
}

// SetLastName sets the "last_name" field.
func (msuo *MvpStaffUpdateOne) SetLastName(s string) *MvpStaffUpdateOne {
	msuo.mutation.SetLastName(s)
	return msuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillableLastName(s *string) *MvpStaffUpdateOne {
	if s != nil {
		msuo.SetLastName(*s)
	}
	return msuo
}

// SetBirthday sets the "birthday" field.
func (msuo *MvpStaffUpdateOne) SetBirthday(t time.Time) *MvpStaffUpdateOne {
	msuo.mutation.SetBirthday(t)
	return msuo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (msuo *MvpStaffUpdateOne) SetNillableBirthday(t *time.Time) *MvpStaffUpdateOne {
	if t != nil {
		msuo.SetBirthday(*t)
	}
	return msuo
}

// ClearBirthday clears the value of the "birthday" field.
func (msuo *MvpStaffUpdateOne) ClearBirthday() *MvpStaffUpdateOne {
	msuo.mutation.ClearBirthday()
	return msuo
}

// Mutation returns the MvpStaffMutation object of the builder.
func (msuo *MvpStaffUpdateOne) Mutation() *MvpStaffMutation {
	return msuo.mutation
}

// Where appends a list predicates to the MvpStaffUpdate builder.
func (msuo *MvpStaffUpdateOne) Where(ps ...predicate.MvpStaff) *MvpStaffUpdateOne {
	msuo.mutation.Where(ps...)
	return msuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (msuo *MvpStaffUpdateOne) Select(field string, fields ...string) *MvpStaffUpdateOne {
	msuo.fields = append([]string{field}, fields...)
	return msuo
}

// Save executes the query and returns the updated MvpStaff entity.
func (msuo *MvpStaffUpdateOne) Save(ctx context.Context) (*MvpStaff, error) {
	return withHooks(ctx, msuo.sqlSave, msuo.mutation, msuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msuo *MvpStaffUpdateOne) SaveX(ctx context.Context) *MvpStaff {
	node, err := msuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (msuo *MvpStaffUpdateOne) Exec(ctx context.Context) error {
	_, err := msuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msuo *MvpStaffUpdateOne) ExecX(ctx context.Context) {
	if err := msuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msuo *MvpStaffUpdateOne) check() error {
	if v, ok := msuo.mutation.Role(); ok {
		if err := mvpstaff.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "MvpStaff.role": %w`, err)}
		}
	}
	return nil
}

func (msuo *MvpStaffUpdateOne) sqlSave(ctx context.Context) (_node *MvpStaff, err error) {
	if err := msuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mvpstaff.Table, mvpstaff.Columns, sqlgraph.NewFieldSpec(mvpstaff.FieldID, field.TypeInt))
	id, ok := msuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MvpStaff.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := msuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mvpstaff.FieldID)
		for _, f := range fields {
			if !mvpstaff.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mvpstaff.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := msuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msuo.mutation.Name(); ok {
		_spec.SetField(mvpstaff.FieldName, field.TypeString, value)
	}
	if value, ok := msuo.mutation.Role(); ok {
		_spec.SetField(mvpstaff.FieldRole, field.TypeString, value)
	}
	if value, ok := msuo.mutation.Email(); ok {
		_spec.SetField(mvpstaff.FieldEmail, field.TypeString, value)
	}
	if msuo.mutation.EmailCleared() {
		_spec.ClearField(mvpstaff.FieldEmail, field.TypeString)
	}
	if value, ok := msuo.mutation.Phone(); ok {
		_spec.SetField(mvpstaff.FieldPhone, field.TypeString, value)
	}
	if msuo.mutation.PhoneCleared() {
		_spec.ClearField(mvpstaff.FieldPhone, field.TypeString)
	}
	if value, ok := msuo.mutation.LastName(); ok {
		_spec.SetField(mvpstaff.FieldLastName, field.TypeString, value)
	}
	if value, ok := msuo.mutation.Birthday(); ok {
		_spec.SetField(mvpstaff.FieldBirthday, field.TypeTime, value)
	}
	if msuo.mutation.BirthdayCleared() {
		_spec.ClearField(mvpstaff.FieldBirthday, field.TypeTime)
	}
	_node = &MvpStaff{config: msuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, msuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mvpstaff.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	msuo.mutation.done = true
	return _node, nil
}
