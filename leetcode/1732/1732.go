package main

import (
	"fmt"
)

type Test struct {
	gain  []int
	want  int
}

func main() {
	tests := []Test{
		{
			gain: []int{-5,1,5,0,-7},
			want: 1,
		},
		{
			gain: []int{-4,-3,-2,-1,4,3,2},
			want: 0,
		},
	}

	for _, tc := range tests {
		got := largestAltitude(tc.gain)
		if got != tc.want {
			fmt.Printf("Input: %s => Output: %s, Expected: %s\n", tc.gain, got, tc.want)
		}
	}
}

func largestAltitude(gain []int) int {
	highestAltitude := 0
	currentAltitude := 0

	for _, v := range gain {
		currentAltitude += v
		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}
