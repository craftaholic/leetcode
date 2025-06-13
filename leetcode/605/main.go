package main

import (
	"fmt"
)

type Test struct {
	flowerbed []int
	n         int
	want      bool
}

func main() {
	tests := []Test{
		{
			flowerbed: []int{1, 0, 0, 0, 0, 0, 1},
			n:         2,
			want:      true,
		},
	}

	for _, tc := range tests {
		got := canPlaceFlowers(tc.flowerbed, tc.n)
		if got != tc.want {
			fmt.Print("Output: %s, Expected: %s\n", got, tc.want)
		}
	}
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	maxPot := 0

	for i := 0; i < len(flowerbed); {
		previousPot := 0
		nextPot := 0
		if i > 0 {
			previousPot = flowerbed[i-1]
		}
		if i < len(flowerbed)-1 {
			nextPot = flowerbed[i+1]
		}

		// if the current pot is empty then check condition
		if flowerbed[i] == 0 {
			if previousPot == 0 && nextPot == 0 {
				// Plant the tree into the pot
				maxPot++
				i += 2 // Skip the pot right after the pot we just have planted
				continue
			}
			i++
		}

		// because current pot is planted -> skip evaluating nextPot
		i += 2
	}

	if maxPot < n {
		return false
	}

	return true
}
