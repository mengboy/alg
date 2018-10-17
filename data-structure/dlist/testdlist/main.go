package main

import "alg/data-structure/dlist"

func main()  {
	dl := dlist.New()

	dl.InsertFirst(1)
	dl.InsertFirst(2)
	dl.InsertFirst(3)
	dl.InsertFirst(4)
	dl.InsertFirst(5)

	dl.Out()
	dl.Reverse()
	dl.Out()

	// 删除
	dl.Del(1)
	dl.Del(5)
	dl.Del(3)
	dl.Out()
}
