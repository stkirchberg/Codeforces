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
		var n, x, s int
		fmt.Fscan(reader, &n, &x, &s)

		var u string
		fmt.Fscan(reader, &u)

		count := make([]int, s+1)
		size := 0

		highPtr := s - 1
		lowPtr := 0

		addSlope := func(v int) {
			count[v]++
			size++
			if v < lowPtr {
				lowPtr = v
			}
			if v < s && v > highPtr {
				highPtr = v
			}
			if size > x {
				for lowPtr <= s && count[lowPtr] == 0 {
					lowPtr++
				}
				if lowPtr <= s {
					count[lowPtr]--
					size--
				}
			}
		}

		for i := 0; i < n; i++ {
			c := u[i]
			if c == 'I' {
				if count[0] > 0 || size < x {
					addSlope(1)
				}
			} else if c == 'E' {
				for highPtr >= 0 && count[highPtr] == 0 {
					highPtr--
				}
				if highPtr >= 0 {
					count[highPtr]--
					count[highPtr+1]++
					if highPtr+1 < s {
						highPtr = highPtr + 1
					}
				}
			} else if c == 'A' {
				for highPtr >= 0 && count[highPtr] == 0 {
					highPtr--
				}
				if highPtr >= 0 {
					count[highPtr]--
					count[highPtr+1]++
					if highPtr+1 < s {
						highPtr = highPtr + 1
					}
				}

				addSlope(0)
			}
		}

		var ans int64 = 0
		for v := 1; v <= s; v++ {
			ans += int64(v) * int64(count[v])
		}
		fmt.Fprintln(writer, ans)
	}
}

func main() {
	solve()
}
