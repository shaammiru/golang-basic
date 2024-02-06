package main

import (
	"fmt"
	"math/rand"
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

	time.Sleep(1 * time.Millisecond)
}

func selectChannel() {
	numbers := []int{3, 1, 9, 7, 2, 3, 5, 2}
	fmt.Println("Numbers :", numbers)

	ch1 := make(chan float64)
	go func(numbers []int, ch chan float64) {
		sum := 0
		for _, e := range numbers {
			sum += e
		}
		ch <- float64(sum) / float64(len(numbers))
	}(numbers, ch1)

	ch2 := make(chan int)
	go func(numbers []int, ch chan int) {
		max := numbers[0]
		for _, e := range numbers {
			if max < e {
				max = e
			}
		}
		ch <- max
	}(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avgCh := <-ch1:
			fmt.Printf("Avg : %.2f\n", avgCh)
		case maxCh := <-ch2:
			fmt.Printf("Max : %d\n", maxCh)
		}
	}
}

func rangeCloseChannel() {
	messages := make(chan string)
	// chan<- untuk mengirim data ke channel (write-only)
	// jadi fungsi ini hanya bisa mengirim data ke channel
	go func(ch chan<- string) {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Message %d", i)
		}
		close(ch)
	}(messages)

	// <-chan untuk menerima data dari channel (read-only)
	// jadi fungsi ini hanya bisa menerima data dari channel
	func(ch <-chan string) {
		for msg := range ch {
			fmt.Println(msg)
		}
	}(messages)
}

func channelTimeout() {
	var messages = make(chan int)

	go func(ch chan<- int) {
		randomizer := rand.New(rand.NewSource(time.Now().Unix()))

		for i := 0; true; i++ {
			ch <- i
			time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
		}
	}(messages)

loop:
	for {
		select {
		case data := <-messages:
			fmt.Println("Receive data", data)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout! No activities for 5 seconds")
			break loop
		}
	}
}
