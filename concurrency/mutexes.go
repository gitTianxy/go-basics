package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

func main() {
	// mutex counter
	mutexCounter(5)
	// concurrent map: by mutex
	concurrMapByMutex()
	// concurrent map: by channel
	concurrMapByCh()
}

func mutexCounter(size int) {
	fmt.Println("*** mutex counter")
	mutex := &sync.Mutex{}
	counter := 0
	for i := 0; i < size; i++ {
		go counts(i, &counter, mutex)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("total count:", counter)
}

func concurrMapByMutex() {
	fmt.Println("*** concurrent map: by mutex")
	mutex := &sync.Mutex{}
	state := map[int]int{}
	for i := 1; i <= 10; i++ {
		go mutexMap([]int{1 * i, 2 * i, 3 * i, 4 * i, 5 * i, 6 * i, 7 * i, 8 * i, 9 * i}, state, mutex)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("final map:", state)
}

func concurrMapByCh() {
	fmt.Println("*** concurrent map: by channel")
	m := make(map[int]int)
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	stop := make(chan bool)
	go rwMap(m, reads, writes, stop)
	go readRoutine(reads)
	go writeRoutine(writes)
	time.Sleep(20*time.Millisecond)
	stop <- true
	fmt.Println("final map:", m)
}

func counts(idx int, counter *int, mutex *sync.Mutex) {
	c := 0
	rd := rand.Intn(100)
	for i := 0; i < rd; i++ {
		rd := rand.Intn(5)
		c += rd
		mutex.Lock()
		*counter += rd
		mutex.Unlock()
		//time.Sleep(time.Millisecond)
	}
	fmt.Println("counter", idx, "get:", c)
}

/**
otherwise, `concurrent map read and map write` would rise
 */
func mutexMap(list []int, m map[int]int, mutex *sync.Mutex) {
	fmt.Println("add", list)
	for i, v := range list {
		mutex.Lock()
		m[i] += v
		mutex.Unlock()
	}
}

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func rwMap(m map[int]int, reads chan *readOp, writes chan *writeOp, stop chan bool) {
	for {
		select {
		// 应答读请求
		case read := <-reads:
			fmt.Println("read:", read.key, m[read.key])
			read.resp <- m[read.key]
			// 应答写请求
		case write := <-writes:
			m[write.key] = write.val
			write.resp <- true
			fmt.Println("write:", write.key, write.val)
		case <-stop:
			break
		}
	}
}

func readRoutine(reads chan *readOp) {
	for {
		read := &readOp{
			key:  rand.Intn(5),
			resp: make(chan int)}
		// 启动读
		reads <- read
		// 等待rwMap的读应答
		<-read.resp
		time.Sleep(time.Millisecond)
	}
}

func writeRoutine(writes chan *writeOp) {
	for {
		write := &writeOp{
			key:  rand.Intn(5),
			val:  rand.Intn(100),
			resp: make(chan bool)}
		// 启动写
		writes <- write
		// 等待rwMap的写应答
		<-write.resp
		time.Sleep(time.Millisecond)
	}
}
