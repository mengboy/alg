package stack

import "alg/data-structure/list"

type Stack struct {
	List *list.List
}

func New() *Stack {
	return &Stack{
		List:new(list.List).Init(),
	}
}

func (s *Stack) Push(v interface{})  {
	s.List.InsertFirst(&list.Node{
		Val:v,
		Next:nil,
	})
}

func (s *Stack) Pop() interface{} {
	return s.List.DelFirst().Val
}

func (s *Stack) Peek() interface{}  {
	return s.List.GetFirst().Val
}

func (s *Stack) Empty() bool {
	if s.List.Next == nil {
		return true
	}
	return false
}
