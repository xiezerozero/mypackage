package main

import (
	"github.com/xiezerozero/mypackage/RPC/simple/service"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 测试: curl localhost:1234/jsonrpc -X POST --data '{"method": "RandomService.Rand", "params":[1217281927asdsadasd1921], "id":0}'
func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))
	rpc.RegisterName("RandomService", new(service.RandomService))

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		// 将http请求转到rpc服务中提供服务
		var conn io.ReadWriteCloser = struct {
			io.ReadCloser
			io.Writer
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		codec := jsonrpc.NewServerCodec(conn)
		rpc.ServeCodec(codec)
	})

	http.ListenAndServe(":1234", nil)
}
