package main

import (
	"fmt"
	"github.com/xiezerozero/mypackage/RPC/stream"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

const PORT = ":50051"

type server struct {
}

//服务端 单向流
func (s *server) GetStream(req *stream.StreamReqData, res stream.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&stream.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

//客户端 单向流
func (this *server) PutStream(cliStr stream.Greeter_PutStreamServer) error {

	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

//客户端服务端 双向流
func (this *server) AllStream(allStr stream.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&stream.StreamResData{Data: "server"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	stream.RegisterGreeterServer(s, &server{})
	//处理链接
	s.Serve(lis)
}
