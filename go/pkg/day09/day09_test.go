package day09_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day09"
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestArea(t *testing.T) {
	cases := []struct {
		p1       day09.Point
		p2       day09.Point
		expected int
	}{
		{
			p1:       day09.Point{X: 2, Y: 5},
			p2:       day09.Point{X: 9, Y: 7},
			expected: 24,
		},
		{
			p1:       day09.Point{X: 7, Y: 1},
			p2:       day09.Point{X: 11, Y: 7},
			expected: 35,
		},
		{
			p1:       day09.Point{X: 7, Y: 3},
			p2:       day09.Point{X: 2, Y: 3},
			expected: 6,
		},
		{
			p1:       day09.Point{X: 2, Y: 5},
			p2:       day09.Point{X: 11, Y: 1},
			expected: 50,
		},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("Area of %d,%d %d,%d should be %d", tc.p1.X, tc.p1.Y, tc.p2.X, tc.p2.Y, tc.expected)
		t.Run(name, func(t *testing.T) {
			area := day09.Area(tc.p1, tc.p2)
			assert.Equal(t, tc.expected, area)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day09.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "50", result)
}

func TestAllInside(t *testing.T) {
	points := []day09.Point{
		{X: 7, Y: 1},
		{X: 11, Y: 1},
		{X: 11, Y: 7},
		{X: 9, Y: 7},
		{X: 9, Y: 5},
		{X: 2, Y: 5},
		{X: 2, Y: 3},
		{X: 7, Y: 3},
	}
	p1 := day09.Point{X: 9, Y: 5}
	p2 := day09.Point{X: 2, Y: 3}

	allInside := day09.AllInside(points, p1, p2, make(map[day09.Point]bool))

	assert.True(t, allInside)
}

func TestPart2(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day09.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "24", result)
}
