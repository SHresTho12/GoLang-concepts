package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Page   int      `json:"person-name"`
	Fruits []string `json:"fruits-Like"`
}

//serialization in go is the process of converting a struct to json or a format to store it on the server, disk , database etc.
//json is the most used format for serialization

func Problem3() {

	person1 := person{
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}

	jsonFormat, _ := json.Marshal(person1)
	fmt.Println(string(jsonFormat))

}
