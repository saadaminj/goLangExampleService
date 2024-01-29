package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			select {
			case value, ok := <-ch1:
				if !ok {
					fmt.Println("ch1 is closed. Exiting.")
					return
				}
				fmt.Printf("Received from ch1: %d\n", value)
			case message, ok := <-ch2:
				if !ok {
					fmt.Println("ch2 is closed. Exiting.")
					return
				}
				fmt.Printf("Received from ch2: %s\n", message)
			case <-time.After(2 * time.Second):
				fmt.Println("Timeout: No communication in the last 2 seconds")
			default:
				// Non-blocking case
				fmt.Println("No communication ready.")
			}
		}
	}()

	// Simulate sending values and messages to the channels
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1) // Close ch1 after sending values
	}()

	go func() {
		for i := 0; i < 2; i++ {
			ch2 <- fmt.Sprintf("Message %d", i+1)
			time.Sleep(2 * time.Second)
		}
		close(ch2) // Close ch2 when done
	}()

	// Wait for goroutines to finish
	time.Sleep(10 * time.Second)
}
