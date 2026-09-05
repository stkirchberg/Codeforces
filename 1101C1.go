package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var n, x, s int
		fmt.Fscan(reader, &n, &x, &s)

		var u string
		fmt.Fscan(reader, &u)

		maxTables := x
		if n < x {
			maxTables = n
		}

		dp := make([]int, maxTables+1)
		for i := 0; i <= maxTables; i++ {
			dp[i] = -1
		}
		dp[0] = 0

		for i := 0; i < n; i++ {
			c := u[i]

			nextDp := make([]int, maxTables+1)
			copy(nextDp, dp)

			for k := 0; k <= maxTables; k++ {
				if (c == 'I' || c == 'A') && k > 0 && dp[k-1] != -1 {
					nextDp[k] = max(nextDp[k], dp[k-1]+1)
				}

				if (c == 'E' || c == 'A') && dp[k] != -1 {
					availableSpace := k*s - dp[k]
					if availableSpace > 0 {
						nextDp[k] = max(nextDp[k], dp[k]+1)
					}
				}
			}
			dp = nextDp
		}

		ans := 0
		for k := 0; k <= maxTables; k++ {
			if dp[k] > ans {
				ans = dp[k]
			}
		}

		fmt.Fprintln(writer, ans)
	}
}

func main() {
	solve()
}
