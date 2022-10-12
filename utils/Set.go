package utils

type Set[T comparable] struct {
	Values map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		Values: make(map[T]struct{}),
	}
}

func (s Set[T]) Add(value T) {
	s.Values[value] = struct{}{}
}

func (s Set[T]) Remove(value T) {
	delete(s.Values, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s.Values[value]
	return ok
}
