package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Test struct {
	input1 *ListNode
	want   *ListNode
}

func main() {
	tests := []Test{
		{
			input1: &ListNode{
				Val:  1,
				Next: nil,
			},
			want: &ListNode{},
		},
	}

	for _, tc := range tests {
		got := reverseList(tc.input1)
		for got.Next != nil {
			fmt.Print(got.Val, " -> ")
		}
	}
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		newNode := &ListNode{
			Val:  head.Val,
			Next: result,
		}
		result = newNode
		head = head.Next
	}

	return result
}
