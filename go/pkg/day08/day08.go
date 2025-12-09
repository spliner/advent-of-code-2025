package day08

import (
	"aoc2025/pkg/set"
	"bufio"
	"errors"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	X int
	Y int
	Z int
}

func (b Box) Distance(other Box) float64 {
	x := math.Pow(float64(b.X-other.X), 2)
	y := math.Pow(float64(b.Y-other.Y), 2)
	z := math.Pow(float64(b.Z-other.Z), 2)
	return math.Sqrt(x + y + z)
}

type Distance struct {
	Box      Box
	Other    Box
	Distance float64
}

func Part1(scanner *bufio.Scanner) (string, error) {
	return Part1WithConnections(scanner, 1000)
}

func Part1WithConnections(scanner *bufio.Scanner, connections int) (string, error) {
	boxes, err := parseBoxes(scanner)
	if err != nil {
		return "", err
	}

	distances := sortDistances(boxes)

	circuits := make(map[Box]*set.Set[Box])
	for _, d := range distances[0:connections] {
		union := set.Union(circuits[d.Box], circuits[d.Other])
		union.Add(d.Box)
		union.Add(d.Other)
		for b := range union.Items() {
			circuits[b] = union
		}
	}

	foo := slices.Collect(maps.Values(circuits))
	slices.SortFunc(foo, func(a, b *set.Set[Box]) int {
		return b.Len() - a.Len()
	})

	product := 1
	var lastCircuit *set.Set[Box]
	take := 0
	for _, f := range foo {
		if lastCircuit == f {
			continue
		}

		product *= f.Len()
		lastCircuit = f

		take++
		if take == 3 {
			break
		}
	}

	return strconv.Itoa(product), nil
}

func parseBoxes(scanner *bufio.Scanner) ([]Box, error) {
	boxes := make([]Box, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var box Box
		split := strings.Split(line, ",")
		if len(split) != 3 {
			return nil, fmt.Errorf("failed to parse point, got len: %d", len(split))
		}

		x, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse x coord: %w", err)
		}
		box.X = x

		y, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse y coord: %w", err)
		}
		box.Y = y

		z, err := strconv.Atoi(split[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse z coord: %w", err)
		}
		box.Z = z

		boxes = append(boxes, box)
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("failed to read input")
	}

	return boxes, nil
}

func sortDistances(boxes []Box) []Distance {
	distances := make([]Distance, 0)
	for i, box := range boxes {
		for _, other := range boxes[i+1:] {
			dist := box.Distance(other)
			distance := Distance{
				Box:      box,
				Other:    other,
				Distance: dist,
			}
			distances = append(distances, distance)
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return int(a.Distance - b.Distance)
	})

	return distances
}

func Part2(scanner *bufio.Scanner) (string, error) {
	boxes, err := parseBoxes(scanner)
	if err != nil {
		return "", err
	}

	distances := sortDistances(boxes)

	circuits := make(map[Box]*set.Set[Box])
	for _, d := range distances {
		union := set.Union(circuits[d.Box], circuits[d.Other])
		union.Add(d.Box)
		union.Add(d.Other)

		if union.Len() == len(boxes) {
			val := d.Box.X * d.Other.X
			return strconv.Itoa(val), nil
		}

		for b := range union.Items() {
			circuits[b] = union
		}
	}

	return "", errors.New("could not build single circuit")
}
