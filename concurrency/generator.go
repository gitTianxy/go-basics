package main

import "fmt"

/**
 * use channel as generator
 */

type XRange struct {
	idx   int
	start int
	end int
	ch    chan int
}

func (xr XRange) HasNext() bool {
	return xr.idx < (xr.end - 1)
}

func (xr XRange) Next() int {
	return <-xr.ch
}

func xrange(start int, end int) chan int {
	if start > end {
		panic("start greater than end")
	}
	ch := make(chan int, end - start + 1)
	go func() {
		for i := start; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	xr := xrange(0, 10)
	for v := range xr {
		fmt.Println(v)
		if len(xr) <= 0 {
			break
		}
	}
}
