package main

import (
	"fmt"
	"time"
	"sync"
)

// First goroutine that generates data
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: sending %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // Signal that no more data will be sent
	fmt.Println("Producer: all data sent, channel closed")
}

// Second goroutine that processes data without using range
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when function returns
	
	// Loop until the channel is closed
	for {
		num, ok := <-ch
		if !ok {
			// Channel is closed
			break
		}
		fmt.Printf("Consumer: received %d\n", num)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Consumer: finished processing")
}

func main() {
	dataChannel := make(chan int)
	var wg sync.WaitGroup
	
	// Add 1 to the wait group for the consumer
	wg.Add(1)
	
	// Launch both goroutines
	go producer(dataChannel)
	go consumer(dataChannel, &wg)
	
	// Wait for the consumer to finish
	wg.Wait()
	fmt.Println("Main: program complete")
}
