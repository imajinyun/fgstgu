package main

import "fmt"

func main() {
	fmt.Println("keywords:", keywords())
	fmt.Println("keywords count:", len(keywords()))
	fmt.Println("built-in const:", []string{"true", "false", "iota", "nil"})
	fmt.Println("built-in type:", builtInTypes())
	fmt.Println("built-in func:", builtInFuncs())
}

func keywords() []string {
	return []string{
		"break", "default", "func", "interface", "select",
		"case", " defer", "go", "map", "struct",
		"chan", " else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}
}

func builtInTypes() []string {
	return []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"complex64", "complex128",
		"bool", "byte", "rune", "string", "error",
	}
}

func builtInFuncs() []string {
	return []string{
		"make", "len", "cap", "new", "append",
		"copy", "close", "delete", "complex",
		"real", "imag", "panic", "recover",
	}
}
