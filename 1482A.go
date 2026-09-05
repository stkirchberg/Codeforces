//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var t int

	_, err := fmt.Scan(&t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var a, b int

		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}

		fmt.Println(a * b)
	}
}
