package stack_test

import (
	"aoc2025/pkg/assert"
	"aoc2025/pkg/stack"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack.New[int]()

	empty := s.IsEmpty()
	assert.True(t, empty)

	_, ok := s.Pop()
	assert.False(t, ok)

	s.Push(10)
	s.Push(20)

	empty = s.IsEmpty()
	assert.False(t, empty)

	item, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 20, item)

	empty = s.IsEmpty()
	assert.False(t, empty)

	item, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 10, item)

	empty = s.IsEmpty()
	assert.True(t, empty)

	_, ok = s.Pop()
	assert.False(t, ok)
}
