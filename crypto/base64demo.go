package main

import (
	"fmt"

	"encoding/base64"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println("original data:", data)
	// standard encoding/decoding
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("standard base64:", sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("standard decoding:", string(sDec))
	// URL-compatible encoding/decoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("url base64:", uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println("url decoding:", string(uDec))
}
