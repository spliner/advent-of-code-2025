package day02

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type IdValidator func(string) bool

func Part1(scanner *bufio.Scanner) (string, error) {
	return run(scanner, IsSplitStringValid)
}

func run(scanner *bufio.Scanner, validator IdValidator) (string, error) {
	if !scanner.Scan() {
		return "", errors.New("failed to scan input")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to scan input: %w", err)
	}

	line := scanner.Text()
	ranges := strings.Split(line, ",")
	var sum int
	for _, r := range ranges {
		start, end, ok := strings.Cut(r, "-")
		if !ok {
			return "", fmt.Errorf("unexpected range: %s", r)
		}

		invalidIds, err := FindInvalidIds(start, end, validator)
		if err != nil {
			return "", fmt.Errorf("failed to find invalid ids: %w", err)
		}

		for _, i := range invalidIds {
			val, _ := strconv.Atoi(i)
			sum += val
		}
	}

	return strconv.Itoa(sum), nil
}

func FindInvalidIds(startId, endId string, validator IdValidator) ([]string, error) {
	invalidIds := make([]string, 0)
	start, err := strconv.Atoi(startId)
	if err != nil {
		return nil, fmt.Errorf("invalid start id %s: %w", startId, err)
	}

	end, err := strconv.Atoi(endId)
	if err != nil {
		return nil, fmt.Errorf("invalid end id %s: %w", endId, err)
	}

	for i := start; i <= end; i++ {
		id := strconv.Itoa(i)
		if !validator(id) {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds, nil
}

func IsSplitStringValid(id string) bool {
	if len(id)%2 != 0 {
		return true
	}

	return !allMatch(id, len(id)/2)
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return run(scanner, IsWindowValid)
}

func IsWindowValid(id string) bool {
	for windowSize := 1; windowSize <= len(id)/2; windowSize++ {
		if len(id)%windowSize != 0 {
			continue
		}

		if allMatch(id, windowSize) {
			return false
		}
	}

	return true
}

func allMatch(id string, windowSize int) bool {
	toMatch := id[0:windowSize]
	for i := range len(id) / windowSize {
		toCompare := id[windowSize*i : windowSize*(i+1)]
		if toMatch != toCompare {
			return false
		}
	}

	return true
}
