package d04

import "testing"

func TestCreateRequest(t *testing.T) {
	str := createRequest()
	t.Log(str)
}

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	reps := processRequestWithBuiltinJSON(reqs)
	t.Log(reps[0])
}

func BenchmarkProcessRequestWithEasyJSON(b *testing.B) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestWithEasyJSON(reqs)
	}
	b.StopTimer()
}

func BenchmarkProcessRequestWithBuiltinJSON(b *testing.B) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestWithBuiltinJSON(reqs)
	}
	b.StopTimer()
}
