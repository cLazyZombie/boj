package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for tc := 0; tc < t; tc++ {
		var m int
		fmt.Scanf("%d", &m)

		root := new(Node)

		for mc := 0; mc < m; mc++ {
			var p int

			for p == 0 {
				fmt.Scanf("%d", &p)
			}

			points := make([]Point, p)
			for pc := 0; pc < p; pc++ {
				fmt.Scanf("%d", &(points[pc].X))
				fmt.Scanf("%d", &(points[pc].Y))
			}

			poly := NewPolygon(points...)
			node := NewNode(poly)

			root.Add(node)
		}

		fmt.Printf("%d\n", root.Depth())
	}
}

type Point struct {
	X, Y int
}

type Polygon struct {
	Points   []Point
	Min, Max Point
}

func NewPolygon(points ...Point) *Polygon {
	poly := new(Polygon)

	max := Point{0, 0}
	min := Point{2000000000, 2000000000}

	poly.Points = points
	for i, _ := range points {
		if min.X > points[i].X {
			min.X = points[i].X
		}

		if min.Y > points[i].Y {
			min.Y = points[i].Y
		}

		if max.X < points[i].X {
			max.X = points[i].X
		}

		if max.Y < points[i].Y {
			max.Y = points[i].Y
		}
	}

	poly.Min, poly.Max = min, max

	// 세로부터 시작하게 순서 수정
	if poly.Points[0].X != poly.Points[1].X {
		poly.Points = append(poly.Points, poly.Points[0])[1:]
	}

	return poly
}

func (poly *Polygon) CanContain(other *Polygon) bool {
	if poly == nil {
		return true
	}

	if other == nil {
		return false
	}

	// boundary check
	if !IsOverlapped(poly.Min.X, poly.Max.X, other.Min.X, other.Max.X) {
		return false
	}

	if !IsOverlapped(poly.Min.Y, poly.Max.Y, other.Min.Y, other.Max.Y) {
		return false
	}

	// real check
	for i := 0; i < len(other.Points); i += 2 {
		next := (i + 1) % len(other.Points)

		if !poly.CheckVertical(other.Points[i], other.Points[next]) {
			return false
		}
	}

	return true
}

func (poly *Polygon) CheckVertical(p1, p2 Point) bool {
	var overlapLeft, overlapRight bool

	for i := 0; i < len(poly.Points); i += 2 {
		next := (i + 1) % len(poly.Points)

		if poly.Points[i].X < p1.X {
			if !overlapLeft && IsOverlapped(p1.Y, p2.Y, poly.Points[i].Y, poly.Points[next].Y) {
				if overlapRight {
					return true
				}

				overlapLeft = true
			}
		} else if poly.Points[i].X > p1.X {
			if !overlapRight && IsOverlapped(p1.Y, p2.Y, poly.Points[i].Y, poly.Points[next].Y) {
				if overlapLeft {
					return true
				}

				overlapRight = true
			}
		}
	}

	return false
}

func IsOverlapped(p1, p2, p3, p4 int) bool {
	start1, end1 := MinMax(p1, p2)
	start2, end2 := MinMax(p3, p4)

	return Max(start1, start2) < Min(end1, end2)
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}

type Node struct {
	Polygon  *Polygon
	Parent   *Node
	Children []*Node
}

func NewNode(p *Polygon) *Node {
	node := new(Node)
	node.Polygon = p
	return node
}

func (n *Node) Add(node *Node) bool {
	// parent
	if node.Polygon.CanContain(n.Polygon) {
		node.Parent = n.Parent
		n.Parent = node
		node.Children = append(node.Children, n)

		if node.Parent != nil {
			node.Parent.Children = append(node.Parent.Children, node)
			for i := range node.Parent.Children {
				if node.Parent.Children[i] == n {
					node.Parent.Children = append(node.Parent.Children[:i], node.Parent.Children[i+1:]...)
					break
				}
			}
		}
		return true
	}

	// child
	if n.Polygon.CanContain(node.Polygon) {
		for i := range n.Children {
			if n.Children[i].Add(node) {
				return true
			}
		}

		// not belong to children
		n.Children = append(n.Children, node)
		node.Parent = n
		return true
	}

	return false
}

func (n *Node) Depth() int {
	maxDepth := 0
	for _, child := range n.Children {
		childDepth := child.Depth()
		if childDepth > maxDepth {
			maxDepth = childDepth
		}
	}

	if n.Parent == nil {
		return maxDepth
	}

	return 1 + maxDepth
}
