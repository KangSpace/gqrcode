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

// Resize0 :
func Resize0(src image.Image, width int, height int) *image.NRGBA {
	return imaging.Resize(src, width, height, imaging.NearestNeighbor)
}

// Grayscale produces a grayscale version of the image.
func Grayscale(src image.Image) *image.NRGBA {
	return imaging.Grayscale(src)
}

// Binarization  : To binary image from grayscale image,
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
