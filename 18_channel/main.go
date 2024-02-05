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
	fmt.Println(">> Buffered Channel Example <<")
	bufferedChannel()
}

