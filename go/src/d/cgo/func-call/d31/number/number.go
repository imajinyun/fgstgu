package number

import "C"

//export average
func average(a, b C.int) C.int {
	return (a + b) / 2
}
