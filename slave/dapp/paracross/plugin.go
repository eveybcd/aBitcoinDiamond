// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paracross

import (
	"github.com/33cn/chain33/pluginmgr"
	_ "github.com/aBitcoinDiamond/slave/dapp/paracross/autotest" // register autotest package
	"github.com/aBitcoinDiamond/slave/dapp/paracross/commands"
	"github.com/aBitcoinDiamond/slave/dapp/paracross/executor"
	"github.com/aBitcoinDiamond/slave/dapp/paracross/rpc"
	"github.com/aBitcoinDiamond/slave/dapp/paracross/types"
	_ "github.com/aBitcoinDiamond/slave/dapp/paracross/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ParaX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.ParcCmd,
		RPC:      rpc.Init,
	})
}
