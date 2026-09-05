package main

import (
	"fmt"
	"io"
	"os"
)

func readInt(b []byte, idx *int) int {
	for *idx < len(b) && (b[*idx] < '0' || b[*idx] > '9') {
		*idx++
	}
	if *idx >= len(b) {
		return 0
	}
	res := 0
	for *idx < len(b) && b[*idx] >= '0' && b[*idx] <= '9' {
		res = res*10 + int(b[*idx]-'0')
		*idx++
	}
	return res
}

func main() {
	b, _ := io.ReadAll(os.Stdin)
	idx := 0
	if idx >= len(b) {
		return
	}
	t := readInt(b, &idx)
	for i := 0; i < t; i++ {
		n := readInt(b, &idx)
		d := readInt(b, &idx)
		adj := make([][]int, n+1)
		for j := 0; j < n-1; j++ {
			u := readInt(b, &idx)
			v := readInt(b, &idx)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		solve(n, d, adj)
	}
}

func solve(n, d int, adj [][]int) {
	var ans int64 = 0

	var dfs func(u, p int) ([]int64, []int64)
	dfs = func(u, p int) ([]int64, []int64) {
		fu := make([]int64, 1)
		fu[0] = 1
		gu := make([]int64, 0)

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			fv, gv := dfs(v, u)

			for i := 0; i < len(gv); i++ {
				j := d - 2 - i
				if j >= 0 && j < len(fu) {
					ans += gv[i] * fu[j]
				}
			}
			for i := 0; i < len(gu); i++ {
				j := d - 2 - i
				if j >= 0 && j < len(fv) {
					ans += gu[i] * fv[j]
				}
			}

			newLenF := len(fu)
			if len(fv)+1 > newLenF {
				newLenF = len(fv) + 1
			}
			if newLenF > d {
				newLenF = d
			}

			newLenG := len(gu)
			if len(gv)+1 > newLenG {
				newLenG = len(gv) + 1
			}
			if len(fu)+len(fv) > newLenG {
				newLenG = len(fu) + len(fv)
			}
			if newLenG > d {
				newLenG = d
			}

			newF := make([]int64, newLenF)
			newG := make([]int64, newLenG)

			copy(newF, fu)
			copy(newG, gu)

			for i := 0; i < len(fu); i++ {
				for j := 0; j < len(fv); j++ {
					if i+j+1 < newLenG {
						newG[i+j+1] += fu[i] * fv[j]
					}
				}
			}

			for i := 0; i < len(gv); i++ {
				if i+1 < newLenG {
					newG[i+1] += gv[i]
				}
			}

			for j := 0; j < len(fv); j++ {
				if j+1 < newLenF {
					newF[j+1] += fv[j]
				}
			}

			fu = newF
			gu = newG
		}
		return fu, gu
	}

	dfs(1, 0)
	fmt.Println(ans)
}
