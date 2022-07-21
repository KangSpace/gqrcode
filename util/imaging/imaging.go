package imaging

import (
	"github.com/KangSpace/gqrcode/util"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

// Using imaging by:
// go get -u github.com/disintegration/imaging

func Resize(src image.Image, newSize int) *image.NRGBA {
	return Resize0(src, newSize, newSize)
}

// Resize0 :TODO Need optimize...
func Resize0(src image.Image, width int, height int) *image.NRGBA {
	return imaging.Resize(src, width, height, imaging.CatmullRom)
}

// Grayscale produces a grayscale version of the image.
func Grayscale(src image.Image) *image.NRGBA {
	return imaging.Grayscale(src)
}

// reference: Wellner 1993, Derek Bradley and Gerhard Roth : Adaptive Thresholding Using the Integral Image
// https://blog.csdn.net/hhygcy/article/details/4280165
// https://www.cnblogs.com/Imageshop/archive/2013/04/22/3036127.html
// http://www.derekbradley.ca/AdaptiveThresholding/index.html
// http://www.scs.carleton.ca/~roth/iit-publications-iti/docs/gerh-50002.pdf
func wellnerAdaptiveThreshold(src image.Image) int {
	//width := src.Bounds().Dx()
	//height := src.Bounds().Dy()
	//s:= width >> 3
	//// static
	//t:= 15
	//i, j:=0
	//sum,count:=0,0
	//index:= 0;
	//x1, y1, x2, y2:= 0,0,0,0
	//s2 := s/2
	return 0
}

// Binarization  To binary image from grayscale image,
func Binarization(src image.Image) *image.NRGBA {
	dst := image.NewNRGBA(src.Bounds())
	srcW := src.Bounds().Dx()
	srcH := src.Bounds().Dy()
	//histogram := imaging.Histogram(src)
	//threshold := wellnerAdaptiveThreshold()
	util.Parallel(0, srcH, func(ys <-chan int) {
		for y := range ys {
			//i := y * dst.Stride
			for x := 0; x < srcW; x++ {
				r, g, b, _ := src.At(x, y).RGBA()
				rgb := uint8(((r >> 8) + (g >> 8) + (b >> 8)) / 3)
				//uint8(threshold)
				if rgb > 40 {
					rgb = 255
				} else {
					rgb = 0
				}
				dst.SetNRGBA(x, y, color.NRGBA{R: rgb, G: rgb, B: rgb, A: 255})
			}
		}
	})
	return dst
}

// ImageBinarization  To binary image by two step,
// 1. Grayscale image
// 2. Binarization image
func ImageBinarization(src image.Image) *image.NRGBA {
	grayImage := Grayscale(src)
	return Binarization(grayImage)
}
