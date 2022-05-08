package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//服务端流模式
	client := proto.NewGreeterClient(conn)
	resp, err := client.GetStream(context.Background(), &proto.StreamReqData{Data: "测试请求"})
	for {
		recv, err := resp.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv)
	}
	//	客户端流模式
	putS, err := client.PostStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("hello%d", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//	双向流模式
	allStr, _ := client.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "我是客户端"})
			time.Sleep(time.Second)

		}
	}()
	wg.Wait()
}
