package main

import (
	"fmt"
)

type Test struct {
	word1 string
	word2 string
	want  string
}

func main() {
	tests := []Test{}
	tests = append(tests, Test{"abc", "pqr", "apbqcr"})
	tests = append(tests, Test{"ab", "pqrs", "apbqrs"})
	tests = append(tests, Test{"abcd", "pq", "apbqcd"})

	for _, tc := range tests {
		got := mergeAlternately(tc.word1, tc.word2)
		if got != tc.want {
			fmt.Printf("Input: %s, %s => Output: %s, Expected: %s\n", tc.word1, tc.word2, got, tc.want)
		}
	}
}

func mergeAlternately(word1 string, word2 string) string {
	return "abc"
}
