package main

import (
	"reflect"
	"fmt"
	"go-basics/reflect"
)

func main() {
	u := ref.User{}
	u.Id = 1
	u.Name = "zhangsan"
	u.Age = 18
	// fields & methods
	info := reflectInfo(u)
	fmt.Println("***", reflect.TypeOf(u).Name(), "fields & methods:")
	for k, v := range info {
		fmt.Printf("%6s: %v\n", k, v)
	}
	// field by index
	mgr := ref.Manager{u, "Jack"}
	reflectByIdx(mgr)
	// set field
	setField(&u, "Name", reflect.ValueOf("lisi"))
	fmt.Println("***user after setting:", u)
	// call method
	callMethod(u, "Hello", reflect.ValueOf("Kevin"))
}

func reflectInfo(o interface{}) map[string]interface{} {
	if reflect.TypeOf(o).Kind() != reflect.Struct {
		panic("input object is not a 'Struct'")
	}

	fm := map[string]interface{}{}
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	// fields reflection
	i := 0
	for i < t.NumField() {
		if v.Field(i).CanInterface() {
			fm[t.Field(i).Name] = v.Field(i).Interface()
		}
		i++
	}
	// methods reflection
	i = 0
	for i < t.NumMethod() {
		m := t.Method(i)
		fm[m.Name] = m.Type
		i++
	}
	return fm
}

func reflectByIdx(o interface{}) {
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Struct {
		panic("input object is not a 'Struct'")
	}
	fmt.Println("***", t.Name())
	v := reflect.ValueOf(o)
	fu := t.FieldByIndex([]int{0})
	fuid := t.FieldByIndex([]int{0, 0})
	funame := t.FieldByIndex([]int{0, 1})
	vuid := v.FieldByIndex([]int{0, 0})
	vuname := v.FieldByIndex([]int{0, 1})
	fmt.Printf("%6s: %v=%v, %v=%v\n", fu.Name, fuid.Name, vuid.Interface(), funame.Name, vuname.Interface())
	ftitle := t.FieldByIndex([]int{1})
	vtitle := v.FieldByIndex([]int{1})
	fmt.Printf("%6s: %v\n", ftitle.Name, vtitle.Interface())
}

func setField(o interface{}, name string, val reflect.Value) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr {
		panic("input not a pointer")
	} else {
		v = v.Elem()
	}
	f := v.FieldByName(name)
	if !f.IsValid() {
		panic(name + " not exists")
	}
	if !f.CanSet() {
		panic(name + " cannot be setted")
	}
	f.Set(val)
}

func callMethod(o interface{}, name string, args...reflect.Value)  {
	v := reflect.ValueOf(o)
	md := v.MethodByName(name)
	fmt.Println("***call", name)
	md.Call(args)
}