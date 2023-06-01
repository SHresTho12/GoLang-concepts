package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func Problem7() {
	//creating a slice of struct

	fmt.Println("Problem 7 ")
	fmt.Println()

	people := []Person{
		{Name: "Mahdi vai", Age: 50, Country: "India"},
		{Name: "Mayank", Age: 20, Country: "India"},
		{Name: "Rahul", Age: 30, Country: "USA"},
		{Name: "Vijay", Age: 40, Country: "France"},
	}

	file, err := os.Create("people.csv")
	if err != nil {
		return
	}
	defer file.Close()

	//writing to a file
	for _, person := range people {
		//I could not find out how to write the age in the file as a number so I converted it to rune
		//but it works only for the first one and the rest of the age is not written in the file
		//file.WriteString(person.Name + "," + string(rune(person.Age)) + "," + person.Country + "\n")

		//the solution is to convert the int to string separately

		ageStr := strconv.Itoa(person.Age) // Convert age to string
		file.WriteString(person.Name + "," + ageStr + "," + person.Country + "\n")
	}
	//reading from a file
	readFile, eror := os.Open("people.csv")
	if eror != nil {
		fmt.Println(eror)
		return
	}

	//reading from a file
	//first initialize a slice of byte to store the data
	//then use the read method to read the data
	data := make([]byte, 100)
	count, readError := readFile.Read(data)
	if readError != nil {
		fmt.Println(readError)
	}
	fmt.Println("Number of bytes read: ", count)
	fmt.Println("Data read: ", string(data))

	defer readFile.Close()

}
