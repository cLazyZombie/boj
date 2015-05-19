package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("%d", Solve(n))
}

var primes []int

func Solve(n int) int {
	makePrimes(n)

	for i := 0; i < len(primes); i++ {
		if primes[i] >= n {
			return n - 1
		}

		if n%primes[i] == 0 {
			b := n / primes[i]
			return n - b
		}
	}

	return n - 1
}

const (
	NOT_INITIALIZED byte = iota
	PRIME
	NOT_PRIME
)

func makePrimes(n int) {
	sqrtN := int(math.Sqrt(float64(n))) + 1
	primes = make([]int, 0, sqrtN+1)
	isPrimes := make([]byte, sqrtN+1)

	isPrimes[0] = NOT_PRIME
	isPrimes[1] = NOT_PRIME

	for i := 2; i <= sqrtN; i++ {
		if isPrimes[i] == NOT_PRIME {
			continue
		}

		isPrimes[i] = PRIME
		primes = append(primes, i)

		for j := 2; i*j <= sqrtN; j++ {
			isPrimes[i*j] = NOT_PRIME
		}
	}
}
