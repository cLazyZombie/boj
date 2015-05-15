package main

import "fmt"

func main() {
	var k int
	fmt.Scanf("%d", &k)

	for i := 0; i < k; i++ {
		fmt.Printf("1")
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("0")
	}
}
