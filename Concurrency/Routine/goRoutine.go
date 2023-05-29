package routine

import (
	"fmt"
	"time"
)

func printNumbers1(num int) {

	for i := 1; i <= num; i++ {
		fmt.Println(i) // Print the current number
	}
}

func GoRoutine() {

	num := 10

	go printNumbers1(num) // Start a new goroutine to print numbers concurrently

	time.Sleep(1 * time.Second) //blocking the main routine so that the printNumbers goroutine can finish printing the numbers

	/*
		while running this program, you will notice that the main routine exits before the printNumbers goroutine finishes printing the numbers.
		This is because the main routine does not wait for the printNumbers goroutine to finish executing.
	*/

	fmt.Println("Main routine finished executing")

}
