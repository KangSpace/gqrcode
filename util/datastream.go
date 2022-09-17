package util

// Bit : use byte to represent bit. value is 0/1
type Bit = byte

// Module : use byte to represent symbol module. value is 0/1
type Module = byte

// Define Data Stream here

type DataStream struct {
	count int
	data  []Bit
}

func NewDataStream(capacity int) *DataStream {
	data := make([]byte, capacity)
	ds := new(DataStream)
	ds.data = data
	ds.count = 0
	return ds
}

func (ds *DataStream) GetCount() int {
	return ds.count
}

func (ds *DataStream) GetData() []Bit {
	return ds.data[:ds.count]

}
func (ds *DataStream) AddBit(b []Bit, count int) {
	bLen := len(b)
	if bLen == 0 && count == 0 {
		return
	}
	ds.autoGrow(count)
	if bLen < count {
		ds.pushInData(GetZeroBytes(count - bLen))
	}
	ds.pushInData(b)
}

// AddIntBit :Add int to []byte into DataStream
func (ds *DataStream) AddIntBit(i int, count int) {
	ds.AddIntBit16(uint16(i), count)
}

// AddIntBit16 :Add int to []byte into DataStream
func (ds *DataStream) AddIntBit16(i uint16, count int) {
	if i == 0 && count == 0 {
		return
	}
	ds.autoGrow(count)
	iBits := IntToBinary(i)
	if iBits == nil {
		return
	}
	iBitsLen := len(iBits)
	if iBitsLen < count {
		ds.pushInData(GetZeroBytes(count - iBitsLen))
	}
	ds.pushInData(iBits)
}

// GetZeroBytes :Get a zero byte array by zeroCount
func GetZeroBytes(zeroCount int) []byte {
	b := make([]byte, zeroCount)
	return b
}

func (ds *DataStream) Clean() {
	ds.data = nil
	ds.count = 0
}

// IteratorByte :Iterate the data bits to Byte(8-bit).
// return byte chan.
func (ds *DataStream) IteratorByte() <-chan byte {
	byteOutChan := make(chan byte)
	bitLen := 8
	go func() {
		bitsCount := ds.count
		for i := 0; bitsCount > 0 && i < bitsCount; i += bitLen {
			max := i + bitLen
			if max > bitsCount {
				max = bitsCount
				bitLen = max - i
			}

			byteOutChan <- BitsToByte(ds.data[i:max], bitLen)
		}
		close(byteOutChan)
	}()
	return byteOutChan
}

func (ds *DataStream) pushInData(bs []byte) {
	for _, b := range bs {
		ds.data[ds.count] = b
		ds.count++
	}
}

func (ds *DataStream) autoGrow(count int) {
	currentDataLength := len(ds.data)
	if currentDataLength < ds.count+count {
		ds.grow(ds.count + count - currentDataLength)
	}
}

// grow up the data capacity of DataStream
func (ds *DataStream) grow(growBy int) {
	currentLength := len(ds.data)
	if growBy == 0 {
		growBy = currentLength
		if growBy < 128 {
			growBy = 128
		} else if growBy >= 1024 {
			growBy = 1024
		}
	} else {
		if growBy/2 > 0 {
			growBy = growBy + 1
		}
	}
	nd := make([]byte, currentLength+growBy)
	copy(nd, ds.data)
	ds.data = nd
}
