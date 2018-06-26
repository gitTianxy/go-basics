package main

import (
	"runtime"
	"time"
	"go-basics/concurrency"
	"strconv"
	"fmt"
	"sync"
)

/**
 * go用channel实现并发线程通信, 分别有写(c <- true)和读(<-c)操作
 */
var ch chan string

func main() {
	// init tasks
	tskNum := 10
	tsks := make([]concurrency.Task, tskNum)
	for i := 0; i < tskNum; i++ {
		tsks[i] = concurrency.Task{"task" + strconv.Itoa(i)}
	}
	// by channel
	channelDemo(tsks)
	// by wait group
	waitGroupDemo(tsks)
	// select between channels: TODO
	selectChannels()
	// ping & pong
	fmt.Println("***ping pong")
	count := 5
	flag := make(chan bool, count)
	ch = make(chan string)
	go pong(flag)
	go ping(count)
	for i := 0; i < count; i++ {
		<-flag
	}
}

func channelDemo(tsks []concurrency.Task) {
	fmt.Println("***channel demo")
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, len(tsks))
	t0 := time.Now()
	// launch task
	for _, t := range tsks {
		go t.ChannelTask(c)
	}
	// wait for finish
	for i := 0; i < len(tsks); i++ {
		<-c
	}
	fmt.Printf("tasks spent: %v secs\n", time.Since(t0).Seconds())
}

func waitGroupDemo(tsks []concurrency.Task) {
	fmt.Println("***waitGroup demo")
	runtime.GOMAXPROCS(runtime.NumCPU())
	tskNum := len(tsks)
	wg := sync.WaitGroup{}
	wg.Add(tskNum)
	t0 := time.Now()
	// launch task
	for _, t := range tsks {
		go t.WaitGroupTask(&wg)
	}
	// wait for finish
	wg.Wait()
	fmt.Printf("tasks spent: %v secs\n", time.Since(t0).Seconds())
}

/**
 * SELECT
 * select {
    case communication clause  :
       statement(s);
	case communication clause  :
       statement(s);
    //任意数量的case
	default :
		statement(s);//可选
	}
 * select通信执行规则
 * 代码执行到 select 时，case语句会按照源代码的顺序被评估，且只评估一次，
 * 评估的结果会出现下面这几种情况:
 * 1. 各case执行顺序:
 * 		如果只有一个 case 语句评估通过，那么就执行这个case里的语句;
 * 		如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个.
 * 2. default case执行:
 * 		有default: 如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句;
 *		没有default: 代码块会被阻塞，直到有一个case通过评估
 */
func selectChannels() {
	fmt.Println("***select channels")
	o := make(chan bool)
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				// 如果c1关闭, ok为false
				if !ok {
					fmt.Println("c1 closed")
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2 closed")
					o <- true
					break
				}
				fmt.Println("c2", v)
			default:
				//fmt.Println("default")
			}
		}
	}()
	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "joe"

	close(c2)
	close(c1)

	<-o
	//for i:=0; i<2;i++  {
	//	<-o
	//}
}

func ping(count int) {
	for i := 0; i < count; i++ {
		ch <- fmt.Sprintf("ping %d", i)
		fmt.Println(<-ch)
	}
}

func pong(flag chan bool) {
	c := 0
	for {
		fmt.Println(<-ch)
		ch <- fmt.Sprintf("pong %d", c)
		flag <- true
		c++
	}
}
