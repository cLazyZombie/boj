package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var cn, cm, n, m int

	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &cn)
		ns := make(map[int]struct{}, cn)

		for c := 0; c < cn; c++ {
			fmt.Scanf("%d", &n)
			ns[n] = struct{}{}
		}

		fmt.Scanf("%d", &cm)
		for c := 0; c < cm; c++ {
			fmt.Scanf("%d", &m)
			if _, ok := ns[m]; ok {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}
	}
}
