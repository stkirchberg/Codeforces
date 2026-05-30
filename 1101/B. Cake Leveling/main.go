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

		var sum int64 = 0
		var minH int64 = -1

		for i := 1; i <= n; i++ {
			var a int64
			fmt.Fscan(reader, &a)

			sum += a

			currentMaxH := sum / int64(i)

			if minH == -1 || currentMaxH < minH {
				minH = currentMaxH
			}

			if i == n {
				fmt.Fprintf(writer, "%d\n", minH)
			} else {
				fmt.Fprintf(writer, "%d ", minH)
			}
		}
	}
}

func main() {
	solve()
}
