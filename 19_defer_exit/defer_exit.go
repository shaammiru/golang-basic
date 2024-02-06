package main

import (
	"fmt"
	"os"
)

func deferExample() {
	// Execute defer after the function is finished
	defer fmt.Println("This is defer line")
	fmt.Println("This is the first line")
	fmt.Println("Hello from deferExample")
}

func iifeDeferExample() {
	func() {
		defer fmt.Println("This is defer line from IIFE")
		fmt.Println("Hello from iifeDeferExample")
	}()

	fmt.Println("This is the last line from iifeDeferExample")
}

func exitExample() {
	fmt.Println("This is the first line from exitExample")
	os.Exit(1)
	fmt.Println("This line will not be executed")
}
