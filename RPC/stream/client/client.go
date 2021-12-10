package main

import (
	"github.com/xiezerozero/mypackage/RPC/stream"
	"google.golang.org/grpc"

	"context"
	_ "google.golang.org/grpc/balancer/grpclb"
	"log"
	"time"
)

const (
	ADDRESS = "localhost:50051"
)

func putstream(c stream.GreeterClient) {
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&stream.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}

func main() {

	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := stream.NewGreeterClient(conn)

	putstream(c)
	return
	//调用服务端推送流
	reqstreamData := &stream.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&stream.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&stream.StreamReqData{Data: "client"})
			time.Sleep(time.Second)
		}
	}()

	select {}

}
