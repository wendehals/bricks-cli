package scripting

import (
	"fmt"
	"log"
)

// stack is a generic stack data structure
type stack[T any] struct {
	elements []T
}

// newStack creates and returns a new stack instance
func newStack[T any]() *stack[T] {
	return &stack[T]{
		elements: make([]T, 0),
	}
}

// push adds an element to the top of the stack
func (s *stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

// pop removes and returns the top element of the stack
func (s *stack[T]) pop() T {
	top := len(s.elements) - 1
	if top < 0 {
		log.Fatalf("pop: no more elements on stack")
	}

	element := s.elements[top]
	// Clear the reference for garbage collection
	var zero T
	s.elements[top] = zero
	s.elements = s.elements[:top]

	return element
}

// peek returns the top element without removing it
func (s *stack[T]) peek() T {
	top := len(s.elements) - 1
	if top < 0 {
		log.Fatalf("peek: no elements on stack")
	}
	return s.elements[top]
}

// isEmpty returns true if the stack has no elements
func (s *stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

// size returns the number of elements in the stack
func (s *stack[T]) size() int {
	return len(s.elements)
}

// clear removes all elements from the stack
func (s *stack[T]) clear() {
	// Clear all references for garbage collection
	var zero T
	for i := range s.elements {
		s.elements[i] = zero
	}
	s.elements = s.elements[:0]
}

// String implements fmt.Stringer for debugging
func (s *stack[T]) String() string {
	return fmt.Sprintf("Stack[%T]%v", *new(T), s.elements)
}
