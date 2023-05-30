package main

import (
	"fmt"
)

func dataSend(ch chan int) {
	ch <- 10
}
func dataReceive(ch chan int) {
	fmt.Println(<-ch)
}

func Chanels() {
	//creating a channel
	ch := make(chan int)
	//sending data to channel
	go func() {
		ch <- 10
	}()
	//receiving data from channel
	fmt.Println(<-ch)
}
func ChanelsBetweenRoutie() {
	//creating a channel
	ch := make(chan int)
	//sending data to channel
	go dataSend(ch)
	//receiving data from channel
	dataReceive(ch)

}
func main() {
	Chanels()
	ChanelsBetweenRoutie()
}
