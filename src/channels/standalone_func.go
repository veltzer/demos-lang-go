package main

import (
	"fmt"
	"time"
)

// Worker function defined outside of main
func sendMessage(ch chan string) {
	time.Sleep(time.Second)
	ch <- "Hello from goroutine!"
}

func main() {
	// Create a channel
	messages := make(chan string)

	// Start a goroutine with the separate function
	go sendMessage(messages)

	// Receive the message from the channel
	msg := <-messages
	fmt.Println(msg)
}
