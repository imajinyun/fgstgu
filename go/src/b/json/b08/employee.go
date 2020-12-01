package b08

/**
 * 1. export GOPATH=$HOME/Workspace/fgstgu/go:$GOPATH
 * 2. cd ~/path/to/go && go get -u github.com/mailru/easyjson/...
 * 3. go build -o bin/easyjson github.com/mailru/easyjson/easyjson
 * 4. cd ~/path/to/workspace && ~/path/to/go/bin/easyjson -all src/c/easyjson/c01/employee.go
 * 5. echo $GOPATH -> /Users/username/Workspace/fgstgu/go:/Users/username/Codes/go
 */

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
