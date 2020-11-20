package b12

import (
	"reflect"
	"testing"
)

/**
 * 结论：
 * 1. 能使用原生代码时，尽量避免反射操作；
 * 2. 提前缓冲反射值对象，对性能有很大的帮助；
 * 3. 避免反射函数调用，实在需要调用时，先提前缓冲函数参数列表，并且尽量少地使用返回值；
 */

type data struct {
	Value int
}

func foo(x int)     {}
func bar(x int) int { return x }

func BenchmarkNativeAssign(b *testing.B) {
	v := data{Value: 200}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Value = 800
	}
}

func BenchmarkReflectAssign(b *testing.B) {
	v := data{Value: 200}
	e := reflect.ValueOf(&v).Elem()
	f := e.FieldByName("Value")
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f.SetInt(400)
	}
}

func BenchmarkReflectFindFieldAndAssign(b *testing.B) {
	v := data{Value: 200}
	e := reflect.ValueOf(&v).Elem()
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.FieldByName("Value").SetInt(300)
	}
}

func BenchmarkNativeCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(100)
	}
}

func BenchmarkReflectCall(b *testing.B) {
	v := reflect.ValueOf(foo)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Call([]reflect.Value{reflect.ValueOf(400)})
	}
}
