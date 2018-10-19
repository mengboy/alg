package main

import "fmt"

func main()  {
	list := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}
	res := rotateRight(&list, 2)

	for res != nil{
		fmt.Print(res.Val, " ")
		res = res.Next
	}

	//res2 := rotateRight2(&list, 2)
	//for res2 != nil{
	//	fmt.Print(res2.Val, " ")
	//	res2 = res2.Next
	//}
}


type ListNode struct {
	Val  int
	Next *ListNode
}

//  复杂度 O(k^2)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	newHead := &ListNode{}
	newHead.Next = head

	// 获取链表长度
	i := 0
	for head != nil{
		i++
		head = head.Next
	}

	// 链表旋转几圈后，需要再次旋转的次数
	var t int
	if k > i{
		t = k % i
	}else {
		t = k
	}

	if t == 0{
		return newHead.Next
	}else {
		// 遍历到第l节点，截断
		l := i - t
		n := 0
		head = newHead.Next
		for head != nil{
			n++
			if n == l{
				needRote := head.Next
				for needRote != nil && needRote.Next != nil{
					needRote = needRote.Next
				}
				needRote.Next = newHead.Next
				// 形成环
				newHead.Next = head.Next
				// 拆掉环
				head.Next = nil
				//th := newHead.Next
				//for j := 0; j < i - 1; j++{
				//	th = th.Next
				//}
				//th.Next = nil
				break
			}
			head = head.Next
		}
	}
	return newHead.Next
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	fast := head
	slow := head
	len := 1
	// 如果链表长度小于k，不需要遍历整个链表长
	for ; fast.Next != nil && len <= k; len++{
		fast = fast.Next
	}
	if len <= k{
		k = k % len
		fast = head
		for i := 0; i < k; i++{
			fast = fast.Next
		}
	}

	for fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
	}

	// 形成环
	fast.Next = head
	head = slow.Next
	// 拆掉环
	slow.Next = nil

	return head

}