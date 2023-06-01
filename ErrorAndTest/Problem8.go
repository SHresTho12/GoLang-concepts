package main

import (
	"fmt"
	"math"
)

func getRoot(num int) (float64, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getRoot the number is negative", r)
		}
	}()

	if num < 0 {
		panic("number cannot be negative")
	}

	root := math.Sqrt(float64(num))
	return root, nil

}

func Problem8() {

	fmt.Println("Problem8")
	fmt.Println()

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	res, err := getRoot(num)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}

}
