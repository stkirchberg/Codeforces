package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int64, n+2)
		var sumNoOps int64 = 0

		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
			sumNoOps += a[i]
		}

		prefAbs := make([]int64, n+2)
		suffA := make([]int64, n+2)

		for i := 1; i <= n; i++ {
			prefAbs[i] = prefAbs[i-1] + abs(a[i])
		}
		for i := n; i >= 1; i-- {
			suffA[i] = suffA[i+1] + a[i]
		}

		maxSum := sumNoOps
		bestM := -1

		for i := 1; i <= n; i++ {
			if a[i] > 0 {
				curSum := prefAbs[i-1] - a[i] + suffA[i+1]
				if curSum > maxSum {
					maxSum = curSum
					bestM = i
				}
			}
		}

		if bestM == -1 {
			fmt.Fprintln(out, 0)
			fmt.Fprintln(out)
			continue
		}

		P := make([]int, n+2)
		for i := 1; i <= n; i++ {
			if i > bestM {
				P[i] = 0
			} else if i == bestM {
				P[i] = 1
			} else {
				if a[i] < 0 {
					P[i] = 1
				} else {
					P[i] = 0
				}
			}
		}

		var S []int
		for i := 1; i <= n; i++ {
			if P[i] != P[i+1] {
				S = append(S, i)
			}
		}

		first := -1
		var rest []int

		for j := len(S) - 1; j >= 0; j-- {
			x := S[j]
			if first == -1 {
				first = x
			} else {
				if a[x] > 0 {
					rest = append(rest, first)
					first = x
				} else {
					rest = append(rest, x)
				}
			}
		}

		for i, j := 0, len(rest)-1; i < j; i, j = i+1, j-1 {
			rest[i], rest[j] = rest[j], rest[i]
		}

		ans := make([]int, 0, len(S))
		if first != -1 {
			ans = append(ans, first)
			ans = append(ans, rest...)
		}

		fmt.Fprintln(out, len(ans))
		for i, v := range ans {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
	}
}
