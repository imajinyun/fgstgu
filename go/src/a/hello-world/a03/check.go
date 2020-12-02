package a03

// -> GOSSAFUNC=outOfRange go build src/a/hello-world/a03/check.go
func outOfRange() int {
	arr := [3]int{1, 2, 3}
	i := 4
	elem := arr[i]
	return elem
}
