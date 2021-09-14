package wechat

import "github.com/xiezerozero/mypackage/interface/send"

type wechat struct {
}

func (w wechat) Send() error {
	return nil
}

func NewWechatSender() send.Sender {
	return wechat{}
}
