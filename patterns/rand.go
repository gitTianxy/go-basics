package main

import "fmt"

func main() {
	for i:=0;i<10 ;i++  {
		fmt.Println(<-rand())
	}
}

func rand() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:

			}
		}
	}()

	return ch
}
