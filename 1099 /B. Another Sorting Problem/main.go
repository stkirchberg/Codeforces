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
	fmt.Fscan(reader, &t)

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		idx := -1
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				idx = i
				break
			}
		}

		if idx == -1 {
			fmt.Fprintln(writer, "YES")
			continue
		}

		k := a[idx] - a[idx+1]

		possible := true
		prev := -1

		for i := 0; i < n; i++ {
			val := a[i]

			if i > idx {
				val += k
			}

			if val < prev {
				possible = false
				break
			}
			prev = val
		}

		if possible {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func main() {
	solve()
}
