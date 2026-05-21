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

		firstIdx := -1
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				firstIdx = i
				break
			}
		}

		if firstIdx == -1 {
			fmt.Println("YES")
			continue
		}

		k := a[firstIdx] - a[firstIdx+1]

		b := make([]int, n)
		copy(b, a)

		for i := firstIdx + 1; i < n; i++ {
			b[i] += k
		}

		possible := true
		for i := 0; i < n-1; i++ {
			if b[i] > b[i+1] {
				possible = false
				break
			}
		}

		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	solve()
}
