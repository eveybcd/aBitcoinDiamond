// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashlock

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/aBitcoinDiamond/slave/executor/hashlock/commands"
	"github.com/aBitcoinDiamond/slave/executor/hashlock/executor"
	"github.com/aBitcoinDiamond/slave/executor/hashlock/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.HashlockX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.HashlockCmd,
		RPC:      nil,
	})
}
