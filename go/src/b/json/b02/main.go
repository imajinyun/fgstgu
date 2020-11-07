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
	line    int
}

func main() {
	info := []DebugInfo{
		{"debug", `File: Page not found`, "/path/to/src/index/page.html", 60},
		{"error", `Panic: An error occurred`, "/path/to/src/acme/util.go", 123},
	}

	if data, err := json.Marshal(info); err == nil {
		fmt.Printf("%s\n", data)
	}

	if data, err := json.MarshalIndent(info, "", "  "); err == nil {
		fmt.Printf("%s\n", data)
	}
}
