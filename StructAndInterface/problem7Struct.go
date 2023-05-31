package main

import "fmt"

func printValues(values interface{}) {

	fmt.Printf("(%v,%T)", values, values)

}

// as we are using an empty interface the type of values can be anything and each type is being printed using the %v and %T format specifiers
func Problem7() {
	fmt.Println()
	fmt.Println("Problem 7 solution starts")
	fmt.Println()

	var values interface{} = 42

	printValues(values)

	values = "abc"
	printValues(values)
}
