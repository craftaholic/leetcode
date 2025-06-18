package main

import (
	"fmt"
)

type Test struct {
	input1 int
	input2 int
	input3 int
	want   int
}

func main() {
	tests := []Test{}

	for i, tc := range tests {
		counter := Constructor()
		got := counter.Ping(1)
		if got != tc.want {
			fmt.Print("tc ", i, ": ")
			fmt.Printf("Input: %d, %d, %d => Output: %d, Expected: %d\n", tc.input1, tc.input2, tc.input3, got, tc.want)
		}
	}
}

type RecentCounter struct {
	pingList              []int
	numberOfRequestInSpan int
	furthestRequest       int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pingList:        []int{},
		furthestRequest: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pingList = append(this.pingList, t)

	i := this.furthestRequest
	for ; i < len(this.pingList); i++ {
		if t-this.pingList[i] <= 3000 {
			this.furthestRequest = i
			break
		}
	}

	return len(this.pingList) - i
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
