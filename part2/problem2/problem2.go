package main

import "fmt"

func BottomUp(n int) int {
	// your code here
	var memoBottomUp = map[int]int{}

	for i := 0; i <= n; i++ {
		if i <= 1 {
			memoBottomUp[i] = i
		} else {
			memoBottomUp[i] = memoBottomUp[i-1] + memoBottomUp[i-2]
		}
	}
	return memoBottomUp[n]
}

func main() {
	fmt.Println(BottomUp(0))  //0
	fmt.Println(BottomUp(1))  //1
	fmt.Println(BottomUp(2))  //1
	fmt.Println(BottomUp(3))  //2
	fmt.Println(BottomUp(5))  //5
	fmt.Println(BottomUp(6))  //8
	fmt.Println(BottomUp(7))  //13
	fmt.Println(BottomUp(9))  //34
	fmt.Println(BottomUp(10)) //55
}
