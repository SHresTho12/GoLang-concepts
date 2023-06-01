package main

import (
	"fmt"
)

func Problem1() {

	// implements defer function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error occured and recovered in Problem1: ", err)
		}
	}()

	fmt.Println("Starting program")
	var myArray [3]int

	//this code does not trigger the recover function as error
	/*
		accessing an array or slice with an index that is out of bounds directly (e.g., myArray[3]) will not cause a runtime panic, which cannot be recovered using the defer and recover mechanism. This type of panic is not recoverable because it occurs at the point of access, before any deferred functions can execute. However, Go does not raise a runtime panic for this scenario. It simply results in undefined behavior, and the program may exhibit unpredictable or inconsistent results.

		so the given code snippet will not trigger the recover function it should be fixed
	*/

	var i = 3
	//myArray[3] = i
	myArray[i] = 1 // this will trigger the recover function

	//also if i use a loop to access the array out of bounds then it will trigger the recover function

	// for i := 0; i < 5; i++ {
	// 	myArray[i] = i
	// }
	fmt.Println("Program finished.")

}
