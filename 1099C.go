package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscan(in, &t)

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		candidates := make([]int, 0, 105)
		candidateIdx := make(map[int]int)

		curr := a[0]
		for steps := 0; steps <= 100; steps++ {
			if _, exists := candidateIdx[curr]; !exists {
				candidateIdx[curr] = len(candidates)
				candidates = append(candidates, curr)
			}

			if curr%2 == 0 {
				curr /= 2
			} else {
				curr += 1
			}
		}

		numCands := len(candidates)
		totalCost := make([]int, numCands)
		validCount := make([]int, numCands)
		localDist := make([]int, numCands)

		for _, val := range a {
			for i := 0; i < numCands; i++ {
				localDist[i] = -1
			}

			curr := val
			for steps := 0; steps <= 100; steps++ {
				if idx, exists := candidateIdx[curr]; exists {
					if localDist[idx] == -1 {
						localDist[idx] = steps
					}
				}
				if curr%2 == 0 {
					curr /= 2
				} else {
					curr += 1
				}
			}

			for i := 0; i < numCands; i++ {
				if localDist[i] != -1 {
					totalCost[i] += localDist[i]
					validCount[i]++
				}
			}
		}

		ans := -1
		for i := 0; i < numCands; i++ {
			if validCount[i] == n {
				if ans == -1 || totalCost[i] < ans {
					ans = totalCost[i]
				}
			}
		}

		fmt.Fprintln(out, ans)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	solve(in, out)
}
