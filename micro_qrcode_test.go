package main

import (
	"fmt"
	"github.com/gqrcode/core/output"
	"strconv"
	"testing"
)

// -----------------------------TestMicroQRCode----------------------------------
// TestNewNumeric1MicroQRCode : test 1 single numeric MicroQRCode
func TestNewNumeric1MicroQRCode(t *testing.T) {
	data := "0123456789012345"
	fmt.Println(len(data))
	//data := "8675309"
	fileNamePrefix := "numeric_micro_qrcode"
	fileName := gqrcodePath + fileNamePrefix+".png"
	out := output.NewPNGOutput(60*4)
	qrcode, err := NewMicroQRCode(data)
	if err != nil{
		t.Fatal(err)
	}
	err = qrcode.Encode(out,fileName)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("version:%v,moduleSize:%d, size:%d, ec:%v mode:%s \n",qrcode.Version,qrcode.Version.GetModuleSize(), out.Size, qrcode.ErrorCorrection,qrcode.Mode.GetMode())
	fmt.Println("SUCCESS,"+fileName)
}

// TestNewNumericAllVersionMicroQRCode :Test numeric all version for Micro QRCode
func TestNewNumericAllVersionMicroQRCode(t *testing.T) {
	// Numeric
	dataArr := []string{
		// M1,5
		"01234",
		// M2,L:10,M:8
		"01234567",
		"0123456789",
		// M3,L:23,M:18
		"01234567890123456789012",
		"012345678901235467",
		// M4,L:35,M:30,Q:21
		"01234567890123456789012345678901234",
		"012345678901234567890123456789",
		"012345678901234567890",
	}
	for _,data:= range dataArr {
		dLen :=  len(data)
		fileNamePrefix := "numeric_micro_qrcode"+ strconv.Itoa(dLen)
		fileName := gqrcodePath + fileNamePrefix + ".png"
		out := output.NewPNGOutput(60 * 4)
		qrcode, err := NewMicroQRCode(data)
		if err != nil {
			t.Fatal(err)
		}
		err = qrcode.Encode(out, fileName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("version:%v,moduleSize:%d, size:%d, ec:%v mode:%s \n", qrcode.Version, qrcode.Version.GetModuleSize(), out.Size, qrcode.ErrorCorrection, qrcode.Mode.GetMode())
		fmt.Println("SUCCESS," + fileName)
	}
}

func TestNewAlphaNumericAllVersionMicroQRCode(t *testing.T) {
	// Alphanumeric
	dataArr := []string{
		// M2,L:6,M:5
		"012AB",
		"012AB+",
		// M3,L:14,M:11
		"012AB+-*C30",
		"012AB+-*C3012AB",
		// M4,L:21,M:18,Q:13
		"012AB+-*C3012AB+-*C30",
		"012AB+-*C3012AB+-*",
		"012AB+-*C3012",
	}
	for _,data:= range dataArr {
		dLen :=  len(data)
		fileNamePrefix := "alphanumeric_micro_qrcode"+ strconv.Itoa(dLen)
		fileName := gqrcodePath + fileNamePrefix + ".png"
		out := output.NewPNGOutput(60 * 4)
		qrcode, err := NewMicroQRCode(data)
		if err != nil {
			t.Fatal(err)
		}
		err = qrcode.Encode(out, fileName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("version:%v,moduleSize:%d, size:%d, ec:%v mode:%s \n", qrcode.Version, qrcode.Version.GetModuleSize(), out.Size, qrcode.ErrorCorrection, qrcode.Mode.GetMode())
		fmt.Println("SUCCESS," + fileName)
	}
}

func TestNewByteAllVersionMicroQRCode(t *testing.T) {
	// byte
	dataArr := []string{
		// M3,L:9,M:7
		"abc123ABC",
		"abc123A",
		// M4,L:15,M:13,Q:9
		"i.kangspace.org",
		"kangspace.org",
		"kangspace",
	}
	for _,data:= range dataArr {
		dLen :=  len(data)
		fileNamePrefix := "byte_micro_qrcode"+ strconv.Itoa(dLen)
		fileName := gqrcodePath + fileNamePrefix + ".png"
		out := output.NewPNGOutput(60 * 5)
		qrcode, err := NewMicroQRCodeAutoQuiet(data)
		if err != nil {
			t.Fatal(err)
		}
		err = qrcode.Encode(out, fileName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("version:%v,moduleSize:%d, size:%d, ec:%v mode:%s \n", qrcode.Version, qrcode.Version.GetModuleSize(), out.Size, qrcode.ErrorCorrection, qrcode.Mode.GetMode())
		fmt.Println("SUCCESS," + fileName)
	}
}
