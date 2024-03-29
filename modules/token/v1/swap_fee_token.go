package v1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocMsgSwapFeeTokenV1 struct {
	FeePaid   models.Coin `bson:"fee_paid"`
	Recipient string      `bson:"recipient"`
	Sender    string      `bson:"sender"`
}

func (m *DocMsgSwapFeeTokenV1) GetType() string {
	return MsgTypeSwapFeeToken
}

func (m *DocMsgSwapFeeTokenV1) BuildMsg(v interface{}) {
	msg := v.(*MsgSwapFeeTokenV1)

	m.FeePaid = models.BuildDocCoin(msg.FeePaid)
	m.Recipient = msg.Recipient
	m.Sender = msg.Sender
}

func (m *DocMsgSwapFeeTokenV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgSwapFeeTokenV1)
	addrs = append(addrs, msg.Recipient, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
