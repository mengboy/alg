package main

import "fmt"

/**
* 两两交换链表节点
 */

func main() {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := swapPairs(&list)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	newHead.Next = head
	pre := newHead

	for pre.Next != nil && pre.Next.Next != nil {
		// 1
		curr := pre.Next
		// 2
		next := curr.Next
		// 1 -> 3
		curr.Next = next.Next
		// 2 -> 1
		next.Next = pre.Next

		// 0 -> 2
		// 0 -> 2 -> 1 -> 3 -> 4
		pre.Next = next

		// 1
		pre = curr
	}
	return newHead.Next
}
