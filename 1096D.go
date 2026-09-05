package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	base = 100003
	mod  = 1000000007
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	_, _ = fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(in, &n)
		size := 2 * n
		a := make([]int, size)

		pos := make([][2]int, n+1)
		for j := 0; j <= n; j++ {
			pos[j][0] = size + 1
			pos[j][1] = -1
		}

		for j := 0; j < size; j++ {
			fmt.Fscan(in, &a[j])
			if j < pos[a[j]][0] {
				pos[a[j]][0] = j
			}
			if j > pos[a[j]][1] {
				pos[a[j]][1] = j
			}
		}

		pow := make([]int64, size+1)
		pow[0] = 1
		for j := 1; j <= size; j++ {
			pow[j] = (pow[j-1] * base) % mod
		}

		h1 := make([]int64, size+1)
		h2 := make([]int64, size+1)
		for j := 0; j < size; j++ {
			h1[j+1] = (h1[j]*base + int64(a[j]+1)) % mod
			h2[j+1] = (h2[j]*base + int64(a[size-1-j]+1)) % mod
		}

		isPal := func(l, r int) bool {
			len := r - l + 1
			val1 := (h1[r+1] - h1[l]*pow[len]%mod + mod) % mod
			val2 := (h2[size-l] - h2[size-r-1]*pow[len]%mod + mod) % mod
			return val1 == val2
		}

		minL, maxR := size, -1
		ans := 0

		for k := 0; k < n; k++ {
			if pos[k][0] < minL {
				minL = pos[k][0]
			}
			if pos[k][1] > maxR {
				maxR = pos[k][1]
			}

			if minL <= maxR && isPal(minL, maxR) {
				ans = k + 1
			} else {
				break
			}
		}

		out.WriteString(strconv.Itoa(ans) + "\n")
	}
}
