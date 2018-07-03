package main

import (
	"time"
	"fmt"
	rand2 "math/rand"
)

/*
tickers are for when you want to do something repeatedly at regular intervals
1. implements:
	a. time.Tick(interval)
	b. time.NewTicker(interval)
2. application
	a. crontab
	b. rate limiter: schedule requests
 */
func main() {
	crontab(5, 500 * time.Millisecond)
	rateLimiter(5, 500 * time.Millisecond)
}

// time.Tick(interval)
func rateLimiter(count int, interval time.Duration)  {
	fmt.Println("***rate limiter")
	requests := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			requests <- i
		}
		close(requests)
	}()
	limiter := time.Tick(interval)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

// time.NewTicker(interval)
func crontab(count int, interval time.Duration) {
	fmt.Println("***crontab")
	c :=0
	ticker := time.NewTicker(interval)
	for t := range ticker.C {
		if c==count {
			break
		}
		fmt.Println("Tick at", t)
		go doJob()
		c++
	}
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func doJob() {
	fmt.Println("do sth...")
	time.Sleep(time.Duration(rand2.Intn(2)) * time.Second)
}
