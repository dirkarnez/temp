package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("hello world v8")
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt...\n\n")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
