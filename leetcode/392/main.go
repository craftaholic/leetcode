package main

import (
	"fmt"
)

type Test struct {
	word1 string
	word2 string
	want  bool
}

func main() {
	tests := []Test{}
	tests = append(tests, Test{"abcd", "akjhbhkc", true})
	tests = append(tests, Test{"abx", "akjhbhkc", false})

	for _, tc := range tests {
		got := isSubsequence(tc.word1, tc.word2)
		if got != tc.want {
			fmt.Printf("Input: %s, %s => Output: %s, Expected: %s\n", tc.word1, tc.word2, got, tc.want)
		}
	}
}

func isSubsequence(s string, t string) bool {
	i := 0
	for cur := 0; cur < len(t); cur++ {
		// Condition to return true when s is iterated
		if i < len(s) {
			if s[i] == t[cur] {
				i++
			}
			continue
		}
	}

	return i == len(s)
}
