package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getOps(val int, target int) int {
	ops := 0
	curr := val
	for curr > target {
		if curr%2 == 0 {
			curr /= 2
		} else {
			curr++
		}
		ops++
	}
	if curr == target {
		return ops
	}
	return 1e9
}

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	minTotalOps := math.MaxInt64

	for k := 0; k <= 30; k++ {
		target := 1 << k
		currentTotalOps := 0
		possible := true

		for _, val := range a {
			cost := getOps(val, target)
			if cost >= 1e9 {
				possible = false
				break
			}
			currentTotalOps += cost
		}

		if possible {
			if currentTotalOps < minTotalOps {
				minTotalOps = currentTotalOps
			}
		}
	}

	fmt.Println(minTotalOps)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[j])
		}

		minTotalOps := math.MaxInt64

		for k := 0; k <= 30; k++ {
			target := 1 << k
			currentTotalOps := 0
			for _, val := range a {
				currentTotalOps += getOps(val, target)
			}
			if currentTotalOps < minTotalOps {
				minTotalOps = currentTotalOps
			}
		}
		fmt.Println(minTotalOps)
	}
}
