package day11

import (
	"aoc2025/pkg/stack"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Device struct {
	Id      string
	Outputs []string
}

func Part1(scanner *bufio.Scanner) (string, error) {
	devices, err := parseDevices(scanner)
	if err != nil {
		return "", err
	}

	deviceStack := stack.New[Device]()
	deviceStack.Push(devices["you"])
	count := 0
	for !deviceStack.IsEmpty() {
		device, _ := deviceStack.Pop()
		for _, output := range device.Outputs {
			if output == "out" {
				count++
			} else {
				outputDevice, ok := devices[output]
				if ok {
					deviceStack.Push(outputDevice)
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func parseDevices(scanner *bufio.Scanner) (map[string]Device, error) {
	devices := make(map[string]Device)
	for scanner.Scan() {
		line := scanner.Text()
		device, err := parseDevice(line)
		if err != nil {
			return nil, err
		}

		devices[device.Id] = device
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return devices, nil
}

func parseDevice(line string) (Device, error) {
	var device Device
	split := strings.Split(line, ":")
	if len(split) != 2 {
		return device, fmt.Errorf("expected len to be 2, got %d", len(split))
	}

	device.Id = split[0]
	device.Outputs = strings.Fields(strings.TrimSpace(split[1]))
	return device, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	devices, err := parseDevices(scanner)
	if err != nil {
		return "", err
	}

	type key struct {
		deviceId string
		dac      bool
		fft      bool
	}
	memo := make(map[key]int)
	var search func(deviceId string, dac, fft bool) int
	search = func(deviceId string, dac, fft bool) int {
		k := key{deviceId, dac, fft}
		n, ok := memo[k]
		if ok {
			return n
		}

		if deviceId == "dac" {
			dac = true
		}
		if deviceId == "fft" {
			fft = true
		}
		if deviceId == "out" {
			if dac && fft {
				memo[k] = 1
				return 1
			}
			memo[k] = 0
			return 0
		}

		d, ok := devices[deviceId]
		if !ok {
			memo[k] = 0
			return 0
		}

		sum := 0
		for _, o := range d.Outputs {
			if ok {
				sum += search(o, dac, fft)
			}
		}
		memo[k] = sum
		return sum
	}

	count := search("svr", false, false)
	return strconv.Itoa(count), nil
}
