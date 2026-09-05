//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		var min, max int
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(in, &a)

			if j == 0 {
				min = a
				max = a
			} else {
				if a < min {
					min = a
				}
				if a > max {
					max = a
				}
			}
		}

		ans := (max - min + 1) / 2
		fmt.Fprintln(out, ans)
	}
}
