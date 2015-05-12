package main

import (
	"github.com/smartystreets/assertions"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var q1 [][]byte = [][]byte{
	[]byte("....."),
	[]byte(".X..."),
	[]byte("....."),
	[]byte(".#..."),
	[]byte("#####"),
}

var a1 [][]byte = [][]byte{
	[]byte("....."),
	[]byte("....."),
	[]byte(".X..."),
	[]byte(".#..."),
	[]byte("#####"),
}

var q2 [][]byte = [][]byte{
	[]byte(".XXXX."),
	[]byte("...X.."),
	[]byte("......"),
	[]byte("....#."),
	[]byte("#...##"),
	[]byte("######"),
}

var a2 [][]byte = [][]byte{
	[]byte("......"),
	[]byte("......"),
	[]byte(".XXXX."),
	[]byte("...X#."),
	[]byte("#...##"),
	[]byte("######"),
}

var q3 [][]byte = [][]byte{
	[]byte("XXX.XXX"),
	[]byte("X.XXX.X"),
	[]byte("X..X..X"),
	[]byte("X.....X"),
	[]byte("......."),
	[]byte(".#...#."),
	[]byte(".##.##."),
	[]byte(".#####."),
	[]byte("#######"),
}

var a3 [][]byte = [][]byte{
	[]byte("......."),
	[]byte("......."),
	[]byte("......."),
	[]byte("......."),
	[]byte("XXX.XXX"),
	[]byte("X#XXX#X"),
	[]byte("X##X##X"),
	[]byte("X#####X"),
	[]byte("#######"),
}

var q4 [][]byte = [][]byte{
	[]byte("..XX"),
	[]byte(".XX."),
	[]byte("XXXX"),
	[]byte("...."),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("####"),
}

var a4 [][]byte = [][]byte{
	[]byte("...."),
	[]byte("..XX"),
	[]byte(".XX."),
	[]byte("XXXX"),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("##.#"),
	[]byte("####"),
}

func TestFall(t *testing.T) {
	Convey("q1, a1", t, func() {
		So(FallingMeteor(len(q1), len(q1[0]), q1), assertions.ShouldResemble, a1)
	})

	Convey("q2, a2", t, func() {
		So(FallingMeteor(len(q2), len(q2[0]), q2), assertions.ShouldResemble, a2)
	})

	Convey("q3, a3", t, func() {
		So(FallingMeteor(len(q3), len(q3[0]), q3), assertions.ShouldResemble, a3)
	})

	Convey("q4, a4", t, func() {
		So(FallingMeteor(len(q4), len(q4[0]), q4), assertions.ShouldResemble, a4)
	})
}

func BenchmarkFall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q5 := make([][]byte, 10000)
		for i := 0; i < len(q5); i++ {
			q5[i] = make([]byte, 3000)
		}

		for y := 0; y < len(q5); y++ {
			for x := 0; x < len(q5[0]); x++ {
				if y == 0 {
					q5[y][x] = 'X'
				} else if y == len(q5)-1 {
					q5[y][x] = '#'
				} else {
					q5[y][x] = '.'
				}
			}
		}
		b.StartTimer()
		FallingMeteor(len(q5), len(q5[0]), q5)
	}
}
