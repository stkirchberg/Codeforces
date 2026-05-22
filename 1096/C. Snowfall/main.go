package main

import (
	"bufio"
	"os"
	"strconv"
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

	t := next()
	for i := 0; i < t; i++ {
		n := next()
		v6 := make([]int, 0)
		v2 := make([]int, 0)
		v3 := make([]int, 0)
		v0 := make([]int, 0)

		for j := 0; j < n; j++ {
			val := next()
			if val%6 == 0 {
				v6 = append(v6, val)
			} else if val%2 == 0 {
				v2 = append(v2, val)
			} else if val%3 == 0 {
				v3 = append(v3, val)
			} else {
				v0 = append(v0, val)
			}
		}

		first := true
		for _, v := range v6 {
			if !first {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(v))
			first = false
		}
		for _, v := range v2 {
			if !first {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(v))
			first = false
		}
		for _, v := range v0 {
			if !first {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(v))
			first = false
		}
		for _, v := range v3 {
			if !first {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(v))
			first = false
		}
		writer.WriteString("\n")
	}
}
