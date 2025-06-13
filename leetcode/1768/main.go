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
	sum := ""
	current := 0 // current at word 1 or 2. 0 is w1, 1 is w2
	w1i := 0
	w2i := 0

	for range len(word1) + len(word2) {
		if w1i >= len(word1) {
			sum += string(word2[w2i])
			w2i++
			current = 1
			continue
		}

		if w2i >= len(word2) {
			sum += string(word1[w1i])
			w1i++
			current = 0
			continue
		}

		if current == 0 {
			current = 1
			sum += string(word1[w1i])
			w1i++
		} else {
			current = 0
			sum += string(word2[w2i])
			w2i++
		}
	}
	return sum
}
