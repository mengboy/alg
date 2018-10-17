package tree

import (
	"fmt"
	"alg/data-structure/queue"
)

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (tree *Tree) Init(val int64) *Tree {
	tree.Root = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	return tree
}

func (tree Tree) Insert(v int64) {
	tree.Root.Insert(v)
}

func (tree Tree) Insert_BST(v int64) {
	tree.Root.Insert_BST(v)
}

func (tree Tree) Del(v int64) {
	tree.Root.del(5)
}

func (tree Tree) LayerOut() {
	tree.Root.layerOut()
}

func (tree Tree) Search(v int64) *TreeNode {
	return tree.Root.search(v)
}


// 递归插入
func (node *TreeNode) Insert(v int64) {
	if v < node.Val {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &TreeNode{v, nil, nil}
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &TreeNode{v, nil, nil}
		}
	}
}

// 非递归插入
func (node *TreeNode) Insert_BST(v int64) {
	t := node
	for {
		if v < t.Val {
			if t.Left != nil {
				t = t.Left
			} else {
				t.Left = &TreeNode{v, nil, nil}
				break
			}
		} else {
			if t.Right != nil {
				t = t.Right
			} else {
				t.Right = &TreeNode{v, nil, nil}
				break
			}
		}
	}
}

// 插入子树
func (node *TreeNode) InsertSubTree(subTree *TreeNode) {
	if subTree == nil {
		return
	}
	t := node
	for {
		if subTree.Val < t.Val {
			if t.Left != nil {
				t = t.Left
			} else {
				t.Left = subTree
				break
			}
		} else {
			if t.Right != nil {
				t = t.Right
			} else {
				t.Right = subTree
				break
			}
		}
	}
}

// 找到要被删除的点
// 被删除点的父子节点
// 被删除点的字节点

// 被删除点是左子节点
// 被删除点的右子节点代替被删除点
// 被删除点的左子节点以右子节点为起始节点遍历，执行插入

// 被删除点是右子节点
// 被删除节点的右子节点代替被删除的节点
// 左子节点以右子节点为起始节点执行插入
func (node *TreeNode) del(v int64) {
	t := node
	parent := node
	for t != nil {
		if v == t.Val {
			break
		}
		parent = t
		if v < t.Val {
			t = t.Left
		}

		if v > t.Val {
			t = t.Right
		}
	}
	if t == nil {
		return
	}
	// 判断t是左子节点or右子节点
	if t == parent.Left {
		// 获取被删除右子节点
		r := t.Right
		l := t.Left
		parent.Left = r
		r.InsertSubTree(l)
		return
	}

	if t == parent.Right {
		// 获取被删除节点的右子节点
		r := t.Right
		l := t.Left
		// 将被删除节点的左子节点代替原来的节点位置
		parent.Right = l
		r.InsertSubTree(r)
	}

}

// 分层打印
func (node *TreeNode) layerOut() {
	q := queue.New()
	q.Push(node)
	for !q.Empty() {
		n := q.Pop().(*TreeNode)
		fmt.Print(n.Val, " ")
		if n.Left != nil {
			q.Push(n.Left)
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
	}
}

// 查询指定节点，不存在返回nil
func (node *TreeNode) search(v int64) *TreeNode {
	t := node
	for t != nil {
		if v == t.Val {
			return t
		}
		if v < t.Val {
			t = t.Left
		}
		if v > t.Val {
			t = t.Right
		}
	}
	return nil
}

func (node *TreeNode) Out() {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	node.Left.Out()
	node.Right.Out()
}

// 中序
func (node *TreeNode) InOrder()  {
	if node == nil {
		return
	}
	node.Left.InOrder()
	fmt.Print(node.Val, " ")
	node.Right.InOrder()
}

func (node *TreeNode) PostOrder()  {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	fmt.Print(node.Val, " ")
}


// 根据前序、后序恢复二叉树
func Revert(preOrder []int64, inOrder []int64) *TreeNode {
	//
	root := revert(preOrder, 0, int64(len(preOrder)) - 1, inOrder, 0, int64(len(inOrder)) - 1)
	return root
}

// 隐含条件如果start == end 直接返回节点
func revert(preOrder []int64, preStart, preEnd int64, inOrder []int64, inStart, inEnd int64) *TreeNode {
	if len(preOrder) != len(inOrder){
		return nil
	}
	if preStart > preEnd || inStart > inEnd {
		return nil
	}


	rootV := preOrder[preStart]

	root := &TreeNode{
		Val:rootV,
	}

	var i int64
	for i = inStart; i < inEnd; i++{
		if inOrder[i] == rootV{
			root.Left = revert(preOrder, preStart + 1, preStart + i - inStart, inOrder, inStart, i - 1)
			root.Right = revert(preOrder, preStart+i-inStart+1, preEnd, inOrder, i + 1, inEnd)
		}
	}
	return root
}

func RevertPostIn(postOrder []int64, inOrder []int64) *TreeNode{
	l := int64(len(postOrder)) - 1
	root := revertPostIn(postOrder, 0, l, inOrder, 0, l)
	return root
}

func revertPostIn(postOrder []int64, postStart, postEnd int64, inOrder []int64, inStart, inEnd int64) *TreeNode {
	if len(postOrder) != len(inOrder){
		return nil
	}
	if postStart > postEnd || inStart > inEnd {
		return nil
	}
	rootV := postOrder[postEnd]
	root := &TreeNode{
		Val:rootV,
	}
	var i int64
	// 根据移动的位数确定边界，而不是根据i的值
	for i = inStart; i < inEnd; i++{
		if inOrder[i] == rootV{
			// postStart+i-inStart-1  （i-inStart 遍历移动的位数， postStart 加上移动的位数 - 1）
			// 2 3 4 5 7 / 10/ 15 16 20 21   i = 5  i-inStart = 5 （2 -- 10）
			// 2 4 3 7 5 / 15 21 20 16 10    取（2 -- 7）所以要减1
			root.Left = revertPostIn(postOrder, postStart, postStart+i-inStart-1, inOrder, inStart, i - 1)
			root.Right = revertPostIn(postOrder, postStart+i-inStart, postEnd - 1, inOrder, i + 1, inEnd)
		}
	}
	return root
}