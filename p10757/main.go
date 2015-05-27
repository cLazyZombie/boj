package main

import (
	"fmt"
	"strconv"
)

func main() {
	var left, right string
	fmt.Scanf("%s %s", &left, &right)
	fmt.Printf(Add(left, right))
}

func Add(left, right string) string {
	maxLen := len(left)
	if len(right) > maxLen {
		maxLen = len(right)
	}

	result := make([]byte, maxLen)
	carry := 0

	for i := 0; i < maxLen; i++ {
		leftIndex := len(left) - i - 1
		rightIndex := len(right) - i - 1

		val := 0

		if leftIndex >= 0 {
			v, _ := strconv.Atoi(string(left[leftIndex]))
			val += v
		}

		if rightIndex >= 0 {
			v, _ := strconv.Atoi(string(right[rightIndex]))
			val += v
		}

		val += carry

		if val > 9 {
			carry = 1
			val -= 10
		} else {
			carry = 0
		}

		result[maxLen-i-1] = strconv.Itoa(val)[0]
	}

	if carry != 0 {
		return "1" + string(result)
	}

	return string(result)
}
