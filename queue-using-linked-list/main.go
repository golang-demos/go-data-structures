package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type NodeList struct {
	first *Node
	last  *Node
}

func (nl *NodeList) Enqueue(item int) {
	nd := Node{}
	nd.value = item
	if nl.last != nil {
		nl.last.next = &nd
	} else {
		nl.first = &nd
		nl.last = &nd
	}
	nl.last = &nd
}

func (nl *NodeList) Dequeue() (int, error) {
	if nl.first == nil {
		fmt.Println("Empty Queue")
		return -1, errors.New("empty")
	}
	itemToDelete := nl.first
	nodeValue := itemToDelete.value
	nl.first = itemToDelete.next
	itemToDelete = nil

	return nodeValue, nil
}

func (nl *NodeList) Show() {
	pointer := nl.first
	for pointer != nil {
		fmt.Print(pointer.value, " |")
		pointer = pointer.next
	}
}

func GetQueue() *NodeList {
	nl := NodeList{}
	nl.first = nil
	nl.last = nil
	return &nl
}

func main() {
	que := GetQueue()

	que.Enqueue(17)
	que.Enqueue(78)
	que.Enqueue(56)
	que.Enqueue(42)
	que.Dequeue()
	que.Enqueue(31)

	que.Show()
}
