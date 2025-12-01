package main

import (
	"aoc2025/pkg/day1"
	"bufio"
	"fmt"
	"os"
)

type solver func(*bufio.Scanner) (string, error)

var solvers = map[string]solver{
	"1-1": day1.Part1,
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
