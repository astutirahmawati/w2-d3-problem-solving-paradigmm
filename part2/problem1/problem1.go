package main

import "fmt"

func TopDown(n int) int {
	// your code here
	var memoTopDown = map[int]int{}
	value, sudahDihitung := memoTopDown[n]
	if sudahDihitung {
		return value
	}
	if n <= 1 {
		memoTopDown[n] = n
	} else {
		memoTopDown[n] = TopDown(n-1) + TopDown(n-2)
	}
	return memoTopDown[n]
}

func main() {
	fmt.Println(TopDown(0))  //0
	fmt.Println(TopDown(1))  //1
	fmt.Println(TopDown(2))  //1
	fmt.Println(TopDown(3))  //2
	fmt.Println(TopDown(5))  //5
	fmt.Println(TopDown(6))  //8
	fmt.Println(TopDown(7))  //13
	fmt.Println(TopDown(9))  //34
	fmt.Println(TopDown(10)) //55
}
