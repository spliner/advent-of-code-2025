package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (item T, ok bool) {
	if len(s.items) == 0 {
		return item, false
	}

	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
