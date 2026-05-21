package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(reader, &t)

	for tc := 0; tc < t; tc++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		days := make([][]int, m)
		for i := 0; i < m; i++ {
			var k int
			fmt.Fscan(reader, &k)
			days[i] = make([]int, k)
			for j := 0; j < k; j++ {
				fmt.Fscan(reader, &days[i][j])
			}
		}

		ans := make([]int, m)
		counts := make([]int, n+1)
		limit := (m + 1) / 2
		possible := true

		for i := 0; i < m; i++ {
			if len(days[i]) == 1 {
				f := days[i][0]
				ans[i] = f
				counts[f]++
				if counts[f] > limit {
					possible = false
					break
				}
			}
		}

		if !possible {
			fmt.Println("NO")
			continue
		}

		for i := 0; i < m; i++ {
			if ans[i] == 0 {
				chosen := false
				for _, f := range days[i] {
					if counts[f] < limit {
						ans[i] = f
						counts[f]++
						chosen = true
						break
					}
				}
				if !chosen {
					possible = false
					break
				}
			}
		}

		if !possible {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for i := 0; i < m; i++ {
				fmt.Printf("%d", ans[i])
				if i < m-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	solve()
}
