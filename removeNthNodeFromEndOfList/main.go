package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l := &ListNode{
	//	Val:1,
	//	Next:&ListNode{
	//		Val:2,
	//		Next:&ListNode{
	//			Val:3,
	//			Next:&ListNode{
	//				Val:4,
	//				Next:&ListNode{
	//					Val:5,
	//					Next:nil,
	//				},
	//			},
	//		},
	//	},
	//}

	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	r := removeNthFromEnd(l, 2)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	i := 0
	q := head
	s := head
	ls := head
	for q != nil {
		q = q.Next
		if n == i {
			ls = s
			s = s.Next
			continue
		}
		if i != n {
			i++
		}
		//删除尾节点
		if i == 1 && q == nil {
			ls.Next = s.Next.Next
			return head
		}
		//删除首节点
		if i == n && q == nil {
			return head.Next
		}
	}

	ls.Next = s.Next
	return head
}
