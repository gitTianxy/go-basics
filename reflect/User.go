package ref

import "fmt"

/**
 * NOTE:
 * only `public fields/methods` can be reflected outside the package
 */
type User struct {
	Id           int
	Name         string
	Age          int
	privateField int
}

func (u User) PublicMethod() {
	fmt.Println("public method")
}

func (u User) privateMethod() {
	fmt.Println("private method")
}

func (u User) Hello(guest string)  {
	fmt.Println("hello", guest, ", I'm", u.Name)
}