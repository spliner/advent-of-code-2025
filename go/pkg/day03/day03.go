package day03

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	return run(scanner, 2)
}

func run(scanner *bufio.Scanner, n int) (string, error) {
	var sum int64
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		maxJoltage := MaxJoltage(line, n)
		sum += maxJoltage
	}

	return strconv.FormatInt(sum, 10), nil
}

func MaxJoltage(input string, n int) int64 {
	index := 0
	runes := []rune(input)
	var result int64

	for i := range n {
		for j := index; j <= len(runes)-(n-i); j++ {
			if runes[index] < runes[j] {
				index = j
			}
		}

		digit := int64(runes[index] - '0')
		result += digit * pow(10, n-i-1)

		index++
	}

	return result
}

func pow(x, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return run(scanner, 12)
}
