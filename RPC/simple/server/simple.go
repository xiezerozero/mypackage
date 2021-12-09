package main

import (
	"github.com/xiezerozero/mypackage/RPC/simple/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))
	rpc.RegisterName("RandomService", new(service.RandomService))

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln("net listen fail: ", err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("accept error: ", err.Error())
		}
		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
