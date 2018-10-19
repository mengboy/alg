package main

import "fmt"

func main()  {
	list := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:6,
				Next:&ListNode{
					Val:4,
					Next:&ListNode{
						Val:5,
						Next:&ListNode{
							Val:6,
							Next:nil,
						},
					},
				},
			},
		},
	}
	out(list)
	r := removeElements(list, 6)
	out(r)
}


//  Definition for singly-linked list.
type ListNode struct {
      Val int
      Next *ListNode
 }

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return nil
	}

	node := &ListNode{
	}
	node.Next = head
	head = node

	step := head.Next
	last := head
	for step != nil{
		if step.Val == val{
			last.Next = step.Next
		}else {
			last = last.Next
		}
		step = step.Next
	}
	return head.Next
}

func out(head *ListNode)  {
	step := head
	for step != nil{
		fmt.Print(step.Val, " ")
		step = step.Next
	}
	fmt.Println()
}