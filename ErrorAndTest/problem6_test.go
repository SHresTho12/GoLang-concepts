package main

import (
	"testing"
)

type customError2 string

func (e customError2) Error() string {
	return string(e)
}
func GetAverage(nums []int) (float64, error) {
	//if there is no element in the sent array return 0 and an error
	if len(nums) <= 0 {
		return 0, customError2("No elements in the array")
	}
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return float64(sum) / float64(len(nums)), nil

}

func TestProblem6(t *testing.T) {

	var tests = []struct {
		input []int
		want  float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2}, 1.5},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for _, test := range tests {
		if got, _ := GetAverage(test.input); got != test.want {
			t.Errorf("GetAverage(%v) = %f", test.input, got)
		}
	}
}
