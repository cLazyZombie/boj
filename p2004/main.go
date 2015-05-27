package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Printf("%d", Solve(n, m))
}

func Solve(n, m int) int {
	twoCount := TwoCount(n) - TwoCount(m) - TwoCount(n-m)
	fiveCount := FiveCount(n) - FiveCount(m) - FiveCount(n-m)
	count := min(twoCount, fiveCount)
	return count
}

func FiveCount(n int) int {
	var count int
	var i int64
	for i = 5; int64(n) >= i; i *= 5 {
		count += n / int(i)
	}

	return int(count)
}

func TwoCount(n int) int {
	var count int
	var i int64
	for i = 2; int64(n) >= i; i *= 2 {
		count += n / int(i)
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
