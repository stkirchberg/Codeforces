//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {

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

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		minCalls := n

		for i := 0; i < n; i++ {
			target := a[i]
			l, r := 0, 0

			for j := 0; j < n; j++ {
				if a[j] < target {
					l++
				} else if a[j] > target {
					r++
				}
			}

			calls := l
			if r > l {
				calls = r
			}

			if calls < minCalls {
				minCalls = calls
			}
		}

		fmt.Fprintln(writer, minCalls)
	}
}

func main() {
	solve()
}
