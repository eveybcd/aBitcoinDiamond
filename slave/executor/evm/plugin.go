// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package evm

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/aBitcoinDiamond/slave/executor/evm/commands"
	"github.com/aBitcoinDiamond/slave/executor/evm/executor"
	"github.com/aBitcoinDiamond/slave/executor/evm/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ExecutorName,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.EvmCmd,
		RPC:      nil,
	})
}
