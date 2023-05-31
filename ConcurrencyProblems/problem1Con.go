package main

import (
	"fmt"
	"time"
)

//as one function generates data and another receives it, we need to make sure that the receiver function is ready to receive data before the sender function sends it.
//This is where channels come into play. Channels are used to synchronize data flow between two goroutines.
//A channel is a communication pipe between goroutines which is used to pass data.
//if the chanel was about to receive a data it blocks the execution of the code until it receives the data.
//The data is sent to the channel using the <- operator.

func dataSend(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}

}
func dataReceive(ch chan int) {

	for i := 0; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}

func Chanels() {
	//creating a channel
	ch := make(chan int)
	//sending data to channel
	go dataSend(ch)
	//receiving data from channel
	go dataReceive(ch)
}

//if we did not use channels to send and receive data the exchange between two routines will be impossible

func Problem1() {
	fmt.Println("Problem1")
	fmt.Println()

	Chanels()
	time.Sleep(1 * time.Second)

}
