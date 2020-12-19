package main

import (
	"math/rand"
	"strings"

	"github.com/pkg/profile"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// -> go run main.go
// -> go tool pprof -http=:9999 /path/to/...
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// concat(100)
	concat2(100)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += random(n)
	}
	return s
}

func concat2(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(random(n))
	}
	return s.String()
}

func random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
