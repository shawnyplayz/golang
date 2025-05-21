package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New();
func ValidationHelper(models any) []string{
errors := []string{}
if err := validate.Struct(models); err !=nil{
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("Field '%s' failed on the '%s' tag", err.Field(), err.Tag()))
	}
}
return errors;
}