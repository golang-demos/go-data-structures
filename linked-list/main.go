package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
	l.length++
}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	}
	for currentHead := l.head; currentHead != nil; {
		if currentHead.next == nil {
			currentHead.next = n
			break
		}
		currentHead = currentHead.next
	}
	l.length++
}

func (l *linkedList) print() {
	for currentHead := l.head; currentHead != nil; {
		fmt.Print(currentHead.data, ", ")
		currentHead = currentHead.next
	}
	fmt.Printf("\n")
}

func (l *linkedList) delete(value int) {
	fmt.Println("\nDeleteing with value : ", value)
	var previousNode *node
	for currentHead := l.head; currentHead != nil; {
		if currentHead.data == value {
			if previousNode == nil {
				l.head = currentHead.next
			} else {
				previousNode.next = currentHead.next
			}
			l.length--
		}
		previousNode = currentHead
		currentHead = currentHead.next
	}
}

func main() {
	list := linkedList{}
	list.prepend(&node{data: 10})
	list.prepend(&node{data: 20})
	list.append(&node{data: 30})
	list.append(&node{data: 16})
	list.append(&node{data: 32})
	list.prepend(&node{data: 27})
	list.prepend(&node{data: 8})
	list.append(&node{data: 41})
	list.append(&node{data: 16})

	list.print()

	// Delete a node
	list.delete(32)
	list.print()

	// Delete a node that occurs twice
	list.delete(16)
	list.print()

	// Delete a node that at the first
	list.delete(8)
	list.print()

	// Delete a node that do not exist
	list.delete(1000)
	list.print()
}
