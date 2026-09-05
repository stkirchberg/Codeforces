//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	if !scanner.Scan() {
		return
	}
	t, _ := strconv.Atoi(scanner.Text())

	for tc := 0; tc < t; tc++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		a := make([]int, n)
		b := make([]int, n)

		maxVal := 0
		for i := 0; i < n; i++ {
			scanner.Scan()
			a[i], _ = strconv.Atoi(scanner.Text())
			if a[i] > maxVal {
				maxVal = a[i]
			}
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			b[i], _ = strconv.Atoi(scanner.Text())
			if b[i] > maxVal {
				maxVal = b[i]
			}
		}

		low := 1
		high := maxVal
		ans := 1

		for low <= high {
			mid := low + (high-low)/2

			twos := 0
			zeroBlocks := 0
			inZeroBlock := false

			for i := 0; i < n; i++ {
				val := 0
				if a[i] >= mid {
					val++
				}
				if b[i] >= mid {
					val++
				}

				if val == 2 {
					twos++
					inZeroBlock = false
				} else if val == 0 {
					if !inZeroBlock {
						zeroBlocks++
						inZeroBlock = true
					}
				} else {
				}
			}

			if twos > zeroBlocks {
				ans = mid
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		fmt.Fprintln(writer, ans)
	}
}
