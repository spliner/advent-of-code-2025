package day05

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start uint64
	End   uint64
}

func Part1(scanner *bufio.Scanner) (string, error) {
	ranges, err := parseRanges(scanner)
	if err != nil {
		return "", fmt.Errorf("failed to parse ranges: %w", err)
	}

	// Ids
	var count int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		id, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse id: %s", line)
		}

		for _, r := range ranges {
			if id >= r.Start && id <= r.End {
				count++
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read ids: %w", err)
	}

	return strconv.Itoa(count), nil
}

func parseRanges(scanner *bufio.Scanner) ([]Range, error) {
	ranges := make([]Range, 0, 100)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		startStr, endStr, ok := strings.Cut(line, "-")
		if !ok {
			return nil, fmt.Errorf("unexpected range: %s", line)
		}

		start, err := strconv.ParseUint(startStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse range start: %s", startStr)
		}

		end, err := strconv.ParseUint(endStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse range end: %s", endStr)
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read ranges: %w", err)
	}

	return ranges, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	ranges, err := parseRanges(scanner)
	if err != nil {
		return "", fmt.Errorf("failed to parse ranges: %w", err)
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Start < b.Start {
			return -1
		} else if a.Start > b.Start {
			return 1
		}
		return 0
	})

	nonOverlappingRanges := make([]Range, 0, len(ranges))
	var previousRange Range
	for _, r := range ranges {
		start := max(r.Start, previousRange.End)
		if start == previousRange.End {
			start += 1
		}
		end := max(r.End, previousRange.End)
		if end < start {
			continue
		}
		nonOverlappingRanges = append(nonOverlappingRanges, Range{Start: start, End: end})
		previousRange = r
	}

	var sum uint64
	for _, r := range nonOverlappingRanges {
		sum += r.End - r.Start + 1
	}
	return strconv.FormatUint(sum, 10), nil
}
