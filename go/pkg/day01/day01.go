package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Dial struct {
	CurrentPosition        int
	StoppedAtZeroCount     int
	PassedThroughZeroCount int
}

func (d *Dial) TurnLeft(times int) {
	fullTurns := times / 100
	d.PassedThroughZeroCount += fullTurns

	times = times % 100
	newPosition := d.CurrentPosition - times
	passedThroughZero := false
	if newPosition < 0 {
		newPosition = 100 + newPosition
		passedThroughZero = d.CurrentPosition != 0
	}

	if newPosition == 0 {
		d.StoppedAtZeroCount++
	} else if passedThroughZero {
		d.PassedThroughZeroCount++
	}

	d.CurrentPosition = newPosition
}

func (d *Dial) TurnRight(times int) {
	fullTurns := times / 100
	d.PassedThroughZeroCount += fullTurns

	times = times % 100
	newPosition := d.CurrentPosition + times
	passedThroughZero := false
	if newPosition > 99 {
		newPosition -= 100
		passedThroughZero = true
	}

	if newPosition == 0 {
		d.StoppedAtZeroCount++
	} else if passedThroughZero {
		d.PassedThroughZeroCount++
	}

	d.CurrentPosition = newPosition
}

func Part1(scanner *bufio.Scanner) (string, error) {
	dial := &Dial{CurrentPosition: 50}
	if err := run(dial, scanner); err != nil {
		return "", fmt.Errorf("failed to run day 1 part 1: %w", err)
	}
	return strconv.Itoa(dial.StoppedAtZeroCount), nil
}

func run(dial *Dial, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := line[0]
		times, err := strconv.Atoi(line[1:])
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", line[1:], err)
		}

		switch direction {
		case 'L':
			dial.TurnLeft(times)
		case 'R':
			dial.TurnRight(times)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	dial := &Dial{CurrentPosition: 50}
	if err := run(dial, scanner); err != nil {
		return "", fmt.Errorf("failed to run day 1 part 2: %w", err)
	}
	return strconv.Itoa(dial.PassedThroughZeroCount + dial.StoppedAtZeroCount), nil
}
