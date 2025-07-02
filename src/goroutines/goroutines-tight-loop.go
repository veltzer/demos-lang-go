package main

import (
	"fmt"
	"sync"
	"runtime"
)

func worker(id int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	sum := 0
	fmt.Printf("Worker %d: Starting\n", id)
	
	// Simple tight loop doing addition
	for i := 0; i < iterations * 1000; i++ {
		sum += i
	}
	
	fmt.Printf("Worker %d: Finished. Sum: %d\n", id, sum)
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	// Start three workers with different workloads
	wg.Add(3)
	go worker(1, 10000000, &wg)   // 10 million iterations
	go worker(2, 20000000, &wg)   // 20 million iterations
	go worker(3, 30000000, &wg)   // 30 million iterations
	fmt.Println("Main: All workers started")
	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("Main: All workers completed")
}
