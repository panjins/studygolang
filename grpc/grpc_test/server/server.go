package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_test/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello" + request.Name}, nil
}

func main() {
	//1.server
	g := grpc.NewServer()

	//	2.注册
	proto.RegisterGreeterServer(g, &Server{})

	//	3.监听
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen" + err.Error())
	}

	err = g.Serve(listen)
	if err != nil {
		panic("run failed")
	}
}
