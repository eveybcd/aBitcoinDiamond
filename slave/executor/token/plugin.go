// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token 创建token
package token

import (
	"github.com/33cn/chain33/pluginmgr"
	_ "github.com/aBitcoinDiamond/slave/executor/token/autotest" // register token autotest package
	"github.com/aBitcoinDiamond/slave/executor/token/commands"
	"github.com/aBitcoinDiamond/slave/executor/token/executor"
	"github.com/aBitcoinDiamond/slave/executor/token/rpc"
	"github.com/aBitcoinDiamond/slave/executor/token/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TokenX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TokenCmd,
		RPC:      rpc.Init,
	})
}
