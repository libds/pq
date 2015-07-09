package pq

import (
	"container/heap"
	"fmt"
)

type priorityQueue []*Item

type Item struct {
	Index    int
	Priority int
	Value    interface{}
}

func (item *Item) String() string {
	return fmt.Sprintf("{Value: %v, Priority: %d, Index: %d}",
		item.Value, item.Priority, item.Index)
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func newMaxPQ() priorityQueue {
	return priorityQueue{}
}

type MaxPQ struct {
	pq *priorityQueue
}

func NewMax() *MaxPQ {
	pq := newMaxPQ()
	heap.Init(&pq)

	return &MaxPQ{&pq}
}

func (p *MaxPQ) Push(value interface{}, priority int) {
	item := &Item{
		Value:    value,
		Priority: priority,
	}
	heap.Push(p.pq, item)
}

func (p *MaxPQ) Size() int {
	return p.pq.Len()
}

func (p *MaxPQ) Peek() *Item {
	if p.Size() == 0 {
		return nil
	}

	return (*p.pq)[0]
}

// update modifies the priority and value of an Item in the queue.
func (pq *MaxPQ) Update(item *Item, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq.pq, item.Index)
}
func (p *MaxPQ) Pop() *Item {
	item := heap.Pop(p.pq).(*Item)
	return item
}
