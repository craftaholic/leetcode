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
	tests = append(tests, Test{"ABCABC", "ABC", "ABC"})
	tests = append(tests, Test{"ABABAB", "ABAB", "AB"})
	tests = append(tests, Test{"LEET", "CODE", ""})

	for _, tc := range tests {
		got := gcdOfStrings(tc.word1, tc.word2)
		if got != tc.want {
			fmt.Printf("Input: %s, %s => Output: %s, Expected: %s\n", tc.word1, tc.word2, got, tc.want)
		}
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	smaller := str1
	divider := ""
	if len(str1) > len(str2) {
		smaller = str2
	}

	for i := 1; i <= len(smaller); i++ {
		// Only check for the length of the dinvider that divides both strings
		if len(str1)%i != 0 || len(str2)%i != 0 {
			continue
		}

		currentDivider := string(smaller[0:i])
		valid := true

		for j := range len(str1) / i {
			if str1[j*i:j*i+i] != currentDivider {
				valid = false
				break
			}
		}

		for j := range len(str2) / i {
			if str2[j*i:j*i+i] != currentDivider {
				valid = false
				break
			}
		}

		if valid {
			divider = currentDivider
		}
	}

	return divider
}
