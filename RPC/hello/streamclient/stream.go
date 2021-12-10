package main

import (
	"context"
	"fmt"
	"github.com/xiezerozero/mypackage/RPC/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9100", grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer func() {
		e := conn.Close()
		fmt.Println("close err", e)
	}()

	client := hello.NewGreeterClient(conn)
	ctx := context.Background()
	//reply, _ := client.SayHello(ctx, &hello.HelloRequest{})
	stream, err := client.SayStreamHello(ctx, &hello.HelloRequest{})
	if err != nil {
		fmt.Println("SayStreamHello error", err.Error())
	}
	for {
		reply, e := stream.Recv()
		if e != nil {
			fmt.Println("error:", e.Error())
			break
		}
		fmt.Println(reply.Message)
	}
}
