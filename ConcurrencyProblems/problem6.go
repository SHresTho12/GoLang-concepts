package main

import (
	"fmt"
)

func printEvenNumbers(ch chan int) {
	for i := 0; i <= 100; i += 2 {
		ch <- i
	}
	close(ch)
}
func printOddNumbers(ch chan int) {
	for i := 1; i <= 100; i += 2 {
		ch <- i
	}
	close(ch)
}
func Problem6() {
	fmt.Println("Problem6")
	fmt.Println()

	evenCh := make(chan int)
	oddCh := make(chan int)

	// run both function as goroutines
	go printEvenNumbers(evenCh)
	go printOddNumbers(oddCh)
	//we are using the blocking nature of the channel to make sure that the execution of the code is blocked until the channel is closed
	// as a result the code will first receive the even number and then the odd number so it turns the printing of the numbers into a synchronized process
	// if we did not use the blocking nature of the channel the execution of the code will be random and the numbers will be printed in a random order
	//approach 1
	// for {
	// 	even, evenOk := <-evenCh
	// 	if evenOk {
	// 		fmt.Println(even)
	// 	}
	// 	odd, oddOk := <-oddCh
	// 	if oddOk {

	// 		fmt.Println(odd)
	// 	}
	// }

	//approach 2
	// we can also use the select statement to make sure that the execution of the code is blocked until the channel is closed
	for {
		select {
		case even, evenOk := <-evenCh:
			if evenOk {
				fmt.Println(even)
			}
		case odd, oddOk := <-oddCh:
			if oddOk {

				fmt.Println(odd)
			}
		}
	}
}
