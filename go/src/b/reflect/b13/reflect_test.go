package b13

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 100
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestBasicType(t *testing.T) {
	var a int = 10
	var b int16 = 100
	var c int32 = 1000
	var d int64 = 9999
	var f1 float32 = 10.00
	var f2 float64 = 1000.9999
	var b1, b2 bool = false, true
	var array = [3]int{}
	var slice = []int{}
	t.Log(CheckType(a))
	t.Log(CheckType(b))
	t.Log(CheckType(c))
	t.Log(CheckType(d))
	t.Log(CheckType(f1))
	t.Log(CheckType(f2))
	t.Log(CheckType(b1))
	t.Log(CheckType(b2))
	t.Log(CheckType(array))
	t.Log(CheckType(slice))
	t.Log(CheckType(struct{}{}))
	t.Log(CheckType(make(map[string]int)))
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log(reflect.DeepEqual(s1, s2))
	t.Log(reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log(c1 == c2)
	t.Log(reflect.DeepEqual(c1, c2))
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field")
	} else {
		t.Log("Tag -> format:", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).
		MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(uint8(18))})
	t.Log("Updated Age:", e)
}

type Customer struct {
	CookieID string
	Name     string
	Age      uint8
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        uint8
}

func (e *Employee) UpdateAge(age uint8) {
	e.Age = age
}

func CheckType(v interface{}) (s string) {
	t := reflect.TypeOf(v)
	fmt.Print(t.Kind())
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		s = "Float"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s = "Integer"
	case reflect.Bool:
		s = "Boolean"
	case reflect.Array:
		s = "Array"
	case reflect.Slice:
		s = "Slice"
	default:
		s = "Unknown"
	}
	return
}
