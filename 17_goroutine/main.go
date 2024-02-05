package main

import (
	"fmt"
	"runtime"
)

/* Goroutine
- Goroutine adalah cara untuk membuat proses yang berjalan secara asynchronous
- Goroutine mirip dengan thread, namun lebih efisien
*/

func main() {
	runtime.GOMAXPROCS(2)

	go sayHello(6, "Goroutine")
	sayHello(6, "Synchronous")

	var input string
	fmt.Scanln(&input)
}

func sayHello(amount int, name string) {
	for i := 0; i < amount; i++ {
		fmt.Println(i, "Hello", name, "!")
	}
}
