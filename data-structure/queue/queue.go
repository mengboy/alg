package queue

import "alg/data-structure/list"

// 先进先出
// list 尾插入
// 头节点删除

type Queue struct {
	List *list.List
}

func (q *Queue) Push(v interface{}) {
	q.List.InsertTailV(v)
}

func (q *Queue) Pop() interface{} {
	return q.List.DelFirst().Val
}

func (q *Queue) Empty() bool {
	if q.List.Next == nil {
		return true
	}
	return false
}

func New() *Queue {
	q := new(Queue)
	l := new(list.List).Init()
	q.List = l
	return q
}

func (q *Queue) Out() {
	q.List.Out()
}

func (q *Queue) Len() int64 {
	return q.List.Length
}
