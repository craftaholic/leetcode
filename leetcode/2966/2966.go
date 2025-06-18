package main

import (
	"fmt"
	"sort"
)

type Test struct {
	input1 []int
	input2 int
	want   int
}

func main() {
	tests := []Test{}

	for _, tc := range tests {
		got := divideArray(tc.input1, tc.input2)
		for _, v := range got {
			for _, vj := range v {
				fmt.Print(vj, ", ")
			}
		}
	}
}

func divideArray(nums []int, k int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := range len(nums) / 3 {
		if nums[i*3+2]-nums[i*3] > k {
			return [][]int{}
		}

		result = append(result, nums[i*3:i*3+3])
	}

	return result
}
