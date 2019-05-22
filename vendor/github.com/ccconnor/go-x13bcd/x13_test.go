package x13

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestX13(t *testing.T) {
	blockHeader, _ := hex.DecodeString("00000060d7954b565e2ede08fadd961e1bc97e0dd4b4f2ab36f8fb280a673862049dd97e42aae4081ff493c65c7b5b8d190a9b2b5332c3a0feaf2bc9a02fa7b0559c35a7316de35c607a431a058c4114")
	x13hash := HashX13sm3(blockHeader)
	hash := hex.EncodeToString(x13hash[:])
	fmt.Println(hash)
	if hash != "c5946f83ed0c9080262837628cd39f49ab87eb05c9e2d9e6080b000000000000" {
		t.Errorf("test failed!")
	}
}
