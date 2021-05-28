package main

import (
	"fmt"
	"github.com/gqrcode/util"
)

func main() {
	b :=util.GetZeroBytes(3)
	fmt.Println(b)
	fmt.Println(byte(1))
	arr := make([]byte,10)
	arr = nil
	fmt.Print("len arr:",len(arr))
}
