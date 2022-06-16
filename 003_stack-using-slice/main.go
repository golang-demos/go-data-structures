package main

import (
	"errors"
	"fmt"

	"github.com/golang-demos/chalk"
)

type Stack struct {
	list  []int
	index int
}

func (s *Stack) Push(item int) {
	fmt.Println("Pushing ", item, " to stack")
	s.index++
	s.list = append(s.list, item)
	s.Show()
}

func (s *Stack) Pop() (int, error) {
	if s.index == -1 {
		fmt.Println("Pop() called : ", chalk.Red(), "No item to retrieve", chalk.Reset())
		return -1, errors.New("empty")
	}
	item := s.list[s.index]
	s.list = append(s.list[:s.index], s.list[s.index+1:]...)
	s.index--
	fmt.Println("Popped value :", item)
	s.Show()
	return item, nil
}

func (s *Stack) Show() {
	fmt.Print("  Stack : ")
	if len(s.list) > 0 {
		fmt.Print(chalk.YellowLight(), s.list, chalk.Reset())
	} else {
		fmt.Print(chalk.Red(), "<empty>", chalk.Reset())
	}
	fmt.Print("\n")
}

func GetStack() *Stack {
	stk := &Stack{}
	stk.index = -1
	return stk
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
