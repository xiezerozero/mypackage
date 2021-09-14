package send

// 该包定义这个包对外暴露的接口
// 下面的子包基于该接口实现相应的代码
type Sender interface {
	Send() error
}
