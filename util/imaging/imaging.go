package imaging

import (
	"github.com/disintegration/imaging"
	"image"
)

// Using imaging by:
// go get -u github.com/disintegration/imaging

func Resize(src image.Image,newSize int) *image.NRGBA{
	return imaging.Resize(src, newSize, newSize, imaging.CatmullRom)
}