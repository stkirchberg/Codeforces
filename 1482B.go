//go:build ignore

package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	if n == 1 {
		fmt.Println(0)
		return
	}

	diffs := make(map[int]bool)
	for i := 1; i < n; i++ {
		diffs[a[i]-a[i-1]] = true
	}

	if len(diffs) == 1 {
		fmt.Println(0)
		return
	}

	if len(diffs) > 2 {
		fmt.Println(-1)
		return
	}

	var pos, neg int
	hasPos, hasNeg := false, false
	for d := range diffs {
		if d >= 0 {
			pos = d
			hasPos = true
		} else {
			neg = d
			hasNeg = true
		}
	}

	if !hasPos || !hasNeg {
		fmt.Println(-1)
		return
	}

	m := pos - neg
	c := pos

	for _, val := range a {
		if val >= m {
			fmt.Println(-1)
			return
		}
	}

	for i := 1; i < n; i++ {
		if (a[i-1]+c)%m != a[i] {
			fmt.Println(-1)
			return
		}
	}

	fmt.Printf("%d %d\n", m, c)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
