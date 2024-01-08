package main

import "fmt"

func main() {
	// With keyword var
	var name string = "John"
	
	// With keyword var and type inference
	age := 20
	
	fmt.Println(name, age)

	// Multiple variable declaration
	var first, second string = "first", "second"
	/* 
		first, second := "first", "second"
		var (
			first string = "first"
			second string = "second"
		)
	*/
	fmt.Println(first, second)
	
	// Underscore variable
	_ = "Underscore variable"

	// Keyword new
	message := new(string)
	*message = "Hello World"

	fmt.Println(message, *message)

	// Constant
	const pi = 3.14

	fmt.Println(pi)
}