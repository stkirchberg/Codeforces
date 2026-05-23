package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

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
		for i := 0; i < n-1; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}

		p := make([]int, n+1)
		children := make([][]int, n+1)

		var dfs_init func(u, parent int)
		dfs_init = func(u, parent int) {
			p[u] = parent
			for _, v := range adj[u] {
				if v != parent {
					children[u] = append(children[u], v)
					dfs_init(v, u)
				}
			}
		}
		dfs_init(n, 0)

		M := make([]int, n+1)
		var dfs_M func(u int)
		dfs_M = func(u int) {
			M[u] = u
			for _, v := range children[u] {
				dfs_M(v)
				if M[v] > M[u] {
					M[u] = M[v]
				}
			}
		}
		dfs_M(n)

		p_cont := make([]int, n+1)
		var dfs_cont func(u, last_valid int)
		dfs_cont = func(u, last_valid int) {
			if u != n && M[u] == u {
				p_cont[u] = last_valid
				last_valid = u
			}
			for _, v := range children[u] {
				dfs_cont(v, last_valid)
			}
		}
		dfs_cont(n, n)

		var S_curr int64 = 0
		W := make([]int64, n+1)
		for u := n - 1; u >= 1; u-- {
			if M[u] != u {
				continue
			}
			p_c := p_cont[u]
			var sum int64
			if p_c == n {
				sum = S_curr
			} else {
				sum = (S_curr - W[p_c] + MOD) % MOD
			}
			if sum == 0 {
				W[u] = 1
			} else {
				W[u] = sum
			}
			S_curr = (S_curr + W[u]) % MOD
		}

		is_leaf := make([]bool, n+1)
		var L_max int = 0
		for u := 1; u <= n; u++ {
			if u != n && len(children[u]) == 0 {
				is_leaf[u] = true
				if u > L_max {
					L_max = u
				}
			}
		}

		in_candidates := make([]bool, n+1)
		has_candidates := false
		for u := 1; u <= n; u++ {
			if is_leaf[u] && u != L_max {
				curr := p[u]
				for curr != n {
					if M[curr] == curr && curr > L_max {
						if !in_candidates[curr] {
							in_candidates[curr] = true
							has_candidates = true
						}
						break
					}
					curr = p[curr]
				}
			}
		}

		var ans int64 = 0
		if !has_candidates {
			ans = 1
		} else {
			for u := 1; u <= n; u++ {
				if in_candidates[u] {
					ans = (ans + W[u]) % MOD
				}
			}
		}

		fmt.Fprintln(writer, ans)
	}
}
