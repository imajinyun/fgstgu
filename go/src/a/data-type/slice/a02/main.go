package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	fmt.Printf("cap(arr) = %d\n", cap(arr))

	s1 := arr[2:10]
	fmt.Printf("s1 type: %T, cap(s1) = %d\n", s1, cap(s1))
	fmt.Println("s1 =", s1)

	s2 := s1[2:8]
	fmt.Printf("s2 type: %T, cap(s2) = %d\n", s2, cap(s2))
	fmt.Println("s2 =", s2)

	s3 := s1[3:8]
	fmt.Printf("s3 type: %T, cap(s3) = %d\n", s3, cap(s3))
	fmt.Println("s3 =", s3)

	s4 := s2[3:8]
	fmt.Printf("s4 type: %T, cap(s2) = %d\n", s4, cap(s4))
	fmt.Println("s4 =", s4)
	fmt.Println()

	s4[0] = "Go"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println()

	s4 = append(s4, "F4")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
