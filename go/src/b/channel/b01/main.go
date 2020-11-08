package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go toSender(ch1)

	// for {
	// 	data := <-ch1
	// 	if data == "" {
	// 		break
	// 	}
	// 	fmt.Println("Read data from channel 1: ", data)
	// }

	// for {
	// 	data, ok := <-ch1
	// 	fmt.Println(ok)
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Read data from channel 2: ", data)
	// }

	for value := range ch1 {
		fmt.Println("Read data form channel 3: ", value)
	}
}

func toSender(ch chan string) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Send data: %d\n", i)
	}
	fmt.Println("🐶 Data sent completed!")
}
