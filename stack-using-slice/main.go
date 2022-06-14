package main

import "fmt"

type Stack struct {
	list  []int
	index int
}

func (s *Stack) Push(item int) {
	s.index++
	s.list = append(s.list, item)
}

func (s *Stack) Pop() int {
	item := s.list[s.index]
	s.list = append(s.list[:s.index], s.list[s.index+1:]...)
	s.index--
	return item
}

func (s *Stack) Show() {
	fmt.Println(s.list)
}

func GetStack() *Stack {
	stk := &Stack{}
	stk.index = -1
	return stk
}

func main() {
	s1 := GetStack()

	s1.Push(12)
	s1.Push(37)
	s1.Pop()
	s1.Push(29)

	s1.Show()

}
