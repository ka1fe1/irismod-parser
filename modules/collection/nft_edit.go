package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type (
	DocMsgUptickNFTEdit struct {
		Sender  string `bson:"sender"`
		Id      string `bson:"id"`
		Denom   string `bson:"denom"`
		URI     string `bson:"uri"`
		Data    string `bson:"data"`
		Name    string `bson:"name"`
		UriHash string `bson:"uri_hash"`
	}
)

func (m *DocMsgUptickNFTEdit) GetType() string {
	return MsgTypeNFTEdit
}

func (m *DocMsgUptickNFTEdit) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickNFTEdit)

	m.Sender = msg.Sender
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
	m.UriHash = msg.UriHash
}

func (m *DocMsgUptickNFTEdit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickNFTEdit)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
