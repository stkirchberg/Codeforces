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

		a := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}

		var sumMax int64 = 0
		var maxMin int64 = 0

		for j := 0; j < n; j++ {
			var b int64
			fmt.Fscan(in, &b)

			maxVal := a[j]
			minVal := b
			if b > a[j] {
				maxVal = b
				minVal = a[j]
			}

			sumMax += maxVal
			if minVal > maxMin {
				maxMin = minVal
			}
		}

		fmt.Fprintln(out, sumMax+maxMin)
	}
}
