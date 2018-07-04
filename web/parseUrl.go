package main

import (
	"net/url"
	"fmt"
	"net"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#fragment"
	fmt.Println("orignial url:", s)
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// scheme
	fmt.Println("scheme:", u.Scheme)
	// user
	fmt.Println("user:", u.User)
	fmt.Println("name:", u.User.Username())
	pwd, _ := u.User.Password()
	fmt.Println("pwd:", pwd)
	// host
	fmt.Println("host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("domain:", host)
	fmt.Println("port:", port)
	// path
	fmt.Println("path:", u.Path)
	fmt.Println(u.Fragment)
	// parse query
	fmt.Println("query:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
