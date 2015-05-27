package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMul5(t *testing.T) {
	Convey("testing mul 5", t, func() {
		So(Mul5([]int{5}), ShouldResemble, []int{25})
		So(Mul5([]int{50000000}), ShouldResemble, []int{50000000, 2})
	})
}

func TestSolver(t *testing.T) {
	Convey("testing solver", t, func() {
		So(Pow5(1), ShouldResemble, []int{5})

		// 5^20 = 95367431640625
		So(Pow5(20), ShouldResemble, []int{31640625, 953674})

		// 5^100 = 7888609052210118054117285652827862296732064351090230047702789306640625
		So(Pow5(100), ShouldResemble, []int{6640625, 77027893, 9023004, 32064351, 78622967, 28565282, 18054117, 90522101, 788860})
	})
}

func TestZeroCount(t *testing.T) {
	Convey("testing zero count", t, func() {
		So(ZeroCount(10, []int{1}), ShouldEqual, 9)
		So(ZeroCount(10, []int{1, 1}), ShouldEqual, 1)
		So(ZeroCount(10, []int{1, 11}), ShouldEqual, 0)
	})
}
