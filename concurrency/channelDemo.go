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
	channelTask()
	bufferedChannel()
	selectChannel()
	pingPong()
}

func pingPong()  {
	fmt.Println("*** ping pong demo")
	count := 10
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
 * 2. default case执行:
 * 		有default: 如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句;
 *		没有default: 代码块会被阻塞，直到有一个case通过评估
 */
func selectChannel() {
	fmt.Println("***select channel demo")
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


/**
 * 1. buffered channel
 * A. 表示只有当写入channel的数据量大于设定的缓冲数量时,对应的goroutine才阻塞挂起,
 * 等待其他goroutine将数据从channel中读出
 * B. 缓冲信道是先进先出的，可以看作一个线程安全的队列
 */
func bufferedChannel() {
	fmt.Println("*** buffered channel demo")
	size := 10
	ch := make(chan int, size)
	// write
	for i:=0;i<size ;i++  {
		ch <- i*100
	}
	close(ch) //不显示关闭, 则后面read完后会让routine挂起产生死锁
	// read
	for v := range ch {
		fmt.Println(v)
	}
}

func channelTask() {
	fmt.Println("***channel task demo")
	// init tasks
	tskNum := 10
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
