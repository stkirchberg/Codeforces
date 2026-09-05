//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Move struct {
	id, from, to int
}

var (
	moves []Move
	pos   []int
	a     []int
)

func reach(m, pegX, countX, pegY, countY int) {
	if m == 0 {
		return
	}
	if pos[m] == pegX && countX > 0 {
		reach(m-1, pegX, countX-1, pegY, countY)
	} else if pos[m] == pegY && countY > 0 {
		reach(m-1, pegX, countX, pegY, countY-1)
	} else {
		targetPeg := pegY
		if countX > 0 {
			targetPeg = pegX
		}
		src := pos[m]
		dst := targetPeg
		aux := 6 - src - dst

		reach(m-1, src, a[m], aux, m-1-a[m])

		moves = append(moves, Move{m, src, dst})
		pos[m] = dst

		if targetPeg == pegX {
			reach(m-1, pegX, countX-1, pegY, countY)
		} else {
			reach(m-1, pegX, countX, pegY, countY-1)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		a = make([]int, n+1)
		pos = make([]int, n+1)
		possible := true

		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &a[j])
			if a[j] >= j {
				possible = false
			}
			pos[j] = 1
		}

		if !possible {
			fmt.Fprintln(out, "NO")
			continue
		}

		moves = make([]Move, 0, 1<<n)
		reach(n, 3, n, 2, 0)

		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, len(moves))
		for _, mv := range moves {
			fmt.Fprintln(out, mv.id, mv.from, mv.to)
		}
	}
}
