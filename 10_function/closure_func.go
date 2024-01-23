package main

import "fmt"

/* Closure Function
Closure function adalah function yang bisa mengembalikan function.
Function yang dikembalikan bisa berupa anonymous function.
*/

// Closure as function variable
func closureAsVar() {
	getName := func() string {
		return "John"
	}

	fmt.Println(getName())
}

// Immediately-invoked function expression (IIFE)
func iife() {
	getName := func() string {
		return "John"
	}()

	fmt.Println(getName)
}

// Closure as return value
func closureAsReturnValue() func() string {
	return func() string {
		return "John Wick"
	}
}

// Alias Closure Schema
type Callback func(string)

// Function as parameter
func functionAsParameter(name string, callback Callback) {
	callback(name)
}
