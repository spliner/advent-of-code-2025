package day11_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day11"
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day11.Part1(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "5", result)
}

func TestPart2(t *testing.T) {
	input := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := day11.Part2(scanner)

	assert.Nil(t, err)
	assert.Equal(t, "2", result)
}
