package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type (
	DocMsgUptickNFTBurn struct {
		Sender string `bson:"sender"`
		Id     string `bson:"id"`
		Denom  string `bson:"denom"`
	}
)

func (m *DocMsgUptickNFTBurn) GetType() string {
	return MsgTypeNFTBurn
}

func (m *DocMsgUptickNFTBurn) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickNFTBurn)

	m.Sender = msg.Sender
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
}

func (m *DocMsgUptickNFTBurn) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickNFTBurn)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
