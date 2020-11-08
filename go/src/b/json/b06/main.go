package main

import (
	"encoding/json"
	"fmt"
)

// DebugInfo struct.
type DebugInfo struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Author  string `json:"-"`
}

func main() {
	data := `
	[
		{"Level":"debug","Message":"Page not found","Author":"Jack"},
		{"Level":"error","Message":"An error occurred","Author":"Jhon"}
	]
	`
	info := []DebugInfo{}
	json.Unmarshal([]byte(data), &info)
	fmt.Println(info)
}

func (d DebugInfo) String() string {
	return fmt.Sprintf("{Level: %s, Message: %s}\n", d.Level, d.Message)
}
