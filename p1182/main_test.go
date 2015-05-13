package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolve(t *testing.T) {
	Convey("1 0, 1 => 2", t, func() {
		So(Solve([]int{1, 0}, 1), ShouldEqual, 2)
	})

	Convey("-7 -3 -2 5 8, 0 => 1", t, func() {
		So(Solve([]int{-7, -3, -2, 5, 8}, 0), ShouldEqual, 1)
	})

	Convey("-7 -3 -2 5 8 -7 -3 -2 5 8, 0 => 33", t, func() {
		So(Solve([]int{-7, -3, -2, 5, 8, -7, -3, -2, 5, 8}, 0), ShouldEqual, 33)
	})
}
