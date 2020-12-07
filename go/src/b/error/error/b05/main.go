package main

type Cat struct{}
type Duck interface {
	Quack()
}

func (c Cat) Quack()  {} // 使用结构体实现接口
// ./main.go:11:5: cannot use Cat literal (type Cat) as type Duck in assignment:
//        Cat does not implement Duck (Quack method has pointer receiver)
// func (c *Cat) Quack() {} // 使用结构体指针实现接口

var d1 Duck = Cat{}  // 使用结构体初始化变量
var d2 Duck = &Cat{} // 使用结构体指针初始化变量

func main() {
}
