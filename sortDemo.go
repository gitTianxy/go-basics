package main

import (
	"sort"
	"fmt"
)

var fruits = []string{"peach", "banana", "kiwi", "mongo"}

/**
1. sort by native order
2. sort by function: self-defined sort rule
 */
func main() {
	fmt.Println("original sequence:", fruits)
	sortByNativeOrder()
	sortByLength()
}

func copyFruits() []string  {
	cp := make([]string, len(fruits))
	copy(cp, fruits)
	return cp
}

/*
sort alphabetically
 */
func sortByNativeOrder() {
	fmt.Println("*** sort alphabetically:")
	cp := copyFruits()
	sort.Strings(cp)
	fmt.Println(cp)
}

/*
sort by word length
 */
func sortByLength() {
	fmt.Println("*** sort by word length:")
	cp := copyFruits()
	sort.Sort(byLength(cp))
	fmt.Println(cp)
}

// implementing 3 methods: Len, Swap, Less
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	//if len(s[i]) == len(s[j]) {
	//	return s[i] < s[j]
	//} else {
	//	return len(s[i]) < len(s[j])
	//}
	return len(s[i]) < len(s[j])
}
