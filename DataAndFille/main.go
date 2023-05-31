package main

import (
	"fmt"

	"encoding/json"
)

func main() {
	var boolean bool = true
	bolB, _ := json.Marshal(boolean)
	fmt.Println(bolB)

	for i := range bolB {
		fmt.Print(string(bolB[i]))
	}

}
