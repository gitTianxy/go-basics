package main

import (
	"strconv"
	"sync"
	"math/rand"
	"fmt"
	"runtime"
	"time"
)

func main()  {
	// init tasks
	tskNum := 10
	tsks := make([]wgTask, tskNum)
	for i := 0; i < tskNum; i++ {
		tsks[i] = wgTask{"task" + strconv.Itoa(i)}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(tskNum)
	// launch task
	t0 := time.Now()
	for _, t := range tsks {
		go t.doTask(&wg)
	}
	// wait for finish
	wg.Wait()
	fmt.Printf("tasks spent: %v secs\n", time.Since(t0).Seconds())

}

type wgTask struct {
	name string
}

func (t wgTask) doTask(wg *sync.WaitGroup) {
	s := 0
	for i := 0; i < rand.Intn(100000000); i++ {
		s += i
	}
	fmt.Println(t.name, "sum:", s)
	wg.Done()
}