package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Part1(scanner *bufio.Scanner) (string, error) {
	points, err := parsePoints(scanner)
	if err != nil {
		return "", err
	}

	maxArea := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			a := Area(p1, p2)
			if a > maxArea {
				maxArea = a
			}
		}
	}
	return strconv.Itoa(maxArea), nil
}

func parsePoints(scanner *bufio.Scanner) ([]Point, error) {
	points := make([]Point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, fmt.Errorf("invalid x coordinate: %s", split[0])
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, fmt.Errorf("invalid y coordinate: %s", split[1])
		}

		points = append(points, Point{X: x, Y: y})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return points, nil
}

func Area(p1, p2 Point) int {
	w := abs(p2.X-p1.X) + 1
	h := abs(p2.Y-p1.Y) + 1
	return w * h
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part2(scanner *bufio.Scanner) (string, error) {
	points, err := parsePoints(scanner)
	if err != nil {
		return "", err
	}

	maxArea := 0
	cache := make(map[Point]bool)

	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			a := Area(p1, p2)
			if a > maxArea {
				if !AllInside(points, p1, p2, cache) {
					continue
				}

				maxArea = a
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}

func AllInside(polygon []Point, p1, p2 Point, cache map[Point]bool) bool {
	minX := min(p1.X, p2.X)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)

	inside := func(p Point) bool {
		inside, ok := cache[p]
		if !ok {
			inside = Inside(polygon, p)
			cache[p] = inside
		}
		return inside
	}

	for x := minX; x <= maxX; x++ {
		// Test ceiling
		p := Point{X: x, Y: minY}
		if !inside(p) {
			return false
		}

		// Test floor
		p.Y = maxY
		if !inside(p) {
			return false
		}
	}

	for y := minY; y <= maxY; y++ {
		// Test left
		p := Point{X: minX, Y: y}
		if !inside(p) {
			return false
		}

		// Test right
		p.X = maxX
		if !inside(p) {
			return false
		}
	}

	return true
}

// https://en.wikipedia.org/wiki/Collinearity
func inEdge(p1, p2, p Point) bool {
	crossProduct := (p.X-p1.X)*(p2.Y-p1.Y) - (p.Y-p1.Y)*(p2.X-p1.X)
	if crossProduct != 0 {
		return false
	}

	minX := min(p1.X, p2.X)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)

	return p.X >= minX && p.X <= maxX && p.Y >= minY && p.Y <= maxY
}

// https://en.wikipedia.org/wiki/Ray_casting
func Inside(polygon []Point, p Point) bool {
	inside := false
	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]

		if inEdge(p1, p2, p) {
			return true
		}

		intersects := (p1.Y > p.Y) != (p2.Y > p.Y) &&
			p.X < (p2.X-p1.X)*(p.Y-p1.Y)/(p2.Y-p1.Y)+p1.X
		if intersects {
			inside = !inside
		}
	}

	return inside
}
