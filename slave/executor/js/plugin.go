package js

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/aBitcoinDiamond/slave/executor/js/executor"
	ptypes "github.com/aBitcoinDiamond/slave/executor/js/types"

	// init auto test
	_ "github.com/aBitcoinDiamond/slave/executor/js/autotest"
	"github.com/aBitcoinDiamond/slave/executor/js/command"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     ptypes.JsX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      command.JavaScriptCmd,
		RPC:      nil,
	})
}
