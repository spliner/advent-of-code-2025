package day04

import "iter"

type Point struct {
	X int
	Y int
}

func (p Point) AdjacentPoints() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		if !yield(Point{X: p.X - 1, Y: p.Y - 1}) {
			return
		}

		if !yield(Point{X: p.X, Y: p.Y - 1}) {
			return
		}

		if !yield(Point{X: p.X + 1, Y: p.Y - 1}) {
			return
		}

		if !yield(Point{X: p.X - 1, Y: p.Y}) {
			return
		}

		if !yield(Point{X: p.X + 1, Y: p.Y}) {
			return
		}

		if !yield(Point{X: p.X - 1, Y: p.Y + 1}) {
			return
		}

		if !yield(Point{X: p.X, Y: p.Y + 1}) {
			return
		}

		yield(Point{X: p.X + 1, Y: p.Y + 1})
	}
}
