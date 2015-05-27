package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPolygon(t *testing.T) {
	Convey("polygon contain test", t, func() {
		p1 := NewPolygon(Point{0, 0}, Point{3, 0}, Point{3, 3}, Point{0, 3})
		p2 := NewPolygon(Point{1, 1}, Point{2, 0}, Point{2, 2}, Point{1, 2})
		p3 := NewPolygon(Point{4, 0}, Point{7, 0}, Point{7, 3}, Point{4, 3})

		So(p1.CanContain(p2), ShouldBeTrue)
		So(p2.CanContain(p1), ShouldBeFalse)
		So(p3.CanContain(p1), ShouldBeFalse)
	})
}

func TestNode(t *testing.T) {
	Convey("node test", t, func() {
		p1 := NewPolygon(Point{0, 0}, Point{3, 0}, Point{3, 3}, Point{0, 3})
		p2 := NewPolygon(Point{1, 1}, Point{2, 0}, Point{2, 2}, Point{1, 2})
		p3 := NewPolygon(Point{4, 0}, Point{7, 0}, Point{7, 3}, Point{4, 3})

		n1 := NewNode(p1)
		n2 := NewNode(p2)
		n3 := NewNode(p3)

		root := new(Node)
		root.Add(n1)
		root.Add(n2)
		root.Add(n3)

		So(root.Depth(), ShouldEqual, 2)
	})

	Convey("complex test", t, func() {
		p1 := NewPolygon(Point{10, 7}, Point{11, 7}, Point{11, 8}, Point{10, 8})
		p2 := NewPolygon(Point{3, 7}, Point{5, 7}, Point{5, 9}, Point{6, 9}, Point{6, 12}, Point{3, 12})
		p3 := NewPolygon(Point{8, 5}, Point{9, 5}, Point{9, 6}, Point{12, 6}, Point{12, 9}, Point{8, 9})
		p4 := NewPolygon(Point{9, 2}, Point{9, 3}, Point{14, 3}, Point{14, 11}, Point{7, 11}, Point{7, 8}, Point{6, 8}, Point{6, 4}, Point{4, 4}, Point{4, 2})

		n1 := NewNode(p1)
		n2 := NewNode(p2)
		n3 := NewNode(p3)
		n4 := NewNode(p4)

		So(p3.CanContain(p1), ShouldBeTrue)
		So(p4.CanContain(p1), ShouldBeTrue)
		So(p4.CanContain(p3), ShouldBeTrue)

		root := new(Node)
		root.Add(n1)
		root.Add(n2)
		root.Add(n3)
		root.Add(n4)

		So(root.Depth(), ShouldEqual, 3)
	})
}
