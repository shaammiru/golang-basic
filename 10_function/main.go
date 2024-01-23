package main

import "fmt"

func main() {
	// Function without parameter
	sayHello()

	// Function with parameter
	addTwoNumber(10, 40)

	// Function with return value
	randDate := getRandomInt()
	fmt.Println(randDate)

	// Multiple return function
	firstName, lastName := getFullName()
	city, country := getCityCountry()
	fmt.Println(firstName, lastName)
	fmt.Println(city, country)

	// Variadic function
	totalSum := sum(10, 20, 30, 40, 50)
	fmt.Println(totalSum)

	// Variadic function with slice
	sliceNumbers := []int{10, 20, 30, 40, 50}
	totalSum = sum(sliceNumbers...)
	fmt.Println(totalSum)

	// Closure as function variable
	closureAsVar()

	// Immediately-invoked function expression (IIFE)
	iife()

	// Closure as return value
	getName := closureAsReturnValue()
	fmt.Println(getName())

	// Function as parameter
	functionAsParameter("John", func(name string) {
		fmt.Println("Hello", name)
	})
}
