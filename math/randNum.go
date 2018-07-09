package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
Go’s math/rand package provides pseudorandom number generation:
当不重设source时, 每次都得到相同的随机序列
 */
func main() {
	sameSourceRandom()

	newSourceRandom()
}

func sameSourceRandom()  {
	fmt.Println("random int within [0, 100]:", rand.Intn(100))
	fmt.Println("random int within [0, 100]:", rand.Intn(100))
	fmt.Println("random float within [0, 100]:", rand.Float64()*100)
	fmt.Println("random float within [0, 100]:", rand.Float64()*100)
}

func newSourceRandom() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("random[new source] int within [0, 100]:", r1.Intn(100))
	fmt.Println("random[new source] int within [0, 100]:", r1.Intn(100))
	fmt.Println("random[new source] float within [0, 100]:", r1.Float64()*100)
	fmt.Println("random[new source] float within [0, 100]:", r1.Float64()*100)
}
