package main

import "fmt"

/**
* 输入：( 2 -> 4 -> 3 -> 5) + (5 -> 6 -> 4)
* 输出：7 -> 0 -> 8 -> 5
* 原因：3425 + 465 = 5807
 */

func main() {
	//l1 := &ListNode{
	//	Val:2,
	//	Next:&ListNode{
	//		Val:4,
	//		Next:&ListNode{
	//			Val:3,
	//			Next:&ListNode{
	//				Val:5,
	//				Next:nil,
	//			},
	//		},
	//},
	//}
	//
	//l2 := &ListNode{
	//	Val:5,
	//	Next:&ListNode{
	//		Val:6,
	//		Next:&ListNode{
	//			Val:4,
	//			Next:nil,
	//	},
	//	},
	//}

	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  0,
		Next: nil,
	}

	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Print(l3.Val, " ")
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var l3 *ListNode = &ListNode{}
	t := l3
	var ten int
	for l1 != nil && l2 != nil {
		t := l1.Val + l2.Val + ten
		if t > 9 {
			ten = t / 10
			one := t - 10
			var tl3 *ListNode = &ListNode{}
			tl3.Val = one
			tl3.Next = nil
			l3.Next = tl3
			l3 = tl3
		} else {
			var tl3 *ListNode = &ListNode{}
			tl3.Val = t
			tl3.Next = nil
			l3.Next = tl3
			l3 = tl3
			ten = 0
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if ten > 0 {
			ten, l3 = add(ten, l3, l1)

		} else {
			l3.Next = l1
			l3 = l1
		}

		l1 = l1.Next
	}

	for l2 != nil {
		if ten > 0 {
			ten, l3 = add(ten, l3, l2)
		} else {
			l3.Next = l2
			l3 = l2
		}
		l2 = l2.Next

	}

	if ten > 0 {
		var tl3 *ListNode = &ListNode{}
		tl3.Val = ten
		tl3.Next = nil
		l3.Next = tl3
		l3 = tl3
	}

	return t.Next
}

func add(ten int, l3 *ListNode, l *ListNode) (int, *ListNode) {
	t := l.Val + ten
	if t > 9 {
		ten = t / 10
		one := t - 10
		var tl3 *ListNode = &ListNode{}
		tl3.Val = one
		tl3.Next = nil
		l3.Next = tl3
		l3 = tl3
	} else {
		var tl3 *ListNode = &ListNode{}
		tl3.Val = t
		tl3.Next = nil
		l3.Next = tl3
		l3 = tl3
		ten = 0
	}
	return ten, l3
}
