package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"runtime"
	"time"
)

var ppCh chan string
var ppFlag chan bool

/**
 * go用channel实现goroutine之间的通信,分别有写(c <- true)和读(<-c)操作
 */
func main() {
	channelTask(5)
	bufferedChannel(5)
	selectChannel()
	pingPong(5)
	mergeChs(5)
	directionalChs()
	closeCh(5)
}

/**
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel’s receivers.
 */
func closeCh(count int) {
	fmt.Println("***close channel demo")
	work := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			w, more := <-work
			if more {
				fmt.Println("todo", w)
			} else {
				done <- true
				break
			}
		}
	}()
	for i := 0; i < count; i++ {
		work <- i
	}

	close(work)
	<-done
	fmt.Println("done")
}

/*
When using channels as function parameters or return value,
you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.
 */
func directionalChs() {
	fmt.Println("***directional channels")
	in := make(chan int)
	out := make(chan int)
	go func() {
		in <- 1
	}()
	go func(in <-chan int, out chan<- int) {
		c := <-in
		fmt.Println("input:", c)
		out <- c
	}(in, out)
	fmt.Println("output:", <-out)
}

func pingPong(count int) {
	fmt.Println("*** ping pong demo")
	ppFlag = make(chan bool)
	ppCh = make(chan string)
	go pong()
	go ping(count)
	for i := 0; i < count; i++ {
		<-ppFlag
	}
}

func ping(count int) {
	for i := 0; i < count; i++ {
		ppCh <- fmt.Sprintf("ping %d", i)
		fmt.Println(<-ppCh)
		ppFlag <- true
	}
}

func pong() {
	c := 0
	for {
		fmt.Println(<-ppCh)
		ppCh <- fmt.Sprintf("pong %d", c)
		c++
	}
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
 * 2a. default case执行:
 * 		有default: 如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句;
 *		没有default: 代码块会被阻塞，直到有一个case通过评估
 * 2b. time.After(Duration): 设置超时
 */
func selectChannel() {
	fmt.Println("***select channel demo")

	closed := make(chan bool)
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for quit := false; !quit; {
			select {
			case v, ok := <-c1:
				// 如果c1关闭, ok为false
				if !ok {
					fmt.Println("c1 closed")
					closed <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2 closed")
					closed <- true
					break
				}
				fmt.Println("c2", v)
			case <-time.After(5 * time.Second):
				fmt.Println("wait timeout")
				closed <- true
				closed <- true
				quit = true
				//default:
				//fmt.Println("default")
			}
		}
	}()
	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "joe"

	//若一直不关闭c1和c2,则最后走timeout通道退出
	//close(c1)
	//close(c2)

	for i := 0; i < 2; i++ {
		<-closed
	}
}

/**
 * 1. buffered channel
 * A. 表示只有当写入channel的数据量大于设定的缓冲数量时,对应的goroutine才阻塞挂起,
 * 等待其他goroutine将数据从channel中读出
 * B. 缓冲信道是先进先出的，可以看作一个线程安全的队列
 */
func bufferedChannel(size int) {
	fmt.Println("*** buffered channel demo")
	ch := make(chan int, size)
	// write
	for i := 0; i < size; i++ {
		ch <- i * 100
	}
	close(ch) //不显示关闭, 则后面read完后会让routine挂起产生死锁
	// read
	for v := range ch {
		fmt.Println(v)
	}
}

func channelTask(tskNum int) {
	fmt.Println("***channel task demo")
	// init tasks
	tsks := make([]chTask, tskNum)
	for i := 0; i < tskNum; i++ {
		tsks[i] = chTask{"task" + strconv.Itoa(i)}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, tskNum)
	t0 := time.Now()
	// launch task
	for _, t := range tsks {
		go t.doTask(c)
	}
	// wait for finish
	for i := 0; i < len(tsks); i++ {
		<-c
	}
	fmt.Printf("tasks spent: %v secs\n", time.Since(t0).Seconds())
}

type chTask struct {
	name string
}

func (t chTask) doTask(c chan bool) {
	s := 0
	for i := 0; i < rand.Intn(100000000); i++ {
		s += i
	}
	fmt.Println(t.name, "sum:", s)
	c <- true
}

/**
 * 通过`ch1<- <-ch2`将一个ch2的内容流入ch1
 */
func mergeChs(num int) {
	fmt.Println("*** merge channels")
	chs := make([]chan int, num)
	for i := 0; i < num; i++ {
		chs[i] = branch(i)
	}

	ch := make(chan int)
	for _, c := range chs {
		go func(c chan int) {
			ch <- <-c
		}(c)
	}

	for i := 0; i < num; i++ {
		fmt.Println(<-ch)
	}
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch <- x
	}()
	return ch
}
