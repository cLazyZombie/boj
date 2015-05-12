package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("Add 1, 1", t, func() {
		So(Add("1", "1"), ShouldEqual, "2")
	})

	Convey("Add 2, 1", t, func() {
		So(Add("2", "1"), ShouldEqual, "3")
	})

	Convey("Add 11, 1", t, func() {
		So(Add("11", "1"), ShouldEqual, "12")
	})

	Convey("Add 19, 1", t, func() {
		So(Add("19", "1"), ShouldEqual, "20")
	})

	Convey("Add 99999999999999999999, 1", t, func() {
		So(Add("99999999999999999999", "1"), ShouldEqual, "100000000000000000000")
	})

	Convey("Add 9223372036854775807,  9223372036854775808", t, func() {
		So(Add("9223372036854775807", "9223372036854775808"), ShouldEqual, "18446744073709551615")
	})
}
