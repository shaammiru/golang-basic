package main

import "fmt"

func main() {
	// Function without parameter
	sayHello()

	// Function with parameter
	sum(10, 40)

	// Function with return value
	randDate := getRandomInt()
	fmt.Println(randDate)
}