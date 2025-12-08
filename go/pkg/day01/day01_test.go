package day01_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/day01"
	"fmt"
	"testing"
)

func TestTurnLeft(t *testing.T) {
	testCases := []struct {
		currentPosition  int
		times            int
		expectedPosition int
	}{
		{
			currentPosition:  50,
			times:            1,
			expectedPosition: 49,
		},
		{
			currentPosition:  0,
			times:            1,
			expectedPosition: 99,
		},
		{
			currentPosition:  10,
			times:            15,
			expectedPosition: 95,
		},
		{
			currentPosition:  50,
			times:            100,
			expectedPosition: 50,
		},
		{
			currentPosition:  50,
			times:            101,
			expectedPosition: 49,
		},
		{
			currentPosition:  50,
			times:            200,
			expectedPosition: 50,
		},
		{
			currentPosition:  50,
			times:            201,
			expectedPosition: 49,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Turning %d times left from %d should stop at %d", tc.times, tc.currentPosition, tc.expectedPosition)
		t.Run(name, func(t *testing.T) {
			dial := day01.Dial{CurrentPosition: tc.currentPosition}
			dial.TurnLeft(tc.times)
			assert.Equal(t, tc.expectedPosition, dial.CurrentPosition)
		})
	}
}

func TestTurnRight(t *testing.T) {
	testCases := []struct {
		currentPosition int
		times           int
		expected        int
	}{
		{
			currentPosition: 50,
			times:           1,
			expected:        51,
		},
		{
			currentPosition: 99,
			times:           1,
			expected:        0,
		},
		{
			currentPosition: 90,
			times:           15,
			expected:        5,
		},
		{
			currentPosition: 50,
			times:           100,
			expected:        50,
		},
		{
			currentPosition: 50,
			times:           101,
			expected:        51,
		},
		{
			currentPosition: 50,
			times:           200,
			expected:        50,
		},
		{
			currentPosition: 50,
			times:           201,
			expected:        51,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Turning %d times right from %d should stop at %d", tc.times, tc.currentPosition, tc.expected)
		t.Run(name, func(t *testing.T) {
			dial := day01.Dial{CurrentPosition: tc.currentPosition}
			dial.TurnRight(tc.times)
			assert.Equal(t, tc.expected, dial.CurrentPosition)
		})
	}
}

func TestStoppedAtZeroCount(t *testing.T) {
	dial := day01.Dial{CurrentPosition: 50}

	// 0
	dial.TurnRight(10)
	assert.Equal(t, 0, dial.StoppedAtZeroCount)

	// +1
	dial.TurnRight(40)
	assert.Equal(t, 1, dial.StoppedAtZeroCount)

	// +1
	dial.TurnRight(200)
	assert.Equal(t, 2, dial.StoppedAtZeroCount)

	// +1
	dial.TurnLeft(200)
	assert.Equal(t, 3, dial.StoppedAtZeroCount)
}

func TestDay1(t *testing.T) {
	dial := day01.Dial{CurrentPosition: 50}

	dial.TurnLeft(68)
	assert.Equal(t, 82, dial.CurrentPosition)
	assert.Equal(t, 0, dial.StoppedAtZeroCount)
	assert.Equal(t, 1, dial.PassedThroughZeroCount)

	dial.TurnLeft(30)
	assert.Equal(t, 52, dial.CurrentPosition)
	assert.Equal(t, 0, dial.StoppedAtZeroCount)
	assert.Equal(t, 1, dial.PassedThroughZeroCount)

	dial.TurnRight(48)
	assert.Equal(t, 0, dial.CurrentPosition)
	assert.Equal(t, 1, dial.StoppedAtZeroCount)
	assert.Equal(t, 1, dial.PassedThroughZeroCount)

	dial.TurnLeft(5)
	assert.Equal(t, 95, dial.CurrentPosition)
	assert.Equal(t, 1, dial.StoppedAtZeroCount)
	assert.Equal(t, 1, dial.PassedThroughZeroCount)

	dial.TurnRight(60)
	assert.Equal(t, 55, dial.CurrentPosition)
	assert.Equal(t, 1, dial.StoppedAtZeroCount)
	assert.Equal(t, 2, dial.PassedThroughZeroCount)

	dial.TurnLeft(55)
	assert.Equal(t, 0, dial.CurrentPosition)
	assert.Equal(t, 2, dial.StoppedAtZeroCount)
	assert.Equal(t, 2, dial.PassedThroughZeroCount)

	dial.TurnLeft(1)
	assert.Equal(t, 99, dial.CurrentPosition)
	assert.Equal(t, 2, dial.StoppedAtZeroCount)
	assert.Equal(t, 2, dial.PassedThroughZeroCount)

	dial.TurnLeft(99)
	assert.Equal(t, 0, dial.CurrentPosition)
	assert.Equal(t, 3, dial.StoppedAtZeroCount)
	assert.Equal(t, 2, dial.PassedThroughZeroCount)

	dial.TurnRight(14)
	assert.Equal(t, 14, dial.CurrentPosition)
	assert.Equal(t, 3, dial.StoppedAtZeroCount)
	assert.Equal(t, 2, dial.PassedThroughZeroCount)

	dial.TurnLeft(82)
	assert.Equal(t, 32, dial.CurrentPosition)
	assert.Equal(t, 3, dial.StoppedAtZeroCount)
	assert.Equal(t, 3, dial.PassedThroughZeroCount)
}
