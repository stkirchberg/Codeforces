package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		scanner.Scan()
		s := scanner.Text()

		open := 0
		close := 0

		for _, char := range s {
			if char == '(' {
				open++
			} else {
				close++
			}
		}

		if open == close {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
