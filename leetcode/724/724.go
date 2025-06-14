package main

import (
	"fmt"

	"github.com/x448/float16"
)

type Test struct {
	nums []int
	want int
}

func main() {
	tests := []Test{
		{
			nums: []int{-1, -1, -1, -1, -1, -1},
			want: -1,
		},
	}

	for _, tc := range tests {
		got := pivotIndex(tc.nums)
		if got != tc.want {
			fmt.Printf("Input: %s => Output: %s, Expected: %s\n", tc.nums, got, tc.want)
		}
	}
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	index := 0
	for _, v := range nums {
		if float32(sum - v)/float32(2) == float32(leftSum) {
			break
		}	

		leftSum += v
		index++
	}
	if index >= len(nums) {
		return -1
	}

	return index
}
