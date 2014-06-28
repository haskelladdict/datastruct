// unit tests for priority queue
package pqueue

import (
	"testing"
)

func Test_queue(t *testing.T) {

	var items []*Item
	one := Item{"one", 1, 0}
	items = append(items, &one)

	pq := New(items)
	pq.Update(&one, "bar", 2)

	if i := pq.Pop(); i.Val != "bar" {
		t.Error("incorrect queue elements")
	}
}
