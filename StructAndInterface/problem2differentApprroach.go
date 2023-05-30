package main

import (
	"fmt"
	"time"
)

type post2Dif interface {
	GenerateThumbnail()
}

type locationP2D struct {
	Lattiude  float64
	Longitude float64
}

// instead of using one struct we have seperated the
type Photo struct {
	ID       int
	Catagory string
	Location locationP2D
	Time     time.Time
}

type Video struct {
	ID       int
	Catagory string
	Location locationP2D
	Time     time.Time
}

//in this approach the generate thumbnail function is implemented in both the structs but the interface is used to generalize the post type
//we can use the interface to call the GenerateThumbnail function without knowing the type of the post
//So the generated thumbnail function is decoupled
//there is no need forr nested if else statements to check the type of the post

func (p Photo) GenerateThumbnail() {

	fmt.Println("Generating photo thumbnail")

}

// both the video and photo struct have the same function name but different implementation
func (v Video) GenerateThumbnail() {

	fmt.Println("Generating video thumbnail")

}

// also as both the photo and video has the same interface we can use the interface to call the GenerateThumbnail function
func thumbnail(p post2Dif) {
	fmt.Println()
	fmt.Println(p)
	p.GenerateThumbnail()

}

func Problem2Dif() {

	fmt.Println()
	fmt.Println("Problem 2 different solution starts")
	fmt.Println()

	loc := locationP2D{
		Lattiude:  37.42,
		Longitude: -122.08,
	}

	photoPost := Photo{
		ID:       1,
		Catagory: "Photo",
		Location: loc,
		Time:     time.Now(),
	}
	thumbnail(photoPost)

	videoPost := Video{
		ID:       2,
		Catagory: "Video",
		Location: loc,
		Time:     time.Now(),
	}
	//same post type struct but as the different implementation of GenerateThumbnail both the structs are different
	thumbnail(videoPost)

}
