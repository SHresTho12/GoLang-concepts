package main

import (
	"fmt"
	"time"
)

// problem 1
type location struct {
	Lattiude  float64
	Longitude float64
}
type post struct {
	ID       int
	Catagory string
	Location location
	Time     time.Time
}

func Problem1Struct() {
	//as both struct have same fields we can use use another struct as a field to use in both feilds. This is called composition

	fmt.Println()
	fmt.Println("Problem 1 solution starts")
	fmt.Println()

	loc := location{
		Lattiude:  37.42,
		Longitude: -122.08,
	}

	post1 := post{
		ID:       1,
		Catagory: "Technology",
		Location: loc,
		Time:     time.Now(),
	}
	post2 := post{
		ID:       2,
		Catagory: "Technology",
		Location: loc,
		Time:     time.Now(),
	}

	fmt.Println(post1)
	fmt.Println(post2)

}
