package main

import (
	"fmt"
	"strings"
	"regexp"
)

/*
1. built-in functions
2. formatting
3. regular expression
 */
func main() {
	builtinfuncs()
	formatting()
	regexpr()
}

func regexpr()  {
	fmt.Println("*** regular expression")
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.FindString("peach punch patch paddle"))
	fmt.Println(r.FindAllString("peach punch patch paddle", -1))
}

func formatting()  {
	fmt.Println("*** formatting")
	fmt.Println(fmt.Sprintf("value: %v", 123))
	fmt.Println(fmt.Sprintf("type: %T", 123))
	fmt.Println(fmt.Sprintf("int: %d", 123))
	fmt.Println(fmt.Sprintf("float: %0.2f", 123.0))
	fmt.Println(fmt.Sprintf("binary: %b", 123))
	fmt.Println(fmt.Sprintf("hex: %x", 123))
	var p = "haha"
	fmt.Println(fmt.Sprintf("pointer: %p", &p))
	fmt.Println(fmt.Sprintf("scientific form: %e", 1234567.0))
}

func builtinfuncs()  {
	fmt.Println("*** built-in functions")
	str := "string sample"

	prefix := "str"
	fmt.Printf("'%v' has prefix '%v': %v\n", str, prefix, strings.HasPrefix(str, prefix))

	letter := "s"
	fmt.Printf("count of '%v' in '%v': %v\n", letter, str, strings.Count(str, letter))
	fmt.Printf("index of '%v' in '%v': %v\n", letter, str, strings.Index(str, letter))
	fmt.Printf("'%v' is contained in '%v': %v\n", letter, str, strings.Contains(str, letter))
	fmt.Println(str,"to UPPER:", strings.ToUpper(str))

	fmt.Println("repeat '", letter,"'", 5, "times:", strings.Repeat(letter, 5))
	fmt.Println("replace ALL 's' by 'S':", strings.Replace(str, "s", "S", -1))
	fmt.Println("replace one 's' by 'S':", strings.Replace(str, "s", "S", 1))

	strs := []string{"A", "B", "C", "D"}
	joiner := ","
	splitter := " "
	fmt.Printf("join '%v' by '%v': %v\n", strs, joiner, strings.Join(strs, joiner))
	fmt.Printf("split '%v' by '%v': %v\n", str, splitter, strings.Split(str, splitter))
}