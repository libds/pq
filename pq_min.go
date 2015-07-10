package pq

import (
	"container/heap"
)

type minPQ []*Item

func (pq minPQ) Len() int {
	return len(pq)
}

func (pq minPQ) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func newMinPQ() minPQ {
	return minPQ{}
}

type MinPQ struct {
	pq *minPQ
}

// Initialize a Heap Minimum Priority Queue
func NewMin() *MinPQ {
	pq := newMinPQ()
	heap.Init(&pq)

	return &MinPQ{&pq}
}

// Push pushes the element x onto the heap.
//
// The complexity is O(log(n)) where n = p.Size().
func (p *MinPQ) Push(value interface{}, priority int) {
	item := &Item{
		Value:    value,
		Priority: priority,
	}
	heap.Push(p.pq, item)
}

// Returns the amount of items in the queue.
//
// Complexity = O(1)
func (p *MinPQ) Size() int {
	return p.pq.Len()
}

// Returns the minimum item in the queue without removing it
//
// The complexity is O(1)
func (p *MinPQ) Peek() *Item {
	if p.Size() == 0 {
		return nil
	}
	return (*p.pq)[0]
}

// update modifies the priority and value of an Item in the queue.
//
// The Complexity is O(1) if you only update the value.
// And O(log(n)) if you change the priority, where n = p.Size()
func (p *MinPQ) Update(item *Item, value interface{}, priority int) {
	item.Value = value
	if priority != item.Priority {
		item.Priority = priority
		heap.Fix(p.pq, item.Index)
	}
}

// Pop removes the minimum element from the heap and returns it.
//
// The complexity is O(log(n)) where n = p.Size().
func (p *MinPQ) Pop() *Item {
	item := heap.Pop(p.pq).(*Item)
	return item
}
