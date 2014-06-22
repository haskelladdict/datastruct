// some sanity tests to make sure our queue works ok
package queue

import (
	"testing"
)

func Test_queue(t *testing.T) {

	as := []int{1, 2, 3, 4}
	aQueue := NewQueue()
	bQueue := NewQueue()
	cQueue := NewQueue()
	for _, a := range as {
		aQueue.Push(a)
		bQueue.Push(a)
		cQueue.Push(a)
	}

	if aQueue.Len() != 4 {
		t.Error("queue has incorrect length")
	}

	for _, a := range as {
		val, err := aQueue.Pop()
		if err != nil {
			t.Error("failed to pop element from queue")
		}
		if a != val.(int) {
			t.Error("elements in queue appear to be out of order")
		}
	}

	bQueue.Clear()
	_, err := bQueue.Pop()
	if bQueue.Len() != 0 || err == nil {
		t.Error("Failed to properly clear queue")
	}

	cQueue.Pop()
	cQueue.Pop()
	cQueue.Push(5)
	cQueue.Push(6)
	comp := []int{3, 4, 5, 6}
	for _, c := range comp {
		val, err := cQueue.Pop()
		if err != nil {
			t.Error("failed to pop element from queue")
		}
		if c != val.(int) {
			t.Error("elements in queue appear to be out of order")
		}
	}
}
