package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	// golang中文件的相对路径从项目底下(即working-directory)开始算
	fpath := "file/data/file-read-example.md"

	checkReadPermission(fpath)

	readByBytes(fpath, 5)
	readByLine(fpath)
	readAll(fpath)

	wlkRoot := "file"
	listFiles(wlkRoot)
	walk(wlkRoot)
	filepathWalk(wlkRoot)
	getDir(fpath)
	exists(fpath)

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readByBytes(fpath string, bytes int) {
	fmt.Println("*** read by bytes")
	f,_ := os.Open(fpath)
	defer f.Close()
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(bytes)
	checkErr(err)
	fmt.Printf("%v bytes: %s\n", bytes, string(b4))
}

func readByLine(fpath string)  {
	fmt.Println("*** read by line")
	f,_ := os.Open(fpath)
	defer f.Close()
	buff := bufio.NewReader(f)
	for i:=0; ;i++  {
		line, err := buff.ReadBytes('\n')
		// 文件已经到达结尾
		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("%d line: %s", i, string(line))
	}
}

func readAll(fpath string)  {
	fmt.Println("*** read all")
	dat,_ := ioutil.ReadFile(fpath)
	fmt.Println(string(dat))
}

func listFiles(root string) {
	fmt.Println("*** files under", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
		}
	}
}

func filepathWalk(rt string) {
	fmt.Println("*** filepath.Walk(", rt, ")")
	err := filepath.Walk(rt, walkFunc)
	checkErr(err)
}

func walkFunc(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fmt.Printf("Visited: %s(size:%v)\n", path, f.Size())
	}
	return nil
}

func walk(rt string) {
	fmt.Println("*** walk through", rt)
	fs, err := ioutil.ReadDir(rt)
	checkErr(err)
	for _, f := range fs {
		if f.IsDir() {
			walk(rt + "/" + f.Name())
		} else {
			fmt.Println(rt + "/" + f.Name())
		}
	}
}

func getDir(fpath string) string {
	fmt.Println("*** dir of", fpath)
	dir := filepath.Dir(fpath)
	fmt.Println(dir)
	return dir
}

func exists(fpath string) {
	fmt.Println("***", fpath, "exists")
	var res bool
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			res = false
		}
	}
	res = true
	fmt.Println(res)
}

func checkReadPermission(fpath string) {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}
