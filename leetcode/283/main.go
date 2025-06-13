package main

import (
	"fmt"
)

type Test struct {
	input []int
	want  []int
}

func main() {
	tests := []Test{}
	tests = append(tests, Test{
		input: []int{0, 1, 0, 3, 12},
		want:  []int{1, 3, 12, 0, 0},
	})
	tests = append(tests, Test{
		input: []int{0},
		want:  []int{0},
	})

	for _, tc := range tests {
		moveZeroes(tc.input)
		for i, v := range tc.input {
			if v != tc.want[i] {
				fmt.Print("Invalid: have ", v, " want ", tc.want[i], "\n")
			}
		}
	}
}

func moveZeroes(nums []int) {
	index := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
			continue
		}
	}

	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
