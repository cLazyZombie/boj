package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolve(t *testing.T) {
	Convey("Test", t, func() {
		So(Solve(0, 0, 13, 40, 0, 37), ShouldEqual, 2)
		So(Solve(0, 0, 3, 0, 7, 4), ShouldEqual, 1)
		So(Solve(0, 0, 2, 0, 7, 4), ShouldEqual, 0)
		So(Solve(1, 1, 1, 1, 1, 5), ShouldEqual, 0)
		So(Solve(1, 1, 1, 1, 1, 1), ShouldEqual, -1)

		So(Solve(0, 0, 10, 8, 0, 1), ShouldEqual, 0)
		So(Solve(0, 0, 10, 8, 0, 2), ShouldEqual, 1)
		So(Solve(0, 0, 10, 8, 0, 3), ShouldEqual, 2)
	})
}
