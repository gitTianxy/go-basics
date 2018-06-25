package interf

import "fmt"

/**
 * Male同时实现`Action`和`Homo`两个接口--通过绑定接口方法的方式
 */
type Male struct {
	name string
	age  int
}

func (m *Male) SetName(name string)  {
	m.name = name
}

func (m Male) Name() string {
	return m.name
}

func (m Male) Sex() string {
	return "male"
}

func (m *Male) SetAge(age int) {
	m.age = age
}

func (m Male) Age() int {
	return m.age
}

func (m Male) SelfIntro() string {
	return fmt.Sprintf("%v[sex:%v, age:%v]", m.Name(), m.Sex(), m.Age())
}
