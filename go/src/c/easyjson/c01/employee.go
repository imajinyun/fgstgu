package c01

// Profile struct.
type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Job struct.
type Job struct {
	Skills []string `json:"skills"`
}

// Employee struct.
type Employee struct {
	Profile Profile `json:"profile"`
	Job     Job     `json:"job"`
}
