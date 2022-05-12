package main

import (
	"fmt"
)

//channel 通道

/*
	通道分为无缓冲通道和有缓冲通道
	 无缓冲通道:
	无缓冲通道(阻塞通道 同步通道） 必须要有接受方

	有缓冲通道：
	有缓冲区的通道


*/

func main() {
	//无缓冲通道栗子
	ch := make(chan int)
	go rev(ch) //启用goroutine 从通道接受值
	ch <- 10
	close(ch)

	//有缓冲通道
	ch2 := make(chan int, 1) //创建一个缓冲区容量为1的缓冲通道
	ch2 <- 10
	fmt.Println("发送成功")

	//	for range 从通道中循环取值
	rangechannel()

}

func rev(c chan int) {
	ret := <-c
	fmt.Println("接受成功", ret)
}

//for range 从通道循环取值

func rangechannel() {
	ch := make(chan int)
	ch1 := make(chan int)

	//向通道中发送值
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	//	接受值

	go func() {
		for {
			i, ok := <-ch //通道关闭后再取值 ok = false
			if !ok {
				break
			}

			ch1 <- i * i

		}
		close(ch1)
	}()

	//主goroutine中从ch1中接受打印值
	for v := range ch1 { //通道关闭后会推出for range循环
		fmt.Println(v)
	}
}
