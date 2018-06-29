package main

import (
	"time"
	"fmt"
)

func main() {
	t := timer{}
	t.wait(2*time.Second)
	t.wait2(2*time.Second)
	if <-timeout(2*time.Second) {
		fmt.Println("time out 3")
	}
}

type timer struct {
}

func (timer timer) wait(dur time.Duration) {
	time.Sleep(dur)
	fmt.Println("time out 1")
}

func (timer timer) wait2(dur time.Duration) {
	ch := make(chan bool)
	go func() {
		time.Sleep(dur)
		ch <- true
	}()
	<-ch
	fmt.Println("time out 2")
}

func timeout(dur time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(dur)
		ch <- true
	}()
	return ch
}

