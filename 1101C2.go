//go:build ignore

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if !scanner.Scan() {
		return
	}
	t, _ := strconv.Atoi(scanner.Text())

	for tc := 0; tc < t; tc++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		u := scanner.Text()

		X := x
		S := 0
		ans := 0
		A_as_E := 0

		for i := 0; i < n; i++ {
			c := u[i]
			if c == 'I' {
				if X > 0 {
					X--
					S += s - 1
					ans++
				}
			} else if c == 'E' {
				if S > 0 {
					S--
					ans++
				} else if A_as_E > 0 && X > 0 {
					A_as_E--
					X--
					S += s
					S--
					ans++
				}
			} else if c == 'A' {
				if S > 0 {
					S--
					A_as_E++
					ans++
				} else if X > 0 {
					X--
					S += s - 1
					ans++
				}
			}
		}
		out.WriteString(strconv.Itoa(ans) + "\n")
	}
}
