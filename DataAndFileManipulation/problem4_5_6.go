package main

import (
	"fmt"
	"os"
)

//there is also a alternative way to create a file using the ioutil package

func Problem4_5_6() {

	fmt.Println("Problem 4 starts here")
	fmt.Println()

	//creating a file

	fille, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fille.Close()

	//writing to a file

	data := []byte("Hello World")
	_, err = fille.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	//opening a file

	file, err2 := os.Open("example.txt")
	testingData := make([]byte, 100)
	count, err := file.Read(testingData)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("Number of bytes read: ", count)
	fmt.Println("Data read: ", string(testingData))

	//deleting a file

	fileDeleteError := os.Remove("example.txt")
	if fileDeleteError != nil {
		fmt.Println(fileDeleteError)
		return
	}
	fmt.Println("File deleted successfully")

}
