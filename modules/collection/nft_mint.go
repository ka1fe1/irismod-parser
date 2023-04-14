package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type DocMsgUptickNFTMint struct {
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
	Denom     string `bson:"denom"`
	Id        string `bson:"id"`
	URI       string `bson:"uri"`
	Data      string `bson:"data"`
	Name      string `bson:"name"`
	UriHash   string `bson:"uri_hash"`
}

func (m *DocMsgUptickNFTMint) GetType() string {
	return MsgTypeNFTMint
}

func (m *DocMsgUptickNFTMint) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickNFTMint)

	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
	m.UriHash = msg.UriHash
}

func (m *DocMsgUptickNFTMint) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickNFTMint)
	addrs = append(addrs, msg.Sender, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
