// Code generated by ent, DO NOT EDIT.

package mvpstaff

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldName, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldRole, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldPhone, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldLastName, v))
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldBirthday, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContainsFold(FieldName, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContainsFold(FieldRole, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContainsFold(FieldPhone, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldContainsFold(FieldLastName, v))
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldEQ(FieldBirthday, v))
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNEQ(FieldBirthday, v))
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIn(FieldBirthday, vs...))
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotIn(FieldBirthday, vs...))
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGT(FieldBirthday, v))
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldGTE(FieldBirthday, v))
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLT(FieldBirthday, v))
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v time.Time) predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldLTE(FieldBirthday, v))
}

// BirthdayIsNil applies the IsNil predicate on the "birthday" field.
func BirthdayIsNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldIsNull(FieldBirthday))
}

// BirthdayNotNil applies the NotNil predicate on the "birthday" field.
func BirthdayNotNil() predicate.MvpStaff {
	return predicate.MvpStaff(sql.FieldNotNull(FieldBirthday))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MvpStaff) predicate.MvpStaff {
	return predicate.MvpStaff(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MvpStaff) predicate.MvpStaff {
	return predicate.MvpStaff(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MvpStaff) predicate.MvpStaff {
	return predicate.MvpStaff(sql.NotPredicates(p))
}