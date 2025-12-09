package set

import (
	"fmt"
	"iter"
	"maps"
	"strconv"
	"strings"
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

func Union[T comparable](s, other *Set[T]) *Set[T] {
	union := New[T]()
	if s != nil {
		for k := range s.items {
			union.Add(k)
		}
	}
	if other != nil {
		for k := range other.items {
			union.Add(k)
		}
	}
	return union
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

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) String() string {
	var sb strings.Builder
	strLen := strconv.Itoa(len(s.items))
	sb.WriteString("(" + strLen + ")")
	sb.WriteString("[")
	for k := range s.items {
		sb.WriteString(fmt.Sprint(k))
	}
	sb.WriteString("]")
	return sb.String()
}
