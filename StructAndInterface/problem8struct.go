package main

import (
	"fmt"
)

// Formatter interface
type Formatter interface {
	Format([]string) string
}

// CSVFormatter struct
type CSVFormatter struct {
}

// Format method for CSVFormatter
func (csv CSVFormatter) Format(data []string) string {

	return CSVFormatterSpecificMethod(data)

}

// HTMLFormatter struct
type HTMLFormatter struct {
}

// Format method for HTMLFormatter

func (html HTMLFormatter) Format(data []string) string {

	return HTMLFormatterSpecificMethod(data)
}

// TestFormat function
func TestFormat(f Formatter, data []string) {

	//The function first checks if the formatter is CSV or it is a html formatter. So based on that we can print instruction to the user
	/* TO SHOW THE user proper instruction we used the type assertion if we did not use it then we had to add the instruction
	code in to the Format method of the CSVFormatter and HTMLFormatter struct which is not a good practice*/
	_, ok := f.(CSVFormatter)
	if !ok {
		fmt.Println(" Html Formatter")
		result := f.Format(data)
		fmt.Println(result)

	} else {
		fmt.Println(" CSV Formatter")
		result := f.Format(data)
		fmt.Println(result)
	}

}

// CSVSpecificMethod for CSVFormatt
func CSVFormatterSpecificMethod(data []string) string {
	result := ""
	for _, v := range data {
		if result != "" {
			result += "," //as the csv data is seperated by , we are adding , after each value
		}
		result += v
	}
	return result

}

// HTMLSpecificMethod for HTMLFormatter
func HTMLFormatterSpecificMethod(data []string) string {
	result := "<ul>"
	for _, v := range data {
		result += "<li>" + v + "</li>" //as the html data is seperated by <li> and </li> we are adding <li> and </li> after each value
	}
	result += "</ul>"
	return result
}

func Problem8() {

	fmt.Println("Problem 8 solution starts")
	fmt.Println()
	// Test CSV formatter
	csvFormatter := CSVFormatter{}
	TestFormat(csvFormatter, []string{"apple", "banana", "orange"})
	fmt.Println()

	// Test HTML formatter
	htmlFormatter := HTMLFormatter{}
	TestFormat(htmlFormatter, []string{"apple", "banana", "orange"})
}
