package main

import (
	"fmt"
	"time"
)

func unbufferedChannel() {
	messages := make(chan string)
	cars := []string{"Honda", "Infiniti", "Toyota", "Lexus", "Renault"}

	for _, each := range cars {
		go func(car string) {
			data := fmt.Sprintf("Car: %s", car)
			messages <- data
		}(each)
	}

	for i := 0; i < len(cars); i++ {
		fmt.Println(<-messages)
	}
}

func bufferedChannel() {
	messages := make(chan string, 2)
	cars := []string{"Honda", "Infiniti", "Toyota", "Lexus", "Renault"}

	go func() {
		for {
			// Blocking
			fmt.Println("Receive data", <-messages)
		}
	}()

	for _, each := range cars {
		fmt.Println("Send data", each)
		// Non-blocking until buffer is full
		messages <- each
	}

	time.Sleep(1 * time.Nanosecond)
}
