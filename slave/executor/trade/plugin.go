// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trade

import (
	"github.com/33cn/chain33/pluginmgr"
	_ "github.com/aBitcoinDiamond/slave/executor/trade/autotest" // register autotest package
	"github.com/aBitcoinDiamond/slave/executor/trade/commands"
	"github.com/aBitcoinDiamond/slave/executor/trade/executor"
	"github.com/aBitcoinDiamond/slave/executor/trade/rpc"
	"github.com/aBitcoinDiamond/slave/executor/trade/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TradeX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TradeCmd,
		RPC:      rpc.Init,
	})
}
