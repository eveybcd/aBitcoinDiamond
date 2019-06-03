// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package retrieve

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/aBitcoinDiamond/slave/executor/retrieve/commands"
	"github.com/aBitcoinDiamond/slave/executor/retrieve/executor"
	"github.com/aBitcoinDiamond/slave/executor/retrieve/rpc"
	"github.com/aBitcoinDiamond/slave/executor/retrieve/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.RetrieveX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.RetrieveCmd,
		RPC:      rpc.Init,
	})
}
