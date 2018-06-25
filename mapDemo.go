package main

import "fmt"

func main() {
	m1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	fmt.Println(m1)

	m2 := make(map[string]int, len(m1))
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
