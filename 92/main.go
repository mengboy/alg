package main

import ."alg/util"

/**
* 两两反转链表
 */

func main()  {
	list := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:&ListNode{
					Val:4,
					Next:&ListNode{
						Val:5,
						Next:nil,
					},
				},
			},
		},
	}
	res := reverseBetween(list, 2, 4)
	Out(res)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	i := 1
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	curr := head
	for i < n {
		if i >= m && i < n{
			next := curr.Next
			nextNext := curr.Next.Next
			pN := pre.Next
			pre.Next = next
			next.Next = pN
			curr.Next = nextNext
		}
		// 查找pre
		if i < m{
			curr = curr.Next
			pre = pre.Next
		}
		i++

	}
	return dummy.Next
}


