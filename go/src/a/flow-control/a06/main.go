package main

func main() {
	a, b := 10, 11
	max := If(a > b, a, b).(int)
	println(max)
}

// If condition is true, return tv, otherwise return fv.
func If(condition bool, tv, fv interface{}) interface{} {
	if condition {
		return tv
	}
	return fv
}
