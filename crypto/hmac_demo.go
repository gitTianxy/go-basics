package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

var (
	key = []byte("ppc")
)

func HmacSha1(content string) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func main() {
	//hmac ,use sha1
	sign := HmacSha1("snoss")
	fmt.Printf("%x\n", sign)
}
