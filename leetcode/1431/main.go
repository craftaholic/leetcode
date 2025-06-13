package main

import (
	"fmt"
)

type Test struct {
	candies      []int
	extraCandies int
	want         []bool
}

func main() {
	tests := []Test{}
	tests = append(tests, Test{
		candies:      []int{2, 3, 5, 1, 3},
		extraCandies: 3,
		want:         []bool{true, true, true, false, true},
	})

	for _, tc := range tests {
		got := kidsWithCandies(tc.candies, tc.extraCandies)
		for i, v := range got {
			if v != tc.want[i] {
				fmt.Print("Output: ", got)
				fmt.Print("Expected: \n", tc.want)
			}
		}
	}
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	result := make([]bool, len(candies))
	candiesWithExtra := make([]int, len(candies))

	for i, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}

		candiesWithExtra[i] = v + extraCandies
	}

	for i, v := range candiesWithExtra {
		if v >= maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
