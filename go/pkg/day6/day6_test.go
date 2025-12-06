package day6_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day6"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day6.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "4277556", result)
}

func TestPart2(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day6.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "3263827", result)
}
