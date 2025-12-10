package day10_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day10"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day10.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "7", result)
}
