// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

import (
	"fmt"

	"github.com/33cn/chain33/types"
	"github.com/aBitcoinDiamond/slave/executor/evm/executor/vm/common"
	evmtypes "github.com/aBitcoinDiamond/slave/executor/evm/types"
)

// BlockData 本文件用来存储硬分叉中需要用到的数据
type BlockData struct {
	blockHeight int64
	testnet     bool
	// 存储多个交易信息
	txs map[string]*TxData
}

// TxData 交易数据
type TxData struct {
	// KV 交易生成的KV数据
	KV map[string][]byte
	// Logs Key 为logType_logIndex
	Logs map[string][]byte
}

var forkData map[int64]*BlockData

func newBlockData(blockHeight int64, testnet bool) *BlockData {
	data := BlockData{blockHeight: blockHeight, testnet: testnet}
	data.txs = make(map[string]*TxData)
	return &data
}

func (block *BlockData) newTxData(txHash string) *TxData {
	data := TxData{}
	data.KV = make(map[string][]byte)
	data.Logs = make(map[string][]byte)
	block.txs[txHash] = &data
	return &data
}

func makeLogReceiptKey(logType int32, logIndex int) string {
	return fmt.Sprintf("%v_%v", logType, logIndex)
}

// InitForkData 初始化硬分叉数据
func InitForkData() {
	forkData = make(map[int64]*BlockData)

	// 556151 高度合约调用引起的状态变更
	data := newBlockData(556151, true)
	txData := data.newTxData("0xc79c9113a71c0a4244e20f0780e7c13552f40ee30b05998a38edb08fe617aaa5")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x1a2039c5e33bc3f4dca79c611d979a87c24aebc7e7f056cdc8845dae0a2c418fd35a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 1)] = common.FromHex("0x1a2039c5e33bc3f4dca79c611d979a87c24aebc7e7f056cdc8845dae0a2c418fd35a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb")
	forkData[556151] = data

	// 556255 高度合约调用引起的状态变更
	data = newBlockData(556255, true)
	txData = data.newTxData("0xa2e3a06322ce7561ec3c7e442dbc0a6b12618f661c80d16a1af0ffda3e8c2dd8")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x1a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 3)] = common.FromHex("0x1a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c80022660a4230783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303112200000000000000000000000000000000000000000000000000000000000002710")
	forkData[556255] = data

	// 556294 高度合约调用引起的状态变更
	data = newBlockData(556294, true)
	txData = data.newTxData("0x206c26b2a00751b13df4f44924f46bd353d4d2b48595687a29c9c9c1c34f6d3f")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x10011a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 1)] = common.FromHex("0x10011a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	forkData[556294] = data
}

// ProcessFork 处理硬分叉逻辑
func ProcessFork(blockHeight int64, txHash []byte, receipt *types.Receipt) {
	if types.IsLocal() {
		return
	}

	// 目前的分叉信息只在测试网中存在
	if !types.IsTestNet() {
		return
	}

	// 首先，查找是否存在分叉所在区块
	if block, ok := forkData[blockHeight]; ok {
		strHash := common.Bytes2Hex(txHash)

		// 然后查找，此区块中需要分叉处理的交易
		if tx, ok := block.txs[strHash]; ok {

			// 替换需要处理的KeyValue中的Value值（保持原顺序不变）
			for i, v := range receipt.KV {
				if val, ok := tx.KV[common.Bytes2Hex(v.Key)]; ok {
					receipt.KV[i].Value = val
				}
			}

			// 替换需要处理的收据信息（根据交易类型和索引号），只替换取值，其它不变
			for i, v := range receipt.Logs {
				if val, ok := tx.Logs[makeLogReceiptKey(v.Ty, i)]; ok {
					receipt.Logs[i].Log = val
				}
			}
		}
	}
}
