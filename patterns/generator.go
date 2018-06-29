package main

import (
	"fmt"
)

/**
 * use channel as generator
 */

type XRange struct {
	idx   int
	start int
	end   int
	ch    chan int
}

func NewXRange(start int, end int) XRange {
	ch := make(chan int)
	go func() {
		for i := start; i < end; i++ {
			ch <- i
		}
	}()
	xr := XRange{idx: start, start: start, end: end, ch: ch}
	return xr
}

func (xr XRange) HasNext() bool {
	return xr.idx < xr.end
}

func (xr *XRange) Next() int {
	xr.idx++
	return <-xr.ch
}

func xrange(start int, end int) chan int {
	if start > end {
		panic("start greater than end")
	}
	ch := make(chan int)
	go func() {
		for i := start; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	start := 0
	end := 10
	xr := xrange(start, end)

	for i := start; i < end; i++ {
		fmt.Println(<-xr)
	}

	xr2 := NewXRange(start, end)
	for xr2.HasNext() {
		fmt.Println(xr2.Next())
	}
}
