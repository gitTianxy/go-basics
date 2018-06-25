package main

import (
	"fmt"
	"go-basics/pkg"
)

func main() {
	a := pkg.ElementA{}
	a.SetName("A")
	fmt.Printf("name of %T: %v\n", a, a.GetName())
	b := pkg.ElementB{}
	b.SetName("B")
	fmt.Printf("name of %T: %v\n", b, b.GetName())
}
