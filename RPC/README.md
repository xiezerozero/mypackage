### 官方rpc
* rpc.Register 注册rpc服务  
* rpc.ServeConn, 对tcp链接进行服务, 默认使用gob数据模式传输  
* rpc.ServeCodec(jsonrpc.NewServerCodec(conn)), 切换为json数据格式传输  
* 可以定义http请求,将http请求转提交到rpc服务  
    1. 参考 RPC/simple/server/http/simple.go
    2. 最后输出由http.ResponseWriter 输出
  

### proto buffer
* RPC/proto目录下执行 protoc -I=./simple --go_out=. ./simple/hello.proto
 -I 指定proto所在的目录  
  --go_out生成go的目录(proto里面定义了go_package会相应使用子目录,这里定义为/pb,会在./pb下生成go文件)
  最后指定proto文件
  
* // 当前目录生成pb.go文件  protoc --go_out=plugins=grpc:./ hello.proto  
  // protoc --go_out=plugins=grpc:./ hello/hello.proto  上级目录执行,生成goods/hello.pb.go, go_package指定包名
  // plugins=grpc, 以grpc插件的形式生成pb.go文件