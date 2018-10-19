package main

import "fmt"

/**
* LRU 缓存机制
* 随机存取map
* 随机插入删除list
* linkedHashMap
 */
 /**

[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],
 [5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],
 [9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],
 [12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
 *
 */

func main()  {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	l.Put(3, 3)
	l.Get(2)
	l.Put(4, 4)
	l.Get(1)
	l.Get(3)
	l.Get(4)
	l.Out()


	l.Put(10, 13)
	l.Put(3, 17)
	l.Put(6, 11)
	l.Put(10, 5)
	l.Put(9,10)
	l.Get(13)
	l.Put(2,19)
	l.Get(2)
	l.Get(3)
	l.Put(5,25)
	l.Get(8)
	l.Put(9,22)
	l.Put(5,5)
	l.Put(1,30)
	l.Get(11)
	l.Put(9,12)
	l.Get(7)
	l.Get(5)
	l.Get(8)
	l.Get(9)
	l.Put(4,30)

	l.Get(4)
	l.Out()
	l.Get(3)
	l.Out()
	l.Get(2)
	l.Out()
	l.Get(1)
	l.Out()
	l.Put(5, 5)
	l.Out()
	l.Put(3, 4)
	l.Put(4, 4)
	l.Out()
	l.Put(5, 4)
	l.Out()
}




type LRUCache struct {
	*Node
	Len int
	Cap int
	NodeMap map[int]*Node
	//   标记尾节点
	Tail *Node
}

type Node struct {
	Key int
	Val int
	Next *Node
	Pre *Node
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:capacity,
		Len:0,
		NodeMap: make(map[int]*Node, 0),
		Tail:nil,
	}
}


func (this *LRUCache) Get(key int) int {
	if node, ok := this.NodeMap[key]; ok{
		pre := node.Pre
		if pre == nil{
			return node.Val
		}

		// 移除node节点
		next := node.Next
		pre.Next = next
		if next != nil{
			next.Pre = pre
		}else {
			// TODO 更新尾节点
			this.Tail = pre
		}

		// 将当前节点放到首节点
		node.Next = this.Node
		this.Node.Pre = node
		this.Node = node
		node.Pre = nil
		return node.Val
	}

	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	newNode := &Node{
		Key:key,
		Val:value,
	}
	// 更新值，并将新的放入链表首部
	if node, ok := this.NodeMap[key]; ok{
		node.Val = value
		pre := node.Pre
		if pre == nil{
			return
		}
		next := node.Next
		if next != nil{
			next.Pre = pre
		}else {
			this.Tail = pre
		}
		pre.Next = next
		node.Next = this.Node
		this.Node.Pre = node
		this.Node = node
		return
	}
	this.NodeMap[key] = newNode
	if this.Len  < this.Cap{
		if this.Len == 0{
			this.Node = newNode
			this.Tail = newNode
			this.Len++
			return
		}
		newNode.Next = this.Node
		this.Node.Pre = newNode
		this.Node = newNode
		this.Len++
	}else {
		// todo 达到容量后，标记尾节点
		newNode.Next = this.Node
		this.Node.Pre = newNode
		this.Node = newNode
		tPre := this.Tail.Pre
		tPre.Next = nil
		delete(this.NodeMap, this.Tail.Key)
		this.Tail = tPre
	}

}



//func (this *LRUCache) DelLast() *Node {
//	n := this.Node
//	for n != nil && n.Next != nil{
//		n = n.Next
//	}
//	if _, ok := this.NodeMap[n.Key]; ok{
//		delete(this.NodeMap, n.Key)
//	}
//	pre := n.Pre
//	pre.Next = nil
//	return n
//}

func (this *LRUCache) Out()  {
	t := this.Node
	fmt.Println("l: ", this.Len, "c: ", this.Cap)
	for t != nil{
		fmt.Print(t.Val, " ")
		t = t.Next
	}
	fmt.Println()
}