package main

import (
	"fmt"
	"reflect"
)

// 使用反射获取结构体的成员类型
func main() {
	type cat struct {
		Name string
		Type int `type:"info" id:"100" message:"Hello World!"`
	}
	ins := cat{"Hello", 100}
	typeOfCat := reflect.TypeOf(ins)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Printf("Name: %v, Tag: %v\n", fieldType.Name, fieldType.Tag)
	}

	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(
			catType.Tag.Get("type"),
			catType.Tag.Get("id"),
			catType.Tag.Get("message"))
	}
}
