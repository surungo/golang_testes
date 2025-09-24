package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string) {
	for i := 0; i < 3; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		ch <- message // Send message to the channel
		fmt.Printf("Producer sent: %s\n", message)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
	close(ch) // Close the channel when done sending
}

func consumer(ch <-chan string) {
	for msg := range ch { // Loop until the channel is closed and empty
		fmt.Printf("Consumer received: %s\n", msg)
		time.Sleep(700 * time.Millisecond) // Simulate some work
	}
	fmt.Println("Consumer finished.")
}

func main() {
	// Create a buffered channel of strings with a capacity of 2
	messages := make(chan string, 2)

	// Start producer and consumer goroutines
	go producer(messages)
	go consumer(messages)

	// Wait for a bit to allow goroutines to complete their work
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine finished.")
}
