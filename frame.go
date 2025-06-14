package main

import (
	"fmt"
)

type Test struct {
	input1 []int
	input2 []int
	want   int
}

func main() {
	tests := []Test{}

	for i, tc := range tests {
		got := findDifference(tc.input1, tc.input2)
		if got != tc.want {
			fmt.Print("tc ", i)
			fmt.Printf("Input: %s, %s => Output: %s, Expected: %s\n", tc.input1, tc.input2, got, tc.want)
		}
	}
}

func findDifference(nums1 []int, nums2 []int) int {
	return 0
}
