package main

import (
"fmt"
"github.com/go-vgo/robotgo"
"time"
)

func main() {
    for {
        robotgo.MouseClick("left", true)
        x, y := robotgo.GetMousePos()
        fmt.Printf("click (%v, %v)\n", x, y)
        time.Sleep(time.Minute)
    }
}
