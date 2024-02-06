package main

import (
	"fmt"
	"runtime"
)

/* Channel
- Channel adalah cara komunikasi antar goroutine
- Channel berfungsi untuk mengirim dan menerima data antar goroutine
- Channel bersifat synchronous, artinya jika kita mengirim data ke channel, maka goroutine lain akan menunggu hingga data tersebut dibaca
*/

func main() {
	runtime.GOMAXPROCS(2)

	// Channel Example
	fmt.Println(">> Unbuffered Channel Example <<")
	unbufferedChannel()

	// Buffered Channel Example
	// Buffered channel adalah channel yang memiliki buffer
	// Ketika membuat buffered channel, kita perlu menentukan kapasitas buffer
	// Ketika buffer penuh, maka channel akan menjadi blocking
	fmt.Println(">> Buffered Channel Example <<")
	bufferedChannel()

	// Select Channel Example
	// Select adalah cara untuk melakukan komunikasi dengan channel secara non-blocking
	fmt.Println(">> Select Channel Example <<")
	selectChannel()

	// Range Close Channel Example
	// Range adalah cara untuk melakukan iterasi data dari channel
	// Close adalah cara untuk memberitahukan bahwa channel sudah selesai digunakan
	fmt.Println(">> Range Close Channel Example <<")
	rangeCloseChannel()

	// Channel Timeout Example
	// Timeout adalah cara untuk menghentikan proses yang terlalu lama
	fmt.Println(">> Channel Timeout Example <<")
	channelTimeout()
}
