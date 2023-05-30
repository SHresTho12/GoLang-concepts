package routine

import (
	"fmt"
	"time"
)

func printNumbersClose(num int, cancel chan bool) {

	//select is switch statement for channels in go
	select {
	//if cancel channel is closed then this case will be executed and return
	case <-cancel:
		fmt.Println("Go routine cancelled")
		return
	default:
		//if cancel channel is not closed then this case will be executed
		for i := 1; i <= num; i++ {
			fmt.Println(i) // Print the current number
		}

	}

}

func CloseGoRoutine() {

	num := 10
	cancel := make(chan bool)         //creating a channel
	go printNumbersClose(num, cancel) // Start a new goroutine to print numbers concurrently

	time.Sleep(1 * time.Second) //blocking the main routine so that the printNumbers goroutine can finish printing the numbers

	//closing the channel
	close(cancel)

	/*
		while running this program, you will notice that the main routine exits before the printNumbers goroutine finishes printing the numbers.
		This is because the main routine does not wait for the printNumbers goroutine to finish executing.
	*/

	fmt.Println("Main routine finished executing")

}
