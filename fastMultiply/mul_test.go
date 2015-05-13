package fastMultiply

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiply(t *testing.T) {
	Convey("8 * 9 = 72", t, func() {
		So(Multiply("8", "9"), ShouldEqual, "72")
	})

	Convey("2 * 12 = 24", t, func() {
		So(Multiply("2", "12"), ShouldEqual, "24")
	})

	Convey("123 * 2 = 246", t, func() {
		So(Multiply("123", "2"), ShouldEqual, "246")
	})

	Convey("12 * 12 = 144", t, func() {
		So(Multiply("12", "12"), ShouldEqual, "144")
	})

	Convey("0 * 123 = 0", t, func() {
		So(Multiply("0", "123"), ShouldEqual, "0")
	})

	Convey("12 * 345 = 4140", t, func() {
		So(Multiply("12", "345"), ShouldEqual, "4140")
	})

	Convey("123456789 * 987654321 = 121932631112635269", t, func() {
		So(Multiply("123456789", "987654321"), ShouldEqual, "121932631112635269")
	})
}

func TestAdd(t *testing.T) {
	Convey("12 + 23 = 35", t, func() {
		So(Add("12", "23"), ShouldEqual, "35")
	})

	Convey("20 + 4 = 24", t, func() {
		So(Add("20", "4"), ShouldEqual, "24")
	})
}

func TestShift(t *testing.T) {
	Convey("12 shift 2 = 1200", t, func() {
		So(shift("12", 2), ShouldEqual, "1200")
	})
}
