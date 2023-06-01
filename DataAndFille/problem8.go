package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Problem8() {
	// Create a map of data
	data := map[string]interface{}{
		"name":    "Raisul",
		"age":     23,
		"country": "BD",
	}
	// Create a new JSON file
	file, _ := os.Create("data.json")
	defer file.Close()
	// Create a new JSON encoder
	//using new encoder is much more efficient than using marshal and also easy to use
	encoder := json.NewEncoder(file)

	// Encode the data to JSON and write it to the file
	encoder.Encode(data)

	fmt.Println("JSON file generated successfully.")
}
