package collection

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type UptickNftClient struct {
}

func NewClient() UptickNftClient {
	return UptickNftClient{}
}

func (nft UptickNftClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {

	switch msg := v.(type) {
	case *MsgUptickNFTMint:
		docMsg := DocMsgUptickNFTMint{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUptickNFTEdit:
		docMsg := DocMsgUptickNFTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUptickNFTTransfer:
		docMsg := DocMsgUptickNFTTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUptickNFTBurn:
		docMsg := DocMsgUptickNFTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUptickIssueDenom:
		docMsg := DocMsgUptickIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUptickTransferDenom:
		docMsg := DocMsgUptickTransferDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
