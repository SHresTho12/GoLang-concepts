package main

import "fmt"

func Problem1() {
	//to allocate a meomory address to a variable we use new() function here we dont know what will be  the value of the variable.
	num1 := new(int)
	num2 := new(int)

	fmt.Println("Enter first number: ")
	fmt.Scanln(num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(num2)

	sum := *num1 + *num2

	fmt.Println("Sum of two numbers is: ", sum)

}
