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

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		var ans []int
		flipped := false

		for i := n - 1; i >= 0; i-- {
			currentVal := a[i]
			if flipped {
				currentVal = -currentVal
			}

			if currentVal > 0 {
				ans = append(ans, i+1)
				flipped = !flipped
			}
		}

		fmt.Fprintln(out, len(ans))
		for i, val := range ans {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, val)
		}
		fmt.Fprintln(out)
	}
}
