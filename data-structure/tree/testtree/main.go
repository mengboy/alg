package main

import (
	"fmt"
	"leetcode-alg/data-structure/tree"
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
	n := t.Search(5)
	fmt.Println(n)
	t.LayerOut()
	fmt.Println()
	t.Del(5)
	t.Root.Out()
}
