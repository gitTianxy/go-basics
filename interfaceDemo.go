package main

import (
	"go-basics/interface"
	"fmt"
)

func main() {
	kevin := interf.Male{}
	kevin.SetName("kevin")
	kevin.SetAge(18)
	fmt.Println(kevin.SelfIntro())

	homo := interf.Homo(kevin)
	fmt.Printf("HOMO[name: %v, sex:%v, age:%v]\n", homo.Name(), homo.Sex(), homo.Age())

	action := interf.Action(kevin)
	fmt.Println("self-introduction:", action.SelfIntro())
}
