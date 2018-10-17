package avltree

import (
	"math"
	"fmt"
)

type AVLTreeNode struct {
	High int64
	Val int64
	Left *AVLTreeNode
	Right *AVLTreeNode
	Freq int64
}

type AVLTree struct {
	Root *AVLTreeNode
}

func (avl AVLTree) Insert(v int64)  {
	avl.Root.insert(v)
}

// 插入
func (avl *AVLTreeNode) insert(v int64)  {
	if avl == nil{
		avl = &AVLTreeNode{
			Val:v,
		}
		return
	}
	if avl.Val > v{
		avl.Left.insert(v)
		if avl.Left.getHigh() - avl.Right.getHigh() == 2{
			if v < avl.Left.Val{
				//ll

			}else {
				// lr

			}
		}
	}else if avl.Val < v{
		avl.Right.insert(v)
		if avl.Right.getHigh() - avl.Left.getHigh() == 2{
			if v < avl.Right.Val{
				// rl

			}else {
				// rr
			}
		}
	}else {
		// already exits
		avl.Freq++
	}
	avl.High = int64(math.Max(float64(avl.Left.getHigh()), float64(avl.Right.getHigh()))) + 1
}

func (avl *AVLTreeNode) llroutate()  {
	lr := avl.Left.Right
	lr.Right = avl
	avl.Left = nil
}

func (avl *AVLTreeNode) lrroute()  {
	avl.Left.rrroutate()
}

func (avl *AVLTreeNode) rrroutate()  {
	rl := avl.Right.Left
	rl.Left = avl
	avl.Right = nil
}

func (avl *AVLTreeNode) rlroute()  {

}

// 获取节点高度
func (avl *AVLTreeNode) getHigh() int64 {
	if avl != nil{
		return avl.High
	}
	return -1
}

func (avl *AVLTreeNode) out()  {
	if avl == nil{
		return
	}
	fmt.Print(avl.Val, " ")
	avl.Left.out()
	avl.Right.out()
}


func (avl *AVLTreeNode) singRotateLeft()  {
	
}