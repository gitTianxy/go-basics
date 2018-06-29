package main

import (
	"fmt"
	"time"
)

func main()  {
	// switch demo 1
	in := "a"
	fmt.Printf("capital of %v: %v\n", in, lower2Capital(in))
	// switch demo 2
	fmt.Println("today is", getWeekDayNum())

	for _, c:= range "hello" {

		fmt.Println(fmt.Sprintf("%c", c))
	}

}

func lower2Capital(lower string) string {
	switch lower {
	case "a", "A":
		return "A"
	case "b", "B":
		return "B"
	default:
		return "OTHERS"
	}
}

func getWeekDayNum() int {
	switch time.Now().Weekday() {
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	case time.Sunday:
		return 7
	default:
		panic("error")
	}
}