package erc20

import (
	"github.com/UptickNetwork/uptick/x/erc20"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		erc20.AppModuleBasic{},
	)
}
