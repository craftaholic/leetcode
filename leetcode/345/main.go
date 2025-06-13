package main

import (
	"fmt"
)

type Test struct {
	word1 string
	want  string
}

func main() {
	tests := []Test{}
	tests = append(tests, Test{"IceCreAm", "AceCreIm"})
	tests = append(tests, Test{"leetcode", "leotcede"})

	for _, tc := range tests {
		got := reverseVowels(tc.word1)
		if got != tc.want {
			fmt.Printf("Input: %s => Output: %s, Expected: %s\n", tc.word1, got, tc.want)
		}
	}
}

func reverseVowels(s string) string {
	result := []rune(s)
	vowels := map[string]any{"a": true, "e": true, "i": true, "o": true, "u": true, "A": true, "E": true, "I": true, "O": true, "U": true}

	vowelList := []rune{}
	indexList := []int{}
	for i, c := range s {
		// if this is a vowels
		if vowels[string(c)] != nil {
			vowelList = append(vowelList, c)
			indexList = append(indexList, i)
		}
	}

	for i, v := range indexList {
		result[v] = vowelList[len(vowelList)-1-i]
	}

	return string(result)
}
