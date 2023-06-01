package main

import (
	"fmt"
)

// error is a built in interface in go
type customError string

// we can define the error here as our need add methods to it and also additional information

func (c customError) Error() string {
	return string(c)
}

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, customError("cannot divide by zero")
	}

	return a / b, nil
}

func Problem3() {

	fmt.Println("Problem3")
	fmt.Println("")

	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&a, &b)

	res, err := divide(a, b)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}

}
