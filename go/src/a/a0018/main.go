package main

import "fmt"

// 可变参数
func main() {
	sum, avg, cnt := score(90, 85.5, 70, 66, 77, 65.5)
	fmt.Printf("cnt = %d, sum = %.2f, avg = %.2f\n", cnt, sum, avg)

	scores := []float64{92, 88, 71.5, 69, 70.5, 88}
	sum, avg, cnt = score(scores...)
	fmt.Printf("cnt = %d, sum = %.2f, avg = %.2f\n", cnt, sum, avg)
}

func score(scores ...float64) (sum, avg float64, cnt int) {
	for _, value := range scores {
		sum += value
		cnt++
	}
	avg = sum / float64(cnt)
	return
}
