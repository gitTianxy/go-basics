package main

import "fmt"

/**
 * 1. defer函数的执行顺序
 * 2. panic/recover
 */
func main() {
	// defer用法
	for i := 0; i<5 ; i++ {
		// 函数: 参数传递(值拷贝)
		//defer fmt.Println(i)
		// 闭包: 引用传递
		defer func() {
			fmt.Println(i)
		}()
	}
	// panic vs recover
	A()
	B()
	C()
}

func A()  {
	fmt.Println("Func A")
}

/**
 * panic: cause exception
 * recover: catch & treat
 */
func B()  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("Recover in B: ", err)
		}
 	}()
	panic("Panic in B")
}

func C()  {
	fmt.Println("Func C")
}