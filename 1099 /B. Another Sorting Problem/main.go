package main

import (
	"bufio"
	"fmt"
	"os"
)

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

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

		if isSorted(a) {
			fmt.Println("YES")
			continue
		}

		l := 0
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				l = i
				break
			}
		}

		k := a[l] - a[l+1] + 1

		b := make([]int, n)
		copy(b, a)

		for i := l + 1; i < n; i++ {
			b[i] += k
		}

		if isSorted(b) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	solve()
}
