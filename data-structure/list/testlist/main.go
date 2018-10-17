package main

import (
	"alg/data-structure/list"
)

type Node struct {
	Value interface{}

	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type SinglyLinkedList struct {
	front  *Node
	length int
}

func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.length = 0
	return s
}
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}
func (s *SinglyLinkedList) Back() *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (s *SinglyLinkedList) Append(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}
	s.length++
}

// 2 -> 3 -> 4 -> 5 -> 3
//2345345345
//2435435

func main() {
	n := new(list.List).Init()
	n.InsertTailV(1)
	n.InsertFirst(&list.Node{
		Val:  2,
		Next: nil,
	})
	n.InsertFirst(&list.Node{
		Val:  3,
		Next: nil,
	})
	n.InsertTail(&list.Node{
		Val:  4,
		Next: nil,
	})
	n.InsertAfterSomeNode(4, &list.Node{
		Val:  5,
		Next: nil,
	})
	//n.Out()

	//n.Next.Next.Next.Next.Next = n.Next
	//fmt.Println(n.IsRing())
	//n.Reverse()
	//n.Out()
	h := list.Reverse2(n.Next)
	n.Node.Next = h
	n.Out()

}
