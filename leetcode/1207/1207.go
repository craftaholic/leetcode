package main

import (
	"fmt"
)

type Test struct {
	input1 []int
	want   bool
}

func main() {
	tests := []Test{
		{
			input1: []int{1,2,2,1,1,3},
			want: true,
		},
		{
			input1: []int{1,2},
			want: false,
		},
		{
			input1: []int{-3,0,1,-3,1,1,1,-3,10,0},
			want: true,
		},
	}

	for i, tc := range tests {
		got := uniqueOccurrences(tc.input1)
		if got != tc.want {
			fmt.Print("tc ", i)
			fmt.Printf("Input: %s => Output: %s, Expected: %s\n", tc.input1, got, tc.want)
		}
	}
}

func uniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int, len(arr))
	occurenceMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		countMap[arr[i]]++	
	}

	for k,v := range countMap {
		if _, exists := occurenceMap[v]; exists {
			return false
		}
		occurenceMap[v] = k
	}

	return true
}
