package main

import "fmt"

//单向通道

/*
chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。

*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)      //数据写入ch1
	go squarer(ch2, ch1) //把ch1中的数据 写入ch2
	printer(ch2)         //输出 ch2中的数据

}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)

}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}

}