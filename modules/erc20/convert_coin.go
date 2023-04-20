package erc20

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocMsgConvertCoin struct {
	Coin     models.Coin `bson:"coin"`
	Receiver string      `bson:"receiver"`
	Sender   string      `bson:"sender"`
}

func (m *DocMsgConvertCoin) GetType() string {
	return MsgTypeConvertCoin
}

func (m *DocMsgConvertCoin) BuildMsg(v interface{}) {
	msg := v.(*MsgConvertCoin)
	m.Coin = models.BuildDocCoin(msg.Coin)
	m.Receiver = msg.Receiver
	m.Sender = msg.Sender
}

func (m *DocMsgConvertCoin) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgConvertCoin)
	addrs = append(addrs, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
