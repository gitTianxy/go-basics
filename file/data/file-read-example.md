# Go by Example: Reading Files
Reading and writing files are basic tasks needed for many Go programs. First we’ll look at some examples of reading files.

## check error
* Reading files requires checking most calls for errors. This helper will streamline our error checks below.
```go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```

## read whole
* Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
```go
dat, err := ioutil.ReadFile("/tmp/dat")
check(err)
fmt.Print(string(dat))
```

## read by part
* You’ll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.
```go
f, err := os.Open("/tmp/dat")
check(err)
```

* Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read.
```go
b1 := make([]byte, 5)
n1, err := f.Read(b1)
check(err)
fmt.Printf("%d bytes: %s\n", n1, string(b1))
```

* You can also Seek to a known location in the file and Read from there.
```go
o2, err := f.Seek(6, 0)
check(err)
b2 := make([]byte, 2)
n2, err := f.Read(b2)
check(err)
fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
```

* The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.
```go
r4 := bufio.NewReader(f)
b4, err := r4.Peek(5)
check(err)
fmt.Printf("5 bytes: %s\n", string(b4))
```

## close after reading
* Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
```go
f.Close()
```
