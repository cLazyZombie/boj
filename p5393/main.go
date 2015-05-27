package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var k int64
		fmt.Scanf("%d", &k)
		fmt.Printf("%d\n", Solve(k))
	}
}

func Solve(k int64) int64 {
	// odd
	m1 := (k - 1) / 3
	if m1%2 == 0 {
		m1 += 1
	} else {
		m1 += 2
	}

	m2 := k
	if m2%2 == 0 {
		m2 -= 1
	}

	oddCount := (m2 - m1 + 1)
	if oddCount%2 == 0 {
		oddCount /= 2
	} else {
		oddCount = oddCount/2 + 1
	}

	// even
	n1 := k/2 + 1
	n2 := k

	evenCount := n2 - n1 + 1

	return oddCount + evenCount
}
