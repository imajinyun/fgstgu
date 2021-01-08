package main

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/redis.v4"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var (
		lockKey    = "counter.lock"
		counterKey = "counter"
	)
	setNXResponse := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := setNXResponse.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result:", lockSuccess)
		return

	}

	getResponse := client.Get(counterKey)
	countValue, err := getResponse.Int64()
	if err == nil {
		countValue++
		setResponse := client.Set(counterKey, countValue, 0)
		if _, err = setResponse.Result(); err != nil {
			println("set value error:", err)
		}
	}
	println("current counter is", countValue)
	delResponse := client.Del(lockKey)
	unlockSuccess, err := delResponse.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success")
	} else {
		println("unlock failed", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
