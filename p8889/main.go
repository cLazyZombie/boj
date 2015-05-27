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
	mx1 := Max(poly.Min.X, other.Min.X)
	mx2 := Min(poly.Max.X, other.Max.X)

	if mx1 > mx2 {
		return false
	}

	my1 := Max(poly.Min.Y, other.Min.Y)
	my2 := Min(poly.Max.Y, other.Max.Y)

	if my1 > my2 {
		return false
	}

	// real check
	for i, _ := range other.Points {
		next := i + 1
		if next == len(other.Points) {
			next = 0
		}

		// vertical line
		if i%2 == 0 {
			if !poly.CheckHorizon(other.Points[i], other.Points[next]) {
				return false
			}
		} else {
			if !poly.CheckVertical(other.Points[i], other.Points[next]) {
				return false
			}
		}
	}

	return true
}

func (poly *Polygon) CheckVertical(p1, p2 Point) bool {
	start := 2000000000
	end := 0

	count := 0

	for i := 0; i < len(poly.Points); i += 2 {
		next := i + 1
		if next == len(poly.Points) {
			next = 0
		}

		yStart, yEnd := MinMax(poly.Points[i].Y, poly.Points[next].Y)
		if p1.Y < yStart || p1.Y > yEnd {
			continue
		}

		start = Min(poly.Points[i].X, start)
		end = Max(poly.Points[i].X, end)
		count++

		if count >= 2 {
			break
		}
	}

	if start > end {
		return false
	}

	pStart, pEnd := MinMax(p1.X, p2.X)

	if Max(start, pStart) >= Min(end, pEnd) {
		return false
	}

	return true
}

func (poly *Polygon) CheckHorizon(p1, p2 Point) bool {
	start := 2000000000
	end := 0

	count := 0

	for i := 1; i < len(poly.Points); i += 2 {
		next := i + 1
		if next == len(poly.Points) {
			next = 0
		}

		xStart, xEnd := MinMax(poly.Points[i].X, poly.Points[next].X)
		if p1.X < xStart || p1.X > xEnd {
			continue
		}

		start = Min(poly.Points[i].Y, start)
		end = Max(poly.Points[i].Y, end)
		count++

		if count >= 2 {
			break
		}
	}

	if start > end {
		return false
	}

	pStart, pEnd := MinMax(p1.Y, p2.Y)

	if Max(start, pStart) >= Min(end, pEnd) {
		return false
	}

	return true
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
