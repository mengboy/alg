package list

import (
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type List struct {
	Node
	Length int64
}

// 插入尾节点
func (list *List) InsertTailV(v interface{}) {
	n := &Node{
		Val:  v,
		Next: nil,
	}
	list.InsertTail(n)
}

// 删除头节点，并返回被删除的节点
func (list *List) DelFirst() *Node {
	if list.Next == nil {
		return nil
	}
	n := list.Next
	list.Next = list.Next.Next
	return n
}

func (list *List) Init() *List {
	list.Length = 0
	return list
}

func (list *List) InsertFirst(n *Node) {
	if n == nil {
		return
	}
	t := list.Next
	list.Next = n
	n.Next = t
	list.Length++
}

func (list *List) InsertTail(n *Node) {
	if n == nil {
		return
	}
	headNext := list.Node.Next
	// list本身0节点
	if headNext == nil {
		list.Node.Next = n
		return
	}
	for headNext.Next != nil {
		headNext = headNext.Next
	}
	headNext.Next = n
	list.Length++
}

func (list *List) InsertAfterSomeNode(exist interface{}, n *Node) {
	if n == nil {
		return
	}
	//t := list.Node
	head := list.Node.Next
	if head == nil {
		return
	}
	for head != nil {
		if head.Val == exist {
			n.Next = head.Next
			head.Next = n
			list.Length++
			break
		}
		head = head.Next
	}
}

// 获取头节点，遍历剩下的节点，使用头插法构造新的链表
func (list *List) ReverseList() {
	if list.Next == nil || list.Next.Next == nil {
		return
	}

	first := list.Node.Next
	next := list.Node.Next.Next
	list.Next.Next = nil
	for next != nil {
		// 3
		nextNext := next.Next
		// 头插法
		next.Next = first
		first = next
		next = nextNext
		t := &List{
			*first,
			4,
		}
		t.Out()
	}
	list.Next = first
}

func (list *List) IsRing() bool {
	if list == nil || list.Next == nil {
		return false
	}

	slow := list.Next
	quick := list.Next

	for slow != nil && quick != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			return true
		}
	}

	return false
}

func (list *List) Out() {
	t := list.Next
	fmt.Println("list len:", list.Length)
	for t != nil {
		fmt.Print(t.Val)
		t = t.Next
		if t != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()

}
