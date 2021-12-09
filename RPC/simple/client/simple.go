package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

var client *rpc.Client

func main() {
	// 默认格式
	//var err error
	//client, err = rpc.Dial("tcp", ":1234")
	//if err != nil {
	//	log.Fatalln("dialing err: ", err.Error())
	//}
	//rand()

	// json格式
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalln("dialing err: ", err.Error())
	}
	client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	rand()
}

func hello() {
	var reply string
	err := client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}

func rand() {
	var result int64
	err := client.Call("RandomService.Rand", time.Now().Unix(), &result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
