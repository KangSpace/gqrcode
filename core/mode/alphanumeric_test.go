package mode

import (
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestChar(t *testing.T) {
	data := "A"
	fmt.Printf("data:%s,data[:1]:%s, data[0]:%s \n", data, data[:1], data[0:])
	data = "AB"
	fmt.Printf("data:%s,data[:1]:%s, data[0]:%s \n", data, data[:1], data[0:])
}

func TestByte(t *testing.T) {
	data := "A"
	fmt.Printf("data:%d \n", data[0])
	fmt.Println("data:" + string(data[0]))
}
func TestKanji(t *testing.T) {
	data := "0日月"
	if a, err := ToShiftJIS(data); err != nil {
		t.Fatal(err)
	} else {
		//aaa,err:= FromShiftJIS(a)
		aaa := strconv.Quote(a)
		for _, v := range strings.Split(aaa[1:len(aaa)-1], "\\x") {
			fmt.Printf("aaa:%v\n", v)
		}
	}

}

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

// Convert a string encoding from ShiftJIS to UTF-8
func FromShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}

// Convert a string encoding from UTF-8 to ShiftJIS
func ToShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())
}
