package main

import (
	"fmt"
)

type Test struct {
	input1 []int
	input2 []int
	want   [][]int
}

func main() {
	tests := []Test{
		{
			input1: []int{1, 2, 3},
			input2: []int{2, 4, 6},
			want: [][]int{
				{1, 3},
				{4, 6},
			},
		},
		{
			input1: []int{1, 2, 3},
			input2: []int{2, 4, 6},
			want: [][]int{
				{1, 3},
				{4, 6},
			},
		},
	}

	for i, tc := range tests {
		got := findDifference(tc.input1, tc.input2)
		fmt.Print("tc ", i)
		fmt.Print(got)
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	maps1 := make(map[int]struct{})
	maps2 := make(map[int]struct{})
	dists1 := []int{}
	dists2 := []int{}

	for _, v := range nums1 {
		maps1[v] = struct{}{}
	}
	for _, v := range nums2 {
		maps2[v] = struct{}{}
	}

	for k := range maps1 {
		if _, exist := maps2[k]; !exist {
			dists1 = append(dists1, k)
		}
	}

	for k := range maps2 {
		if _, exist := maps1[k]; !exist {
			dists2 = append(dists2, k)
		}
	}

	return [][]int{
		dists1,
		dists2,
	}
}
