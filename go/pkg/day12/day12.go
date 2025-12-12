package day12

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Shape struct {
	Items [][]rune
}

func (s *Shape) RequiredSpace() int {
	var sum int
	for _, l := range s.Items {
		var count int
		for _, r := range l {
			if r == '#' {
				count++
			}
		}
		sum += count
	}
	return sum
}

type Region struct {
	Width    int
	Height   int
	Presents []int
}

func (r *Region) Area() int {
	return r.Width * r.Height
}

func Part1(scanner *bufio.Scanner) (string, error) {
	shapes, err := parseShapes(scanner)
	if err != nil {
		return "", err
	}

	regions, err := parseRegions(scanner)
	if err != nil {
		return "", err
	}
	fmt.Println(regions)

	count := 0
	for i, r := range regions {
		availableArea := int(r.Area())
		for i, count := range r.Presents {
			shape := shapes[i]
			availableArea -= shape.RequiredSpace() * count
		}
		if availableArea >= 0 {
			count++
		}

		fmt.Println("region", i, "avail", availableArea)
	}

	return strconv.Itoa(count), nil
}

func parseShapes(scanner *bufio.Scanner) ([]Shape, error) {
	n := 6
	shapes := make([]Shape, 0, n)
	for range n {
		rawShape := make([][]rune, 0, 3)
		for j := 0; j < 5 && scanner.Scan(); j++ {
			line := scanner.Text()
			if line == "" {
				continue
			}

			r := []rune(line)
			if unicode.IsDigit(r[0]) {
				continue
			}

			rawShape = append(rawShape, r)
		}

		shape := Shape{Items: rawShape}
		shapes = append(shapes, shape)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return shapes, nil
}

func parseRegions(scanner *bufio.Scanner) ([]Region, error) {
	regions := make([]Region, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		before, after, _ := strings.Cut(line, ": ")
		strWidth, strHeight, _ := strings.Cut(before, "x")
		width, _ := strconv.Atoi(strWidth)
		height, _ := strconv.Atoi(strHeight)

		strPresents := strings.Fields(after)
		presents := make([]int, 0, len(strPresents))
		for _, strCount := range strPresents {
			count, _ := strconv.Atoi(strCount)
			presents = append(presents, count)
		}

		region := Region{
			Width:    width,
			Height:   height,
			Presents: presents,
		}
		regions = append(regions, region)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}
	return regions, nil
}
