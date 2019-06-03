package slave

import (
	_ "github.com/33cn/chain33/system"
	"github.com/33cn/chain33/util/cli"                  //cli init
	_ "github.com/aBitcoinDiamond/slave/consensus/init" //consensus init
	_ "github.com/aBitcoinDiamond/slave/dapp/init"      //dapp init
	_ "github.com/aBitcoinDiamond/slave/mempool/init"   //mempool init
)

func StartSlave() {

	cli.RunChain33("slaveChain")
}
