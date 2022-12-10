package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validatorFns map[string]validator.Func

func init() {
	validatorFns = make(map[string]validator.Func)
	// register validators here
	validatorFns["validatePhone"] = validatePhone
}

func addValidatorFn(customName string, fn validator.Func) {

	validatorFns[customName] = fn

}

func validatePhone(f1 validator.FieldLevel) bool {

	re := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)
	matches := re.FindAllString(f1.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func SignUpValidator(user interface{}) (err error) {

	validate := validator.New()

	registerCustomValidators(validate, []string{"validatePhone"})

	err = validate.Struct(user)

	return err

}

func loginValidator(user interface{}) (err error) {

	validate := validator.New()
	// register custom validator
	registerCustomValidators(validate, []string{"validatePhone"})
	return validateError(validate, user)

}

func validateError(validate *validator.Validate, entity interface{}) (err error) {

	err = validate.Struct(entity)

	validationErr := err.(validator.ValidationErrors)

	return validationErr

}

func registerCustomValidators(validate *validator.Validate, customValidators []string) {

	for _, fnName := range customValidators {
		validate.RegisterValidation(fnName, validatorFns[fnName])
	}

}
