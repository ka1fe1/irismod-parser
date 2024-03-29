package v1beta1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocMsgTransferTokenOwner struct {
	SrcOwner string `bson:"src_owner"`
	DstOwner string `bson:"dst_owner"`
	Symbol   string `bson:"symbol"`
}

func (m *DocMsgTransferTokenOwner) GetType() string {
	return MsgTypeTransferTokenOwner
}

func (m *DocMsgTransferTokenOwner) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferTokenOwner)

	m.Symbol = msg.Symbol
	m.SrcOwner = msg.SrcOwner
	m.DstOwner = msg.DstOwner
}

func (m *DocMsgTransferTokenOwner) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransferTokenOwner)
	addrs = append(addrs, msg.SrcOwner, msg.DstOwner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
