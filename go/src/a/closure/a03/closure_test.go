package a03

import "testing"

var v int = 10

func TestInc(t *testing.T) {
	// 10
	t.Log(v)
	v = inc()

	// 101
	t.Log(v)
}

func TestClosureReferenceProblem(t *testing.T) {
	// 闭包的这种以引用方式访问外部变量的行为可能会导致一些隐含的问题
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

func TestClosureReferenceFixed1(t *testing.T) {
	// 在循环体内部再定义一个局部变量，这样每次迭代 defer 语句的闭包函数捕获的都是不同的变量，这些变量的值对应迭代时的值
	for i := 0; i < 3; i++ {
		i := i
		defer func() { println(i) }()
	}
}

func TestClosureReferenceFixed2(t *testing.T) {
	// 将迭代变量通过闭包函数的参数传入，defer 语句会马上对调用参数求值
	for i := 0; i < 3; i++ {
		defer func(i int) { println(i) }(i)
	}
}

func inc() (v int) {
	defer func() { v++ }()
	return 100
}
