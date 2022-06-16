package main

import (
	"errors"
	"fmt"

	"github.com/golang-demos/chalk"
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
	fmt.Println("Adding ", item, " to queue")
	nd := Node{}
	nd.value = item
	if nl.last != nil {
		nl.last.next = &nd
	} else {
		nl.first = &nd
		nl.last = &nd
	}
	nl.last = &nd
	nl.Show()
}

func (nl *NodeList) Dequeue() (int, error) {
	if nl.first == nil {
		fmt.Println(chalk.Red(), "No item to retrieve", chalk.Reset())
		return -1, errors.New("empty")
	}
	itemToDelete := nl.first
	nodeValue := itemToDelete.value
	nl.first = itemToDelete.next
	if nl.first == nil {
		nl.last = nl.first
	}
	itemToDelete = nil
	fmt.Println("Retrieved value :", nodeValue)
	nl.Show()
	return nodeValue, nil
}

func (nl *NodeList) Show() {
	fmt.Print("  Queue : ")
	pointer := nl.first
	if pointer != nil {
		fmt.Print(chalk.YellowLight())
		for pointer != nil {
			fmt.Print(pointer.value, " |")
			pointer = pointer.next
		}
		fmt.Print(chalk.Reset())
	} else {
		fmt.Print(chalk.Red(), "<empty>", chalk.Reset())
	}
	fmt.Print("\n")
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
	que.Dequeue()
	que.Dequeue()
	que.Enqueue(78)
	que.Enqueue(56)
	que.Enqueue(42)
	que.Dequeue()
	que.Enqueue(31)
}
