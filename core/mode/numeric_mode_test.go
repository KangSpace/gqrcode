package mode

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"reflect"
	"strconv"
	"testing"
)
var home = os.Getenv("HOME")
var desktop = home+"/Desktop/gqrcode/"
func TestImageToImg(t *testing.T) {
	img1 := image.NewRGBA(image.Rect(0,0,10,10))
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			img1.Set(i,j,image.Black)
		}
	}
	file1,_:=os.Create(desktop+"img.png")
	png.Encode(file1,img1)

	img2 := image.NewRGBA(image.Rect(0,0,100,100))
	draw.Draw(img2, img2.Bounds(), &image.Uniform{C: image.White}, image.ZP, draw.Src)

	//(0,0)-(10,10) blue
	draw.Draw(img2,
		image.Rectangle{image.ZP,image.Point{X: 10, Y: 10}},
		&image.Uniform{C:  color.RGBA{0,0,255,255}}, image.ZP, draw.Src)

	//(10,10)-(20,20) black
	r:= image.Rectangle{Min: image.Point{X: 10, Y: 10}, Max: image.Point{X: 20, Y: 20}}
	draw.Draw(img2,r,img1,image.Pt(0,0), draw.Src)

	file2,_:=os.Create(desktop+"img2.png")
	png.Encode(file2,img2)
}

// See qrcode_test.go
func TestNumericNewQRCode(t *testing.T) {
}

func TestBoolDefaultVal(t *testing.T){
	var def bool
	init:=false
	fmt.Println("default:"+ strconv.FormatBool( reflect.ValueOf(def).IsValid()))
	fmt.Printf("default:%v\n", reflect.ValueOf(def).Kind())
	fmt.Printf("default:%v\n", reflect.ValueOf(def).IsZero())
	fmt.Println("init:"+ strconv.FormatBool(reflect.ValueOf(init).IsValid()))
	fmt.Printf("init:%v\n", reflect.ValueOf(init).Kind())
	fmt.Printf("init:%v\n", reflect.ValueOf(init).IsZero())
}