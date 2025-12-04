package set

import (
	"iter"
	"maps"
)

var mark = struct{}{}

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable]() *Set[T] {
	items := make(map[T]struct{})
	return &Set[T]{
		items: items,
	}
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Add(item T) bool {
	if s.Contains(item) {
		return false
	}

	s.items[item] = mark

	return true
}

func (s *Set[T]) Remove(item T) bool {
	if !s.Contains(item) {
		return false
	}

	delete(s.items, item)
	return true
}

func (s *Set[T]) Items() iter.Seq[T] {
	return maps.Keys(s.items)
}
