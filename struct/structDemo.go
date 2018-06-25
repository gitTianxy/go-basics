package main

import "fmt"

/**
 * 1. 命名结构
 * 2. 匿名结构
 * 3. struct嵌套
 * NOTE
 * 1. go中没有class, 没有继承
 */

func main() {
	// 命名结构
	kevin := &Person{
		name: "kevin",
		age: 28,
	}
	kevin.sex = "male"
	//kevin.Human.sex = "male"
	fmt.Println(*kevin)
	setAge(kevin, 18)
	fmt.Println(*kevin)
	// 嵌套struct
	kevin.contact.phone = "136"
	kevin.contact.city="shanghai"
	fmt.Println(*kevin)
	// 匿名结构
	csBk := &struct {
		title string
		price float32
	}{
		title: "computer science",
		price: 25.0,
	}
	fmt.Println(*csBk)
}

/**
 * 因为go函数时值传递, 所以需要传入引用来实现set操作
 */
func setAge(p *Person, age int)  {
	p.age = age
}

type Human struct {
	sex string
}

type Person struct {
	name string
	age int
	contact struct{
		phone string
		city string
	}
	Human
}
