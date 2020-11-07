package main

import (
	"encoding/json"
	"fmt"
)

// DebugInfo struct.
type DebugInfo struct {
	Level   string
	Message string
	file    string
}

func main() {
	data := `
	[
		{"Level":"debug","Message":"Page not found","file":"a.go"},
		{"Level":"error","Message":"An error occurred","file":"b.go"}
	]
	`
	fmt.Println("The JSON will be parsed as a map:")
	info := []map[string]string{}
	json.Unmarshal([]byte(data), &info)
	fmt.Println(info)

	fmt.Println()

	fmt.Println("Fields not exported will not be parsed by JSON:")
	debug := []DebugInfo{}
	json.Unmarshal([]byte(data), &debug)
	fmt.Println(debug)
}

func (d DebugInfo) String() string {
	return fmt.Sprintf("{Level: %s, Message: %s}", d.Level, d.Message)
}
