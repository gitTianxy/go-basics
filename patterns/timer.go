package main

import (
	"time"
	"fmt"
)

/*
a `timer` is used to execute GO code after a certain period
 */
func main() {
	t := timer{}
	t.wait(2 * time.Second)
	t.wait2(2 * time.Second)
	if <-timeout(2 * time.Second) {
		fmt.Println("time out, by channel blk")
	}
	builtinTimer(2 * time.Second)
}

type timer struct {
}

func (timer timer) wait(dur time.Duration) {
	time.Sleep(dur)
	fmt.Println("time out, by time.sleep")
}

func (timer timer) wait2(dur time.Duration) {
	ch := make(chan bool)
	go func() {
		time.Sleep(dur)
		ch <- true
	}()
	<-ch
	fmt.Println("time out, by channel blk")
}

func timeout(dur time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(dur)
		ch <- true
	}()
	return ch
}

/**
1. get a channel timer
2. do stop before expiration
 */
func builtinTimer(ws time.Duration) {
	timer := time.NewTimer(ws)
	<-timer.C
	fmt.Println(ws, "expired")
}
