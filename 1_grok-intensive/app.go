package main

import "fmt"

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func findMax(a []int)  (m int) {
	m = a[0]
	for i := 1; i < len(a); i++ {
		m = max(m, a[i])
	}
	return
}

func main1() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(findMax(arr))
}
