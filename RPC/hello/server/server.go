package main

import (
	"context"
	"fmt"
	"github.com/xiezerozero/mypackage/RPC/hello"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var count int

type Server struct {
}

func (s Server) SayStreamHello(request *hello.HelloRequest, server hello.Greeter_SayStreamHelloServer) error {
	go func() {
		fmt.Println("new request")
		for i := 0; i < 10; i++ {
			e := server.Send(&hello.HelloReply{Message: strconv.Itoa(i)})
			if e != nil {
				fmt.Println("error", e.Error())
			}
		}
	}()
	return nil
}

func (s Server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	count++
	return &hello.HelloReply{Message: "hello" + strconv.Itoa(count)}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9100")
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &Server{})
	log.Println("start gRPC listen on port " + "9100")
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
