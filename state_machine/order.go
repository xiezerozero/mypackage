package state_machine

import (
	"errors"
	"math/rand"
	"strconv"
)

// 订单流程: 初始化->确认->付款->收货->结束
// 取消: 初始化->确认->取消
// (退货等流程)

type State uint8

const (
	STATE_UNKNOWN = iota
	STATE_INITIAL
	STATE_CONFIRM
	STATE_CANCEL
	STATE_PAIED
	STATE_GET
	STATE_DONE
)

type Order struct {
	OrderId string
	state   State
}

func NewOrder() *Order {
	return &Order{
		OrderId: strconv.Itoa(rand.Intn(1000)),
		state:   STATE_INITIAL,
	}
}

func (o *Order) Confirm() error {
	if o.state != STATE_INITIAL {
		return errors.New("状态错误")
	}
	o.state = STATE_CONFIRM
	return nil
}

func (o *Order) Cancel() error {
	if o.state != STATE_INITIAL {
		return errors.New("状态错误")
	}
	o.state = STATE_CANCEL
	return nil
}

func (o *Order) Pay() error {
	if o.state != STATE_CONFIRM {
		return errors.New("状态错误")
	}
	o.state = STATE_PAIED
	return nil
}

func (o *Order) Get() error {
	if o.state != STATE_PAIED {
		return errors.New("状态错误")
	}
	o.state = STATE_GET
	return nil
}

func (o *Order) Done() error {
	if o.state != STATE_GET {
		return errors.New("状态错误")
	}
	o.state = STATE_DONE
	return nil
}

// ...
