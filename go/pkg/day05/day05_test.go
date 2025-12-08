package day05_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day05"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day05.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "3", result)
}

func TestPart2(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day05.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "14", result)
}
