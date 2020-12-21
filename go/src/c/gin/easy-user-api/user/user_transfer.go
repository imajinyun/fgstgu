package user

// UserTransfer struct.
type UserTransfer struct {
	ID   int64  `json:"id,string,omitempty"`
	Name string `json:"name"`
	Age  uint8  `json:"age,string"`
}
