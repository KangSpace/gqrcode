package output

import (
	"bytes"
	"errors"
	"github.com/gqrcode/core/model"
	"github.com/gqrcode/util"
	"github.com/gqrcode/util/imaging"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Define PNG output Here

type ImageOutput struct {
	image *image.NRGBA
	*BaseOutput
}

func NewPNGOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}

// NewPNGOutput0 :Output a new PNG image by auto size.
func NewPNGOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}

func NewJPGOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}

// NewJPGOutput0 :Output a new JPG image by auto size.
func NewJPGOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}
func NewGIFOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}

// NewGIFOutput0 :Output a new GIF image by auto size.
func NewGIFOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}

func (out *ImageOutput) initImage(size int){
	// Point range : (0,0),(size-1,size-1)
	out.image = image.NewNRGBA(image.Rect(0,0,size,size))
	out.modules = make([][]*bool,size+1)
	for i := range out.modules{
		out.modules[i] = make([]*bool, size + 1)
	}
}

// Init :init for output when size is AUTO_SIZE
func (out *ImageOutput) Init(version *model.Version,qz *model.QuietZone){
	if out.Size == AUTO_SIZE{
		out.Size = version.GetDefaultPixelSize() + qz.GetDefaultPixelSize()
		out.initImage(out.Size)
	}
}

// Write : write data
func (out *ImageOutput) Write(x int,y int, black bool) {
	setColor := image.White
	if black {
		setColor = image.Black
	}
	out.image.Set(x,y,setColor)
	out.modules[x][y] = &black
}

// WriteModule : write data
func (out *ImageOutput) WriteModule(x int,y int, black bool,pixelSize int){
	setColor := image.White
	if black {
		setColor = image.Black
	}
	out.WriteModuleColor(x,y,black,setColor,pixelSize)
}

func (out *ImageOutput) WriteModuleColor(x int,y int,dark bool,setColor color.Color,pixelSize int){
	out.modules[x][y] = &dark
	x = x * pixelSize
	y = y * pixelSize
	for i:=0; i<pixelSize; i++{
		for j:=0; j<pixelSize; j++ {
			out.image.Set(x+i, y+j, setColor)
		}
	}
}

func (out *ImageOutput) IsModuleSet(x int,y int) bool{
	return out.BaseOutput.IsModuleSet(x,y)
}
func (out *ImageOutput) GetModule(x int,y int) bool{
	return out.BaseOutput.GetModule(x,y)
}

// Save : save file
func (out *ImageOutput) Save(fileName string) error{
	if file,err:= os.Create(fileName);err == nil{
		defer file.Close()
		switch out.BaseOutput.Type {
		case JPG:
			return jpeg.Encode(file,out.image,nil)
		case PNG:
			return png.Encode(file,out.image)
		case GIF:
			return gif.Encode(file,out.image,nil)
		}
		return errors.New("not supported \""+ string(out.BaseOutput.Type)+"\"")
		
	}else{
		return err
	}
}

func (out *ImageOutput) SaveToBase64() (base64Str string, err error){
	imageBytes := bytes.NewBuffer(nil)
	var base64UrlImageType util.Base64URLImageType
	switch out.BaseOutput.Type {
	case JPG:
		err = jpeg.Encode(imageBytes,out.image,nil)
		base64UrlImageType = util.JpegType
	case PNG:
		err = png.Encode(imageBytes,out.image)
		base64UrlImageType = util.PngType
	case GIF:
		err = gif.Encode(imageBytes,out.image,nil)
		base64UrlImageType = util.GifType
	default:
		return "",errors.New("not supported \""+ string(out.BaseOutput.Type)+"\"")
	}
	if err != nil {
		return "",err
	}
	return util.ImageToBase64Url(base64UrlImageType,imageBytes.Bytes()),nil
}

func (out *ImageOutput) GetBaseOutput() *BaseOutput{
	return out.BaseOutput

}

func (out *ImageOutput) GetImage() *image.NRGBA{
	return out.image
}

func (out *ImageOutput) drawNewImage(minPoint image.Point,maxPoint image.Point,imageSize int){
	newImg := image.NewNRGBA(image.Rect(0,0,imageSize,imageSize))
	draw.Draw(newImg,newImg.Bounds(),image.White, image.Pt(0,0), draw.Src)
	r:= image.Rectangle{Min: minPoint, Max: maxPoint}
	draw.Draw(newImg,r,out.image,image.Pt(0,0), draw.Over)
	//logo image handle
	newImg = out.drawLogoImage(newImg)
	out.image = newImg
}

func (out *ImageOutput) drawLogoImage(srcImage *image.NRGBA) *image.NRGBA{
	logoOption := out.BaseOutput.containLogoOption()
	if logoOption != nil{
		logoFilePath := logoOption.Value
		var logoFile *os.File;var err error
		if logoFile,err = os.Open(logoFilePath);err != nil {
			panic(err)
		}
		defer logoFile.Close()
		var logoImg image.Image
		if logoImg,err =png.Decode(logoFile);err != nil {
			if logoImg,err =jpeg.Decode(logoFile);err != nil {
				if logoImg,err =gif.Decode(logoFile);err != nil {
					panic(err)
				}
			}
		}
		width := logoImg.Bounds().Dx()
		height := logoImg.Bounds().Dy()
		logoSize := width * height
		remainSize := int(0.1 * float32(out.Size * out.Size))
		// remain 10%
		if  remainSize < logoSize {
			remainRate := float32(remainSize)/float32(logoSize)
			width = int( float32(width) * remainRate)
			height = int( float32(height) * remainRate)
		}
		x := (out.Size - width) / 2 - 1
		y := (out.Size - height) /2 - 1
		logoImg = imaging.Resize0(logoImg,width,height)
		r:= image.Rectangle{Min: image.Pt(x,y), Max: image.Pt(x+width,y+height)}
		draw.Draw(srcImage,r,logoImg,image.Pt(0,0), draw.Over)
	}
	return srcImage
}

func (out *ImageOutput) drawIntoNewImage(minPoint image.Point,maxPoint image.Point){
	out.drawNewImage(minPoint,maxPoint,out.Size)
}

func (out *ImageOutput)	ResizeToFit(moduleSize int, quietZoneSize int, pixelSize int) {
	modulePixels := moduleSize * pixelSize
	quietZonePixels := quietZoneSize * pixelSize
	if out.Size == modulePixels{
		return
	}else if out.Size == modulePixels + quietZonePixels{
		imgX:= quietZonePixels/2 ; imgY := quietZonePixels/2
		out.drawIntoNewImage(image.Pt(imgX,imgY),image.Point{X: imgX + moduleSize * pixelSize, Y: imgY + moduleSize * pixelSize})
		return
	}else {
		maxImageSize := modulePixels+quietZonePixels
		imgX := 0; imgY := 0
		if quietZoneSize > 0{
			imgX = quietZonePixels/2 ; imgY = quietZonePixels/2
		}
		out.drawNewImage(image.Pt(imgX,imgY),image.Point{X: imgX + moduleSize * pixelSize , Y: imgY + moduleSize * pixelSize},maxImageSize)
		out.image = imaging.Resize(out.image,out.Size)
	}
	return
}

// Clone : Shallow copy BaseOutput and modules from output, init new image instance
func (out *ImageOutput) Clone() Output{
	 clone := &ImageOutput{BaseOutput:&BaseOutput{Type: out.Type,Size: out.Size,Options: out.Options}}
	 clone.initImage(out.Size)
	return clone
}