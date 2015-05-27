package main

import "fmt"

func main() {
	var n, s int
	fmt.Scanf("%d %d", &n, &s)

	values := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &(values[i]))
		if err != nil {
			fmt.Printf("0")
			return
		}
	}

	fmt.Printf("%d", Solve(values, s))
}

var (
	sets []int
	goal int
)

func Solve(s []int, g int) int {
	sets = s
	goal = g
	return solve(0, 0, false)
}

func solve(now, sum int, chosen bool) int {
	if now == len(sets)-1 {
		answer := 0

		if sum+sets[now] == goal {
			answer++
		}

		if sum == goal && chosen {
			answer++
		}

		return answer
	}

	return solve(now+1, sum+sets[now], true) + solve(now+1, sum, chosen)
}
