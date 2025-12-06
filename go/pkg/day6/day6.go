package day6

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	numbers := make([][]int, 0)
	operators := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		split := strings.Fields(line)
		if strings.HasPrefix(line, "*") || strings.HasPrefix(line, "+") {
			for _, s := range split {
				operator := strings.TrimSpace(s)
				operators = append(operators, operator)
			}

			break
		}

		row := make([]int, 0, len(split))
		for _, s := range split {
			trimmed := strings.TrimSpace(s)
			n, err := strconv.Atoi(trimmed)
			if err != nil {
				return "", fmt.Errorf("failed to parse \"%s\" as number: %w", s, err)
			}

			row = append(row, n)
		}

		numbers = append(numbers, row)
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	columns := len(numbers[0])
	rows := len(numbers)

	var sum int
	for c := range columns {
		operator := operators[c]
		var acc int
		var reducer func(acc, curr int) int
		switch operator {
		case "*":
			acc = 1
			reducer = func(acc, curr int) int {
				return acc * curr
			}
		case "+":
			reducer = func(acc, curr int) int {
				return acc + curr
			}
		}
		for r := range rows {
			n := numbers[r][c]
			acc = reducer(acc, n)
		}
		sum += acc
	}

	return strconv.Itoa(sum), nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	lines := make([][]rune, 0)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	lastLine := lines[len(lines)-1]
	start := 0
	delimiter := findDelimiter(lastLine, start)
	var sum int
	for start <= len(lastLine) {
		numbers := make([]int, 0)
		for col := delimiter; col >= start; col-- {
			var numberBuilder strings.Builder
			for row := range len(lines) - 1 {
				candidate := []rune(lines[row])[col]
				if candidate != ' ' {
					numberBuilder.WriteRune(candidate)
				}
			}

			n, err := strconv.Atoi(numberBuilder.String())
			if err != nil {
				return "", err
			}
			numbers = append(numbers, n)
		}

		operator := lastLine[start]
		var acc int
		var reducer func(acc, curr int) int
		switch operator {
		case '*':
			acc = 1
			reducer = func(acc, curr int) int {
				return acc * curr
			}
		case '+':
			reducer = func(acc, curr int) int {
				return acc + curr
			}
		}
		for _, n := range numbers {
			acc = reducer(acc, n)
		}
		sum += acc

		start = delimiter + 2
		delimiter = findDelimiter(lastLine, start)
	}

	return strconv.Itoa(sum), nil
}

func findDelimiter(line []rune, start int) int {
	for i := start + 1; i < len(line); i++ {
		if line[i] == '*' || line[i] == '+' {
			// Accounting for operator index + whitespace
			return i - 2
		}
	}

	return len(line) - 1
}
