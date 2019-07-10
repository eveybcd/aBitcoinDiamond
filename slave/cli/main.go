// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package main

import (
	_ "github.com/33cn/chain33/system"
	"github.com/aBitcoinDiamond/slave/cli/buildflags"
	//_ "github.com/aBitcoinDiamond/slave"

	"github.com/33cn/chain33/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8601"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName)

}
