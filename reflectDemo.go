package main

import (
	"reflect"
	"go-basics/reflect"
	"fmt"
)

func main() {
	u := ref.User{}
	u.Id = 1
	u.Name = "zhangsan"
	u.Age = 18
	info := reflectInfo(u)
	fmt.Println(reflect.TypeOf(u), "fields & methods:")
	for k, v := range info {
		fmt.Printf("%6s: %v\n", k, v)
	}
}

func reflectInfo(o interface{}) map[string]interface{} {
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
