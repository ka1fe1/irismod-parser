package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type (
	DocMsgUptickNFTTransfer struct {
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
		URI       string `bson:"uri"`
		Name      string `bson:"name"`
		Denom     string `bson:"denom"`
		Id        string `bson:"id"`
		Data      string `bson:"data"`
		UriHash   string `bson:"uri_hash"`
	}
)

func (m *DocMsgUptickNFTTransfer) GetType() string {
	return MsgTypeNFTTransfer
}

func (m *DocMsgUptickNFTTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickNFTTransfer)

	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
	m.UriHash = msg.UriHash
}

func (m *DocMsgUptickNFTTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickNFTTransfer)
	addrs = append(addrs, msg.Sender, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
