package main

import (
	"fmt"
	"alg/data-structure/queue"
)

func main() {
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Out()
	fmt.Println(q.Pop())
	q.Out()
	q.Pop()
	q.Pop()
	fmt.Println(q.Empty())
}
