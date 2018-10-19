package util

import "fmt"

func Out(node *ListNode)  {
	for node != nil{
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}
type ListNode struct {
	Val int
	Next *ListNode
}
