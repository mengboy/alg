package util

import (
	"alg/data-structure/queue"
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

func (t *TreeNode) InOrder()  {
	if t == nil{
		return
	}
	t.Left.InOrder()
	fmt.Print(t.Val, " ")
	t.Right.InOrder()
}

func (t *TreeNode) PreOrder()  {
	if t == nil{
		return
	}
	fmt.Print(t.Val, " ")
	t.Left.PreOrder()
	t.Right.PreOrder()
}

func (t *TreeNode) PostOrder()  {
	if t == nil{
		return
	}
	
	t.Left.PostOrder() 
	t.Right.PostOrder()
	fmt.Print(t.Val, " ")
}

func (t *TreeNode) Layer()  {
	queue := queue.New()

}