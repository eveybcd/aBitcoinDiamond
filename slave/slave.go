package slave

import (
	_ "github.com/33cn/chain33/system"
	"github.com/33cn/chain33/util/cli"
	_ "github.com/aBitcoinDiamond/slave/consensus/init"
	_ "github.com/aBitcoinDiamond/slave/dapp/evm"
	_ "github.com/aBitcoinDiamond/slave/dapp/ticket"
)

func StartSlave() {

	cli.RunChain33("slaveChain")
}
