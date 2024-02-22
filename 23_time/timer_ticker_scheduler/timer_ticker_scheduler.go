package timer_ticker_scheduler

import (
	"fmt"
	"os"
	"time"
)

/* Timer, Ticker, Shceduler
- Ada beberapa fungsi dalam package time yang bisa digunakan untuk mengatur waktu
- Timer: digunakan untuk mengatur waktu tunggu sebelum menjalankan suatu fungsi
- Ticker: digunakan untuk menjalankan suatu fungsi secara berkala
- Scheduler: digunakan untuk menjalankan suatu fungsi secara berkala dengan waktu yang sudah ditentukan
*/

func init() {
	// Sleep
	fmt.Println(">> Sleep <<")
	timeSleep()
	fmt.Println()

	// Timer
	fmt.Println(">> Timer <<")
	timer()
	fmt.Println()

	// Timer AfterFunc
	fmt.Println(">> Timer AfterFunc <<")
	timerAfterFunc()
	fmt.Println()

	// Time After
	fmt.Println(">> Time After <<")
	timeAfter()
	fmt.Println()

	// Ticker Scheduler
	fmt.Println(">> Ticker Scheduler <<")
	tickerScheduler()
	fmt.Println()

	// Timer and Goroutine
	fmt.Println(">> Timer and Goroutine <<")
	timerAndGoroutine()
	fmt.Println()
}

func timeSleep() {
	fmt.Println("Start")
	time.Sleep(time.Second * 2)
	fmt.Println("End")
}

func timer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Start")
	<-timer.C
	fmt.Println("End")
}

func timerAfterFunc() {
	ch := make(chan bool)

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Expired")
		ch <- true
	})

	fmt.Println("Start")
	<-ch
	fmt.Println("Finish")
}

func timeAfter() {
	<-time.After(2 * time.Second)
	fmt.Println("Expired")
}

func tickerScheduler() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello!!", t)
		}
	}
}

func timerAndGoroutine() {
	timeout := time.Second * 5
	ch := make(chan bool)

	go func(timeout time.Duration, ch chan<- bool) {
		time.AfterFunc(timeout, func() {
			ch <- true
		})
	}(timeout, ch)

	go func(timeout time.Duration, ch <-chan bool) {
		<-ch
		fmt.Println("\nTimeout!, no answer more than", timeout)
		os.Exit(0)
	}(timeout, ch)

	var input string
	fmt.Print("What is 1 + 1? ")
	fmt.Scanln(&input)

	if input == "2" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong!")
	}
}
