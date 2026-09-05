//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(reader, &n)

		adj := make([][]int, n+1)
		deg := make([]int, n+1)
		for i := 0; i < n-1; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
			deg[u]++
			deg[v]++
		}

		X_0 := 0
		for i := 1; i <= n; i++ {
			if deg[i] == 1 {
				X_0 = max(X_0, i)
			}
		}

		M := make([]int, n+1)
		var dfs_M func(u, p int)
		dfs_M = func(u, p int) {
			M[u] = u
			for _, v := range adj[u] {
				if v != p {
					dfs_M(v, u)
					M[u] = max(M[u], M[v])
				}
			}
		}
		dfs_M(n, 0)

		max_child := make([]int, n+1)
		var dfs_calc func(u, p, last_valid int)
		dfs_calc = func(u, p, last_valid int) {
			next_valid := last_valid
			if M[u] == u {
				if last_valid != 0 {
					max_child[last_valid] = max(max_child[last_valid], u)
				}
				next_valid = u
			}
			for _, v := range adj[u] {
				if v != p {
					dfs_calc(v, u, next_valid)
				}
			}
		}
		dfs_calc(n, 0, 0)

		dp := make([]int64, n+1)
		pref := make([]int64, n+1)

		dp[X_0] = 1
		for i := 1; i <= n; i++ {
			if M[i] == i && i > X_0 {
				var L, R int
				if i == n {
					L = max_child[n]
					R = n - 1
				} else {
					L = max_child[i] + 1
					R = i - 1
				}
				if L <= R {
					sum := (pref[R] - pref[L-1] + MOD) % MOD
					dp[i] = sum
				}
			}
			pref[i] = (pref[i-1] + dp[i]) % MOD
		}

		fmt.Fprintln(writer, dp[n])
	}
}
