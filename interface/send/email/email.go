package email

import "github.com/xiezerozero/mypackage/interface/send"

type emailSender struct {
}

func (e emailSender) Send() error {
	// todo
	return nil
}

func NewEmailSender() send.Sender {
	return emailSender{}
}
