package output

import (
	"fmt"
	"github.com/gqrcode/core/model"
	"image"
	"image/color"
	"math"
)

// Define Output here

type Type = int

// defined OutputType constants
const (
	JPG Type= iota // 0
	PNG // 1
	GIF // 2
	SVG // 3
)


const (
	// AUTO_SIZE : set auto size for qrcode , per module fill by 4 pixels
	AUTO_SIZE = 0
)

type BaseOutput struct {
	// image type
	Type Type
	// image width/height
	Size int
	modules [][]*bool
}

func (out *BaseOutput) GetModule(x int,y int) bool{
	//TODO 此处需返回是否有值,不是返0/1
	return out.modules[x][y] != nil
}

// Output : the output interface for qrcode print
type Output interface {
	GetBaseOutput() *BaseOutput
	// Init :init for output when size is AUTO_SIZE
	Init(version *model.Version,qz *model.QuietZone)
	Write(x int,y int, black bool)
	// WriteModule :write per module by pixelSize
	WriteModule(x int,y int, black bool,pixelSize int)
	WriteModuleColor(x int,y int, setColor color.Color,pixelSize int)
	// GetModule : x,y is module axes , not pixel axes.
	GetModule(x int,y int) bool
	GetImage() *image.RGBA
	Save(fileName string) error
	Clone() Output
	DrawIntoNewImage(minPoint image.Point,maxPoint image.Point)
	Resize()
}


// EvalPenalty :Evaluate Penalty.
// TODO 重新检查惩罚值计算
func (out *BaseOutput) EvalPenalty(moduleSize int) uint {
	return out.evalPenaltyRule1(moduleSize) +
		out.evalPenaltyRule2(moduleSize) +
		out.evalPenaltyRule3(moduleSize) +
		out.evalPenaltyRule4(moduleSize)
}

//TODO need implement Micro QR Code penalty

func (out *BaseOutput) evalPenaltyRule1(moduleSize int) uint {
	var result uint
	for x := 0; x < moduleSize; x++ {
		checkForX := false
		var cntX uint
		checkForY := false
		var cntY uint
		for y := 0; y <moduleSize; y++ {
			if out.GetModule(x, y) == checkForX {
				cntX++
			} else {
				checkForX = !checkForX
				if cntX >= 5 {
					result += cntX - 2
				}
				cntX = 1
			}
			if out.GetModule(x, y) == checkForY {
				cntY++
			} else {
				checkForY = !checkForY
				if cntY >= 5 {
					result += cntY - 2
				}
				cntY = 1
			}
		}
		if cntX >= 5 {
			result += cntX - 2
		}
		if cntY >= 5 {
			result += cntY - 2
		}
	}
	fmt.Printf("%p evalPenaltyRule1 %d\n:",out,result)
	return result
}

func (out *BaseOutput) evalPenaltyRule2(moduleSize int) uint {
	var result uint
	for x := 0; x <moduleSize-1; x++ {
		for y := 0; y <moduleSize-1; y++ {
			val := out.GetModule(x, y)
			if out.GetModule(x, y+1) == val && out.GetModule(x+1, y) == val && out.GetModule(x+1, y+1) == val {
				result += 3
			}
		}
	}
	fmt.Printf("%p evalPenaltyRule2 %d\n:",out,result)
	return result
}

func (out *BaseOutput) evalPenaltyRule3(moduleSize int) uint {
	pattern1 := []bool{true, false, true, true, true, false, true, false, false, false, false}
	pattern2 := []bool{false, false, false, false, true, false, true, true, true, false, true}
	var result uint
	for x := 0; x <=moduleSize - len(pattern1); x++ {
		for y := 0; y <moduleSize; y++ {
			pattern1XFound := true
			pattern2XFound := true
			pattern1YFound := true
			pattern2YFound := true

			for i := 0; i < len(pattern1); i++ {
				iv := out.GetModule(x+i, y)
				if iv != pattern1[i] {
					pattern1XFound = false
				}
				if iv != pattern2[i] {
					pattern2XFound = false
				}
				iv = out.GetModule(y, x+i)
				if iv != pattern1[i] {
					pattern1YFound = false
				}
				if iv != pattern2[i] {
					pattern2YFound = false
				}
			}
			if pattern1XFound || pattern2XFound {
				result += 40
			}
			if pattern1YFound || pattern2YFound {
				result += 40
			}
		}
	}
	fmt.Printf("%p evalPenaltyRule2 %d\n:",out,result)
	return result
}

func (out *BaseOutput) evalPenaltyRule4(moduleSize int) uint {
	blankCount := 0
	for  row:= 0; row < moduleSize; row++ {
		for col := 0; col < moduleSize; col++ {
			if out.GetModule(row,col) {
				blankCount++
			}
		}
	}
	percDark := float64(blankCount) * 100 / float64(moduleSize)
	floor := math.Abs(math.Floor(percDark/5) - 10)
	ceil := math.Abs(math.Ceil(percDark/5) - 10)
	fmt.Printf("%p evalPenaltyRule4 floor:%f ceil:%f \n:",out,floor,ceil)
	return uint(math.Min(floor, ceil) * 10)
}