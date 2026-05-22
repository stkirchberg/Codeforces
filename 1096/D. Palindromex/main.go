package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	base = 100003
	mod1 = 1000000007
	mod2 = 1000000009
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	next := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	pow1 := make([]int64, 200005)
	pow2 := make([]int64, 200005)
	pow1[0], pow2[0] = 1, 1
	for i := 1; i < 200005; i++ {
		pow1[i] = (pow1[i-1] * base) % mod1
		pow2[i] = (pow2[i-1] * base) % mod2
	}

	t := next()
	for i := 0; i < t; i++ {
		n := next()
		size := 2 * n
		a := make([]int, size)
		pos := make([][2]int, n)
		seen := make([]bool, n)
		for j := 0; j < size; j++ {
			a[j] = next()
			if !seen[a[j]] {
				pos[a[j]][0] = j
				seen[a[j]] = true
			} else {
				pos[a[j]][1] = j
			}
		}

		h1f := make([]int64, size+1)
		h2f := make([]int64, size+1)
		h1b := make([]int64, size+1)
		h2b := make([]int64, size+1)

		for j := 0; j < size; j++ {
			h1f[j+1] = (h1f[j]*base + int64(a[j])) % mod1
			h2f[j+1] = (h2f[j]*base + int64(a[j])) % mod2
		}
		for j := size - 1; j >= 0; j-- {
			h1b[j] = (h1b[j+1]*base + int64(a[j])) % mod1
			h2b[j] = (h2b[j+1]*base + int64(a[j])) % mod2
		}

		check := func(l, r int) bool {
			len := r - l + 1
			f1 := (h1f[r+1] - h1f[l]*pow1[len]%mod1 + mod1) % mod1
			f2 := (h2f[r+1] - h2f[l]*pow2[len]%mod2 + mod2) % mod2
			b1 := (h1b[l] - h1b[r+1]*pow1[len]%mod1 + mod1) % mod1
			b2 := (h2b[l] - h2b[r+1]*pow2[len]%mod2 + mod2) % mod2
			return f1 == b1 && f2 == b2
		}

		minL, maxR := size, -1
		ans := 0
		for k := n; k >= 1; k-- {
			val := k - 1
			if pos[val][0] < minL {
				minL = pos[val][0]
			}
			if pos[val][1] < minL {
				minL = pos[val][1]
			}
			if pos[val][0] > maxR {
				maxR = pos[val][0]
			}
			if pos[val][1] > maxR {
				maxR = pos[val][1]
			}

			if check(minL, maxR) {
				ans = k
				break
			}
		}
		writer.WriteString(strconv.Itoa(ans) + "\n")
	}
}
