package main

import (
	"fmt"
	"time"
)

// First goroutine that generates data
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: sending %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // Signal that no more data will be sent
	fmt.Println("Producer: channel closed")
}

// Second goroutine that processes data
func consumer(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Printf("Consumer: received %d\n", num)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Consumer: finished")
	done <- true // Signal that consumer is done
}

func main() {
	dataChannel := make(chan int)
	doneChannel := make(chan bool)
	
	// Launch both goroutines
	go producer(dataChannel)
	go consumer(dataChannel, doneChannel)
	
	// Wait for the consumer to finish
	<-doneChannel
	fmt.Println("Main: program complete")
}
