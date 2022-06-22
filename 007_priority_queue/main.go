package main

import (
	"errors"
	"fmt"

	"github.com/golang-demos/chalk"
)

type NodeQueue struct {
	head  *Node
	last  *Node
	order string
}

type Node struct {
	value    int
	priority int
	next     *Node
}

func (nq *NodeQueue) Enqueue(item int, priorityParam ...int) {
	priority := item
	if len(priorityParam) == 1 {
		priority = priorityParam[0]
	}

	nn := &Node{
		item, priority, nil,
	}
	fmt.Println("Pushing ", item, " to Queue")
	if nq.last == nil {
		nq.head = nn
		nq.last = nq.head
	} else {
		cursor := nq.head
		var prevCursor *Node
		for cursor != nil {
			if cursor.priority >= priority {
				if prevCursor == nil {
					nn.next = cursor
					nq.head = nn
				} else {
					prevCursor.next = nn
					nn.next = cursor
				}
				break
			}
			prevCursor = cursor
			cursor = cursor.next
			if cursor == nil {
				prevCursor.next = nn
				nq.last = nn
			}
		}
	}
	nq.Show()
}

func (nq *NodeQueue) Dequeue() (int, error) {
	if nq.head == nil {
		fmt.Println("Dequeue() called. Queue is empty")
		return -1, errors.New("empty")
	}

	rn := nq.head
	nq.head = nq.head.next

	if nq.head == nil {
		nq.last = nil
	}
	fmt.Println("Retrieving ", rn.value, " from queue")
	nq.Show()
	return rn.value, nil
}

func (nq *NodeQueue) Show() {
	fmt.Print("   Queue : ", chalk.YellowLight())
	cursor := nq.head
	if cursor == nil {
		fmt.Print(chalk.Red(), "<empty>", chalk.Reset())
	} else {
		for cursor != nil {
			fmt.Print("[", cursor.value, ",", cursor.priority, "], ")
			cursor = cursor.next
		}
	}
	fmt.Println(chalk.Reset())
}

func getPirorityQueue(order ...string) *NodeQueue {
	nq := &NodeQueue{}
	if len(order) == 1 {
		nq.order = order[0]
	} else {
		nq.order = "asc"
	}
	return nq
}

func main() {
	fmt.Println("Ascending Priority Queue")
	que := getPirorityQueue()

	que.Enqueue(17)
	que.Dequeue()
	que.Dequeue()
	que.Enqueue(78)
	que.Enqueue(56)
	que.Enqueue(42)
	que.Dequeue()
	que.Enqueue(31)

	fmt.Println("Descending Priority Queue")
	que1 := getPirorityQueue()

	que1.Enqueue(17)
	que1.Dequeue()
	que1.Dequeue()
	que1.Enqueue(78)
	que1.Enqueue(56)
	que1.Enqueue(42)
	que1.Dequeue()
	que1.Enqueue(31)
}
