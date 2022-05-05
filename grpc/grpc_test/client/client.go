package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic("connect failed " + err.Error())
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangsan"})

	fmt.Println(r.Message)

}
