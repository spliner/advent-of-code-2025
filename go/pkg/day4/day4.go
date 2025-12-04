package day4

import (
	"aoc2025/pkg/set"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	rolls, err := parse(scanner)
	if err != nil {
		return "", err
	}

	reachable := ReachableRolls(rolls)
	return strconv.Itoa(len(reachable)), nil
}

func parse(scanner *bufio.Scanner) (*set.Set[Point], error) {
	rolls := set.New[Point]()
	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		for x, r := range line {
			if r == '@' {
				rolls.Add(Point{X: x, Y: y})
			}
		}

		y++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return rolls, nil
}

func ReachableRolls(rolls *set.Set[Point]) []Point {
	toRemove := make([]Point, 0)
	for roll := range rolls.Items() {
		var count int
		for p := range roll.AdjacentPoints() {
			if rolls.Contains(p) {
				count++
			}
		}

		if count < 4 {
			toRemove = append(toRemove, roll)
		}
	}

	return toRemove
}

func Part2(scanner *bufio.Scanner) (string, error) {
	rolls, err := parse(scanner)
	if err != nil {
		return "", err
	}

	var sum int
	for {
		reachable := ReachableRolls(rolls)
		if len(reachable) == 0 {
			break
		}

		for _, r := range reachable {
			sum++
			rolls.Remove(r)
		}
	}

	return strconv.Itoa(sum), nil
}
