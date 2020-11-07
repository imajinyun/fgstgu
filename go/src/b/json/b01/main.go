package main

import (
	"encoding/json"
	"fmt"
)

const indentString = "  "

func main() {
	m := map[string][]string{
		"level":   {"debug"},
		"message": {"Apple", "Banana", "Cherry", "Lemon"},
	}
	fmt.Println("No formatting:")
	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
	}

	fmt.Println("\nFormatting:")
	if data, err := json.MarshalIndent(m, "", indentString); err == nil {
		fmt.Printf("%s\n", data)
	}
}
