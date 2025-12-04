package set_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/set"
	"testing"
)

func TestAdd(t *testing.T) {
	s := set.New[int]()

	success := s.Add(10)
	assert.True(t, success)

	success = s.Add(10)
	assert.False(t, success)
}

func TestContains(t *testing.T) {
	s := set.New[int]()

	contains := s.Contains(10)
	assert.False(t, contains)

	s.Add(10)
	contains = s.Contains(10)
	assert.True(t, contains)
}

func TestRemove(t *testing.T) {
	s := set.New[int]()

	success := s.Remove(10)
	assert.False(t, success)

	s.Add(10)
	success = s.Remove(10)
	assert.True(t, success)

	success = s.Remove(10)
	assert.False(t, success)
}
