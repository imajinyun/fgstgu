package main

import "fmt"

func main() {
	numbers := make([]int, 0, 20)
	dump("numbers", numbers)

	numbers = append(numbers, 1)
	dump("numbers", numbers)

	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
	dump("numbers", numbers)

	s1 := []int{10, 20, 30, 40, 50, 60, 70}
	numbers = append(numbers, s1...)
	dump("numbers", numbers)

	// Remove first element
	numbers = numbers[1:]
	dump("numbers", numbers)

	// Remove last element
	numbers = numbers[:len(numbers)-1]
	dump("numbers", numbers)

	// Remove middle element
	pos := int(len(numbers) / 2)
	numbers = append(numbers[:pos], numbers[pos+1:]...)
	dump("numbers", numbers)

	newnumbers := make([]int, len(numbers), (cap(numbers) * 2))
	count := copy(newnumbers, numbers)
	fmt.Printf("Count of copies from slice numbers: %v\n", count)
	dump("newnumbers", newnumbers)

	numbers[len(numbers)-1] = 88
	newnumbers[0] = 99
	dump("numbers", numbers)
	dump("newnumbers", newnumbers)
}

func dump(name string, x []int) {
	fmt.Print(name, ": ")
	fmt.Printf("addr=%p\tlen=%d\tcap=%d\tslice=%v\n", x, len(x), cap(x), x)
}
