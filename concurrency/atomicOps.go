package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

/**
1. atomic write: atomic.addXXX(&seed, num)
2. atomic read: atomic.loadXXX(&seed)
 */
func main() {
	// atomic counter
	var counter uint32
	wc := 5
	concurrentJob(wc, &counter)
	r := waitResult(5, &counter)
	fmt.Println("final result:", r)
}

func concurrentJob(wc int, counter *uint32) {
	for i := 0; i < wc; i++ {
		go func(i int) {
			fmt.Println("start job", i)
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			fmt.Println("finish job", i)
			atomic.AddUint32(counter, 1)
		}(i)
	}
}

func waitResult(slpSecs int, counter *uint32) uint32 {
	time.Sleep(time.Duration(slpSecs) * time.Second)
	return atomic.LoadUint32(counter)
}
