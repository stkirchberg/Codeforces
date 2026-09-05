package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		maxDiff := 0
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				diff := a[i] - a[i+1]
				if diff > maxDiff {
					maxDiff = diff
				}
			}
		}

		if maxDiff == 0 {
			fmt.Fprintln(out, "YES")
			continue
		}

		k := maxDiff
		possible := true
		prev := a[0]

		for i := 1; i < n; i++ {
			if a[i] >= prev {
				prev = a[i]
			} else if a[i]+k >= prev {
				prev = a[i] + k
			} else {
				possible = false
				break
			}
		}

		if possible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	solve(in, out)
}
