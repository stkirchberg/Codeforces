package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)

	var result []int

	for i := 1; i <= 2*n; i += 2 {
		result = append(result, i)
	}

	for i := 2; i <= 2*n; i += 2 {
		result = append(result, i)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
