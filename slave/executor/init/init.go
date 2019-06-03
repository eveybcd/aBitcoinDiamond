package init

import (
	_ "github.com/aBitcoinDiamond/slave/executor/ticket" //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/token"  //auto gen

	_ "github.com/aBitcoinDiamond/slave/executor/evm"       //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/hashlock"  //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/js"        //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/paracross" //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/retrieve"  //auto gen
	_ "github.com/aBitcoinDiamond/slave/executor/trade"     //auto gen
)
