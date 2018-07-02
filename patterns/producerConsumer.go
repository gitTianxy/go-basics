package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go consumer(jobs, done)
	producer(jobs, 20)
	close(jobs)
	<-done
	fmt.Println("all jobs done")
}

func producer(jobs chan int, count int) {
	for i := 0; i < count; i++ {
		jobs <- i
		fmt.Println("produce job", i)
	}
}

func consumer(jobs <-chan int, done chan<- bool) {
	for {
		job, more := <-jobs
		if more {
			fmt.Println("do job", job)
		} else {
			done <- true
			break
		}
	}
}
