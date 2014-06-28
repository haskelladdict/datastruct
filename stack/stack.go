package stack

import (
	"fmt"
)

// Stack is a generic stack structure
type Stack struct {
	stack []interface{}
}

// NewStack creates and returns a new stack structure
func New() *Stack {
	return &Stack{}
}

// Pop removes the last element from a stack
func (s *Stack) Pop() (interface{}, error) {
	var item interface{}
	if len(s.stack) == 0 {
		return item, fmt.Errorf("stack is empty")
	}

	item = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return item, nil
}

// Push element to the stack
func (s *Stack) Push(i interface{}) {
	s.stack = append(s.stack, i)
}

// Len returns the size of the stack
func (s *Stack) Len() int {
	return len(s.stack)
}

// Clear empties the stack
func (s *Stack) Clear() {
	s.stack = make([]interface{}, 0)
}
