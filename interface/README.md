### accept interfaces, return structs
1. 对参数接收呢,要自由化
2. 对返回呢,要保守
3. 使用方决定限制,比如可以用interface类型接收返回值,限制了使用方的行为, 例如:
    使用方只需要使用Write方法,可以使用io.Writer类型限制,虽然返回os.File,但是被当做Writer使用
```go
var w io.Writer = GetFile()
```

### accept interfaces, return interface
* 返回和上面不一致, 主要是为了控制使用方, 这一点可以参考context.Context, 通过返回值限制调用方的行为  
* 返回接口类型一般来说是包已经封装好了逻辑,其他包调用方不需要再实现了,直接调用包暴露的接口方法即可 
* 例如,send包封装好了消息发送,暴露了Sender接口, 对于外部包来说,只需要调用Send方法,在send包下面定义多个子包,分别通过
wechat,email,phone,dingding等实现Send接口,这个实现对于外部包来说,只需要提供各个实现的定义返回Sender接口即可.
  (甚至这些实现的struct都不需要对外暴露)
```go
type Sender interface{
	Send() error
}
type email struct

func (e email)Send() error {}

func NewEmailSender() Sender{
	return email{}
}


```
