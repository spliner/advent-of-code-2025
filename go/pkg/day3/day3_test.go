package day3_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day3"
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestMaxJoltage(t *testing.T) {
	cases := []struct {
		input    string
		n        int
		expected int64
	}{
		{
			input:    "987654321111111",
			n:        2,
			expected: 98,
		},
		{
			input:    "987654321111111",
			n:        12,
			expected: 987654321111,
		},
		{
			input:    "811111111111119",
			n:        2,
			expected: 89,
		},
		{
			input:    "811111111111119",
			n:        12,
			expected: 811111111119,
		},
		{
			input:    "234234234234278",
			n:        2,
			expected: 78,
		},
		{
			input:    "234234234234278",
			n:        12,
			expected: 434234234278,
		},
		{
			input:    "818181911112111",
			n:        2,
			expected: 92,
		},
		{
			input:    "818181911112111",
			n:        12,
			expected: 888911112111,
		},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%s with %d digits should return %d", tc.input, tc.n, tc.expected)
		t.Run(name, func(t *testing.T) {
			joltage := day3.MaxJoltage(tc.input, tc.n)
			assert.Equal(t, tc.expected, joltage)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day3.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "357", result)
}

func TestPart2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day3.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "3121910778619", result)
}
