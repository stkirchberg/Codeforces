//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(in, &n)

		var s string
		fmt.Fscan(in, &s)

		a := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}

		c := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &c[i])
		}

		possible := true
		for i := 2; i <= n; i++ {
			if c[i] < c[i-1] {
				possible = false
				break
			}
		}

		if s[0] == '1' && a[1] != c[1] {
			possible = false
		}

		if !possible {
			fmt.Fprintln(out, "No")
			continue
		}

		b := make([]int64, n+1)
		i := 1

		for i <= n {
			L := i
			R := i
			for R+1 <= n && s[R] == '1' {
				R++
			}

			offset := make([]int64, R-L+1)
			offset[0] = 0
			for k := L + 1; k <= R; k++ {
				offset[k-L] = offset[k-1-L] + a[k]
			}

			forcedCount := 0
			var base int64

			for k := L; k <= R; k++ {
				if k == 1 || c[k] > c[k-1] {
					reqBase := c[k] - offset[k-L]

					if forcedCount == 0 {
						base = reqBase
						forcedCount++
					} else if base != reqBase {
						possible = false
						break
					}
				}
			}

			if !possible {
				break
			}

			if forcedCount == 0 {
				base = math.MaxInt64
				for k := L; k <= R; k++ {
					val := c[k] - offset[k-L]
					if val < base {
						base = val
					}
				}
			}

			for k := L; k <= R; k++ {
				b[k] = base + offset[k-L]
				if b[k] > c[k] {

					possible = false
					break
				}
			}

			if !possible {
				break
			}

			i = R + 1
		}

		if !possible {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
			ans := make([]int64, n)
			ans[0] = b[1]
			for k := 2; k <= n; k++ {
				ans[k-1] = b[k] - b[k-1]
			}

			for k := 0; k < n; k++ {
				if k > 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, ans[k])
			}
			fmt.Fprintln(out)
		}
	}
}

func main() {
	solve()
}
