package main

import (
	"fmt"
	"strings"
)

/*
I. built-in operations
append
iteration
join/split

II. self-defined functions
index
include
any
all
filter
map
 */
func main() {
	// built-in: append slice
	fmt.Println("***append slice")
	s := make([]int, 0)
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	// built-in: iteration
	s2 := []string{"A", "B", "C", "D"}
	for idx, v := range s2 {
		fmt.Println(idx, v)
	}

	// self-defined: index
	t := "B"
	fmt.Printf("index of '%v' in %v: %v\n", t, s2, index(s2, t))
	t = "b"
	fmt.Printf("index of '%v' in %v: %v\n", t, s2, index(s2, t))
	// self-defined: include
	t = "A"
	fmt.Printf("'%v' is included in %v: %v\n", t, s2, include(s2, t))
	t = "a"
	fmt.Printf("'%v' is included in %v: %v\n", t, s2, include(s2, t))
	// self-defined: any
	strs := []string{"peach", "apple", "pear", "plum"}
	fmt.Printf("any %v has prefix 'p': %v\n", strs, any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	// self-defined: all
	fmt.Printf("all %v has prefix 'p': %v\n", strs, all(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	// self-defined: filter
	fmt.Printf("%v has prefix 'p': %v\n", strs, filter(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	// self-defined: mapping
	fmt.Printf("%v to UPPER CASE: %v\n", strs, mapping(strs, func(s string) string {
		return strings.ToUpper(s)
	}))
}

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []string, t string) bool  {
	return index(vs, t) >= 0
}

/*
Any returns true if one of the strings in the slice satisfies the predicate f
 */
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/*
All returns true if all of the strings in the slice satisfy the predicate f.
 */
func all(vs []string, f func(string) bool) bool  {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filter(vs []string, f func(string) bool) []string {
	rs := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			rs = append(rs, v)
		}
	}
	return rs
}

func mapping(vs []string, f func(string) string) []string {
	rs := make([]string, len(vs))
	for i, v := range vs {
		rs[i] = f(v)
	}
	return rs
}