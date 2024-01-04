package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func libraryEnqueue() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue("string")

	for q.Len() > 0 {
		val := q.Dequeue()
		fmt.Println("Dequeue", val)
	}
}

/*
	Difference between any and interface{}?
*/
type Queue struct {
	items []any
}

func (q *Queue) enqueue(value any) {
	q.items = append(q.items, value)
}

func (q *Queue) dequeue() any {
	if len(q.items) == 0 {
		return nil
	}
	rm := q.items[0]
	q.items = q.items[1:]
	return rm
}

func (q *Queue) printQueue() {
	for _, e := range q.items {
		fmt.Println(e)
	}
}
