package main

import (
	"fmt"
)

/**
 * 1. 函数也是类型: go语言中一切皆类型
 * 2. 多返回值
 * 3. 变长参数
 * 4. 值传递
 * 5. 匿名函数: 不能作为顶层函数--必需包在其他函数体中
 */

func main() {
	// 多返回值, 变长参数
	s, m := calcs("calculating sum & multiplication of the sequence", 1, 2, 3, 4, 5)
	fmt.Printf("sum:%v, multiplication product:%v\n", s, m)
	// 值传递: value vs reference
	a := 1
	valRef(a, &a)
	fmt.Println("value after change:", a)
	// 匿名函数
	nf := func() {
		fmt.Println("anonymous function")
	}
	nf()
	// 闭包
	cm := make(map[string]string)
	mbd := mapBuilder(cm)
	mbd("name", "kevin")
	mbd("age", "28")
	fmt.Println(cm)
}

/**
 * 多返回值, 变长参数
 */
func calcs(msg string, p ...int) (s, m int) {
	fmt.Printf("%v: %v\n", msg, p)
	m = 1
	for _, a := range p {
		s += a
		m *= a
	}
	return
}

/**
 * go语言的函数参数都是值传递, 无论传递的是值还是引用值
 */
func valRef(v int, r *int) {
	fmt.Printf("value: %v, refer: %v\n", v, r)
	*r = 2
}

/**
 * 闭包
 * 闭包=函数+引用环境; '函数'是一些可执行的代码，这些代码在函数被定义后就确定了，不会在执行时发生变化，
 * 所以一个函数只有一个实例。'闭包'在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
 */
func mapBuilder(m map[string]string) func(string, string) {
	return func(k string, v string) {
		m[k] = v
	}
}
