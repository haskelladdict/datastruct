package pqueue

import "container/heap"

// Item represents a single element of a Priority queue
type Item struct {
	Val      interface{} // Value stored in heap
	Priority int         // the Priority of the item in the heap
	Index    int         // the Index of the item in the heap
}

// Pqueue implements a priority queue leveraging go's heap implementation
type Pqueue priorityQueue

// New creates a new priority queue initialized by items
func New(items []*Item) *Pqueue {
	pq := priorityQueue(items)
	heap.Init(&pq)
	return (*Pqueue)(&pq)
}

// Pop removes and returns the highest priority item of the queue
func (pq *Pqueue) Pop() *Item {
	return heap.Pop((*priorityQueue)(pq)).(*Item)
}

// Push adds a new item to the priority queue.
func (pq *Pqueue) Push(item *Item) {
	heap.Push((*priorityQueue)(pq), item)
}

// Update updates the Priority and Value of an item in the Priority queue
func (pq *Pqueue) Update(item *Item, Value interface{}, Priority int) {
	item.Val = Value
	item.Priority = Priority
	heap.Fix((*priorityQueue)(pq), item.Index)
}

// PriorityQueue implements a Priority queue
type priorityQueue []*Item

// Len provides the length of the PriorityQueue
func (p priorityQueue) Len() int {
	return len(p)
}

// Less implements less than for a PriorityQueue
func (p priorityQueue) Less(i, j int) bool {
	return p[i].Priority < p[j].Priority
}

// Swap implements swapping for a PriorityQueue
func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].Index = i
	p[j].Index = j
}

// Push adds an element to the PriorityQueue
func (p *priorityQueue) Push(i interface{}) {
	n := len(*p)
	item := i.(*Item)
	item.Index = n
	*p = append(*p, item)
}

// Pop removes the top element from the PriorityQueue
func (p *priorityQueue) Pop() interface{} {
	tmp := *p
	n := len(*p)
	item := tmp[n-1]
	item.Index = -1
	*p = tmp[0 : n-1]
	return item
}
