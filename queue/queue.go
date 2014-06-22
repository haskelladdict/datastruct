package queue

import (
	"fmt"
)

// Queue is a generic queue structure
type Queue struct {
	queue []interface{}
}

// NewQueue creates and returns a new queue structure
func NewQueue() *Queue {
	return &Queue{}
}

// Pop removes the first element from the queue
func (s *Queue) Pop() (interface{}, error) {
	var item interface{}
	if len(s.queue) == 0 {
		return item, fmt.Errorf("queue is empty")
	}

	item = s.queue[0]
	s.queue = s.queue[1:len(s.queue)]
	return item, nil
}

// Push an element to a queue
func (s *Queue) Push(i interface{}) {
	s.queue = append(s.queue, i)
}

// Len returns the size of the queue
func (s *Queue) Len() int {
	return len(s.queue)
}

// Clear empties the queue
func (s *Queue) Clear() {
	s.queue = make([]interface{}, 0)
}
