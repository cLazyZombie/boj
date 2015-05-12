package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//	var row, column int
	//	fmt.Scanf("%d %d", &row, &column)
	//
	//	pic := make([][]byte, row)
	//	for i := 0; i < row; i++ {
	//		var s string
	//		fmt.Scanf("%s", &s)
	//		pic[i] = []byte(s)
	//	}
	//
	//	result := FallingMeteor(row, column, pic)
	//	for i := 0; i < row; i++ {
	//		fmt.Println(string(result[i]))
	//	}

	FallingMeteor2()
}

const (
	METEOR byte = byte('X')
	GROUND byte = byte('#')
	AIR    byte = byte('.')
)

type Pos struct {
	X, Y int
}

type MinMax struct {
	Min, Max int
}

func FallingMeteor(row, column int, pic [][]byte) [][]byte {
	meteorMinMax, maxLine := getMeteorMinMax(row, column, pic)

	groundMax := getGroundMax(row, column, pic, maxLine)

	// just return pic if no meteor
	if len(meteorMinMax) == 0 {
		return pic
	}

	// get min distance
	minDistance := row
	for x, meteor := range meteorMinMax {
		distance := groundMax[x] - meteor.Max - 1
		if distance < minDistance {
			minDistance = distance
		}
	}

	// falling
	for x, meteor := range meteorMinMax {
		startY := meteor.Min
		endY := meteor.Max

		for y := endY; y >= startY; y-- {
			if pic[y][x] == METEOR {
				pic[y+minDistance][x] = METEOR
				pic[y][x] = AIR
			}
		}
	}

	return pic
}

func getMeteorMinMax(row, column int, pic [][]byte) (map[int]*MinMax, int) {
	result := make(map[int]*MinMax)
	maxLine := 0
	traversed := make([]bool, row*column)

	// get meteor height
	stack := make([]Pos, 0)

Loop:
	for y := 0; y < row; y++ {
		for x := 0; x < column; x++ {
			if pic[y][x] == METEOR {
				stack = append(stack, Pos{x, y})
				break Loop
			}
		}
	}

	// set traversed
	traversed[stack[0].X+stack[0].Y*column] = true

	for len(stack) > 0 {
		// pop
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// set min max
		minMax, ok := result[pos.X]
		if !ok {
			minMax = &MinMax{pos.Y, pos.Y}
			result[pos.X] = minMax
		} else {
			if pos.Y < minMax.Min {
				minMax.Min = pos.Y
			}

			if pos.Y > minMax.Max {
				minMax.Max = pos.Y
			}
		}

		// set meteor line
		if pos.Y > maxLine {
			maxLine = pos.Y
		}

		// check left
		if pos.X > 0 {
			x, y := pos.X-1, pos.Y
			if !traversed[x+y*column] && pic[y][x] == METEOR {
				traversed[x+y*column] = true
				stack = append(stack, Pos{x, y})
			}
		}

		// check right
		if pos.X < column-1 {
			x, y := pos.X+1, pos.Y
			if !traversed[x+y*column] && pic[y][x] == METEOR {
				traversed[x+y*column] = true
				stack = append(stack, Pos{x, y})
			}
		}

		// check up
		if pos.Y > 0 {
			x, y := pos.X, pos.Y-1
			if !traversed[x+y*column] && pic[y][x] == METEOR {
				traversed[x+y*column] = true
				stack = append(stack, Pos{x, y})
			}
		}

		// check down
		if pos.Y < row-1 {
			x, y := pos.X, pos.Y+1
			if !traversed[x+y*column] && pic[y][x] == METEOR {
				traversed[x+y*column] = true
				stack = append(stack, Pos{x, y})
			}
		}
	}

	return result, maxLine
}

func getGroundMax(row, column int, pic [][]byte, meteorMaxLine int) []int {
	result := make([]int, column)

	for x := 0; x < column; x++ {
		for y := meteorMaxLine + 1; y < row; y++ {
			if pic[y][x] == GROUND {
				result[x] = y
				break
			}
		}
	}

	return result
}

func FallingMeteor2() {
	var row, column int
	fmt.Scanf("%d %d", &row, &column)

	meteors := make([]int, column)
	grounds := make([]int, column)
	for i := 0; i < column; i++ {
		meteors[i] = -1
		grounds[i] = row
	}

	r := bufio.NewReader(os.Stdin)
	mpic := make([][]byte, row)
	opic := make([][]byte, row)
	for i := 0; i < row; i++ {
		mpic[i] = make([]byte, column)
		opic[i] = make([]byte, column)
	}

	for y := 0; y < row; y++ {
		bytes, _, _ := r.ReadLine()

		for x := 0; x < column; x++ {
			b := bytes[x]
			switch b {
			case METEOR:
				mpic[y][x] = b
				if y > meteors[x] {
					meteors[x] = y
				}
				opic[y][x] = AIR

			case GROUND:
				opic[y][x] = b
				if y < grounds[x] {
					grounds[x] = y
				}

			case AIR:
				opic[y][x] = b
			}
		}
	}

	minDistance := row
	for x := 0; x < column; x++ {
		if meteors[x] == -1 {
			continue
		}

		d := grounds[x] - meteors[x] - 1
		if d < minDistance {
			minDistance = d
		}
	}

	// print
	for y := 0; y < row; y++ {
		for x := 0; x < column; x++ {
			offY := y - minDistance
			if offY >= 0 && mpic[offY][x] == METEOR {
				fmt.Printf("%c", METEOR)
			} else {
				fmt.Printf("%c", opic[y][x])
			}
		}
		fmt.Println()
	}
}
