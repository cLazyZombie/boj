package main

import (
	//. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//func TestSolver(t *testing.T) {
//	Convey("solver test", t, func() {
//		Convey("case #1", func() {
//			blocks := []Block{
//				Block{Left: 1, Right: 3, Height: 2},
//				Block{Left: 2, Right: 4, Height: 3},
//			}
//			answer := Solve(blocks)
//			So(answer, ShouldEqual, "1 2 2 3 4 0 ")
//		})
//
//		Convey("case #2", func() {
//			blocks := []Block{
//				Block{Left: 1, Right: 3, Height: 2},
//				Block{Left: 3, Right: 5, Height: 3},
//			}
//			answer := Solve(blocks)
//			So(answer, ShouldEqual, "1 2 3 3 5 0 ")
//		})
//
//		Convey("case #3", func() {
//			blocks := []Block{
//				Block{Left: 1, Right: 3, Height: 2},
//				Block{Left: 3, Right: 5, Height: 2},
//			}
//			answer := Solve(blocks)
//			So(answer, ShouldEqual, "1 2 5 0 ")
//		})
//
//		Convey("case #4", func() {
//			blocks := []Block{
//				Block{Left: 1, Right: 3, Height: 2},
//				Block{Left: 1, Right: 5, Height: 3},
//			}
//			answer := Solve(blocks)
//			So(answer, ShouldEqual, "1 3 5 0 ")
//		})
//	})
//}

func BenchmarkSolver(b *testing.B) {
	blocks := []Block{
		Block{Left: 1, Height: 11, Right: 5},
		Block{Left: 2, Height: 6, Right: 7},
		Block{Left: 3, Height: 13, Right: 9},
		Block{Left: 12, Height: 7, Right: 16},
		Block{Left: 14, Height: 3, Right: 25},
		Block{Left: 19, Height: 18, Right: 22},
		Block{Left: 23, Height: 13, Right: 29},
		Block{Left: 24, Height: 4, Right: 28},
	}

	for i := 0; i < b.N; i++ {
		Solve(blocks)
	}
}
