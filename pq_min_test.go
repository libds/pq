package pq

import (
	"fmt"
	"testing"
)

func TestMinPQ(t *testing.T) {
	fmt.Println("Test Min PQ")
	pq := NewMin()
	if pq.Peek() != nil {
		t.Fatal("Peek should return nil when empty")
	}
	items := []*Item{
		{Value: 22, Priority: 200},
		{Value: 52, Priority: 100},
	}

	for _, item := range items {
		pq.Push(item.Value, item.Priority)
	}

	job := pq.Peek()
	if job.Value != items[len(items)-1].Value {
		t.Fatalf("Job value must equal %s", items[len(items)-1].Value)
	}

	pq.Update(job, job.Value, 250)

	if job.Priority != 250 {
		t.Fatalf("%s priority should be 250", job)
	}

	if j := pq.Peek(); j.Value == job.Value {
		t.Fatalf("%s shouldn't be on the top", job)
	}
}
