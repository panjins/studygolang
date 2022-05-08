package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"stream_grpc_test/proto"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	//服务端流模式
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

func (s server) PostStream(cliStr proto.Greeter_PostStreamServer) error {
	for {
		recv, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(recv.Data)
		}

	}
	return nil
}

func (s server) AllStream(allStr proto.Greeter_AllStreamServer) error {
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
			_ = allStr.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)

		}
	}()
	wg.Wait()
	return nil
}

//func (s *Server) GetStream(ctx context.Context, req *proto.StreamReqData) (*proto.StreamResData, error) {
//	return nil, nil
//}

func main() {
	newServer := grpc.NewServer()
	proto.RegisterGreeterServer(newServer, &server{})

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	err = newServer.Serve(listen)
	if err != nil {
		panic(err)
	}

}
