package main

import (
	"encoding/json"
	"fmt"
)

func Problem2() {

	type person struct {
		Page   int      `json:"person-name"`
		Fruits []string `json:"fruits-Like"`
	}
	str := `
{
    "person-name": 1,
    "fruits-Like": ["apple", "peach"]
}`
	res := person{}

	//turing the json in to struct
	//the struct that we are converting to we must keep an eye on the json tags because if the json tags are not same as the struct tags then the json will not be converted to struct
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

}
