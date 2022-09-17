package util

import (
	"fmt"
	"testing"
)

func TestIntToBinaryArray(t *testing.T) {
	defer catchPanic(t)
	is := []int{1, 2, 3, 4, 10, 11, 12, 13, 14, 15, 16, 0, -1, 67}
	for _, i := range is {
		bs := IntTo8BitArray(i)
		fmt.Println(i, " binary: ", bs)
	}
	fmt.Println("SUCCESS")

}

func TestIntTo8BitArray(t *testing.T) {
	defer catchPanic(t)
	is := []int{1, 2, 3, 4, 10, 11, 12, 13, 14, 15, 16, 0, -1, 67, 255, 257}
	for _, i := range is {
		bs := IntTo8BitArray(i)
		fmt.Println(i, " binary: ", bs)
	}
	fmt.Println("SUCCESS")

}

func TestBitsToByte(t *testing.T) {
	defer catchPanic(t)
	bitArray1 := []byte{0, 0, 0, 0, 0, 0, 0, 1}
	bitArray10 := IntTo8BitArray(10)
	bitArray255 := IntTo8BitArray(255)
	fmt.Println(bitArray1, "from 1 to :", Bits8ToByte(bitArray1))
	fmt.Println(bitArray10, "from 10 to :", Bits8ToByte(bitArray10))
	fmt.Println(bitArray255, "from 255 to :", Bits8ToByte(bitArray255))
}

func TestTemp1(t *testing.T) {
	fmt.Println(byte(1))
}

func catchPanic(t *testing.T) {
	if err_ := recover(); err_ != nil {
		//err = errors.New(err_.(string))
		//fmt.Println("err:",reflect.TypeOf(err_))
		//if reflect.TypeOf(err).String() == "*errors.errorString" {
		//	// do something
		//}
		fmt.Print(err_)
		t.Fail()
	}
}
