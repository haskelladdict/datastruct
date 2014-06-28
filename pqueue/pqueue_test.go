// unit tests for priority queue
package pqueue

import (
	"math/rand"
	"testing"
	"time"
)

const maxint = 2147483647

func Test_queue(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var items []*Item
	pq := New(items)

	for i := 0; i < 1000000; i++ {
		item := &Item{i, r.Int(), i}
		pq.Push(item)
		items = append(items, item)
	}

	val := -1
	for i := 0; i < 1000000; i++ {
		newVal := pq.Pop().Priority
		if newVal < val {
			t.Error("incorrect ordering in priority queue pq")
		}
	}

	pq1 := New(items)
	for i := 0; i < 1000000; i++ {
		pq1.Update(items[i], r.Int(), i)
	}

	val = -1
	for i := 0; i < 1000000; i++ {
		newVal := pq1.Pop().Priority
		if newVal < val {
			t.Error("incorrect ordering in priority queue pq1")
		}
	}

}
