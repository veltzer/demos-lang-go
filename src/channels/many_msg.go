package main

import (
	"fmt"
	"time"
)

// Worker function defined outside of main
func sendMessages(ch chan string, id int) {
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Message %d from goroutine %d", i, id)
		time.Sleep(100 * time.Millisecond)
		ch <- message
	}
}

func main() {
	// Create a channel
	messages := make(chan string)
	
	// Start a goroutine with the separate function
	go sendMessages(messages, 1)
	
	// Receive all ten messages from the channel
	for i := 1; i <= 10; i++ {
		msg := <-messages
		fmt.Println(msg)
	}
	
	fmt.Println("All messages received!")
}
