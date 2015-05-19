package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var (
		x1, y1, r1 int
		x2, y2, r2 int
	)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %d %d %d %d", &x1, &y1, &r1, &x2, &y2, &r2)

		fmt.Printf("%d", Solve(x1, y1, r1, x2, y2, r2))
		if i != t-1 {
			fmt.Println()
		}
	}
}

func Solve(x1, y1, r1, x2, y2, r2 int) int {
	dx := x1 - x2
	dy := y1 - y2

	dsq := Sq(dx) + Sq(dy)

	if dsq == 0 {
		if r1 == r2 {
			return -1
		} else {
			return 0
		}
	}

	if d2 := Sq(r2 - r1); dsq <= d2 { // inner
		if dsq == d2 {
			return 1
		} else {
			return 0
		}
	} else { // outer
		rsq := Sq(r1 + r2)

		if dsq < rsq {
			return 2
		} else if dsq == rsq {
			return 1
		}
	}

	return 0
}

func Sq(a int) int {
	return a * a
}
