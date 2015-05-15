package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var b Block
	blocks := make([]Block, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &b.Left, &b.Height, &b.Right)
		blocks = append(blocks, b)
	}

	answer := Solve(blocks)
	fmt.Print(answer)
}

type Block struct {
	Left, Height, Right int
}

var lines map[int][]int

func Solve(buildings []Block) string {
	answer := new(bytes.Buffer)

	lines = make(map[int][]int)
	for _, b := range buildings {
		Append(b.Left, b.Height)
		Append(b.Right, -b.Height)
	}

	points := make([]int, 0, len(buildings)*2)
	for i, _ := range lines {
		points = append(points, i)
	}
	sort.Ints(points)

	cur := make([]int, 1, len(buildings)) // insert 0 for ground

	h := 0
	for _, pos := range points {
		line := lines[pos]
		for _, v := range line {
			if v > 0 {
				cur = append(cur, v)
				sort.Ints(cur)
			} else {
				i := sort.SearchInts(cur, -v)
				cur = append(cur[:i], cur[i+1:]...)
			}
		}

		// get max height from current and check changed from prev height
		if cur[len(cur)-1] != h {
			h = cur[len(cur)-1]
			fmt.Fprintf(answer, "%d %d ", pos, h)
		}
	}

	return answer.String()
}

func Append(pos, height int) {
	line, ok := lines[pos]
	if !ok {
		line = make([]int, 0, 10)
	}
	line = append(line, height)
	lines[pos] = line
}
