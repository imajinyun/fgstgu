package main

/*
struct A {
	// a: type 是 Go 语言的关键字
	int type;

	// b: 这种情况下，Go 语言关键字命名的成员将无法访问（被屏蔽）
	float _type;
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A

	// a 情况下：_type 对应结构体 A 的 type 字段
	// b 情况下：_type 对应结构体 A 的 _type 字段
	fmt.Println(a._type)
}
