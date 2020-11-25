package b14

import (
	"errors"
	"reflect"
	"testing"
)

func TestFillUserInfo(t *testing.T) {
	settings := map[string]interface{}{"Name": "Tom", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Logf("%p, %T, %v", &e, e, e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Logf("%p, %T, %v", c, *c, *c)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func fillBySettings(p interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(p).Kind() != reflect.Ptr {
		return errors.New("The first param should be a pointer type")
	}
	if reflect.TypeOf(p).Elem().Kind() != reflect.Struct {
		return errors.New("The first param should be a struct type")
	}
	if settings == nil {
		return errors.New("Setting is nil")
	}
	for k, v := range settings {
		field, ok := (reflect.ValueOf(p)).Elem().Type().FieldByName(k)
		if !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			value := reflect.ValueOf(p)
			elem := value.Elem()
			elem.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
