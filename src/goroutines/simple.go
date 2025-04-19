package main

import (
	"fmt"
	"time"
)

// Function to be executed by multiple goroutines
func worker(id int) {
	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d: Finished\n", id)
}

func main() {
	// Launch three goroutines running the same function
	go worker(1)
	go worker(2)
	go worker(3)
	
	// Print from main
	fmt.Println("Main: All workers started")
	
	// Wait for all goroutines to complete
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Main: Exiting")
}
