package main

import (
	"fmt"
)

type Test struct {
	nums []int
	k    int
	want float64
}

func main() {
	tests := []Test{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75000,
		},
		{
			nums: []int{5},
			k:    1,
			want: 5.00000,
		},
	}

	for _, tc := range tests {
		got := findMaxAverage(tc.nums, tc.k)
		if got != tc.want {
			fmt.Printf("Input: %s, %s => Output: %s, Expected: %s\n", tc.nums, tc.k, got, tc.want)
		}
	}
}

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0.00000
	currentSum := 0.00000

	// Calculate the first window values
	for i := range k {
		if i < len(nums) {
			maxSum += float64(nums[i])
		}
	}

	// If length < k -> return the avg based on len of nums
	if len(nums) < k {
		return maxSum / float64(len(nums))
	}

	// If k > len nums -> sliding the window
	currentSum = maxSum
	// Sliding the window for the rest of item
	for i := range len(nums) - k {
		currentSum = currentSum + float64(nums[i+k]) - float64(nums[i])
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum / float64(k)
}
