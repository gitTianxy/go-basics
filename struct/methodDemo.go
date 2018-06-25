package main

import (
	"fmt"
	"strconv"
)

func main() {
	d := Duck{"duckA"}
	d.swim()
	fmt.Println("name[origin]:", d.getName())
	d.setName("duckA'")
	fmt.Println("name[after setting]:", d.getName())
	// 包装一个Integer
	intObj := Integer{5}
	fmt.Println(intObj.toString())
	cmp := 4
	fmt.Println(intObj, "equals to", cmp, ":", intObj.equals(cmp))
	// 绑定到类型上的方法可以有两种调用方式: method value vs expression
	fmt.Printf("method value call: ")
	d.swim()
	fmt.Printf("method expression call: ")
	(Duck).swim(d)
}

type Duck struct {
	name string
}

/**
 * 方法定义
 * 1. receiver: d
 * 2. operations: swim() {...}
 */
func (d Duck) swim() {
	fmt.Println(d.name, "is swimming")
}

/** !!!go中无/不允许方法重载
func (d Duck) swim(secs int) {
	fmt.Println(d.name, "swims for", secs, "seconds")
}
 */

func (d Duck) getName() string {
	return d.name
}

/**
 * 操作对象, 必需传入引用
 */
func (d *Duck) setName(name string) {
	d.name = name
}

type Integer struct {
	val int
}

func (i Integer) toString() string {
	return strconv.Itoa(i.val)
}

func (i Integer) equals(b int) bool  {
	return i.val == b
}
