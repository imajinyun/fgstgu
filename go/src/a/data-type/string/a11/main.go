package main

import (
	"bytes"
	"fmt"
)

// 遍历可变参数列表——获取每一个参数的值
func main() {
	fmt.Println(joinStrings("Apple", "Orangle", "Pear"))
	fmt.Println(joinStrings("Tom", "Jack", "Jhon"))
}

func joinStrings(list ...string) string {
	var b bytes.Buffer
	for _, s := range list {
		b.WriteString(s)
	}
	return b.String()
}
