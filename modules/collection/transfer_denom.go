package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type DocMsgUptickTransferDenom struct {
	Id        string `bson:"id"`
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
}

func (m *DocMsgUptickTransferDenom) GetType() string {
	return MsgTypeTransferDenom
}

func (m *DocMsgUptickTransferDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickTransferDenom)

	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
	m.Id = strings.ToLower(msg.Id)
}

func (m *DocMsgUptickTransferDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickTransferDenom)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
