package day04_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day04"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day04.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "13", result)
}

func TestPart2(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day04.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "43", result)
}
