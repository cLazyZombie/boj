package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolver(t *testing.T) {
	Convey("general case", t, func() {
		So(Solve(1), ShouldEqual, 2)
		So(Solve(12), ShouldEqual, 10)
		So(Solve(13), ShouldEqual, 12)
		So(Solve(240), ShouldEqual, 200)
		So(Solve(3600), ShouldEqual, 3000)
	})
}
