package day10

import (
	"bufio"
	"fmt"
	"iter"
	"math"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		target := parseTarget(line)
		buttons, err := parseButtons(line)
		if err != nil {
			return "", fmt.Errorf("failed to parse input: %w", err)
		}

		_ = make([]bool, len(target))

		minClicks := math.MaxInt
		for buttons := range permutations(buttons) {
			clicks, ok := apply(target, buttons, minClicks)
			if ok && clicks < minClicks {
				minClicks = clicks
			}
		}
		sum += minClicks
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strconv.Itoa(sum), nil
}

func parseTarget(input string) []bool {
	start := strings.Index(input, "[")
	end := strings.LastIndex(input, "]")
	target := make([]bool, end-start-1)
	runeInput := []rune(input)
	for i, r := range runeInput[start+1 : end] {
		toggled := false
		if r == '#' {
			toggled = true
		}
		target[i] = toggled
	}
	return target
}

func parseButtons(input string) ([][]uint, error) {
	start := strings.Index(input, "(")
	end := strings.LastIndex(input, ")")
	runeInput := []rune(input)
	substr := runeInput[start : end+1]
	split := strings.Fields(string(substr))
	buttons := make([][]uint, len(split))
	for i, s := range split {
		// (0,1,2,4,5)
		substr := s[1 : len(s)-1]
		strNumbers := strings.Split(substr, ",")
		button := make([]uint, len(strNumbers))
		for j, strNumber := range strNumbers {
			n, err := strconv.ParseUint(strNumber, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("failed to parse button: %w", err)
			}
			button[j] = uint(n)
		}
		buttons[i] = button
	}
	return buttons, nil
}

func permutations(buttons [][]uint) iter.Seq[[][]uint] {
	return func(yield func([][]uint) bool) {
		current := make([][]uint, len(buttons))
		var dfs func(int, int) bool
		dfs = func(i, position int) bool {
			if i == len(buttons) {
				if position > 0 {
					if !yield(current[:position]) {
						return false
					}
				}
				return true
			}

			if !dfs(i+1, position) {
				return false
			}

			current[position] = buttons[i]
			return dfs(i+1, position+1)
		}

		dfs(0, 0)
	}
}

func apply(target []bool, buttons [][]uint, minClicks int) (int, bool) {
	clicks := 0
	state := make([]bool, len(target))
	for _, b := range buttons {
		clicks++
		if clicks >= minClicks {
			return clicks, false
		}

		push(state, b)
		if equal(state, target) {
			return clicks, true
		}
	}
	return clicks, false
}

func push(state []bool, button []uint) {
	for _, n := range button {
		state[n] = !state[n]
	}
}

func equal(state []bool, target []bool) bool {
	for i := range len(state) {
		if state[i] != target[i] {
			return false
		}
	}

	return true
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return "lol", nil
}
