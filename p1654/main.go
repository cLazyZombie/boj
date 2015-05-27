package main

import "fmt"

func main() {
	var k, n int
	fmt.Scanf("%d %d", &k, &n)

	lines := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &(lines[i]))
	}

	fmt.Printf("%d", Solve(lines, n))
}

func Solve(lines []int, count int) int {
	min, max := 1, 1
	for _, line := range lines {
		if line > max {
			max = line + 1
		}
	}

	for {
		l := (min + max) / 2

		if canCut(lines, l, count) {
			min = l
		} else {
			max = l
		}

		if min >= max-1 {
			return min
		}
	}

	return 0
}

func canCut(lines []int, l int, count int) bool {
	curCount := 0

	for _, line := range lines {
		curCount += line / l
		if curCount >= count {
			return true
		}
	}

	return curCount >= count
}
