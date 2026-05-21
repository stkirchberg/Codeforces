package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func solve() {
	var n, d int
	fmt.Fscan(reader, &n, &d)

	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := int64(0)

	for m := 1; m <= n; m++ {
		var branches [][]int
		for _, neighbor := range adj[m] {
			distCounts := make([]int, n+1)
			var dfs func(u, p, depth int)
			dfs = func(u, p, depth int) {
				distCounts[depth]++
				for _, v := range adj[u] {
					if v != p {
						dfs(v, u, depth+1)
					}
				}
			}
			dfs(neighbor, m, 1)
			branches = append(branches, distCounts)
		}

		pairs := make([]int64, n+1)
		singles := make([]int64, n+1)

		for _, b := range branches {
			for s1 := 1; s1 <= n; s1++ {
				if pairs[s1] == 0 {
					continue
				}
				for s2 := 1; s2 <= n; s2++ {
					if b[s2] == 0 {
						continue
					}
					if s1+s2+1 == d {
						ans += pairs[s1] * int64(b[s2])
					}
				}
			}

			for s1 := 1; s1 <= n; s1++ {
				if b[s1] == 0 {
					continue
				}
				for s2 := 1; s2 <= n; s2++ {
					if singles[s2] == 0 {
						continue
					}
					pairs[s1+s2] += singles[s2] * int64(b[s1])
				}
			}

			for s := 1; s <= n; s++ {
				singles[s] += int64(b[s])
			}
		}

		if d == 1 {
			continue
		}
	}

	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
