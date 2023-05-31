package main

import "fmt"

func oddEvenDetector(cha chan int, res chan bool) {
	num := <-cha

	if num%2 == 0 {
		res <- true

	} else {
		res <- false
	} //closing the channel to indicate that the goroutine has completed its execution and no further data is available to be sent to the channels
	close(res)

}

func Problem10() {

	fmt.Println("Problem10")

	fmt.Println()

	var a int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&a)

	//creating a channel
	ch := make(chan int)
	res := make(chan bool)

	//sending data to channel
	go oddEvenDetector(ch, res)
	ch <- a

	//receiving data from channel
	if <-res {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
