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
	defer conn.Close()

	client := hello.NewGreeterClient(conn)
	ctx := context.Background()
	reply, _ := client.SayHello(ctx, &hello.HelloRequest{})
	fmt.Println(reply.Message)
}
