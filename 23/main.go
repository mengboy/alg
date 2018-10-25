package main

import (
	"fmt"
	"math"
)

/**
* 合并有序链表
 */

func main(){
	list1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:4,
			Next:&ListNode{
				Val:5,
				Next:nil,
			},
		},
	}
	list2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}
	list3 := &ListNode{
		Val:2,
		Next:&ListNode{
			Val:6,
			Next:nil,
		},
	}
	lists := []*ListNode{
		list1, list2, list3,
	}

	mergedList := mergeKLists(lists)
	out(mergedList)

}

func out(list *ListNode)  {
	for list != nil{
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || len(lists) == 1{
		if len(lists) == 1{
			return lists[0]
		}
		return nil
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[0:mid])
	right := mergeKLists(lists[mid:len(lists)])
	return merge2Lists(left, right)
}


func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode  {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil || list2 != nil{
		var a, b int
		if list1 == nil{
			a = math.MaxInt32
		}else {
			a = list1.Val
		}

		if list2 == nil{
			b = math.MaxInt32
		}else {
			b = list2.Val
		}

		if a < b {
			newNode := &ListNode{
				Val:a,
				Next: nil,
			}
			head.Next = newNode
			if list1 != nil{
				list1 = list1.Next
			}
		}else {
			newNode := &ListNode{
				Val:b,
				Next: nil,
			}
			head.Next = newNode
			if list2 != nil{
				list2 = list2.Next
			}
		}
		head = head.Next
	}
	return dummy.Next
}