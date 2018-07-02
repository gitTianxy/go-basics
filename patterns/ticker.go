package main

import (
	"time"
	"fmt"
	rand2 "math/rand"
)

/*
tickers are for when you want to do something repeatedly at regular intervals
 */
func main() {
	ticker1(5, 500 * time.Millisecond)
	ticker2(5, 500 * time.Millisecond)
}

func ticker2(count int, interval time.Duration)  {
	fmt.Println("***ticker 2")
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

func ticker1(count int, interval time.Duration) {
	fmt.Println("***ticker 1")
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
