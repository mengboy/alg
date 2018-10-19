package main

/**
* 求两链表交叉点
* 遍历分别获取两链表长度
* 相减得长比短多n节点
* 长得先走n
* 一起走，第一次相遇即交叉点，没有相遇不存在交叉点
 */

import (
	. "alg/util"
	"fmt"
)

func main()  {
	node := &ListNode{
		Val:4,
		Next:&ListNode{
			Val:5,
			Next:nil,
		},
	}

	list1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:node,
			},
		},
	}

	list2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:node,
		},
	}

	fmt.Println(IntersectionNode(list1, list2).Val)

}

func IntersectionNode(list1 *ListNode, list2 *ListNode) *ListNode {
	if list2  == nil || list1 == nil{
		return nil
	}

	t1 := list1
	t2 := list2

	l1 := 0
	l2 := 0
	for t1 != nil{
		l1++
		t1 = t1.Next
	}
	for t2 != nil{
		l2++
		t2 = t2.Next
	}
	if l2 > l1{
		n := l2 - l1
		for i := 0; i < n; i++{
			list2 = list2.Next
		}
		for list1 != nil && list2 != nil{
			if list1 == list2{
				return list2
			}
			list2 = list2.Next
			list1 = list1.Next
		}
	}else {
		n := l1 - l2
		for i := 0; i < n; i++ {
			list1 = list1.Next
		}
		for list1 != nil && list2 != nil{
			if list1 == list2{
				return list2
			}
			list2 = list2.Next
			list1 = list1.Next
		}
	}

	return nil
}
