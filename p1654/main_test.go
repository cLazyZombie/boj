package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolve(t *testing.T) {
	Convey("Test", t, func() {
		So(Solve([]int{802, 743, 457, 539}, 11), ShouldEqual, 200)

		So(Solve([]int{100}, 10), ShouldEqual, 10)

		So(Solve([]int{100}, 10), ShouldEqual, 10)
	})
}
