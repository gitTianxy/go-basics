package main

import "fmt"

func main() {
	buffer := 5
	jobs := make(chan int, buffer)
	done := make(chan bool)

	go consumer(jobs, done)
	// 同步进程, 和consumer进程沟通
	producer(jobs, 20)
	// 关闭job, <-jobs将输出'0,false'
	close(jobs)
	<-done
	fmt.Println("all jobs done")
}

func producer(jobs chan<- int, count int) {
	// 计数从1开始, 以区别close(jobs)时<-jobs返回0
	for i := 1; i <= count; i++ {
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
