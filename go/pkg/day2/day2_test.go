package day2_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day2"
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestIsSplitStringValid(t *testing.T) {
	cases := []struct {
		id      string
		isValid bool
	}{
		{
			id:      "1",
			isValid: true,
		},
		{
			id:      "101",
			isValid: true,
		},
		{
			id:      "11",
			isValid: false,
		},
		{
			id:      "12",
			isValid: true,
		},
		{
			id:      "1010",
			isValid: false,
		},
		{
			id:      "10101",
			isValid: true,
		},
		{
			id:      "1188511885",
			isValid: false,
		},
		{
			id:      "222222",
			isValid: false,
		}, {
			id:      "2222222",
			isValid: true,
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("IsIdValid for %s should be %v", tc.id, tc.isValid)
		t.Run(name, func(t *testing.T) {
			isValid := day2.IsSplitStringValid(tc.id)
			assert.Equal(t, tc.isValid, isValid)
		})
	}
}

func TestPart1(t *testing.T) {
	line := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	scanner := bufio.NewScanner(strings.NewReader(line))

	result, err := day2.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "1227775554", result)
}

func TestIsWindowValid(t *testing.T) {
	cases := []struct {
		id      string
		isValid bool
	}{
		{
			id:      "11",
			isValid: false,
		},
		{
			id:      "12",
			isValid: true,
		},
		{
			id:      "111",
			isValid: false,
		},
		{
			id:      "112",
			isValid: true,
		},
		{
			id:      "1188511885",
			isValid: false,
		},
		{
			id:      "2121212121",
			isValid: false,
		},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("Foo for %s should be %v", tc.id, tc.isValid)
		t.Run(name, func(t *testing.T) {
			isValid := day2.IsWindowValid(tc.id)
			assert.Equal(t, tc.isValid, isValid)
		})
	}
}

func TestPart2(t *testing.T) {
	line := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	scanner := bufio.NewScanner(strings.NewReader(line))

	result, err := day2.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "4174379265", result)
}
