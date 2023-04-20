package collection

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type DocMsgUptickIssueDenom struct {
	Id               string `bson:"id"`
	Name             string `bson:"name"`
	Sender           string `bson:"sender"`
	Schema           string `bson:"schema"`
	Symbol           string `bson:"symbol"`
	MintRestricted   bool   `bson:"mint_restricted"`
	UpdateRestricted bool   `bson:"update_restricted"`
	Description      string `bson:"description"`
	Uri              string `bson:"uri"`
	UriHash          string `bson:"uri_hash"`
	Data             string `bson:"data"`
}

func (m *DocMsgUptickIssueDenom) GetType() string {
	return MsgTypeIssueDenom
}

func (m *DocMsgUptickIssueDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgUptickIssueDenom)

	m.Sender = msg.Sender
	m.Schema = msg.Schema
	m.Id = strings.ToLower(msg.Id)
	m.Name = msg.Name
	m.Symbol = msg.Symbol
	m.MintRestricted = msg.MintRestricted
	m.UpdateRestricted = msg.UpdateRestricted
	m.Description = msg.Description
	m.Uri = msg.Uri
	m.UriHash = msg.UriHash
	m.Data = msg.Data
}

func (m *DocMsgUptickIssueDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUptickIssueDenom)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
