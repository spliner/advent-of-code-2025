package main

import (
	"aoc2025/pkg/day1"
	"aoc2025/pkg/day2"
	"aoc2025/pkg/day3"
	"aoc2025/pkg/day4"
	"aoc2025/pkg/day5"
	"aoc2025/pkg/day6"
	"bufio"
	"fmt"
	"os"
)

type solver func(*bufio.Scanner) (string, error)

var solvers = map[string]solver{
	"1-1": day1.Part1,
	"1-2": day1.Part2,
	"2-1": day2.Part1,
	"2-2": day2.Part2,
	"3-1": day3.Part1,
	"3-2": day3.Part2,
	"4-1": day4.Part1,
	"4-2": day4.Part2,
	"5-1": day5.Part1,
	"5-2": day5.Part2,
	"6-1": day6.Part1,
	"6-2": day6.Part2,
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Invalid arguments")
		os.Exit(1)
	}

	day := os.Args[1]
	part := os.Args[2]

	f, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file:", err)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to close file:", err)
		}
	}()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	solver, ok := solvers[day+"-"+part]
	if !ok {
		fmt.Fprintln(os.Stderr, "Could not find solver for day "+day+" part "+part)
		os.Exit(1)
	}

	result, err := solver(scanner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to run day "+day+" part "+part+":", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
