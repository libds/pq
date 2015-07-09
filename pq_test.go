package pq

import (
	"fmt"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	fmt.Println("Test Max PQ")
	pq := NewMax()
	items := []*Item{
		{Value: 22, Priority: 200},
		{Value: 52, Priority: 100},
	}

	for _, item := range items {
		pq.Push(item.Value, item.Priority)
	}

	job := pq.Peek()
	if job.Value != items[0].Value {
		t.Fatalf("Job value must equal %s", items[0].Value)
	}

	pq.Update(job, job.Value, 50)

	if job.Priority != 50 {
		t.Fatalf("%s priority should be 50", job)
	}
	if j := pq.Peek(); j.Value == job.Value {
		t.Fatalf("%s shouldn't be on the top", job)
	}
}
