package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func errorUsage() {
	input := ""
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
	} else {
		fmt.Println(number, "is number")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func customError() {
	input := ""
	fmt.Print("Input your name: ")
	fmt.Scanln(&input)

	if valid, err := validate(input); valid {
		fmt.Println("Hello", input)
	} else {
		fmt.Println(err.Error())
	}
}

func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured:", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()

	input := ""
	fmt.Print("Input your name: ")
	fmt.Scanln(&input)

	if valid, err := validate(input); valid {
		fmt.Println("Hello", input)
	} else {
		panic(err.Error())
	}
}
