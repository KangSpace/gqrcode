package output

import (
	"errors"
	"github.com/gqrcode/core/model"
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
	image *image.RGBA
	*BaseOutput
}

func NewPNGOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}
func NewPNGOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}

func NewJPGOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}
func NewJPGOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}
func NewGIFOutput(size int) *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: size}}
	out.initImage(size)
	return out
}
func NewGIFOutput0() *ImageOutput{
	out:= &ImageOutput{BaseOutput: &BaseOutput{Type: PNG, Size: AUTO_SIZE}}
	return out
}

func (out *ImageOutput) initImage(size int){
	// Point range : (0,0),(size-1,size-1)
	out.image = image.NewRGBA(image.Rect(0,0,size,size))
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
	out.WriteModuleColor(x,y,setColor,pixelSize)
}

func (out *ImageOutput) WriteModuleColor(x int,y int, setColor color.Color,pixelSize int){
	defTrue := true
	out.modules[x][y] = &defTrue
	x = x * pixelSize
	y = y * pixelSize
	for i:=0; i<pixelSize; i++{
		for j:=0; j<pixelSize; j++ {
			out.image.Set(x+i, y+j, setColor)
		}
	}
}

func (out *ImageOutput) GetModule(x int,y int) bool{
	return out.BaseOutput.GetModule(x,y)
}

// Save : save file
func (out *ImageOutput) Save(fileName string) error{
	if file,err:= os.Create(fileName);err == nil{
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

func (out *ImageOutput) GetBaseOutput() *BaseOutput{
	return out.BaseOutput

}

func (out *ImageOutput) GetImage() *image.RGBA{
	return out.image
}

func (out *ImageOutput) DrawIntoNewImage(minPoint image.Point,maxPoint image.Point){
	imageSize := out.Size
	newImg := image.NewRGBA(image.Rect(0,0,imageSize,imageSize))
	draw.Draw(newImg,newImg.Bounds(),image.White, image.Pt(0,0), draw.Src)
	r:= image.Rectangle{Min: minPoint, Max: maxPoint}
	draw.Draw(newImg,r,out.image,image.Pt(0,0), draw.Over)
	out.image = newImg
}

func (out *ImageOutput) Resize(){

}

// Clone : Shallow copy BaseOutput and modules from output, init new image instance
func (out *ImageOutput) Clone() Output{
	 clone := &ImageOutput{BaseOutput:&BaseOutput{Type: out.Type,Size: out.Size}}
	 clone.initImage(out.Size)
	return clone
}