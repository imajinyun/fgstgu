package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomTests = []struct {
	n interface{}
}{
	{rand.Int()},
	{rand.Intn(5)},
	{rand.Int31()},
	{rand.Int31n(5)},
	{rand.Int63()},
	{rand.Int63n(5)},
	{rand.Uint32()},
	{rand.Uint64()},
}

func main() {
	for i := 0; i < 2; i++ {
		for _, v := range randomTests {
			fmt.Printf("%v\n", v.n)
		}
		fmt.Println()
	}
	fmt.Println("The following will be printed randomly:")

	rand.Seed(time.Now().Local().UnixNano())
	fmt.Printf("%v\n", rand.Intn(5))

	fmt.Println("Generate a random number in a specified range:")
	// [25, 75)
	fmt.Printf("%v\n", rand.Intn(50)+25)
}
