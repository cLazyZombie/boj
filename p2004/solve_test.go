package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolver(t *testing.T) {
	Convey("testing", t, func() {
		Convey("25 C 12 => 2", func() {
			So(Solve(25, 12), ShouldEqual, 2)
		})

		Convey("2000000000 C 12 => 8", func() {
			So(Solve(2000000000, 12), ShouldEqual, 8)
		})

		Convey("2000000000 C 1000000000 => 8", func() {
			So(Solve(2000000000, 1000000000), ShouldEqual, 1)
		})
	})
}
