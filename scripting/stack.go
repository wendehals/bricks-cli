package scripting

import (
	"log"
)

type scriptStack struct {
	elements []interface{}
}

func newScriptStack() *scriptStack {
	return &scriptStack{}
}

func (s *scriptStack) push(data interface{}) {
	s.elements = append(s.elements, data)
}

func (s *scriptStack) pop() interface{} {
	top := len(s.elements) - 1
	if top < 0 {
		log.Fatalf("pop: no more elements on stack")
	}

	element := s.elements[top]
	s.elements[top] = nil
	s.elements = s.elements[:top]

	return element
}
