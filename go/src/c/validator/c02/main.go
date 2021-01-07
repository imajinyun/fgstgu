package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// MyStruct struct.
type MyStruct struct {
	String string `validate:"is-awesome"`
}

// use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyValue)

	s := MyStruct{String: "awesome"}
	if err := validate.Struct(s); err != nil {
		fmt.Printf("errors:\n%+v\n", err)
	}
	s.String = "not awesome"
	if err := validate.Struct(s); err != nil {
		fmt.Printf("errors:\n%+v\n", err)
	}
}

// ValidateMyValue implements validator.Func.
func ValidateMyValue(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
