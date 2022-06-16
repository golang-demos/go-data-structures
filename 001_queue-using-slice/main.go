package main

import (
	"errors"
	"fmt"

	"github.com/golang-demos/chalk"
)

type Queue struct {
	list []int
}

func (q *Queue) enqueue(item int) {
	fmt.Println("Adding ", item, " to queue")
	q.list = append(q.list, item)
	q.Show()
}

func (q *Queue) dequeue() (int, error) {
	length := len(q.list)
	var err error
	returnItem := -1
	if length > 0 {
		item := q.list[0]
		q.list = q.list[1:]
		returnItem = item
		fmt.Println("Retrieved value :", returnItem)
	} else {
		fmt.Println(chalk.Red(), "No item to retrieve", chalk.Reset())
		err = errors.New("empty queue")
	}
	q.Show()
	return returnItem, err
}

func (q *Queue) Show() {
	fmt.Print("  Queue : ")
	length := len(q.list)
	if length > 0 {
		fmt.Print(chalk.YellowLight())
		for _, item := range q.list {
			fmt.Print(item, " | ")
		}
		fmt.Print(chalk.Reset())
	} else {
		fmt.Print(chalk.Red(), "<empty>", chalk.Reset())
	}
	fmt.Print("\n")
}

func getQueue() *Queue {
	q := Queue{}
	return &q
}

func main() {
	que := getQueue()

	que.enqueue(17)
	que.dequeue()
	que.dequeue()
	que.enqueue(78)
	que.enqueue(56)
	que.enqueue(42)
	que.dequeue()
	que.enqueue(31)
}
