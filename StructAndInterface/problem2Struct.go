package main

import (
	"fmt"
	"time"
)

type locationP2 struct {
	Lattiude  float64
	Longitude float64
}
type postP2 struct {
	ID       int
	Catagory string
	Location locationP2
	Time     time.Time
}

// we are using interface to indicate that any post can have a thumbnail but we are not specifying how to generate it
// both the structs can have different implementation of GenerateThumbnail
// with out interface we would have to write two different functions for each struct
// or we could not generalize the post category and have two different structs
type ThumbnailGenerator interface {
	GenerateThumbnail()
}

func (p postP2) GenerateThumbnail() {

	if p.Catagory == "Photo" {
		fmt.Println("Generating photo thumbnail")
	} else if p.Catagory == "Video" {
		fmt.Println("Generating video thumbnail")
	} else {
		fmt.Println("No thumbnail generated")
	}
}

func Problem2() {

	fmt.Println()
	fmt.Println("Problem 2 solution starts")
	fmt.Println()

	loc := locationP2{
		Lattiude:  37.42,
		Longitude: -122.08,
	}

	photoPost := postP2{
		ID:       1,
		Catagory: "Photo",
		Location: loc,
		Time:     time.Now(),
	}
	photoPost.GenerateThumbnail()

	videoPost := postP2{
		ID:       2,
		Catagory: "Video",
		Location: loc,
		Time:     time.Now(),
	}
	//same post type struct but as the different implementation of GenerateThumbnail both the structs are different
	videoPost.GenerateThumbnail()

}
