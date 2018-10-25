package main

import "fmt"

/**
* 链表插入排序
* 新建一个链表，将原链表进行插入
 */

func main()  {
	list := &ListNode{
		Val:4,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:1,
				Next:&ListNode{
					Val:3,
					Next:nil,
				},
			},
		},
	}

	res := insertionSortList2(list)
	out(res)

}

type ListNode struct {
 	Val int
 	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	newhead := &ListNode{}
	newhead.Next = head
	curr := head.Next
	// curr 上一个节点
	clast := head
	for curr != nil{
		step := newhead.Next
		// step 上一个节点
		pre := newhead
		for step != nil && step != curr{
			if step.Val > curr.Val{
				pre.Next = curr
				nextNext := curr.Next
				curr.Next = step
				clast.Next = nextNext
				break
			}
			step = step.Next
			pre = pre.Next
		}
		clast = step
		curr = step.Next
	}

	return newhead.Next
}

func out(node *ListNode)  {
	for node != nil{
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}


// 相当于重新插入节点
func insertionSortList2(head *ListNode) *ListNode {
	newhead := &ListNode{
	}
	curr := head
	for curr != nil{
		pre := newhead
		// 查找比curr大的节点，查找要插入节点的位置
		for pre.Next != nil && pre.Next.Val < curr.Val{
			pre = pre.Next
		}
		temp := curr.Next
		curr.Next = pre.Next
		pre.Next = curr
		curr = temp
	}
	return newhead.Next
}