// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ticket

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/aBitcoinDiamond/slave/executor/ticket/commands"
	"github.com/aBitcoinDiamond/slave/executor/ticket/executor"
	"github.com/aBitcoinDiamond/slave/executor/ticket/rpc"
	"github.com/aBitcoinDiamond/slave/executor/ticket/types"
	// init wallet
	_ "github.com/aBitcoinDiamond/slave/executor/ticket/wallet"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TicketX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TicketCmd,
		RPC:      rpc.Init,
	})
}
