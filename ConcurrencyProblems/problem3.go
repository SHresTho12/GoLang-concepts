package main

import (
	"fmt"
	"time"
)

// using buffered channels enables us to send data to the channel without blocking the execution of the code
// we can send data to the channel until the buffer is full and then the execution will be blocked
// we can receive data from the channel until the buffer is empty and then the execution will be blocked
// It helps us to simulate asynchronous execution of the code.
func dataSend2(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}

}
func dataReceive2(ch chan int) {

	for i := 0; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}

func Problem3() {

	fmt.Println("Problem3")
	fmt.Println()

	//creating a channel
	ch := make(chan int, 5)
	//sending data to channel
	go dataSend2(ch)
	//receiving data from channel
	go dataReceive2(ch)
	time.Sleep(1 * time.Second)
}
