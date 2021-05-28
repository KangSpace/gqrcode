package util

import (
	"errors"
	"sort"
)


// StrIn :Check target string is or not in strArray
func StrIn(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// IntIn :Check target int is or not in intArray
func IntIn(target int, intArray []int) bool {
	sort.Ints(intArray)
	index := sort.SearchInts(intArray, target)
	if index < len(intArray) && intArray[index] == target {
		return true
	}
	return false
}

// IntToBinary : int change to binary array in []byte , array not full 8-bit
func IntToBinary(i int) []Bit{
	if i < 0 {
		return nil
	}

	bLen := 0
	for c:=i; c!=0;{
		c = c>>1
		bLen++
	}
	bytes := make([]Bit, bLen)
	for ; i > 0; i, bLen = i / 2 , bLen-1 {
		lsb := i % 2
		bytes[bLen-1] = byte(lsb)
	}
	return bytes
}

// IntTo8BitArray : int change to binary array in []byte, and byte is 8-bit
func IntTo8BitArray(i int) []Bit{
	if i < 0 {
		return nil
	}else if i == 0{
		return []Bit{0,0,0,0,0,0,0,0}
	}

	bLen := 0
	for c:=i; c!=0;{
		c = c>>1
		bLen++
	}
	raminBitLen := bLen%8
	if raminBitLen >0 {
		bLen += 8 - raminBitLen
	}
	bytes := make([]byte, bLen)
	for ; i > 0; i, bLen = i / 2 , bLen-1 {
		lsb := i % 2
		bytes[bLen-1] = byte(lsb)
	}
	return bytes
}

// Bits8ToByte : convert 8 bits to 1 byte
func Bits8ToByte(bitArray []Bit) byte{
	if len(bitArray) != 8 {
		panic(errors.New("input bit length is not 8"))
	}
	var bitsByte byte = 0
	for i:=0; i <8 ; i++ {
		bitsByte += bitArray[i]<<(7-i)
	}
	return bitsByte
}

func ByteArrayToIntArray(data []byte) (dataInts []int){
	dataInts = make([]int,len(data))
	for i := 0; i < len(data); i++ {
		dataInts[i] = int(data[i])
	}
	return dataInts
}

// IntArrayToByteArray :convert int array to byte array
func IntArrayToByteArray(data []int) (dataBytes []byte){
	dataBytes = make([]byte,len(data))
	for i := 0; i < len(data); i++ {
		dataBytes[i] = byte(data[i])
	}
	return dataBytes
}

// ByteArrayTo8BitArray :convert byte array to 8 bit array(1 byte split to 8 bit save in Bit array)
// e.g. :
// input : {byte(1)} result: {0,0,0,0,0,0,0,1}
func ByteArrayTo8BitArray(data []byte) (dataBits []Bit){
	dataBitsLen := len(data) * 8
	return ByteArrayTo8BitArrayWithCount(data,dataBitsLen)
}

// ByteArrayTo8BitArrayWithCount :convert byte array to 8 bit array(1 byte split to 8 bit save in Bit array)
// e.g. :
// input : {byte(1)} result: {0,0,0,0,0,0,0,1}
//
// param data: input byte array
// param bitsCount: total bit count for return
func ByteArrayTo8BitArrayWithCount(data []byte,bitsCount int) (dataBits []Bit){
	dataLen := len(data)
	if dataLen == 0 || bitsCount == 0{
		return nil
	}
	dataBits = make([]Bit,0,bitsCount)
	for i := 0; i < dataLen; i++ {
		arr:= IntTo8BitArray(int(data[i]))
		dataBits = append(dataBits,arr...)
	}
	return dataBits
}