package fastMultiply

import "strconv"

// a * b = (a1 * 10^n + a2) * (b1 * 10^m + b2)
// = a1 * b1 * 10^(n+m) + a1 * b2 * 10^n + a2 * b1 * 10^m + a2 * b2
// n = a 자릿수 / 2, m = b 자릿수 / 2
func Multiply(left, right string) string {
	// 항상 right 자릿수가 더 크게
	if len(left) > len(right) {
		left, right = right, left
	}

	leftLen := len(left)
	rightLen := len(right)

	if leftLen == 0 || rightLen == 0 {
		panic("len == 0")
	}

	if leftLen == 1 && rightLen == 1 {
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		return strconv.Itoa(l * r)
	} else if leftLen == 1 {
		m := rightLen / 2
		b1, b2 := split(right, m)

		first := shift(Multiply(left, b1), m)
		second := Multiply(left, b2)
		return Add(first, second)
	} else {
		n := leftLen / 2
		a1, a2 := split(left, n)

		m := rightLen / 2
		b1, b2 := split(right, m)

		first := shift(Multiply(a1, b1), n+m)
		second := shift(Multiply(a1, b2), n)
		third := shift(Multiply(a2, b1), m)
		fourth := Multiply(a2, b2)

		return AddSlice(first, second, third, fourth)
	}

	return ""
}

func shift(v string, n int) string {
	if v == "0" {
		return "0"
	}

	buf := []byte(v)
	for i := 0; i < n; i++ {
		buf = append(buf, '0')
	}

	return string(buf)
}

func AddSlice(values ...string) string {
	if len(values) < 2 {
		panic("len(values) < 2")
	}

	acc := values[0]
	for i := 1; i < len(values); i++ {
		acc = Add(acc, values[i])
	}

	return acc
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

func byteToInt(b byte) int {
	return int(b - '0')
}

func split(s string, pos int) (string, string) {
	if pos < 1 {
		panic("pos <= 1")
	}

	index := len(s) - pos

	return s[:index], s[index:]
}
