package point

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(pt *Point, x float64, y float64) *Point {
	if pt == nil {
		pt = &Point{}
	}

	pt.x = x
	pt.y = y

	return pt
}

func (pt Point) X() float64 {
	return pt.x
}

func (pt Point) Y() float64 {
	return pt.y
}

func DeletePoint(pt *Point) {}

func Distance(p *Point, q *Point) float64 {
	dx := p.X() - q.X()
	dy := p.Y() - q.Y()

	return math.Sqrt(dx*dx + dy*dy)
}
