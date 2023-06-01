package main

import (
	"fmt"
	"os"
)

//This function takes the name / path of the file we are trying to read and returns the data in the file as a string
//if the file is not opened or some error occurs, the program panics so to inidicate the error we are using the recover function
// if we did not use the recover function we will just see the program ending unexpectedly

func ReadFile(filename string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unable to read the file ", err)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		panic(err)

	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	return string(data[:count]), nil
}

func Problem7() {

	fmt.Println(
		"Problem7",
	)
	fmt.Println("---------------------------------------------")

	data, err := ReadFile("test.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Data: ", data)
	}

}
