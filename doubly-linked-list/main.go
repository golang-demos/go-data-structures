package main

import "fmt"

type node struct {
	data     int
	next     *node
	previous *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	if l.head != nil {
		l.head.previous = n
	}
	n.next = l.head
	l.head = n
	l.length++
}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		for currentHead := l.head; currentHead != nil; {
			if currentHead.next == nil {
				currentHead.next = n
				n.previous = currentHead
				break
			}
			currentHead = currentHead.next
		}
	}
	l.length++
}

func (l *linkedList) delete(value int) {
	fmt.Println("\nDeleting by value : ", value)
	fmt.Print("Before deleting : ")
	l.print()
	var previousNode *node
	for currentHead := l.head; currentHead != nil; {
		if currentHead.data == value {
			if previousNode == nil {
				l.head = l.head.next
				if l.head != nil {
					l.head.previous = nil
				}
				l.length--
			} else {
				previousNode.next = currentHead.next
				if currentHead.next != nil {
					currentHead.next.previous = previousNode
				}
				l.length--
			}
		}

		previousNode = currentHead
		currentHead = currentHead.next
	}
	fmt.Print("\nAfter deleting  : ")
	l.print()
	fmt.Printf("\n")
}

func (l *linkedList) print() {
	fmt.Print("l(", l.length, ") : ")
	for currentHead := l.head; currentHead != nil; {
		fmt.Print(currentHead.data, ", ")
		currentHead = currentHead.next
	}
}

func main() {
	myList := linkedList{}
	myList.prepend(&node{data: 10})
	myList.prepend(&node{data: 20})
	myList.append(&node{data: 30})
	myList.prepend(&node{data: 40})
	myList.prepend(&node{data: 50})
	myList.append(&node{data: 60})
	myList.append(&node{data: 40})

	myList.print()

	// Delete a node
	myList.delete(20)

	// Delete a node that occurs twice
	myList.delete(40)

	// Delete head node
	myList.delete(50)

	// Try to delete a node that does not exist
	myList.delete(500)

	// TODO: Add more test cases
}
