package collection

import (
	nft "github.com/UptickNetwork/uptick/x/collection/module"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
