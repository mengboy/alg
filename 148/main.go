package main

import (
	"fmt"
	"math"
)

/**
* 链表排序
* 1. 从新创建一个新的链表，将原链表打散从新插入
* 2. 归并排序
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
	res := sortList(list)
	out(res)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	mid := findMid(head)
	head1 := head
	head2 := mid.Next
	mid.Next = nil
	left := sortList(head1)
	right := sortList(head2)
	return merge(left, right)
}

// 查找中间节点
func findMid(head *ListNode) *ListNode  {
	if head == nil{
		return nil
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//
func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for head1 != nil || head2 != nil {
		var a, b int
		if head1 == nil{
			a = math.MaxInt32
		}else {
			a = head1.Val
		}
		if head2 == nil{
			b = math.MaxInt32
		}else {
			b = head2.Val
		}
		if a < b{
			head.Next = &ListNode{
				Val:a,
				Next:nil,
			}
			if head1 != nil{
				head1 = head1.Next
			}
		}else {
			head.Next = &ListNode{
				Val:b,
				Next:nil,
			}
			if head2 != nil{
				head2 = head2.Next
			}
		}
		head = head.Next
	}
	return dummy.Next
}

func out(node *ListNode)  {
	for node != nil{
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}