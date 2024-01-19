package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Function
Function adalah blok kode yang diberi nama dan bisa dipanggil di bagian lain program.
Function biasanya digunakan untuk mengelompokkan kode agar lebih rapi dan mudah di-maintain.
*/

// Function without parameter
func sayHello() {
	fmt.Println("Hello")
}

// Function with parameter
// func sum(first, second int)
func addTwoNumber(first int, second int) {
	fmt.Println(first + second)
}

// Function with return value
func getRandomInt() int {
	return rand.New(rand.NewSource(time.Now().Unix())).Int()
}