package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"bufio"
	"strings"
)

func main() {
	var fpath = "file/data/file-write-example.txt"

	mkFile(fpath)

	checkWritePermission(fpath)

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writeBytes(f, []byte("hello, bytes\r"))
	writeString(f, "hello, string\r")
	buffWriter(f)

	rename(fpath, strings.Replace(fpath, "file-write-example", "file-write-new", -1))

}

func rename(old string, new string) {
	fmt.Println("*** rename", old)
	err := os.Rename(old, new)
	if err != nil {
		log.Fatal(err)
	}
}

func mkFile(fpath string) {
	// check & make dir
	dir := filepath.Dir(fpath)
	_, err := os.Stat(fpath)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	// create file
	file, err := os.Create(fpath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(file)
}

func writeBytes(f *os.File, bytes []byte) {
	_, err := f.Write(bytes)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeString(f *os.File, content string) {
	_, err := f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func buffWriter(f *os.File)  {
	fmt.Println("*** buffered writer")
	//f,_ := os.Create(fpath)
	//defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	n, _ :=w.Write([]byte("buffered bytes\r"))
	fmt.Println("write", n, "bytes into file")
	n, _ = w.WriteString("buffered string\r")
	w.Flush()
	fmt.Println("write", n, "bytes into file")
}

func checkWritePermission(fpath string) {
	//Write permission
	file, err := os.OpenFile(fpath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()
}