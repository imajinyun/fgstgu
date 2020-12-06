// -> go test -bench=. -cpuprofile=cpu.prof
// -> go tool pprof cpu.prof
// -> go test -bench=. -memprofile=mem.prof
// -> go tool pprof mem.prof

package d04

// Request struct.
type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad       []int  `json:"payload"`
}

// Response struct.
type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
