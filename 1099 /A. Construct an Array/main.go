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

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		res := make([]int, n)
		left := 1
		right := 2 * n

		for j := 0; j < n; j++ {
			if j%2 == 0 {
				res[j] = left
				left++
			} else {
				res[j] = right
				right--
			}
		}

		for j := 0; j < n; j++ {
			fmt.Fprintf(writer, "%d", res[j])
			if j < n-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	solve()
}
