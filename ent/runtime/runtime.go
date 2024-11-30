// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/mikestefanello/pagoda/ent/decision"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/mvpplannedroute"
	"github.com/mikestefanello/pagoda/ent/mvpstaff"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
	"github.com/mikestefanello/pagoda/ent/schema"
	"github.com/mikestefanello/pagoda/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	decisionFields := schema.Decision{}.Fields()
	_ = decisionFields
	// decisionDescURL is the schema descriptor for url field.
	decisionDescURL := decisionFields[0].Descriptor()
	// decision.URLValidator is a validator for the "url" field. It is called by the builders before save.
	decision.URLValidator = decisionDescURL.Validators[0].(func(string) error)
	lawyerFields := schema.Lawyer{}.Fields()
	_ = lawyerFields
	// lawyerDescFullName is the schema descriptor for full_name field.
	lawyerDescFullName := lawyerFields[0].Descriptor()
	// lawyer.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	lawyer.FullNameValidator = lawyerDescFullName.Validators[0].(func(string) error)
	// lawyerDescAverageRating is the schema descriptor for average_rating field.
	lawyerDescAverageRating := lawyerFields[9].Descriptor()
	// lawyer.DefaultAverageRating holds the default value on creation for the average_rating field.
	lawyer.DefaultAverageRating = lawyerDescAverageRating.Default.(float64)
	// lawyerDescReviewCount is the schema descriptor for review_count field.
	lawyerDescReviewCount := lawyerFields[10].Descriptor()
	// lawyer.DefaultReviewCount holds the default value on creation for the review_count field.
	lawyer.DefaultReviewCount = lawyerDescReviewCount.Default.(int)
	mvpplannedrouteFields := schema.MvpPlannedRoute{}.Fields()
	_ = mvpplannedrouteFields
	// mvpplannedrouteDescStatus is the schema descriptor for status field.
	mvpplannedrouteDescStatus := mvpplannedrouteFields[2].Descriptor()
	// mvpplannedroute.DefaultStatus holds the default value on creation for the status field.
	mvpplannedroute.DefaultStatus = mvpplannedrouteDescStatus.Default.(string)
	// mvpplannedroute.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	mvpplannedroute.StatusValidator = mvpplannedrouteDescStatus.Validators[0].(func(string) error)
	mvpstaffFields := schema.MvpStaff{}.Fields()
	_ = mvpstaffFields
	// mvpstaffDescRole is the schema descriptor for role field.
	mvpstaffDescRole := mvpstaffFields[1].Descriptor()
	// mvpstaff.RoleValidator is a validator for the "role" field. It is called by the builders before save.
	mvpstaff.RoleValidator = mvpstaffDescRole.Validators[0].(func(string) error)
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	rmluserFields := schema.RMLUser{}.Fields()
	_ = rmluserFields
	// rmluserDescName is the schema descriptor for name field.
	rmluserDescName := rmluserFields[0].Descriptor()
	// rmluser.NameValidator is a validator for the "name" field. It is called by the builders before save.
	rmluser.NameValidator = rmluserDescName.Validators[0].(func(string) error)
	// rmluserDescEmail is the schema descriptor for email field.
	rmluserDescEmail := rmluserFields[1].Descriptor()
	// rmluser.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	rmluser.EmailValidator = rmluserDescEmail.Validators[0].(func(string) error)
	// rmluserDescPassword is the schema descriptor for password field.
	rmluserDescPassword := rmluserFields[2].Descriptor()
	// rmluser.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	rmluser.PasswordValidator = rmluserDescPassword.Validators[0].(func(string) error)
	// rmluserDescRole is the schema descriptor for role field.
	rmluserDescRole := rmluserFields[3].Descriptor()
	// rmluser.DefaultRole holds the default value on creation for the role field.
	rmluser.DefaultRole = rmluserDescRole.Default.(string)
	// rmluserDescReviewCount is the schema descriptor for review_count field.
	rmluserDescReviewCount := rmluserFields[4].Descriptor()
	// rmluser.DefaultReviewCount holds the default value on creation for the review_count field.
	rmluser.DefaultReviewCount = rmluserDescReviewCount.Default.(int)
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescRating is the schema descriptor for rating field.
	reviewDescRating := reviewFields[0].Descriptor()
	// review.RatingValidator is a validator for the "rating" field. It is called by the builders before save.
	review.RatingValidator = func() func(int) error {
		validators := reviewDescRating.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(rating int) error {
			for _, fn := range fns {
				if err := fn(rating); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reviewDescCreatedAt is the schema descriptor for created_at field.
	reviewDescCreatedAt := reviewFields[2].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the created_at field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[3].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)