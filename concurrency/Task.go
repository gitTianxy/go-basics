package concurrency

import (
	"math/rand"
	"fmt"
	"sync"
)

type Task struct {
	Name string
}

func (t Task) ChannelTask(c chan bool)  {
	s := 0
	for i := 0; i < rand.Intn(100000000); i++ {
		s += i
	}
	fmt.Println(t.Name, "sum:", s)
	c <- true
}

func (t Task) WaitGroupTask(wg *sync.WaitGroup) {
	s := 0
	for i := 0; i < rand.Intn(100000000); i++ {
		s += i
	}
	fmt.Println(t.Name, "sum:", s)
	wg.Done()
}