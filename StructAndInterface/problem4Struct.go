package main

import "fmt"

type Location4 struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Post4 struct {
	Category string    `json:"category"`
	Time     int64     `json:"time"`
	Location Location4 `json:"location"`
}

//The problem was that the receiver was not a pointer type it was a value type so the change was not reflected in the original struct
//that is why the location was not being set
//The problem is solved simply by changing the receiver to a pointer type

func (post *Post4) setLocation() {
	post.Location = Location4{
		Latitude:  23.99,
		Longitude: 90.45,
	}
}

func Problem4() {
	post := Post4{}
	post.setLocation()
	fmt.Println(post.Location)
}
