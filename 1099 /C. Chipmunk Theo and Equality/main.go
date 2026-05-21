package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getOps(val int, target int) int64 {
	ops := int64(0)
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
	return 1e15
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

		minTotalOps := int64(math.MaxInt64)

		for k := 0; k <= 30; k++ {
			target := 1 << k
			currentTotalOps := int64(0)
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
