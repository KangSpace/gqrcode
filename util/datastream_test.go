package util

import (
	"fmt"
	"testing"
)

func TestNewDataStream(t *testing.T) {
	nds := NewDataStream(10)
	fmt.Println("new nds: ", &nds)
	bBits := []byte{1, 1, 1, 1}
	nds.AddBit(bBits, len(bBits))
	//0b101
	i := 5
	count := 5
	nds.AddIntBit(i, count)
	fmt.Println("final nds: ", &nds, " ", nds)
	fmt.Println("final nds data: ", nds.GetData())
}
