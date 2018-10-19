package main

import (
	"alg/data-structure/tree"
	"fmt"
)

func main() {
	t := new(tree.Tree).Init(10)
	t.Insert(5)
	t.Insert(16)
	t.Insert_BST(20)
	t.Insert_BST(3)
	t.Insert_BST(2)
	t.Insert_BST(7)
	t.Insert(21)
	t.Insert(15)
	t.Insert(4)

	t.LayerOut()

	//t.LayerOut()
	//fmt.Println()
	//t.Del(5)
	t.Root.Out()
	fmt.Println()
	t.Root.PreOrder2()
	fmt.Println()
	t.Root.InOrder()
	fmt.Println()
	t.Root.InOrder2()
	fmt.Println()
	t.Root.PostOrder()
	fmt.Println()
	t.Root.PostOrder2()
	fmt.Println()

	preOrder := []int64{
		10, 5, 3, 2, 4, 7, 16, 15, 20, 21,
	}

	inderOrder := []int64{
		2, 3, 4, 5, 7, 10, 15, 16, 20, 21,
	}

	postOrder := []int64{
		2, 4, 3, 7, 5, 15, 21, 20, 16, 10,
	}

	t2 := tree.Revert(preOrder, inderOrder)

	t2.Out()
	fmt.Println()

	t3 := tree.RevertPostIn(postOrder, inderOrder)
	t3.Out()
}
