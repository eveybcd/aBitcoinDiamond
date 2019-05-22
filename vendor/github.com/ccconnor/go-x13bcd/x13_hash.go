package x13

/*
#include "x13.h"
*/
import "C"
import (
	"unsafe"
)

func HashX13sm3(b []byte) []byte {
	var blockHash [32]byte
	C.x13_hash((*C.char)(unsafe.Pointer(&b[0])), (*C.char)(unsafe.Pointer(&blockHash[0])))
	return blockHash[:]
}
