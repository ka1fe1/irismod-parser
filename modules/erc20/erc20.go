package erc20

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type Erc20Client struct {
}

func NewClient() Erc20Client {
	return Erc20Client{}
}

func (erc20 Erc20Client) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgConvertCoin:
		docMsg := DocMsgConvertCoin{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgConvertERC20:
		docMsg := DocMsgConvertERC20{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}
}
