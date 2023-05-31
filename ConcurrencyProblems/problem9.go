package main

import (
	"fmt"
	"sort"
	"time"
)

//this can be done using two goroutines but the problem is that the execution of the code will not be in order
//there are multiple implementations for this problem using go routine and channel to implement the merge sort algorithm

func sortArray(arr [10]int, ch chan []int) {
	//sorting the array
	sort.Ints(arr[:])
	//sending the sorted array to the channel
	ch <- arr[:]
}

func Problem9() {

	fmt.Println("Problem9")
	fmt.Println()
	arr1 := [10]int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}

	ch := make(chan []int)
	go sortArray(arr1, ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)

}

//there is a sample code i found on github that implements the merge sort algorithm using go routine and channel
/*

func Merge(ldata []float64, rdata []float64) (result []float64) {
	result = make([]float64, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}


func MultiMergeSort(data []float64, res chan []float64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []float64)
	rightChan := make(chan []float64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
  go MultiMergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(ldata, rdata)
	return
}

func RunMultiMergeSort(data []float64) (multiResult []float64) {
	res := make(chan []float64)
	go MultiMergeSort(data, res)
	multiResult = <-res
	return
}

*/
