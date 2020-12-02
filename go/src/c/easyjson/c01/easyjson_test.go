package c01

/**
 * 1. export GOPATH=$HOME/Workspace/fgstgu/go:$GOPATH
 * 2. cd ~/path/to/go && go get -u github.com/mailru/easyjson/...
 * 3. go build -o bin/easyjson github.com/mailru/easyjson/easyjson
 * 4. cd ~/path/to/workspace && ~/path/to/go/bin/easyjson -all src/c/easyjson/c01/employee.go
 * 5. echo $GOPATH -> /Users/username/Workspace/fgstgu/go:/Users/username/Codes/go
 */

import (
	"encoding/json"
	"testing"
)

var jsons = `{
	"profile": {
	  	"name": "Jack",
		"age": 40
	},
	"job": {
		"skills": ["Java", "Go", "C", "PHP"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsons), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*e)
	if v, err := json.Marshal(e); err != nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsons))
	t.Log(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		t.Log(string(v))
	}
}

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsons), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsons))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
