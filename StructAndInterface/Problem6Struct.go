package main

import (
	"fmt"
)

type calculator interface {
	add() float64
	sub() float64
	mul() float64
	div() float64
}

type BasicCalculator struct {
	a, b float64
}

type ScientificCalculator struct {
	a, b float64
}

func (b BasicCalculator) add() float64 {
	return b.a + b.b
}

func (b BasicCalculator) sub() float64 {
	return b.a - b.b
}

func (b BasicCalculator) mul() float64 {
	return b.a * b.b
}

func (b BasicCalculator) div() float64 {
	return b.a / b.b
}

func (s ScientificCalculator) add() float64 {
	return s.a + s.b
}

func (s ScientificCalculator) sub() float64 {
	return s.a - s.b
}

func (s ScientificCalculator) mul() float64 {
	return s.a * s.b
}

func (s ScientificCalculator) div() float64 {
	return s.a / s.b
}

func Problem6() {

	fmt.Println()
	fmt.Println("Problem 6 solution starts")
	fmt.Println()

	b := BasicCalculator{a: 10, b: 5}
	s := ScientificCalculator{a: 10, b: 5}

	fmt.Println(b.add())
	fmt.Println(b.sub())
	fmt.Println(b.mul())
	fmt.Println(b.div())

	fmt.Println(s.add())
	fmt.Println(s.sub())
	fmt.Println(s.mul())
	fmt.Println(s.div())

}
