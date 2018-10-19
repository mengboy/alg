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
				Next: nil,
			},
		},
	}

	//res := swapPairs(&list)

	//for res != nil {
	//	fmt.Println(res.Val)
	//	res = res.Next
	//}

	res2 := swapList(&list)
	for res2 != nil {
		fmt.Println(res2.Val)
		res2 = res2.Next
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

func swapList(head *ListNode) *ListNode {
	// 节点数小于2，直接返回
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{
	}
	newHead.Next = head
	pre := newHead
	// 但前节点
	curr := head
	for curr != nil && curr.Next != nil{
		// 当前节点的下一节点
		next := curr.Next
		// 下下一个节点
		nextNext := curr.Next.Next
		// 交换当前节点和当前节点的下一节点
		pre.Next = next
		pre.Next.Next = curr
		// 更新下次交换的当前节点的前置节点
		pre = pre.Next.Next
		// 更新当前节点下一节点
		curr.Next = nextNext
		// 更新当前节点
		curr = curr.Next
	}
	return newHead.Next
}