package main

import (
	"fmt"
)

func Frog(jumps []int) int {
	// your code here
	var memoBottomUp = map[int]int{}
	memoBottomUp[0] = jumps[0]

	for i := 1; i <= len(jumps); i++ {
		memoBottomUp[i] == memoBottomUp[i+1] || memoBottomUp[i] == memoBottomUp[i+2]

		pathfrog = append(pathfrog, memoBottomUp[i])
	}

	for i := 1; i <= len(pathfrog); i++ {
		for j := 2; j <= len(pathfrog); j++ {
			cost += math.abs(hi - hj)
		}
	}

	return cost
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
