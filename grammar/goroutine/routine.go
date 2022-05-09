package main

import (
	"fmt"
	"sync"
)

//go 语言协程是一种轻量级线程，GO运行时调度

func Test(n int) {
	defer wg.Done() //相当于wg.Add(-1)
	fmt.Println(n)
}

//WaitGroup 会阻塞主线程 等所有的goroutine 执行完成
var wg sync.WaitGroup

func main() {
	fmt.Println("并发编程")
	for i := 0; i < 1000; i++ {
		wg.Add(1) //添加或减少goroutine 的数量
		go Test(i)
	}
	wg.Wait() //执行阻塞 等待所有协程执行完毕

}
