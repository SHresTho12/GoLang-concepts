package main

import (
	"fmt"
	"time"
)

func printSquares(ch chan int) {
	//as we are using a chanel it will clock the execution here and wait for the data to be sent to the channel
	value := <-ch
	// if !ok {
	// 	fmt.Println("Channel closed")
	// 	return
	// }
	fmt.Print("Square of ", value, " is :")
	fmt.Println(value * value)

}
func Problem2() {
	fmt.Println("Problem2")
	fmt.Println()

	//creating a channel
	ch := make(chan int)

	var data int
	fmt.Println("Enter data")
	fmt.Scanln(&data)

	//sending data to channel
	go printSquares(ch)
	ch <- data

	//using sleep to top the execution of the main go routine for a while so that the other go routine can execute
	time.Sleep(1 * time.Second)

}
