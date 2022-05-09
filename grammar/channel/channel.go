package main

import "fmt"

//channel 通道

func main() {
	//	1.无缓冲的通道 同步通道
	/*
		var ch chan int //int 表示通道中的元素是int 类型
		fmt.Println("未初始化", ch)

		//	初始化通道
		ch = make(chan int)
		fmt.Println("初始化通道", ch)

	*/
	ch := make(chan int) //定义并初始化通道

	go func() {
		//从通道接受数据
		fmt.Println("从通道中接受数据ch:", <-ch)
	}()

	//向通道发送数据
	ch <- 100

	//	2.缓冲通道，不是同步通道

	/*
		1.不能向已关闭的通道中读取数据
		2. 缓冲区满的时候不能发送数据
		3. 缓冲区为空时不能读取数据

	*/

	ch2 := make(chan int, 3) //3表示缓冲区可以存放3个元素
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	for i := 0; i < 3; i++ {
		fmt.Println("从通道中读取的数据", <-ch2)
	}

}
