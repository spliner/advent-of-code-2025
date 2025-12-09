package day08_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day08"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day08.Part1WithConnections(scanner, 10)

	assert.Nil(t, err)
	assert.Equal(t, "40", result)
}

func TestPart2(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day08.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "25272", result)
}
