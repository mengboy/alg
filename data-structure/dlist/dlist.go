package dlist

import "fmt"

type Node struct {
	Val interface{}
	Next *Node
	Prev *Node
}

type DList struct {
	*Node
	Length int64
}

func New() *DList {
	return &DList{
		&Node{},
		0,
	}
}

func (dl *DList ) InsertFirst(v interface{})  {
	n := &Node{
		Val:v,
		Next: nil,
		Prev:nil,
	}
	if dl.Length == 0{
		dl.Node = n
		dl.Length++
		return
	}

	n.Next = dl.Node
	dl.Node.Prev = n
	dl.Node = n
	dl.Length++
}

func (dl *DList) Reverse()  {
	if dl.Node == nil || dl.Next == nil{
		return
	}

	first := dl.Node
	next := dl.Next
	dl.Next = nil

	for next != nil{
		nextNext := next.Next
		next.Next = first
		next.Prev = nil
		first.Prev = next
		first = next
		next = nextNext
	}
	dl.Node = first
}

func (dl *DList) Del(v interface{}) *Node {
	if dl.Node == nil{
		return nil
	}
	t := dl.Node

	for t != nil{
		if v == t.Val{
			pre := t.Prev
			if pre == nil{
				dl.Node = t.Next
				t.Next.Prev = nil
				break
			}
			next := t.Next
			if next == nil{
				pre.Next = nil
				break
			}
			pre.Next = next
			next.Prev = pre
			break
		}
		t = t.Next
	}
	return t
}

func (dl *DList) Out()  {
	t := dl.Node
	for t != nil{
		fmt.Print(t.Val, " ")
		t = t.Next
	}
	fmt.Println()
}


