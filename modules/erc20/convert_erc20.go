package erc20

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocMsgConvertERC20 struct {
	ContractAddress string `bson:"contract_address"`
	Amount          string `bson:"amount"`
	Receiver        string `bson:"receiver"`
	Sender          string `bson:"sender"`
}

func (m *DocMsgConvertERC20) GetType() string {
	return MsgTypeConvertERC20
}

func (m *DocMsgConvertERC20) BuildMsg(v interface{}) {
	msg := v.(*MsgConvertERC20)
	m.ContractAddress = msg.ContractAddress
	m.Amount = msg.Amount.String()
	m.Receiver = msg.Receiver
	m.Sender = msg.Sender
}

func (m *DocMsgConvertERC20) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgConvertERC20)
	addrs = append(addrs, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
