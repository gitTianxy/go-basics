package main

import "fmt"

func main() {
	// append slice
	fmt.Println("***append slice")
	s := make([]int, 0)
	fmt.Println(s)
	s = appendSlice(s, 2)
	fmt.Println(s)

	// iteration
	s2 := []string{"A", "B", "C", "D"}
	for idx, v := range s2 {
		fmt.Println(idx, v)
	}
}

/**
 * 由于在slice扩容过程中会重新分配地址, 所以不用引用而是返回值来取得扩容后结果
 */
func appendSlice(s []int, e int) []int {
	return append(s, e)
}
