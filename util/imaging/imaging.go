package imaging

import (
	"github.com/disintegration/imaging"
	"image"
)

// Using imaging by:
// go get -u github.com/disintegration/imaging

func Resize(src image.Image,newSize int) *image.NRGBA{
	return Resize0(src, newSize, newSize)
}
// Resize0 :TODO Need optimize...
func Resize0(src image.Image,width int,height int) *image.NRGBA{
	return imaging.Resize(src, width, height,imaging.CatmullRom)
}