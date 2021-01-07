package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// DbBackedUser User struct.
type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt32  `validate:"required"`
}

// use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func main() {
	validate = validator.New()

	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc.
	validate.RegisterCustomTypeFunc(
		ValidateValuer,
		sql.NullString{},
		sql.NullInt32{},
		sql.NullBool{},
		sql.NullFloat64{},
	)

	// build object for validation.
	x := DbBackedUser{
		Name: sql.NullString{String: "", Valid: true},
		Age:  sql.NullInt32{Int32: 0, Valid: false},
	}
	if err := validate.Struct(x); err != nil {
		fmt.Printf("errors:\n%+v\n", err)
	}
}

// ValidateValuer implements validator.CustomTypeFunc.
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err != nil {
			return val
		}
		// handle the error how you want.
	}
	return nil
}
