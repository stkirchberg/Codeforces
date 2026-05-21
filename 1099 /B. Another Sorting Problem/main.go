package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)
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
			fmt.Println("YES")
			continue
		}

		k := a[idx] - a[idx+1]

		if k < 0 {
			k = -k
		}

		b := make([]int, n)
		copy(b, a)

		threshold := a[idx+1]
		for i := 0; i <= idx; i++ {
			if b[i] < threshold {
				b[i] += k
			}
		}

		ok := true
		for i := 0; i < n-1; i++ {
			if b[i] > b[i+1] {
				ok = false
				break
			}
		}

		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	solve()
}
