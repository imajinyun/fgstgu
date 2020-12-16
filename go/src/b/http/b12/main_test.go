package main

import "testing"

// Student struct.
type Student struct {
	Name string
	Age  int
}

// Subscriber struct.
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

var executeTemplateTests = []struct {
	tmpl string
	text interface{}
}{
	{"🎉 Hello World: {{.}}\n", "This is a test"},
	{"🔥 Action: {{.}}\n", 1024},
	{"🐶 Hello Go: {{.}}\n", true},
	{"❯ start {{if .}}Dot is true!{{end}} finished\n", true},
	{"❯ start {{if .}}Dot is true!{{end}} finished\n", false},
	{"❯ Name: {{.Name}}, Age: {{.Age}}\n", Student{Name: "Jack", Age: 88}},
	{"❯ Name: {{.Name}}, {{if .Active}}Rate: ${{.Rate}}{{end}}\n", Subscriber{"Jhon Burg", 4.99, true}},
	{"❯ Name: {{.Name}}, {{if .Active}}Rate: ${{.Rate}}{{end}}\n", Subscriber{"Jhon Burg", 4.99, false}},
	{"❯ Before loop: {{.}}\n{{range .}}  In loop: {{.}}\n{{end}}❯ After loop: {{.}}\n", []string{"a", "b", "c", "d"}},
	{"❯ Prices: \n{{range .}}${{.}}\n{{end}}", []float64{}},
	{"❯ Prices: \n{{range .}}${{.}}\n{{end}}", nil},
}

func TestExecuteTemplate(t *testing.T) {
	for _, v := range executeTemplateTests {
		executeTemplate(v.tmpl, v.text)
	}
}
