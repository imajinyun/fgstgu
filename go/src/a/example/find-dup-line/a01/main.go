/*
格式化输出占位符：
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true 或 false
%c          字符（rune），Unicode 码点
%s          字符串
%q          带双引号的字符串 "abc" 或带单引号的字符 'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "EOF" {
			fmt.Printf("✔️ your input `EOF` quit\n")
			break
		}
		fmt.Fprintf(os.Stdout, "✔️ your input: "+input.Text()+"\n")
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	fmt.Printf("✔️ your input count result:\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
