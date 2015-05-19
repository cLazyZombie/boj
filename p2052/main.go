package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	answer := Pow5(n)

	leng := len(answer)
	zeroCount := ZeroCount(n, answer)

	fmt.Printf("0.")

	for i := 0; i < zeroCount; i++ {
		fmt.Printf("0")
	}

	for i := leng - 1; i >= 0; i-- {
		if i == leng-1 {
			fmt.Printf("%d", answer[i])
		} else {
			fmt.Printf("%08d", answer[i])
		}
	}
}

const LIMIT_VALUE = 100000000

func Pow5(n int) []int {
	result := make([]int, 1, n/8+1)
	result[0] = 5

	for i := 1; i < n; i++ {
		result = Mul5(result)
	}

	return result
}

func Mul5(numbers []int) []int {
	carry := 0
	for i := 0; i < len(numbers); i++ {
		v := numbers[i]*5 + carry
		if v >= LIMIT_VALUE {
			carry = v / LIMIT_VALUE
			v = v % LIMIT_VALUE
		} else {
			carry = 0
		}
		numbers[i] = v
	}

	if carry > 0 {
		numbers = append(numbers, carry)
	}

	return numbers
}

func ZeroCount(n int, numbers []int) int {
	count := 0

	last := numbers[len(numbers)-1]
	for last != 0 {
		count++
		last = last / 10
	}

	count += (len(numbers) - 1) * 8

	return n - count
}
