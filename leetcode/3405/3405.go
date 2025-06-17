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
	tests := []Test{
		{
			input1: 3,
			input2: 2,
			input3: 1,
			want: 4,
		},
		{
			input1: 4,
			input2: 2,
			input3: 2,
			want: 6,
		},
		{
			input1: 5,
			input2: 2,
			input3: 0,
			want: 2,
		},
	}

	for i, tc := range tests {
		got := countGoodArrays(tc.input1, tc.input2, tc.input3)
		if got != tc.want {
			fmt.Print("tc ", i, ": ")
			fmt.Printf("Input: %d, %d, %d => Output: %d, Expected: %d\n", tc.input1, tc.input2, tc.input3, got, tc.want)
		}
	}
}

func countGoodArrays(n int, m int, k int) int {
	// To solve this, we can imagine arr[i-1], and arr[i] is 1 number j
	// Then the combination is all distributions of the array where 
	// j belongs [1, m]
	// total number in the imagine array is n - k elements

	return 0
}
