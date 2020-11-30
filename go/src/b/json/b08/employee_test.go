package b08

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
