package main

import (
	"fmt"
)

func SumOfSquares(c, quit chan int) {
	// your code here
	//this loop sends the square values sequentially  to the channel
	for i := 1; i <= 6; i++ {
		c <- i * i
	}
	close(c)
	//this indicates that all the square values have been sent to the channel and the goroutine now has completed its execution
	quit <- 0

}
func Problem7() {
	fmt.Println("Problem7")
	fmt.Println()

	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0

	go func() {
		fmt.Println("Sum of squares")
		for i := 0; i < 6; i++ {

			sum += <-mychannel

		}
		fmt.Println(sum)
		//this indicates that the main goroutine has received all the square values from the channel and the goroutine now has completed its execution
		<-quitchannel

	}()
	SumOfSquares(mychannel, quitchannel)

}
