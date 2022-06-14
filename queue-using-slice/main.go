package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	list []int
}

func (q *Queue) enqueue(item int) {
	q.list = append(q.list, item)
}

func (q *Queue) dequeue() (int, error) {
	length := len(q.list)
	if length > 0 {
		item := q.list[0]
		q.list = q.list[1:]
		return item, nil
	} else {
		fmt.Println("Empty Queue")
		return -1, errors.New("empty queue")
	}
}

func (q *Queue) Show() {
	for _, item := range q.list {
		fmt.Print(item, " | ")
	}
}

func getQueue() *Queue {
	q := Queue{}
	return &q
}

func main() {
	que := getQueue()

	que.enqueue(17)
	que.enqueue(78)
	que.enqueue(56)
	que.enqueue(42)
	que.dequeue()
	que.enqueue(31)

	que.Show()
}
