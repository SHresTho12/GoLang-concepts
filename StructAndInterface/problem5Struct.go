package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Salary int
}

// The way to do it is just simply pass the pointer to the struct
// This way the change will be reflected in the original struct
// If we don't pass the pointer then the change will not be reflected in the original struct and we have to initialize the struct with the new value
func (e *Employee) GiveRaise() {
	e.Salary += 1000
}

func Problem5() {

	fmt.Printf("\n\nProblem 5 solution starts\n\n")

	emp := Employee{
		Name:   "John Doe",
		Age:    30,
		Salary: 3000,
	}

	fmt.Println("Before Raise")
	fmt.Println(emp)

	emp.GiveRaise()

	fmt.Println("After Raise")
	fmt.Println(emp)

}
