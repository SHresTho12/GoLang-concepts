package main

import (
	"fmt"
	"os"
)

// we can return multiple values from a function in go here we are returning an int and an error
// the error will indicate that something went wrong
func addition(num1 int, num2 int) (int, error) {

	result := num1 + num2
	return result, nil

}

func Problem2() {

	fmt.Println("Problem2")

	fmt.Println()

	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&a, &b)

	res, err := addition(a, b)

	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	fmt.Println("Result: ", res)

}
