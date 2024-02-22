package main

import "fmt"

/* Timer, Ticker, Shceduler
- Ada beberapa fungsi dalam package time yang bisa digunakan untuk mengatur waktu
- Timer: digunakan untuk mengatur waktu tunggu sebelum menjalankan suatu fungsi
- Ticker: digunakan untuk menjalankan suatu fungsi secara berkala
- Scheduler: digunakan untuk menjalankan suatu fungsi secara berkala dengan waktu yang sudah ditentukan
*/

func main() {
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
