package scripting

import (
	"log"

	"github.com/wendehals/bricks/model"
)

type stack struct {
	elements []*model.Collection
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(collection *model.Collection) {
	s.elements = append(s.elements, collection)
}

func (s *stack) pop() *model.Collection {
	top := len(s.elements) - 1
	if top < 0 {
		log.Fatalf("pop: no more elements on stack")
	}

	element := s.elements[top]
	s.elements[top] = nil
	s.elements = s.elements[:top]

	return element
}
