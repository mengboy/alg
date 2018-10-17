package main

import (
	"fmt"
	"runtime"
	"sync"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// &stu 指向的是同以地址
	// 遍历过程中回将这个数据拷贝到这个地址
	// 最终m所有k指向的是同一v
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m{
		fmt.Println(k,":", v)
	}
}

func selectTest() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}



func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {

	//ua.Lock()
	if ua.ages == nil{
		ua.ages = map[string]int{}
	}
	//defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//func (ua *UserAges) Get(name string) int {
//	ua.Lock()
//	defer ua.Unlock()
//	if age, ok := ua.ages[name]; ok {
//		return age
//	}
//	return -1
//}

func main() {
	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++{
		go func(x int) {
			ua := new(UserAges)
			ua.Add("xx" , x)
			fmt.Println(ua.Get("xx" ))
			wg.Done()
		}(i)
	}
	wg.Wait()

}
//func main()  {
//	//pase_student()
//	selectTest()
//}
