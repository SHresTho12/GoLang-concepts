package routine

import (
	"fmt"
	"time"
)

func printNumbers(num int) {

	for i := 1; i <= num; i++ {
		fmt.Println(i) // Print the current number
	}
}
func printLetters(char rune) {

	for i := 'a'; i <= char; i++ {
		fmt.Printf("%c\n", i) // Print the current letter
	}
}

func MultipleGoRoutine() {

	num := 10
	char := 'j'

	go printNumbers(num)  // Start a new goroutine to print numbers concurrently
	go printLetters(char) // Start a new goroutine to print letters concurrently

	time.Sleep(1 * time.Second) //blocking the main routine so that the printNumbers goroutine can finish printing the numbers

	/*
		THE OUTPUT IS NOT DETERMINISTIC AND MAY VARY FROM RUN TO RUN
		1 2 3 4 5 6 7 8 9 10 a b c d e f g h i j
		1 a 2 b 3 c 4 d 5 e 6 f 7 g 8 h 9 i 10 j
		as both the functions are running concurrently not sequentially
	*/

	fmt.Println("Main routine finished executing")

}
