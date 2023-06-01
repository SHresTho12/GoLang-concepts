package main

import (
	"strconv"
	"testing"
)

func ConvertStringToInt(str string) (int, error) {

	convertedInt, err := strconv.Atoi(str)

	if err != nil {
		return 0, err
	}
	return convertedInt, nil
}

func TestProblem5(t *testing.T) {
	//test data
	var tests = []struct {
		input string
		want  int
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"a", 0},
		{"b", 0},   //this test case should fail as we are sending a string with caracters that are not valid integers
		{"12d", 0}, //this test case should fail as we are expecting an error
	}

	for _, test := range tests {
		if got, _ := ConvertStringToInt(test.input); got != test.want {
			t.Errorf("ConvertStringToInt(%s) = %d", test.input, got)
		}
	}
}
