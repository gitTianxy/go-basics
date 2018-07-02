package main

import (
	"encoding/json"
	"fmt"
)

/*
Go offers built-in support for JSON encoding and decoding,
including to and from built-in and custom data types.
 */
func main() {
	// convert obj 2 json
	obj2json()
	// parse json 2 obj
	json2obj()
}

func obj2json() {
	fmt.Println("*** obj 2 json")
	// bool 2 string
	bolB, err := json.Marshal(true)
	if err == nil {
		fmt.Println(string(bolB))
	}
	// array 2 string
	arrB, err := json.Marshal([]int{1, 2, 3, 4, 5})
	if err == nil {
		fmt.Println(string(arrB))
	}
	// map 2 json
	mapB, err := json.Marshal(map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5})
	if err == nil {
		fmt.Println(string(mapB))
	}
	// self-defined obj 2 json
	o := obj{
		Id:   1,
		Name: "obj",
		Sub: subObj{
			Id:   10,
			Name: "sub obj",
		},
		privateField: "field-ignored-in-json",
	}
	oB, err := json.Marshal(o)
	if err == nil {
		fmt.Println(string(oB))
	}
}

func json2obj() {
	fmt.Println("*** json 2 obj")
	// str 2 bool
	var b bool
	boolStr := `true`
	if err := json.Unmarshal([]byte(boolStr), &b); err == nil {
		fmt.Println(b)
	}
	// string 2 array
	var arr []int
	arrStr := `[1,2,3,4,5]`
	if err := json.Unmarshal([]byte(arrStr), &arr); err == nil {
		fmt.Println(arr)
	}
	// string 2 map
	var m map[string]int
	mapStr := `{"A":1,"B":2,"C":3,"D":4,"E":5}`
	if err := json.Unmarshal([]byte(mapStr), &m); err == nil {
		fmt.Println(m)
	}
	// string 2 obj
	var o obj
	objStr := `{"ID":1,"name":"obj","sub":{"ID":10,"name":"sub obj"}}`
	if err := json.Unmarshal([]byte(objStr), &o); err == nil {
		fmt.Println(o)
	}
}

type obj struct {
	Id           int    `json:"ID"`
	Name         string `json:"name"`
	Sub          subObj `json:"sub"`
	privateField string //private field will ignored
}

type subObj struct {
	Id   int    `json:"ID"`
	Name string `json:"name"`
}
