package main

import (
	"fmt"
)

type Test struct {
	input1 int
	input2 int
	input3 int
	want   int
}

func main() {
	tests := []Test{}

	for i, tc := range tests {
		got := countGoodArrays(tc.input1, tc.input2, tc.input3)
		if got != tc.want {
			fmt.Print("tc ", i, ": ")
			fmt.Printf("Input: %d, %d, %d => Output: %d, Expected: %d\n", tc.input1, tc.input2, tc.input3, got, tc.want)
		}
	}
}

func countGoodArrays(n int, m int, k int) int {
	return 0
}
