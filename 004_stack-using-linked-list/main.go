package main

import (
	"errors"
	"fmt"

	"github.com/golang-demos/chalk"
)

type Node struct {
	item int
	next *Node
}

type NodeList struct {
	first *Node
	last  *Node
}

func (nl *NodeList) Push(item int) {
	fmt.Println("Pushing", item, "to stack")
	nn := &Node{
		item, nil,
	}
	if nl.last == nil {
		nl.first = nn
		nl.last = nl.first
	} else {
		nl.last.next = nn
		nl.last = nl.last.next
	}
	nl.Show()
}

func (nl *NodeList) Pop() (int, error) {
	if nl.last == nil {
		fmt.Println("Pop() Called : No item to retrieve")
		nl.Show()
		return -1, errors.New("empty")
	}
	poppedNode := nl.last
	if nl.first == nl.last {
		nl.first = nil
		nl.last = nil
	}

	pointer := nl.first
	for pointer != nil {
		if pointer.next == nl.last {
			pointer.next = nil
			nl.last = pointer
			break
		}
		pointer = pointer.next
	}
	fmt.Println("Popped Value : ", poppedNode.item)
	nl.Show()
	return poppedNode.item, nil
}

func (nl *NodeList) Show() {
	fmt.Print("  Stack : ", chalk.YellowLight())
	if nl.last == nil {
		fmt.Print(chalk.Red(), "<empty>", chalk.Reset())
	}
	pointer := nl.first
	for pointer != nil {
		fmt.Print(pointer.item, " | ")
		pointer = pointer.next
	}
	fmt.Println(chalk.Reset())
}

func GetStack() *NodeList {
	nl := &NodeList{}
	return nl
}

func main() {
	s1 := GetStack()

	s1.Push(12)
	s1.Pop()
	s1.Pop()
	s1.Push(37)
	s1.Push(82)
	s1.Pop()
	s1.Push(29)
}
