package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes

	// Simulate some work
	time.Sleep(time.Second)
	fmt.Println("Worker: Done")
}

func main() {
	var wg sync.WaitGroup

	// Start a worker goroutine
	wg.Add(1) // Increment the counter
	go worker(&wg)

	// Wait for all added goroutines to finish
	wg.Wait() // Block until the counter becomes zero

	fmt.Println("Main: All workers are done.")
}
