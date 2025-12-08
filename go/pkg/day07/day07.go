package day07

import (
	"bufio"
	"errors"
	"strconv"
	"unicode/utf8"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	beams, err := parseBeams(scanner)
	if err != nil {
		return "", err
	}

	splits := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, val := range beams {
			if val != 0 && line[i] == '^' {
				splits++
				beams[i-1] += val
				beams[i+1] += val
				beams[i] = 0
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", errors.New("failed to read input")
	}

	return strconv.Itoa(splits), nil
}

func parseBeams(scanner *bufio.Scanner) ([]int, error) {
	if !scanner.Scan() {
		return nil, errors.New("unexpected end of scanner")
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("failed to read input")
	}

	firstLine := scanner.Text()

	beams := make([]int, utf8.RuneCountInString(firstLine))
	for i, r := range firstLine {
		val := 0
		if r == 'S' {
			val = 1
		}
		beams[i] = val
	}

	return beams, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	beams, err := parseBeams(scanner)
	if err != nil {
		return "", err
	}

	timelines := 1
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, val := range beams {
			if val != 0 && line[i] == '^' {
				timelines += val
				beams[i-1] += val
				beams[i+1] += val
				beams[i] = 0
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", errors.New("failed to read input")
	}

	return strconv.Itoa(timelines), nil
}
