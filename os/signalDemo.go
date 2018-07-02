package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

/**
test steps:
1. go run os/signalDemo.go
2. ctrl+c
 */
func main()  {
	done := make(chan bool, 1)
	// register channel on listening specific signals
	// syscall.SIGINT: stop input
	// syscall.SIGTERM: shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
