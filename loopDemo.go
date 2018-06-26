package main

import "fmt"

func main() {
	s := 0
	// for
	for i := 0; i < 10; i++ {
		s += i
	}
	fmt.Println("sum:", s)
	// for while
	i := 0
	s = 0
	for i < 10 {
		s += i
		i++
	}
	fmt.Println("sum:", s)
	// for each: map
	m := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	// for each: slice
	slice := []string{"A", "B", "C", "D"}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}
