package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolver(t *testing.T) {
	Convey("10 => 5", t, func() {
		So(Solve(2), ShouldEqual, 1)
		So(Solve(4), ShouldEqual, 2)
		So(Solve(6), ShouldEqual, 3)
		So(Solve(10), ShouldEqual, 5)
		So(Solve(11), ShouldEqual, 10)
	})

	Convey("1000000000 => 500000000", t, func() {
		So(Solve(1000000000), ShouldEqual, 500000000)
	})
}
