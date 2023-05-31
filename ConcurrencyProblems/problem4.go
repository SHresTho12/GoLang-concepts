package main

import (
	"fmt"
	"sync"
	"time"
)

// wait groups are used to wait for a collection of goroutines to finish executing
// it is a alternative to the time.Sleep() function.
func downloadFile(fileNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Download started for file", fileNum)
	//for now we are assuming that the download takes 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Download finished for file", fileNum)
}

// use of wait groups enables us to run the download function for multiple files at the same time and make the main function to wait for each of them to finish
func Problem4() {
	fmt.Println("Problem4")
	fmt.Println()

	var wg sync.WaitGroup

	//number of files to be downloaded
	numFiles := 4
	wg.Add(numFiles)

	for i := 0; i < numFiles; i++ {
		//for each file create a go routine
		go downloadFile(i, &wg)
	}

	wg.Wait()

	fmt.Println("All files downloaded successfully " + fmt.Sprint(numFiles) + " files")

}
