package evm

import (
	evm "github.com/evmos/ethermint/x/evm/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

// MsgEthereumTx encapsulates an Ethereum transaction as an SDK message.
type DocMsgEthereumTx struct {
	Data  string  `bson:"data"`
	Size_ float64 `bson:"size"`
	Hash  string  `bson:"hash"`
	From  string  `bson:"from"`
}

func (doctx *DocMsgEthereumTx) GetType() string {
	return MsgTypeEthereumTx
}

func (doctx *DocMsgEthereumTx) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgEthereumTx)
	doctx.Size_ = msg.Size_
	doctx.Hash = msg.Hash
	doctx.From = msg.From
	if txData, err := evm.UnpackTxData(msg.Data); err == nil {
		doctx.Data = utils.MarshalJsonIgnoreErr(txData)
	}
}

func (m *DocMsgEthereumTx) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
	)
	msg, ok := v.(*MsgEthereumTx)
	if ok {
		addrs = append(addrs, msg.From)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
